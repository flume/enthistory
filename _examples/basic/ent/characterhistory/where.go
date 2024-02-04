// Code generated by ent, DO NOT EDIT.

package characterhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/flume/enthistory"
	"github.com/flume/enthistory/_examples/basic/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldRef, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldAge, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldName, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v enthistory.OpType) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v enthistory.OpType) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...enthistory.OpType) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...enthistory.OpType) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotNull(FieldRef))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldAge, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldContainsFold(FieldName, v))
}

// NicknamesIsNil applies the IsNil predicate on the "nicknames" field.
func NicknamesIsNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIsNull(FieldNicknames))
}

// NicknamesNotNil applies the NotNil predicate on the "nicknames" field.
func NicknamesNotNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotNull(FieldNicknames))
}

// InfoIsNil applies the IsNil predicate on the "info" field.
func InfoIsNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldIsNull(FieldInfo))
}

// InfoNotNil applies the NotNil predicate on the "info" field.
func InfoNotNil() predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.FieldNotNull(FieldInfo))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CharacterHistory) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CharacterHistory) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CharacterHistory) predicate.CharacterHistory {
	return predicate.CharacterHistory(sql.NotPredicates(p))
}
