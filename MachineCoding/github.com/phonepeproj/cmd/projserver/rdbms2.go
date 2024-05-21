package main

import (
	"errors"
	"fmt"
	"sync"
)

// DataType represents the type of the column
type DataType string

const (
	StringType DataType = "string"
	IntType    DataType = "int"
)

// Column represents a column in a table
type Column struct {
	Name     string
	Type     DataType
	Required bool
}

// Record represents a record in a table
type Record map[string]interface{}

// Table represents a table in the database
type Table struct {
	Name    string
	Columns []Column
	Records []Record
}

// Database represents the in-memory database
type Database struct {
	Tables map[string]*Table
}

var dbInstance *Database
var once sync.Once

// GetDatabaseInstance returns the singleton instance of the database
func GetDatabaseInstance() *Database {
	once.Do(func() {
		dbInstance = &Database{
			Tables: make(map[string]*Table),
		}
	})
	return dbInstance
}

// TableFactory is a factory for creating tables
type TableFactory struct{}

func (tf *TableFactory) CreateTable(name string, columns []Column) (*Table, error) {
	return &Table{
		Name:    name,
		Columns: columns,
		Records: []Record{},
	}, nil
}

// RecordBuilder is a builder for creating records
type RecordBuilder struct {
	record Record
}

func NewRecordBuilder() *RecordBuilder {
	return &RecordBuilder{record: make(Record)}
}

func (rb *RecordBuilder) AddField(name string, value interface{}) *RecordBuilder {
	rb.record[name] = value
	return rb
}

func (rb *RecordBuilder) Build() Record {
	return rb.record
}

// FilterStrategy defines the strategy for filtering records
type FilterStrategy interface {
	Filter(table *Table, columnName string, value interface{}) []Record
}

// EqualityFilterStrategy filters records by equality
type EqualityFilterStrategy struct{}

func (efs *EqualityFilterStrategy) Filter(table *Table, columnName string, value interface{}) []Record {
	var filteredRecords []Record
	for _, record := range table.Records {
		if record[columnName] == value {
			filteredRecords = append(filteredRecords, record)
		}
	}
	return filteredRecords
}

// Methods for Database operations
func (db *Database) AddTable(table *Table) error {
	if _, exists := db.Tables[table.Name]; exists {
		return errors.New("table already exists")
	}
	db.Tables[table.Name] = table
	return nil
}

func (db *Database) DeleteTable(name string) error {
	if _, exists := db.Tables[name]; !exists {
		return errors.New("table does not exist")
	}
	delete(db.Tables, name)
	return nil
}

func (db *Database) InsertRecord(tableName string, record Record) error {
	table, exists := db.Tables[tableName]
	if !exists {
		return errors.New("table does not exist")
	}
	for _, column := range table.Columns {
		value, present := record[column.Name]
		if !present && column.Required {
			return fmt.Errorf("missing required field: %s", column.Name)
		}
		if present {
			switch column.Type {
			case StringType:
				strValue, ok := value.(string)
				if !ok || len(strValue) > 20 {
					return fmt.Errorf("invalid value for column %s: must be a string of max length 20", column.Name)
				}
			case IntType:
				intValue, ok := value.(int)
				if !ok || intValue < -1024 || intValue > 1023 {
					return fmt.Errorf("invalid value for column %s: must be an int between -1024 and 1023", column.Name)
				}
			default:
				return fmt.Errorf("unknown type for column %s", column.Name)
			}
		}
	}
	table.Records = append(table.Records, record)
	return nil
}

func (db *Database) PrintAllRecords(tableName string) error {
	table, exists := db.Tables[tableName]
	if !exists {
		return errors.New("table does not exist")
	}
	for _, record := range table.Records {
		fmt.Println(record)
	}
	return nil
}

func (db *Database) FilterRecords(strategy FilterStrategy, tableName, columnName string, value interface{}) error {
	table, exists := db.Tables[tableName]
	if !exists {
		return errors.New("table does not exist")
	}
	filteredRecords := strategy.Filter(table, columnName, value)
	for _, record := range filteredRecords {
		fmt.Println(record)
	}
	return nil
}

func main() {
	db := GetDatabaseInstance()

	tableFactory := &TableFactory{}
	table, err := tableFactory.CreateTable("users", []Column{
		{Name: "id", Type: IntType, Required: true},
		{Name: "name", Type: StringType, Required: true},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.AddTable(table)
	if err != nil {
		fmt.Println(err)
		return
	}

	recordBuilder := NewRecordBuilder()
	record1 := recordBuilder.AddField("id", 1).AddField("name", "Alice").Build()
	record2 := recordBuilder.AddField("id", 2).AddField("name", "Bob").Build()

	err = db.InsertRecord("users", record1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.InsertRecord("users", record2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("All Records:")
	err = db.PrintAllRecords("users")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Filtered Records (name=Alice):")
	equalityFilter := &EqualityFilterStrategy{}
	err = db.FilterRecords(equalityFilter, "users", "name", "Alice")
	if err != nil {
		fmt.Println(err)
		return
	}
}
