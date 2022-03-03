// Code generated by entc, DO NOT EDIT.

package ent

import (
	"digidrop/ent/filemiddleware"
	"digidrop/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// FileMiddleware is the model entity for the FileMiddleware schema.
type FileMiddleware struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// URLID holds the value of the "url_id" field.
	URLID string `json:"url_id,omitempty"`
	// FilePath holds the value of the "file_path" field.
	FilePath string `json:"file_path,omitempty"`
	// Accessed holds the value of the "accessed" field.
	Accessed bool `json:"accessed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileMiddlewareQuery when eager-loading is set.
	Edges                                   FileMiddlewareEdges `json:"edges"`
	file_middleware_file_middleware_to_user *uuid.UUID
}

// FileMiddlewareEdges holds the relations/edges for other nodes in the graph.
type FileMiddlewareEdges struct {
	// FileMiddlewareToUser holds the value of the FileMiddlewareToUser edge.
	FileMiddlewareToUser *User `json:"FileMiddlewareToUser,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FileMiddlewareToUserOrErr returns the FileMiddlewareToUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileMiddlewareEdges) FileMiddlewareToUserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.FileMiddlewareToUser == nil {
			// The edge FileMiddlewareToUser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.FileMiddlewareToUser, nil
	}
	return nil, &NotLoadedError{edge: "FileMiddlewareToUser"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileMiddleware) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case filemiddleware.FieldAccessed:
			values[i] = new(sql.NullBool)
		case filemiddleware.FieldURLID, filemiddleware.FieldFilePath:
			values[i] = new(sql.NullString)
		case filemiddleware.FieldID:
			values[i] = new(uuid.UUID)
		case filemiddleware.ForeignKeys[0]: // file_middleware_file_middleware_to_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type FileMiddleware", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileMiddleware fields.
func (fm *FileMiddleware) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case filemiddleware.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fm.ID = *value
			}
		case filemiddleware.FieldURLID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_id", values[i])
			} else if value.Valid {
				fm.URLID = value.String
			}
		case filemiddleware.FieldFilePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_path", values[i])
			} else if value.Valid {
				fm.FilePath = value.String
			}
		case filemiddleware.FieldAccessed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field accessed", values[i])
			} else if value.Valid {
				fm.Accessed = value.Bool
			}
		case filemiddleware.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field file_middleware_file_middleware_to_user", values[i])
			} else if value.Valid {
				fm.file_middleware_file_middleware_to_user = new(uuid.UUID)
				*fm.file_middleware_file_middleware_to_user = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFileMiddlewareToUser queries the "FileMiddlewareToUser" edge of the FileMiddleware entity.
func (fm *FileMiddleware) QueryFileMiddlewareToUser() *UserQuery {
	return (&FileMiddlewareClient{config: fm.config}).QueryFileMiddlewareToUser(fm)
}

// Update returns a builder for updating this FileMiddleware.
// Note that you need to call FileMiddleware.Unwrap() before calling this method if this FileMiddleware
// was returned from a transaction, and the transaction was committed or rolled back.
func (fm *FileMiddleware) Update() *FileMiddlewareUpdateOne {
	return (&FileMiddlewareClient{config: fm.config}).UpdateOne(fm)
}

// Unwrap unwraps the FileMiddleware entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fm *FileMiddleware) Unwrap() *FileMiddleware {
	tx, ok := fm.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileMiddleware is not a transactional entity")
	}
	fm.config.driver = tx.drv
	return fm
}

// String implements the fmt.Stringer.
func (fm *FileMiddleware) String() string {
	var builder strings.Builder
	builder.WriteString("FileMiddleware(")
	builder.WriteString(fmt.Sprintf("id=%v", fm.ID))
	builder.WriteString(", url_id=")
	builder.WriteString(fm.URLID)
	builder.WriteString(", file_path=")
	builder.WriteString(fm.FilePath)
	builder.WriteString(", accessed=")
	builder.WriteString(fmt.Sprintf("%v", fm.Accessed))
	builder.WriteByte(')')
	return builder.String()
}

// FileMiddlewares is a parsable slice of FileMiddleware.
type FileMiddlewares []*FileMiddleware

func (fm FileMiddlewares) config(cfg config) {
	for _i := range fm {
		fm[_i].config = cfg
	}
}
