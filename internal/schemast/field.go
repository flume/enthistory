// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schemast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"entgo.io/ent/schema/field"
)

var replacer = strings.NewReplacer("interface {}", "any", "interface{}", "any")

// Field converts a *field.Descriptor back into an *ast.CallExpr of the ent field package that can be used
// to construct it.
func Field(desc *field.Descriptor) (*ast.CallExpr, error) {
	switch t := desc.Info.Type; {
	case t.Numeric(), t == field.TypeString, t == field.TypeBool, t == field.TypeTime, t == field.TypeBytes:
		return fromSimpleType(desc, true)
	case t == field.TypeUUID:
		call, err := fromSimpleType(desc, false)
		if err != nil {
			return nil, err
		}
		return fromComplexType(
			call,
			structLit(
				&ast.SelectorExpr{
					X:   ast.NewIdent("uuid"),
					Sel: ast.NewIdent("UUID"),
				},
			))
	case t == field.TypeJSON:
		expr := "struct{}{}"
		if desc.Info != nil && desc.Info.RType != nil {
			expr = desc.Info.RType.Ident + "{}"
			if desc.Info.RType.Kind == reflect.Pointer {
				expr = "&" + expr
			}
		}
		expr = replacer.Replace(expr)
		exp, err := parser.ParseExpr(expr)
		if err != nil {
			return nil, fmt.Errorf("schemast: json field %s generation error %w", desc.Name, err)
		}
		if c, ok := exp.(*ast.CompositeLit); ok {
			if v, ok := c.Type.(*ast.StructType); ok {
				v.Fields = &ast.FieldList{
					Opening: 1,
					Closing: 1,
				}
			}
		}
		call, err := fromSimpleType(desc, false)
		if err != nil {
			return nil, err
		}
		return fromComplexType(
			call,
			exp,
		)
	case t == field.TypeEnum:
		return fromEnumType(desc)
	default:
		return nil, fmt.Errorf("schemast: unsupported type %s", t.ConstName())
	}
}

// AppendField adds a field to the returned values of the Fields method of type typeName.
func (c *Context) AppendField(typeName string, desc *field.Descriptor) error {
	newField, err := Field(desc)
	if err != nil {
		return err
	}

	// wrap ast.CallExpression with a newline at the beginning
	return c.appendReturnItem(kindField, typeName, newField)
}

// RemoveField removes a field from the returned values of the Fields method of type typeName.
func (c *Context) RemoveField(typeName string, fieldName string) error {
	stmt, err := c.returnStmt(typeName, "Fields")
	if err != nil {
		return err
	}
	returned, ok := stmt.Results[0].(*ast.CompositeLit)
	if !ok {
		return fmt.Errorf("schemast: unexpected AST component type %T", stmt.Results[0])
	}
	for i, item := range returned.Elts {
		call, ok := item.(*ast.CallExpr)
		if !ok {
			return fmt.Errorf("schemast: expected return statement elements to be call expressions")
		}
		name, err := extractFieldName(call)
		if err != nil {
			return err
		}
		if name == fieldName {
			returned.Elts = append(returned.Elts[:i], returned.Elts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("schemast: could not find field %q in type %q", fieldName, typeName)
}

func newFieldCall(desc *field.Descriptor) *builderCall {
	return &builderCall{
		curr: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("\nfield"),
				Sel: ast.NewIdent(fieldConstructor(desc)),
			},
			Args: []ast.Expr{
				strLit(desc.Name),
			},
		},
	}
}

func fromEnumType(desc *field.Descriptor) (*ast.CallExpr, error) {
	info := desc.Info
	_ = info
	call, err := fromSimpleType(desc, false)
	if err != nil {
		return nil, err
	}
	modifier := "Values"
	for _, pair := range desc.Enums {
		if len(strings.Split(info.Ident, ".")) > 1 {
			modifier = "GoType"
			break
		}
		if pair.N != pair.V {
			modifier = "NamedValues"
			break
		}
	}
	args := make([]ast.Expr, 0)
	if modifier == "GoType" {
		parts := strings.Split(info.Ident, ".")
		args = append(args, fnCall(selectorLit(parts[0], parts[1]), strLit("")))
	} else {
		for _, pair := range desc.Enums {
			args = append(args, strLit(pair.N))
			if modifier == "NamedValues" {
				args = append(args, strLit(pair.V))
			}
		}
	}
	builder := &builderCall{curr: call}
	builder.method(modifier, args...)
	return builder.curr, nil
}

func fromComplexType(call *ast.CallExpr, filedType ast.Expr) (*ast.CallExpr, error) {
	callExpr := call
	// Loop through calls to find the base and append the filedType there
	for {
		if selectExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			if prevExpr, ok := selectExpr.X.(*ast.CallExpr); ok {
				callExpr = prevExpr
			} else {
				break
			}
		} else {
			break
		}
	}
	// Append the filedType to the args of the initial *ast.CallExpr
	callExpr.Args = append(callExpr.Args, filedType)
	return call, nil
}

func fromSimpleType(desc *field.Descriptor, rtype bool) (*ast.CallExpr, error) {
	builder := newFieldCall(desc)
	t := desc.Info.Type
	if rtype && desc.Info.RType != nil {
		var defaultValue ast.Expr
		if t.Numeric() {
			defaultValue = intLit(0)
			if desc.Default != nil {
				if def, ok := desc.Default.(int); ok {
					defaultValue = intLit(def)
				}
			}
		} else if t == field.TypeString {
			defaultValue = strLit("")
			if desc.Default != nil {
				if def, ok := desc.Default.(string); ok {
					defaultValue = strLit(def)
				}
			}
		} else if t == field.TypeBool {
			defaultValue = boolLit(false)
			if desc.Default != nil {
				if def, ok := desc.Default.(bool); ok {
					defaultValue = boolLit(def)
				}
			}
		} else if t == field.TypeTime {
			defaultValue = structLit(&ast.SelectorExpr{
				X:   ast.NewIdent("time"),
				Sel: ast.NewIdent("Time"),
			})
		} else if t == field.TypeBytes {
			defaultValue = boolSliceLit([]byte{})
			if desc.Default != nil {
				if def, ok := desc.Default.([]byte); ok {
					defaultValue = boolSliceLit(def)
				}
			}
		} else {
			return nil, fmt.Errorf("schemast: unsupported type %s", t.ConstName())
		}
		builder.method("GoType", initLit(desc.Info.RType.Ident, []ast.Expr{defaultValue}))
	}
	if desc.Nillable {
		builder.method("Nillable")
	}
	if desc.Optional {
		builder.method("Optional")
	}
	if desc.Unique {
		builder.method("Unique")
	}
	if desc.Sensitive {
		builder.method("Sensitive")
	}
	if desc.Immutable {
		builder.method("Immutable")
	}
	if desc.Comment != "" {
		builder.method("Comment", strLit(desc.Comment))
	}
	if desc.Tag != "" {
		builder.method("StructTag", strLit(desc.Tag))
	}
	if desc.StorageKey != "" {
		builder.method("StorageKey", strLit(desc.StorageKey))
	}
	if len(desc.SchemaType) > 0 {
		builder.method("SchemaType", strMapLit(desc.SchemaType))
	}
	if len(desc.Annotations) != 0 {
		annots, err := toAnnotASTs(desc.Annotations)
		if err != nil {
			return nil, err
		}
		builder.annotate(annots...)
	}
	if desc.Default != nil {
		hasDefaultFunc := t.Numeric() || t == field.TypeString || t == field.TypeBytes
		method, expr, err := defaultExpr(desc.Default, hasDefaultFunc)
		if err != nil {
			return nil, err
		}
		builder.method(method, expr)
	}
	if desc.UpdateDefault != nil {
		_, expr, err := defaultExpr(desc.UpdateDefault, false)
		if err != nil {
			return nil, err
		}
		builder.method("UpdateDefault", expr)
	}

	// Unsupported features
	var unsupported error
	if len(desc.Validators) != 0 {
		unsupported = combineUnsupported(unsupported, "Descriptor.Validators")
	}
	if unsupported != nil {
		return nil, unsupported
	}
	return builder.curr, nil
}

func fieldConstructor(dsc *field.Descriptor) string {
	cn := dsc.Info.ConstName()
	if dsc.Info.Type == field.TypeFloat64 {
		cn = strings.TrimSuffix(cn, "64")
	}
	return strings.TrimPrefix(cn, "Type")
}

const (
	builderMethodDefault     = "Default"
	builderMethodDefaultFunc = "DefaultFunc"
)

func defaultExpr(d any, hasDefaultFunc bool) (string, ast.Expr, error) {
	v := reflect.ValueOf(d)
	switch v.Kind() {
	case reflect.String:
		return builderMethodDefault, strLit(d.(string)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		lit := &ast.BasicLit{
			Kind:  token.INT,
			Value: fmt.Sprintf("%d", d),
		}
		return builderMethodDefault, lit, nil
	case reflect.Float32, reflect.Float64:
		lit := &ast.BasicLit{
			Kind:  token.FLOAT,
			Value: fmt.Sprintf("%#v", d),
		}
		return builderMethodDefault, lit, nil
	case reflect.Bool:
		lit := &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.FormatBool(d.(bool)),
		}
		return builderMethodDefault, lit, nil
	case reflect.Func:
		f := runtime.FuncForPC(v.Pointer()).Name()
		pkg := strings.Split(f, "/")
		if len(pkg) > 1 {
			f = pkg[len(pkg)-1]
		}
		parts := strings.Split(f, ".")
		if len(parts) == 2 {
			// Named function like uuid.New
			selector := selectorLit(parts[0], parts[1])
			if !hasDefaultFunc {
				return builderMethodDefault, selector, nil
			}
			return builderMethodDefaultFunc, selector, nil
		}
		// Anonymous function - try to extract from source
		funcLit, err := extractFuncLitFromSource(v.Pointer())
		if err != nil {
			return "", nil, fmt.Errorf("schemast: failed to extract anonymous function: %w", err)
		}
		if !hasDefaultFunc {
			return builderMethodDefault, funcLit, nil
		}
		return builderMethodDefaultFunc, funcLit, nil
	default:
		return "", nil, fmt.Errorf("schemast: unsupported default field kind: %q", v.Kind())
	}
}

// extractFuncLitFromSource extracts an anonymous function literal from its source file
// by using runtime information to locate the function definition.
func extractFuncLitFromSource(ptr uintptr) (*ast.FuncLit, error) {
	fn := runtime.FuncForPC(ptr)
	if fn == nil {
		return nil, fmt.Errorf("could not get function info")
	}

	file, line := fn.FileLine(ptr)
	if file == "" {
		return nil, fmt.Errorf("could not get source file location")
	}

	// Read and parse the source file
	src, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read source file %s: %w", file, err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("could not parse source file %s: %w", file, err)
	}

	// Find the function literal at or near the given line
	var foundFuncLit *ast.FuncLit
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if funcLit, ok := n.(*ast.FuncLit); ok {
			pos := fset.Position(funcLit.Pos())
			// Check if this function literal is at or near our target line
			// We check a range because the reported line might be the func keyword
			// or the opening brace
			if pos.Line >= line-1 && pos.Line <= line+1 {
				foundFuncLit = funcLit
				return false // Stop searching
			}
		}
		return true
	})

	if foundFuncLit == nil {
		return nil, fmt.Errorf("could not find function literal at %s:%d", file, line)
	}

	// Create a clean copy of the function literal without position information
	// This ensures it can be properly printed in the generated code
	return cleanFuncLit(foundFuncLit), nil
}

// cleanFuncLit creates a copy of a FuncLit with cleaned position information
func cleanFuncLit(fl *ast.FuncLit) *ast.FuncLit {
	return &ast.FuncLit{
		Type: cleanFuncType(fl.Type),
		Body: cleanBlockStmt(fl.Body),
	}
}

func cleanFuncType(ft *ast.FuncType) *ast.FuncType {
	return &ast.FuncType{
		Params:  cleanFieldList(ft.Params),
		Results: cleanFieldList(ft.Results),
	}
}

func cleanFieldList(fl *ast.FieldList) *ast.FieldList {
	if fl == nil {
		return nil
	}
	newList := &ast.FieldList{}
	for _, field := range fl.List {
		newList.List = append(newList.List, cleanField(field))
	}
	return newList
}

func cleanField(f *ast.Field) *ast.Field {
	newField := &ast.Field{
		Type: f.Type, // Keep type as-is for now
	}
	for _, name := range f.Names {
		newField.Names = append(newField.Names, ast.NewIdent(name.Name))
	}
	return newField
}

func cleanBlockStmt(bs *ast.BlockStmt) *ast.BlockStmt {
	if bs == nil {
		return nil
	}
	newBlock := &ast.BlockStmt{}
	newBlock.List = append(newBlock.List, bs.List...)
	return newBlock
}

func extractFieldName(fd *ast.CallExpr) (string, error) {
	sel, ok := fd.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", fmt.Errorf("schemast: unexpected type %T", fd.Fun)
	}
	if inner, ok := sel.X.(*ast.CallExpr); ok {
		return extractFieldName(inner)
	}
	if final, ok := sel.X.(*ast.Ident); ok && final.Name != "field" {
		return "", fmt.Errorf(`schemast: expected field AST to be of form field.<Type>("name")`)
	}
	if len(fd.Args) == 0 {
		return "", fmt.Errorf("schemast: expected field constructor to have at least name arg")
	}
	name, ok := fd.Args[0].(*ast.BasicLit)
	if !ok && name.Kind == token.STRING {
		return "", fmt.Errorf("schemast: expected field name to be a string literal")
	}
	return strconv.Unquote(name.Value)
}
