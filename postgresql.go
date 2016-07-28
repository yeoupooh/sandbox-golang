package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

type jsonobject struct {
    DriverName string
		Url string
}

func main() {
	// read config file
	file, e := ioutil.ReadFile("./postgresql.config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	// parse config
	var jsontype jsonobject
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
	fmt.Printf("Driver Name: %v\n", jsontype.DriverName)
	fmt.Printf("URL: %v\n", jsontype.Url)

	// open db
	db, err := sql.Open(jsontype.DriverName, jsontype.Url)
	if err != nil {
		log.Fatal(err)
	}

	// execute query
	rows, err := db.Query("SELECT * FROM E2_IMG")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(rows)
	}
}
