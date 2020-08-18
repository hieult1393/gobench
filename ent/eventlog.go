// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/ent/eventlog"
)

// EventLog is the model entity for the EventLog schema.
type EventLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Message holds the value of the "message" field.
	Message string `json:"message"`
	// Level holds the value of the "level" field.
	Level string `json:"level"`
	// Source holds the value of the "source" field.
	Source string `json:"source"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventLogQuery when eager-loading is set.
	Edges                  EventLogEdges `json:"edges"`
	application_event_logs *int
}

// EventLogEdges holds the relations/edges for other nodes in the graph.
type EventLogEdges struct {
	// Applications holds the value of the applications edge.
	Applications *Application
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplicationsOrErr returns the Applications value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventLogEdges) ApplicationsOrErr() (*Application, error) {
	if e.loadedTypes[0] {
		if e.Applications == nil {
			// The edge applications was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: application.Label}
		}
		return e.Applications, nil
	}
	return nil, &NotLoadedError{edge: "applications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EventLog) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // message
		&sql.NullString{}, // level
		&sql.NullString{}, // source
		&sql.NullTime{},   // created_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*EventLog) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // application_event_logs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventLog fields.
func (el *EventLog) assignValues(values ...interface{}) error {
	if m, n := len(values), len(eventlog.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	el.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		el.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field message", values[1])
	} else if value.Valid {
		el.Message = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field level", values[2])
	} else if value.Valid {
		el.Level = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field source", values[3])
	} else if value.Valid {
		el.Source = value.String
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[4])
	} else if value.Valid {
		el.CreatedAt = value.Time
	}
	values = values[5:]
	if len(values) == len(eventlog.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field application_event_logs", value)
		} else if value.Valid {
			el.application_event_logs = new(int)
			*el.application_event_logs = int(value.Int64)
		}
	}
	return nil
}

// QueryApplications queries the applications edge of the EventLog.
func (el *EventLog) QueryApplications() *ApplicationQuery {
	return (&EventLogClient{config: el.config}).QueryApplications(el)
}

// Update returns a builder for updating this EventLog.
// Note that, you need to call EventLog.Unwrap() before calling this method, if this EventLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (el *EventLog) Update() *EventLogUpdateOne {
	return (&EventLogClient{config: el.config}).UpdateOne(el)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (el *EventLog) Unwrap() *EventLog {
	tx, ok := el.config.driver.(*txDriver)
	if !ok {
		panic("ent: EventLog is not a transactional entity")
	}
	el.config.driver = tx.drv
	return el
}

// String implements the fmt.Stringer.
func (el *EventLog) String() string {
	var builder strings.Builder
	builder.WriteString("EventLog(")
	builder.WriteString(fmt.Sprintf("id=%v", el.ID))
	builder.WriteString(", name=")
	builder.WriteString(el.Name)
	builder.WriteString(", message=")
	builder.WriteString(el.Message)
	builder.WriteString(", level=")
	builder.WriteString(el.Level)
	builder.WriteString(", source=")
	builder.WriteString(el.Source)
	builder.WriteString(", created_at=")
	builder.WriteString(el.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EventLogs is a parsable slice of EventLog.
type EventLogs []*EventLog

func (el EventLogs) config(cfg config) {
	for _i := range el {
		el[_i].config = cfg
	}
}
