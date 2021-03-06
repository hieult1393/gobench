// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/gobench-io/gobench/ent/migrate"

	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/ent/counter"
	"github.com/gobench-io/gobench/ent/gauge"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/group"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/metric"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Application is the client for interacting with the Application builders.
	Application *ApplicationClient
	// Counter is the client for interacting with the Counter builders.
	Counter *CounterClient
	// Gauge is the client for interacting with the Gauge builders.
	Gauge *GaugeClient
	// Graph is the client for interacting with the Graph builders.
	Graph *GraphClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Histogram is the client for interacting with the Histogram builders.
	Histogram *HistogramClient
	// Metric is the client for interacting with the Metric builders.
	Metric *MetricClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Application = NewApplicationClient(c.config)
	c.Counter = NewCounterClient(c.config)
	c.Gauge = NewGaugeClient(c.config)
	c.Graph = NewGraphClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Histogram = NewHistogramClient(c.config)
	c.Metric = NewMetricClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Application: NewApplicationClient(cfg),
		Counter:     NewCounterClient(cfg),
		Gauge:       NewGaugeClient(cfg),
		Graph:       NewGraphClient(cfg),
		Group:       NewGroupClient(cfg),
		Histogram:   NewHistogramClient(cfg),
		Metric:      NewMetricClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:      cfg,
		Application: NewApplicationClient(cfg),
		Counter:     NewCounterClient(cfg),
		Gauge:       NewGaugeClient(cfg),
		Graph:       NewGraphClient(cfg),
		Group:       NewGroupClient(cfg),
		Histogram:   NewHistogramClient(cfg),
		Metric:      NewMetricClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Application.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Application.Use(hooks...)
	c.Counter.Use(hooks...)
	c.Gauge.Use(hooks...)
	c.Graph.Use(hooks...)
	c.Group.Use(hooks...)
	c.Histogram.Use(hooks...)
	c.Metric.Use(hooks...)
}

// ApplicationClient is a client for the Application schema.
type ApplicationClient struct {
	config
}

// NewApplicationClient returns a client for the Application from the given config.
func NewApplicationClient(c config) *ApplicationClient {
	return &ApplicationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `application.Hooks(f(g(h())))`.
func (c *ApplicationClient) Use(hooks ...Hook) {
	c.hooks.Application = append(c.hooks.Application, hooks...)
}

// Create returns a create builder for Application.
func (c *ApplicationClient) Create() *ApplicationCreate {
	mutation := newApplicationMutation(c.config, OpCreate)
	return &ApplicationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Application entities.
func (c *ApplicationClient) CreateBulk(builders ...*ApplicationCreate) *ApplicationCreateBulk {
	return &ApplicationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Application.
func (c *ApplicationClient) Update() *ApplicationUpdate {
	mutation := newApplicationMutation(c.config, OpUpdate)
	return &ApplicationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApplicationClient) UpdateOne(a *Application) *ApplicationUpdateOne {
	mutation := newApplicationMutation(c.config, OpUpdateOne, withApplication(a))
	return &ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApplicationClient) UpdateOneID(id int) *ApplicationUpdateOne {
	mutation := newApplicationMutation(c.config, OpUpdateOne, withApplicationID(id))
	return &ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Application.
func (c *ApplicationClient) Delete() *ApplicationDelete {
	mutation := newApplicationMutation(c.config, OpDelete)
	return &ApplicationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ApplicationClient) DeleteOne(a *Application) *ApplicationDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ApplicationClient) DeleteOneID(id int) *ApplicationDeleteOne {
	builder := c.Delete().Where(application.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApplicationDeleteOne{builder}
}

// Query returns a query builder for Application.
func (c *ApplicationClient) Query() *ApplicationQuery {
	return &ApplicationQuery{config: c.config}
}

// Get returns a Application entity by its id.
func (c *ApplicationClient) Get(ctx context.Context, id int) (*Application, error) {
	return c.Query().Where(application.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApplicationClient) GetX(ctx context.Context, id int) *Application {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// QueryGroups queries the groups edge of a Application.
func (c *ApplicationClient) QueryGroups(a *Application) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(application.Table, application.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, application.GroupsTable, application.GroupsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ApplicationClient) Hooks() []Hook {
	return c.hooks.Application
}

// CounterClient is a client for the Counter schema.
type CounterClient struct {
	config
}

// NewCounterClient returns a client for the Counter from the given config.
func NewCounterClient(c config) *CounterClient {
	return &CounterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `counter.Hooks(f(g(h())))`.
func (c *CounterClient) Use(hooks ...Hook) {
	c.hooks.Counter = append(c.hooks.Counter, hooks...)
}

// Create returns a create builder for Counter.
func (c *CounterClient) Create() *CounterCreate {
	mutation := newCounterMutation(c.config, OpCreate)
	return &CounterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Counter entities.
func (c *CounterClient) CreateBulk(builders ...*CounterCreate) *CounterCreateBulk {
	return &CounterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Counter.
func (c *CounterClient) Update() *CounterUpdate {
	mutation := newCounterMutation(c.config, OpUpdate)
	return &CounterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CounterClient) UpdateOne(co *Counter) *CounterUpdateOne {
	mutation := newCounterMutation(c.config, OpUpdateOne, withCounter(co))
	return &CounterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CounterClient) UpdateOneID(id int) *CounterUpdateOne {
	mutation := newCounterMutation(c.config, OpUpdateOne, withCounterID(id))
	return &CounterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Counter.
func (c *CounterClient) Delete() *CounterDelete {
	mutation := newCounterMutation(c.config, OpDelete)
	return &CounterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CounterClient) DeleteOne(co *Counter) *CounterDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CounterClient) DeleteOneID(id int) *CounterDeleteOne {
	builder := c.Delete().Where(counter.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CounterDeleteOne{builder}
}

// Query returns a query builder for Counter.
func (c *CounterClient) Query() *CounterQuery {
	return &CounterQuery{config: c.config}
}

// Get returns a Counter entity by its id.
func (c *CounterClient) Get(ctx context.Context, id int) (*Counter, error) {
	return c.Query().Where(counter.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CounterClient) GetX(ctx context.Context, id int) *Counter {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// QueryMetric queries the metric edge of a Counter.
func (c *CounterClient) QueryMetric(co *Counter) *MetricQuery {
	query := &MetricQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(counter.Table, counter.FieldID, id),
			sqlgraph.To(metric.Table, metric.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, counter.MetricTable, counter.MetricColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CounterClient) Hooks() []Hook {
	return c.hooks.Counter
}

// GaugeClient is a client for the Gauge schema.
type GaugeClient struct {
	config
}

// NewGaugeClient returns a client for the Gauge from the given config.
func NewGaugeClient(c config) *GaugeClient {
	return &GaugeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gauge.Hooks(f(g(h())))`.
func (c *GaugeClient) Use(hooks ...Hook) {
	c.hooks.Gauge = append(c.hooks.Gauge, hooks...)
}

// Create returns a create builder for Gauge.
func (c *GaugeClient) Create() *GaugeCreate {
	mutation := newGaugeMutation(c.config, OpCreate)
	return &GaugeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Gauge entities.
func (c *GaugeClient) CreateBulk(builders ...*GaugeCreate) *GaugeCreateBulk {
	return &GaugeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Gauge.
func (c *GaugeClient) Update() *GaugeUpdate {
	mutation := newGaugeMutation(c.config, OpUpdate)
	return &GaugeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GaugeClient) UpdateOne(ga *Gauge) *GaugeUpdateOne {
	mutation := newGaugeMutation(c.config, OpUpdateOne, withGauge(ga))
	return &GaugeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GaugeClient) UpdateOneID(id int) *GaugeUpdateOne {
	mutation := newGaugeMutation(c.config, OpUpdateOne, withGaugeID(id))
	return &GaugeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Gauge.
func (c *GaugeClient) Delete() *GaugeDelete {
	mutation := newGaugeMutation(c.config, OpDelete)
	return &GaugeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GaugeClient) DeleteOne(ga *Gauge) *GaugeDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GaugeClient) DeleteOneID(id int) *GaugeDeleteOne {
	builder := c.Delete().Where(gauge.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GaugeDeleteOne{builder}
}

// Query returns a query builder for Gauge.
func (c *GaugeClient) Query() *GaugeQuery {
	return &GaugeQuery{config: c.config}
}

// Get returns a Gauge entity by its id.
func (c *GaugeClient) Get(ctx context.Context, id int) (*Gauge, error) {
	return c.Query().Where(gauge.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GaugeClient) GetX(ctx context.Context, id int) *Gauge {
	ga, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ga
}

// QueryMetric queries the metric edge of a Gauge.
func (c *GaugeClient) QueryMetric(ga *Gauge) *MetricQuery {
	query := &MetricQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gauge.Table, gauge.FieldID, id),
			sqlgraph.To(metric.Table, metric.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, gauge.MetricTable, gauge.MetricColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GaugeClient) Hooks() []Hook {
	return c.hooks.Gauge
}

// GraphClient is a client for the Graph schema.
type GraphClient struct {
	config
}

// NewGraphClient returns a client for the Graph from the given config.
func NewGraphClient(c config) *GraphClient {
	return &GraphClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `graph.Hooks(f(g(h())))`.
func (c *GraphClient) Use(hooks ...Hook) {
	c.hooks.Graph = append(c.hooks.Graph, hooks...)
}

// Create returns a create builder for Graph.
func (c *GraphClient) Create() *GraphCreate {
	mutation := newGraphMutation(c.config, OpCreate)
	return &GraphCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Graph entities.
func (c *GraphClient) CreateBulk(builders ...*GraphCreate) *GraphCreateBulk {
	return &GraphCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Graph.
func (c *GraphClient) Update() *GraphUpdate {
	mutation := newGraphMutation(c.config, OpUpdate)
	return &GraphUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GraphClient) UpdateOne(gr *Graph) *GraphUpdateOne {
	mutation := newGraphMutation(c.config, OpUpdateOne, withGraph(gr))
	return &GraphUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GraphClient) UpdateOneID(id int) *GraphUpdateOne {
	mutation := newGraphMutation(c.config, OpUpdateOne, withGraphID(id))
	return &GraphUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Graph.
func (c *GraphClient) Delete() *GraphDelete {
	mutation := newGraphMutation(c.config, OpDelete)
	return &GraphDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GraphClient) DeleteOne(gr *Graph) *GraphDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GraphClient) DeleteOneID(id int) *GraphDeleteOne {
	builder := c.Delete().Where(graph.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GraphDeleteOne{builder}
}

// Query returns a query builder for Graph.
func (c *GraphClient) Query() *GraphQuery {
	return &GraphQuery{config: c.config}
}

// Get returns a Graph entity by its id.
func (c *GraphClient) Get(ctx context.Context, id int) (*Graph, error) {
	return c.Query().Where(graph.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GraphClient) GetX(ctx context.Context, id int) *Graph {
	gr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gr
}

// QueryGroup queries the group edge of a Graph.
func (c *GraphClient) QueryGroup(gr *Graph) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(graph.Table, graph.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, graph.GroupTable, graph.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMetrics queries the metrics edge of a Graph.
func (c *GraphClient) QueryMetrics(gr *Graph) *MetricQuery {
	query := &MetricQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(graph.Table, graph.FieldID, id),
			sqlgraph.To(metric.Table, metric.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, graph.MetricsTable, graph.MetricsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GraphClient) Hooks() []Hook {
	return c.hooks.Graph
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	gr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gr
}

// QueryApplication queries the application edge of a Group.
func (c *GroupClient) QueryApplication(gr *Group) *ApplicationQuery {
	query := &ApplicationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, group.ApplicationTable, group.ApplicationColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGraphs queries the graphs edge of a Group.
func (c *GroupClient) QueryGraphs(gr *Group) *GraphQuery {
	query := &GraphQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(graph.Table, graph.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.GraphsTable, group.GraphsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// HistogramClient is a client for the Histogram schema.
type HistogramClient struct {
	config
}

// NewHistogramClient returns a client for the Histogram from the given config.
func NewHistogramClient(c config) *HistogramClient {
	return &HistogramClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `histogram.Hooks(f(g(h())))`.
func (c *HistogramClient) Use(hooks ...Hook) {
	c.hooks.Histogram = append(c.hooks.Histogram, hooks...)
}

// Create returns a create builder for Histogram.
func (c *HistogramClient) Create() *HistogramCreate {
	mutation := newHistogramMutation(c.config, OpCreate)
	return &HistogramCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Histogram entities.
func (c *HistogramClient) CreateBulk(builders ...*HistogramCreate) *HistogramCreateBulk {
	return &HistogramCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Histogram.
func (c *HistogramClient) Update() *HistogramUpdate {
	mutation := newHistogramMutation(c.config, OpUpdate)
	return &HistogramUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HistogramClient) UpdateOne(h *Histogram) *HistogramUpdateOne {
	mutation := newHistogramMutation(c.config, OpUpdateOne, withHistogram(h))
	return &HistogramUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HistogramClient) UpdateOneID(id int) *HistogramUpdateOne {
	mutation := newHistogramMutation(c.config, OpUpdateOne, withHistogramID(id))
	return &HistogramUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Histogram.
func (c *HistogramClient) Delete() *HistogramDelete {
	mutation := newHistogramMutation(c.config, OpDelete)
	return &HistogramDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HistogramClient) DeleteOne(h *Histogram) *HistogramDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HistogramClient) DeleteOneID(id int) *HistogramDeleteOne {
	builder := c.Delete().Where(histogram.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HistogramDeleteOne{builder}
}

// Query returns a query builder for Histogram.
func (c *HistogramClient) Query() *HistogramQuery {
	return &HistogramQuery{config: c.config}
}

// Get returns a Histogram entity by its id.
func (c *HistogramClient) Get(ctx context.Context, id int) (*Histogram, error) {
	return c.Query().Where(histogram.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HistogramClient) GetX(ctx context.Context, id int) *Histogram {
	h, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return h
}

// QueryMetric queries the metric edge of a Histogram.
func (c *HistogramClient) QueryMetric(h *Histogram) *MetricQuery {
	query := &MetricQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(histogram.Table, histogram.FieldID, id),
			sqlgraph.To(metric.Table, metric.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, histogram.MetricTable, histogram.MetricColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HistogramClient) Hooks() []Hook {
	return c.hooks.Histogram
}

// MetricClient is a client for the Metric schema.
type MetricClient struct {
	config
}

// NewMetricClient returns a client for the Metric from the given config.
func NewMetricClient(c config) *MetricClient {
	return &MetricClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `metric.Hooks(f(g(h())))`.
func (c *MetricClient) Use(hooks ...Hook) {
	c.hooks.Metric = append(c.hooks.Metric, hooks...)
}

// Create returns a create builder for Metric.
func (c *MetricClient) Create() *MetricCreate {
	mutation := newMetricMutation(c.config, OpCreate)
	return &MetricCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Metric entities.
func (c *MetricClient) CreateBulk(builders ...*MetricCreate) *MetricCreateBulk {
	return &MetricCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Metric.
func (c *MetricClient) Update() *MetricUpdate {
	mutation := newMetricMutation(c.config, OpUpdate)
	return &MetricUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MetricClient) UpdateOne(m *Metric) *MetricUpdateOne {
	mutation := newMetricMutation(c.config, OpUpdateOne, withMetric(m))
	return &MetricUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MetricClient) UpdateOneID(id int) *MetricUpdateOne {
	mutation := newMetricMutation(c.config, OpUpdateOne, withMetricID(id))
	return &MetricUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Metric.
func (c *MetricClient) Delete() *MetricDelete {
	mutation := newMetricMutation(c.config, OpDelete)
	return &MetricDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MetricClient) DeleteOne(m *Metric) *MetricDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MetricClient) DeleteOneID(id int) *MetricDeleteOne {
	builder := c.Delete().Where(metric.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MetricDeleteOne{builder}
}

// Query returns a query builder for Metric.
func (c *MetricClient) Query() *MetricQuery {
	return &MetricQuery{config: c.config}
}

// Get returns a Metric entity by its id.
func (c *MetricClient) Get(ctx context.Context, id int) (*Metric, error) {
	return c.Query().Where(metric.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MetricClient) GetX(ctx context.Context, id int) *Metric {
	m, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return m
}

// QueryGraph queries the graph edge of a Metric.
func (c *MetricClient) QueryGraph(m *Metric) *GraphQuery {
	query := &GraphQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(metric.Table, metric.FieldID, id),
			sqlgraph.To(graph.Table, graph.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, metric.GraphTable, metric.GraphColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHistograms queries the histograms edge of a Metric.
func (c *MetricClient) QueryHistograms(m *Metric) *HistogramQuery {
	query := &HistogramQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(metric.Table, metric.FieldID, id),
			sqlgraph.To(histogram.Table, histogram.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metric.HistogramsTable, metric.HistogramsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCounters queries the counters edge of a Metric.
func (c *MetricClient) QueryCounters(m *Metric) *CounterQuery {
	query := &CounterQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(metric.Table, metric.FieldID, id),
			sqlgraph.To(counter.Table, counter.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metric.CountersTable, metric.CountersColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGauges queries the gauges edge of a Metric.
func (c *MetricClient) QueryGauges(m *Metric) *GaugeQuery {
	query := &GaugeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(metric.Table, metric.FieldID, id),
			sqlgraph.To(gauge.Table, gauge.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metric.GaugesTable, metric.GaugesColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MetricClient) Hooks() []Hook {
	return c.hooks.Metric
}
