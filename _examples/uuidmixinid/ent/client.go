// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"_examples/uuidmixinid/ent/migrate"

	"_examples/uuidmixinid/ent/menuitem"
	"_examples/uuidmixinid/ent/menuitemhistory"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// MenuItem is the client for interacting with the MenuItem builders.
	MenuItem *MenuItemClient
	// MenuItemHistory is the client for interacting with the MenuItemHistory builders.
	MenuItemHistory *MenuItemHistoryClient
	// historyActivated determines if the history hooks have already been activated
	historyActivated bool
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.MenuItem = NewMenuItemClient(c.config)
	c.MenuItemHistory = NewMenuItemHistoryClient(c.config)
}

// withHistory adds the history hooks to the appropriate schemas - generated by enthistory
func (c *Client) WithHistory() {
	if !c.historyActivated {

		// MenuItem hooks
		c.MenuItem.Use(enthistory.HistoryTriggerInsert[*MenuItemMutation]())

		c.historyActivated = true
	}
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		MenuItem:        NewMenuItemClient(cfg),
		MenuItemHistory: NewMenuItemHistoryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		MenuItem:        NewMenuItemClient(cfg),
		MenuItemHistory: NewMenuItemHistoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		MenuItem.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.MenuItem.Use(hooks...)
	c.MenuItemHistory.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.MenuItem.Intercept(interceptors...)
	c.MenuItemHistory.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MenuItemMutation:
		return c.MenuItem.mutate(ctx, m)
	case *MenuItemHistoryMutation:
		return c.MenuItemHistory.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MenuItemClient is a client for the MenuItem schema.
type MenuItemClient struct {
	config
}

// NewMenuItemClient returns a client for the MenuItem from the given config.
func NewMenuItemClient(c config) *MenuItemClient {
	return &MenuItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `menuitem.Hooks(f(g(h())))`.
func (c *MenuItemClient) Use(hooks ...Hook) {
	c.hooks.MenuItem = append(c.hooks.MenuItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `menuitem.Intercept(f(g(h())))`.
func (c *MenuItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.MenuItem = append(c.inters.MenuItem, interceptors...)
}

// Create returns a builder for creating a MenuItem entity.
func (c *MenuItemClient) Create() *MenuItemCreate {
	mutation := newMenuItemMutation(c.config, OpCreate)
	return &MenuItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MenuItem entities.
func (c *MenuItemClient) CreateBulk(builders ...*MenuItemCreate) *MenuItemCreateBulk {
	return &MenuItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MenuItemClient) MapCreateBulk(slice any, setFunc func(*MenuItemCreate, int)) *MenuItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MenuItemCreateBulk{err: fmt.Errorf("calling to MenuItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MenuItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MenuItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MenuItem.
func (c *MenuItemClient) Update() *MenuItemUpdate {
	mutation := newMenuItemMutation(c.config, OpUpdate)
	return &MenuItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MenuItemClient) UpdateOne(mi *MenuItem) *MenuItemUpdateOne {
	mutation := newMenuItemMutation(c.config, OpUpdateOne, withMenuItem(mi))
	return &MenuItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MenuItemClient) UpdateOneID(id uuid.UUID) *MenuItemUpdateOne {
	mutation := newMenuItemMutation(c.config, OpUpdateOne, withMenuItemID(id))
	return &MenuItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MenuItem.
func (c *MenuItemClient) Delete() *MenuItemDelete {
	mutation := newMenuItemMutation(c.config, OpDelete)
	return &MenuItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MenuItemClient) DeleteOne(mi *MenuItem) *MenuItemDeleteOne {
	return c.DeleteOneID(mi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MenuItemClient) DeleteOneID(id uuid.UUID) *MenuItemDeleteOne {
	builder := c.Delete().Where(menuitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MenuItemDeleteOne{builder}
}

// Query returns a query builder for MenuItem.
func (c *MenuItemClient) Query() *MenuItemQuery {
	return &MenuItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMenuItem},
		inters: c.Interceptors(),
	}
}

// Get returns a MenuItem entity by its id.
func (c *MenuItemClient) Get(ctx context.Context, id uuid.UUID) (*MenuItem, error) {
	return c.Query().Where(menuitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MenuItemClient) GetX(ctx context.Context, id uuid.UUID) *MenuItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MenuItemClient) Hooks() []Hook {
	return c.hooks.MenuItem
}

// Interceptors returns the client interceptors.
func (c *MenuItemClient) Interceptors() []Interceptor {
	return c.inters.MenuItem
}

func (c *MenuItemClient) mutate(ctx context.Context, m *MenuItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MenuItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MenuItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MenuItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MenuItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown MenuItem mutation op: %q", m.Op())
	}
}

// MenuItemHistoryClient is a client for the MenuItemHistory schema.
type MenuItemHistoryClient struct {
	config
}

// NewMenuItemHistoryClient returns a client for the MenuItemHistory from the given config.
func NewMenuItemHistoryClient(c config) *MenuItemHistoryClient {
	return &MenuItemHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `menuitemhistory.Hooks(f(g(h())))`.
func (c *MenuItemHistoryClient) Use(hooks ...Hook) {
	c.hooks.MenuItemHistory = append(c.hooks.MenuItemHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `menuitemhistory.Intercept(f(g(h())))`.
func (c *MenuItemHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.MenuItemHistory = append(c.inters.MenuItemHistory, interceptors...)
}

// Create returns a builder for creating a MenuItemHistory entity.
func (c *MenuItemHistoryClient) Create() *MenuItemHistoryCreate {
	mutation := newMenuItemHistoryMutation(c.config, OpCreate)
	return &MenuItemHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MenuItemHistory entities.
func (c *MenuItemHistoryClient) CreateBulk(builders ...*MenuItemHistoryCreate) *MenuItemHistoryCreateBulk {
	return &MenuItemHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MenuItemHistoryClient) MapCreateBulk(slice any, setFunc func(*MenuItemHistoryCreate, int)) *MenuItemHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MenuItemHistoryCreateBulk{err: fmt.Errorf("calling to MenuItemHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MenuItemHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MenuItemHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MenuItemHistory.
func (c *MenuItemHistoryClient) Update() *MenuItemHistoryUpdate {
	mutation := newMenuItemHistoryMutation(c.config, OpUpdate)
	return &MenuItemHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MenuItemHistoryClient) UpdateOne(mih *MenuItemHistory) *MenuItemHistoryUpdateOne {
	mutation := newMenuItemHistoryMutation(c.config, OpUpdateOne, withMenuItemHistory(mih))
	return &MenuItemHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MenuItemHistoryClient) UpdateOneID(id uuid.UUID) *MenuItemHistoryUpdateOne {
	mutation := newMenuItemHistoryMutation(c.config, OpUpdateOne, withMenuItemHistoryID(id))
	return &MenuItemHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MenuItemHistory.
func (c *MenuItemHistoryClient) Delete() *MenuItemHistoryDelete {
	mutation := newMenuItemHistoryMutation(c.config, OpDelete)
	return &MenuItemHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MenuItemHistoryClient) DeleteOne(mih *MenuItemHistory) *MenuItemHistoryDeleteOne {
	return c.DeleteOneID(mih.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MenuItemHistoryClient) DeleteOneID(id uuid.UUID) *MenuItemHistoryDeleteOne {
	builder := c.Delete().Where(menuitemhistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MenuItemHistoryDeleteOne{builder}
}

// Query returns a query builder for MenuItemHistory.
func (c *MenuItemHistoryClient) Query() *MenuItemHistoryQuery {
	return &MenuItemHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMenuItemHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a MenuItemHistory entity by its id.
func (c *MenuItemHistoryClient) Get(ctx context.Context, id uuid.UUID) (*MenuItemHistory, error) {
	return c.Query().Where(menuitemhistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MenuItemHistoryClient) GetX(ctx context.Context, id uuid.UUID) *MenuItemHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MenuItemHistoryClient) Hooks() []Hook {
	return c.hooks.MenuItemHistory
}

// Interceptors returns the client interceptors.
func (c *MenuItemHistoryClient) Interceptors() []Interceptor {
	return c.inters.MenuItemHistory
}

func (c *MenuItemHistoryClient) mutate(ctx context.Context, m *MenuItemHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MenuItemHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MenuItemHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MenuItemHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MenuItemHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown MenuItemHistory mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		MenuItem, MenuItemHistory []ent.Hook
	}
	inters struct {
		MenuItem, MenuItemHistory []ent.Interceptor
	}
)