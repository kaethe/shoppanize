// database.go
package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func handleUserDB(params []string, mode int) error {
	db, err := sql.Open("sqlite3", "./shop.db")
	if err != nil {
		fmt.Errorf("Error opening Database")
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Errorf("Error starting db transaction")
	}

	switch mode {
	case 0:
		{
			//add a new user

			//open the database

			stmt, err := tx.Prepare("INSERT INTO users(username, password) VALUES (?, ?)")
			defer stmt.Close()
			if err != nil {
				fmt.Errorf("Error preparing statement")
			}
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(params[2]), bcrypt.DefaultCost)
			if err != nil {
				fmt.Errorf("Error encoding password")
			}
			_, err = stmt.Exec(params[1], passwordHash)
			if err != nil {
				fmt.Println("Error executing statement", err.Error())
			}
			tx.Commit()
			return nil
		}
	case 1:
		{
			fmt.Println("Removing", params)
			stmt, err := tx.Prepare("DELETE FROM users WHERE uuid = ?;")
			defer stmt.Close()
			if err != nil {
				fmt.Errorf("Error preparing statement")
			}
			uuid, err := strconv.ParseInt(params[1], 10, 64)
			if err != nil {
				fmt.Println("Error parsing String", err.Error())
			}
			fmt.Println("uuid", uuid, params[1])
			_, err = stmt.Exec(uuid)
			if err != nil {
				fmt.Println("Error executing statement", err.Error())
			}
			tx.Commit()
		}
	}
	return nil
}

func handleDatabaseRequest(s string) error {
	params := strings.Split(s, " ")
	for i := 0; i < len(params); i++ {
		params[i] = strings.TrimSuffix(params[i], "\n")
	}
	fmt.Println(params)

	switch params[0] {
	case "uadd":
		{
			handleUserDB(params, 0)
		}
	case "ladd":
		{

		}
	case "iadd":
		{

		}
	case "lremove":
		{

		}
	case "uremove":
		{
			handleUserDB(params, 1)
		}
	case "iremove":
		{

		}

	}
	return nil
}
