// Package internal contains internal code for the hello service.
package internal

type DB int

// HelloDBConn returns a database connection for the hello service only.
func HelloDBConn() *DB {
	return new(DB)
}
