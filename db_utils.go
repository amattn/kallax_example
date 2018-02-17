package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/amattn/deeperror"
)

// from https://github.com/src-d/go-kallax/blob/master/benchmarks/bench_test.go
func envOrDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		v = def
	}
	return v
}

// from https://github.com/src-d/go-kallax/blob/master/benchmarks/bench_test.go
func dbURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		envOrDefault("DBUSER", "default"),
		envOrDefault("DBPASS", "default"),
		envOrDefault("DBHOST", "0.0.0.0"),
		envOrDefault("DBPORT", "5432"),
		envOrDefault("DBNAME", "default"),
	)
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL())
	if err != nil {
		derr := deeperror.New(570261409, "sql.Open failure", err)
		derr.AddDebugField("dbURL", dbURL())
		return nil, derr
	}
	return db, nil
}
