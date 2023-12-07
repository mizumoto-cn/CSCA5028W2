package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: out <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO data (name, count) VALUES (?, ?)", data.Name, data.Count)
	if err != nil {
		log.Fatal(err)
	}
}
