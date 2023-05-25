// Routines for dealing with any postgres dbms bureaucracy

package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Server struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"databaseName"`
}

type Config struct {
	Credentials Credentials `json:"credentials"`
	Server      Server      `json:"server"`
}

func OpenConnection(config Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Server.Host, config.Server.Port, config.Credentials.Username, config.Credentials.Password, config.Server.DatabaseName)
	return sql.Open("postgres", psqlInfo)
}

// Load the config for open a connection to a given db
func loadConfig() Config {
	var config Config = Config{
		Credentials: Credentials{
			Username: "animal",
			Password: "4n1m4l",
		},
		Server: Server{
			Host:         "localhost",
			Port:         5432,
			DatabaseName: "animals_db",
		},
	}
	return config
}

var Conn *sql.DB = nil

// el init se invoca antes del main asi que queda establecida la conexion... y si hay tests que usan este paquete tmb se correrá el init ANTES de los tests
func init() {
	config := loadConfig()
	var err error
	Conn, err = OpenConnection(config)
	if err != nil {
		panic(err)
	}
	// verificamos que hay conexión...
	err = Conn.Ping()
	if err != nil {
		panic(err)
	}
	_, err = Conn.Exec("SELECT 1")
	if err != nil {
		panic(err)
	}
}

// Helper rutines

// Interface for things that performs an scan over a given row.
// Actually it is a common interface for https://pkg.go.dev/database/sql#Rows.Scan and https://pkg.go.dev/database/sql#Row.Scan
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// Scans a single row from a given query
type RowScanFunc func(rows RowScanner) (interface{}, error)

// Scans multiples rows using a scanner function in order to build a new "scanable" struct
func ScanMultiples(rows *sql.Rows, rowScanFunc RowScanFunc) ([]interface{}, error) {
	scaneables := []interface{}{}
	for rows.Next() {
		scanable, err := rowScanFunc(rows)
		if scanable == nil {
			return nil, err
		}
		scaneables = append(scaneables, scanable)
	}
	err := rows.Err()
	if err != nil {
		return nil, err
	}
	return scaneables, nil
}
