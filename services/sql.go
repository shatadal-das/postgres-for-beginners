package services

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
)

type SQL struct {
	Db *sql.DB
}

func NewSQLDb(db_uri string) *SQL {
	db, err := sql.Open("postgres", db_uri)
	if err != nil {
		fmt.Println("Error opening database")
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging database")
		panic(err)
	}

	return &SQL{Db: db}
}

func (s *SQL) Close() error {
	return s.Db.Close()
}

func (s *SQL) CreateTable(createTableQuery string) {
	_, err := s.Db.Exec(string(createTableQuery))
	if err != nil {
		fmt.Println("Error creating table")
		panic(err)
	}
}

func (s *SQL) CreateTableFromFilename(filename string) {
	createTableQuery, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading createTable.sql file")
		panic(err)
	}

	_, err = s.Db.Exec(string(createTableQuery))
	if err != nil {
		fmt.Println("Error creating table")
		panic(err)
	}
}

func (s *SQL) Insert(query string, args ...any) {
	_, err := s.Db.Exec(query, args...)
	if err != nil {
		fmt.Println("Error inserting data")
		panic(err)
	}
}

func (s *SQL) InsertFromFilename(filename string, args ...any) {
	insertQuery, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading insert.sql file")
		panic(err)
	}

	_, err = s.Db.Exec(string(insertQuery), args...)
	if err != nil {
		fmt.Println("Error inserting data")
		panic(err)
	}
}

func (s *SQL) FindOne(query string, obj interface{}, args ...interface{}) {

	val := reflect.ValueOf(obj)
	// Ensure obj is a pointer to a struct
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		panic("obj must be a pointer to a struct")
	}

	// Prepare a slice to hold the values to be scanned into the struct fields
	var values []interface{}

	// Loop through the struct fields
	for i := 0; i < val.Elem().NumField(); i++ {
		// Create a pointer to each field's value and add it to the values slice
		field := val.Elem().Field(i)
		values = append(values, field.Addr().Interface())
	}

	err := s.Db.QueryRow(query, args...).Scan(values...)
	if err != nil {
		fmt.Println("Error finding one")
		panic(err)
	}
}
