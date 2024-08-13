// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/file"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
	// FileExtension holds the value of the "file_extension" field.
	FileExtension string `json:"file_extension,omitempty"`
	// FileSize holds the value of the "file_size" field.
	FileSize int `json:"file_size,omitempty"`
	// ContentType holds the value of the "content_type" field.
	ContentType string `json:"content_type,omitempty"`
	// StoreKey holds the value of the "store_key" field.
	StoreKey string `json:"store_key,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Annotation holds the value of the "annotation" field.
	Annotation string `json:"annotation,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	user_files   *string
	selectValues sql.SelectValues
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Organization holds the value of the organization edge.
	Organization []*Organization `json:"organization,omitempty"`
	// Entity holds the value of the entity edge.
	Entity []*Entity `json:"entity,omitempty"`
	// Group holds the value of the group edge.
	Group []*Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedOrganization map[string][]*Organization
	namedEntity       map[string][]*Entity
	namedGroup        map[string][]*Group
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) OrganizationOrErr() ([]*Organization, error) {
	if e.loadedTypes[1] {
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// EntityOrErr returns the Entity value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) EntityOrErr() ([]*Entity, error) {
	if e.loadedTypes[2] {
		return e.Entity, nil
	}
	return nil, &NotLoadedError{edge: "entity"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) GroupOrErr() ([]*Group, error) {
	if e.loadedTypes[3] {
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldTags:
			values[i] = new([]byte)
		case file.FieldFileSize:
			values[i] = new(sql.NullInt64)
		case file.FieldID, file.FieldCreatedBy, file.FieldUpdatedBy, file.FieldDeletedBy, file.FieldMappingID, file.FieldFileName, file.FieldFileExtension, file.FieldContentType, file.FieldStoreKey, file.FieldCategory, file.FieldAnnotation:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt, file.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case file.ForeignKeys[0]: // user_files
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case file.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				f.CreatedBy = value.String
			}
		case file.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				f.UpdatedBy = value.String
			}
		case file.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				f.DeletedAt = value.Time
			}
		case file.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				f.DeletedBy = value.String
			}
		case file.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				f.MappingID = value.String
			}
		case file.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &f.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case file.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				f.FileName = value.String
			}
		case file.FieldFileExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_extension", values[i])
			} else if value.Valid {
				f.FileExtension = value.String
			}
		case file.FieldFileSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_size", values[i])
			} else if value.Valid {
				f.FileSize = int(value.Int64)
			}
		case file.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_type", values[i])
			} else if value.Valid {
				f.ContentType = value.String
			}
		case file.FieldStoreKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_key", values[i])
			} else if value.Valid {
				f.StoreKey = value.String
			}
		case file.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				f.Category = value.String
			}
		case file.FieldAnnotation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field annotation", values[i])
			} else if value.Valid {
				f.Annotation = value.String
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_files", values[i])
			} else if value.Valid {
				f.user_files = new(string)
				*f.user_files = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the File entity.
func (f *File) QueryUser() *UserQuery {
	return NewFileClient(f.config).QueryUser(f)
}

// QueryOrganization queries the "organization" edge of the File entity.
func (f *File) QueryOrganization() *OrganizationQuery {
	return NewFileClient(f.config).QueryOrganization(f)
}

// QueryEntity queries the "entity" edge of the File entity.
func (f *File) QueryEntity() *EntityQuery {
	return NewFileClient(f.config).QueryEntity(f)
}

// QueryGroup queries the "group" edge of the File entity.
func (f *File) QueryGroup() *GroupQuery {
	return NewFileClient(f.config).QueryGroup(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("generated: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(f.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(f.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(f.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(f.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(f.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", f.Tags))
	builder.WriteString(", ")
	builder.WriteString("file_name=")
	builder.WriteString(f.FileName)
	builder.WriteString(", ")
	builder.WriteString("file_extension=")
	builder.WriteString(f.FileExtension)
	builder.WriteString(", ")
	builder.WriteString("file_size=")
	builder.WriteString(fmt.Sprintf("%v", f.FileSize))
	builder.WriteString(", ")
	builder.WriteString("content_type=")
	builder.WriteString(f.ContentType)
	builder.WriteString(", ")
	builder.WriteString("store_key=")
	builder.WriteString(f.StoreKey)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(f.Category)
	builder.WriteString(", ")
	builder.WriteString("annotation=")
	builder.WriteString(f.Annotation)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOrganization returns the Organization named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedOrganization(name string) ([]*Organization, error) {
	if f.Edges.namedOrganization == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedOrganization[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedOrganization(name string, edges ...*Organization) {
	if f.Edges.namedOrganization == nil {
		f.Edges.namedOrganization = make(map[string][]*Organization)
	}
	if len(edges) == 0 {
		f.Edges.namedOrganization[name] = []*Organization{}
	} else {
		f.Edges.namedOrganization[name] = append(f.Edges.namedOrganization[name], edges...)
	}
}

// NamedEntity returns the Entity named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedEntity(name string) ([]*Entity, error) {
	if f.Edges.namedEntity == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedEntity[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedEntity(name string, edges ...*Entity) {
	if f.Edges.namedEntity == nil {
		f.Edges.namedEntity = make(map[string][]*Entity)
	}
	if len(edges) == 0 {
		f.Edges.namedEntity[name] = []*Entity{}
	} else {
		f.Edges.namedEntity[name] = append(f.Edges.namedEntity[name], edges...)
	}
}

// NamedGroup returns the Group named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedGroup(name string) ([]*Group, error) {
	if f.Edges.namedGroup == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedGroup[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedGroup(name string, edges ...*Group) {
	if f.Edges.namedGroup == nil {
		f.Edges.namedGroup = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		f.Edges.namedGroup[name] = []*Group{}
	} else {
		f.Edges.namedGroup[name] = append(f.Edges.namedGroup[name], edges...)
	}
}

// Files is a parsable slice of File.
type Files []*File
