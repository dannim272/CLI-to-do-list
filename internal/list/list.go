package list

import (
	"fmt"
	"database/sql"
	_ "github.com/glebarez/go-sqlite" 
	"strings"
)

type Task struct {
	Id int
	Name string
	Date string
}

func List() {
	db, err := sql.Open("sqlite","./cmd/checklist.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	rows, err := db.Query(`SELECT id, Task, Created FROM checklist`)
	if err != nil {
		fmt.Println("error fetching rows")
	}
	defer rows.Close()
	var task Task 
	for rows.Next() {
		rows.Scan(&task.Id, &task.Name, &task.Date)
		task.Date = strings.TrimSpace(task.Date)
		task.Name = strings.TrimSpace(task.Name)
		fmt.Println(task.Id,":   ",task.Name,"        (",task.Date,")")
	}
}
