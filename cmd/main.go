package main

import (
	"fmt"
	"database/sql"
	_ "github.com/glebarez/go-sqlite" 
	"os"
	"bufio"
	"to-do-list/internal/test"
	"to-do-list/internal/list"
	"to-do-list/internal/add"
	"to-do-list/internal/delete"
)

func deleteAll() {
	/*db, err := sql.Open("sqlite","./checklist.db")
	if err != nil {
		fmt.Println("first err")
		fmt.Println(err)
		return
	}*/
}

func main() {
	fmt.Println("    Welcome to your personal to-do list    ")
	fmt.Println("-------------------------------------------")
	db, err := sql.Open("sqlite","./checklist.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS checklist (
		id INTEGER PRIMARY KEY,
		Task VARCHAR(32),
		Created VARCHAR(32))`)
	if err != nil {
		fmt.Println("Error in creating table")
	} 
	statement.Exec()
	db.Close()

	command := os.Args[1]

	switch command {
	case "list":
		list.List()
		break
		
	case "add":
		fmt.Println("Creat a new task")
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		newTask, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("something's wrong")
		}
		add.Add(newTask)
		break

	case "delete":
		fmt.Println("Delete - All or id?")
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		req, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("something's wrong")
		}
		if req == "All" || req == "all" {
			delete.DeleteAll()
		}
		break

	case "test":
		test.Test()
		break
	}
}
