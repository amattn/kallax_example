package main

import (
	"log"
	"time"
)

// basic usage:
// defer trace("name_of_function", time.Now())
func trace(function_name string, start_time time.Time) {
	elapsed := time.Since(start_time)
	log.Printf("%s exiting ~ %s elapsed", function_name, elapsed)
}
