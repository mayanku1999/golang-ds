package main

import (
	"errors"
	"fmt"
)

// Column represents a column in a table
type Column struct {
	Name     string
	Type     string // Assume "string" or "int" types
	Required bool   // Indicates if the column is mandatory
}

// Record represents a record in a table
type Record map[string]interface{}

// RecordValidator validates records against column definitions
type RecordValidator interface {
	Validate(record Record, columns []Column) error
}

// DefaultRecordValidator validates records based on column types and constraints
type DefaultRecordValidator struct{}

// Validate validates a record against column definitions
func (rv DefaultRecordValidator) Validate(record Record, columns []Column) error {
	for _, column := range columns {
		value, ok := record[column.Name]
		if !ok && column.Required {
			return errors.New("missing required field: " + column.Name)
		}

		if ok {
			switch column.Type {
			case "string":
				strValue, ok := value.(string)
				if !ok {
					return errors.New("invalid type for column: " + column.Name)
				}
				if len(strValue) > 20 {
					return errors.New("string value exceeds maximum length for column: " + column.Name)
				}
			case "int":
				intValue, ok := value.(int)
				if !ok {
					return errors.New("invalid type for column: " + column.Name)
				}
				if intValue < -1024 || intValue > 1023 {
					return errors.New("int value out of range for column: " + column.Name)
				}
			}
		}
	}
	return nil
}

// Table represents a table in the database
type Table struct {
	Name    string
	Columns []Column
	Records []Record
}

// TableBuilder builds table instances
type TableBuilder struct {
	name    string
	columns []Column
}

// NewTableBuilder creates a new TableBuilder instance
func NewTableBuilder(name string) *TableBuilder {
	return &TableBuilder{name: name}
}

// WithColumn adds a column to the table being built
func (tb *TableBuilder) WithColumn(name, columnType string, required bool) *TableBuilder {
	tb.columns = append(tb.columns, Column{Name: name, Type: columnType, Required: required})
	return tb
}

// Build creates a new table instance
func (tb *TableBuilder) Build() *Table {
	return &Table{
		Name:    tb.name,
		Columns: tb.columns,
		Records: make([]Record, 0),
	}
}

// Database represents an in-memory SQL-like database
type Database struct {
	tables map[string]*Table
}

// Singleton instance of the Database
var instance *Database

// GetInstance returns the singleton instance of the Database
func GetInstance() *Database {
	if instance == nil {
		instance = &Database{
			tables: make(map[string]*Table),
		}
	}
	return instance
}

// CreateTable creates a new table in the database
func (db *Database) CreateTable(table *Table) error {
	if _, exists := db.tables[table.Name]; exists {
		return errors.New("table already exists")
	}

	db.tables[table.Name] = table

	return nil
}

// InsertRecord inserts a new record into the specified table
func (db *Database) InsertRecord(tableName string, record Record, validator RecordValidator) error {
	table, exists := db.tables[tableName]
	if !exists {
		return errors.New("table not found")
	}

	// Validate record against table columns
	if err := validator.Validate(record, table.Columns); err != nil {
		return err
	}

	// Add record to the table
	table.Records = append(table.Records, record)
	return nil
}

// PrintTable prints all records in a table
func (db *Database) PrintTable(tableName string) error {
	table, exists := db.tables[tableName]
	if !exists {
		return errors.New("table not found")
	}

	fmt.Printf("Records in table %s:\n", tableName)
	for _, record := range table.Records {
		fmt.Println(record)
	}
	return nil
}

// FilterAndPrint filters records in a table based on column values and prints them
func (db *Database) FilterAndPrint(tableName, columnName string, value interface{}) error {
	table, exists := db.tables[tableName]
	if !exists {
		return errors.New("table not found")
	}

	fmt.Printf("Filtered records in table %s for column %s with value %v:\n", tableName, columnName, value)
	for _, record := range table.Records {
		if val, ok := record[columnName]; ok && val == value {
			fmt.Println(record)
		}
	}
	return nil
}

func main() {
	// Example usage
	db := GetInstance()

	// Create a table
	tableBuilder := NewTableBuilder("people").
		WithColumn("name", "string", true).
		WithColumn("age", "int", false)
	peopleTable := tableBuilder.Build()

	err := db.CreateTable(peopleTable)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Create a record validator
	validator := DefaultRecordValidator{}

	// Insert records
	err = db.InsertRecord("people", Record{"name": "Alice", "age": 25}, validator)
	if err != nil {
		fmt.Println("Error inserting record:", err)
		return
	}
	err = db.InsertRecord("people", Record{"name": "Bob", "age": 30}, validator)
	if err != nil {
		fmt.Println("Error inserting record:", err)
		return
	}

	// Print all records in the table
	err = db.PrintTable("people")
	if err != nil {
		fmt.Println("Error printing table:", err)
		return
	}

	// Filter and print records
	err = db.FilterAndPrint("people", "name", "Alice")
	if err != nil {
		fmt.Println("Error filtering and printing records:", err)
		return
	}
}
