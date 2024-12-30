package main

import (
	"database/sql"
	"fmt"
	"os"
)

func getQuery() string {
	content, err := os.ReadFile("query.sql")
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return ""
	}
	return string(content)
}

func getWordsFromDB() (map[string]int, error) {


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := getQuery()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	words := make(map[string]int)

	for rows.Next() {
		var word string
		var frequency int
		if err := rows.Scan(&word, &frequency); err != nil {
			return nil, err
		}
		words[word] = frequency
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
