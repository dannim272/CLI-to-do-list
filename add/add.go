package add

import (
	"fmt"
	"database/sql"
	_ "github.com/glebarez/go-sqlite" 
	"time"
)

func Add(newTask string) {
	db, err := sql.Open("sqlite","./checklist.db")
	if err != nil {
		fmt.Println("first err")
		fmt.Println(err)
		return
	}

	defer db.Close()
	now := time.Now()
	var date string = now.Format("01-02-2006")
	sql, err := db.Prepare(`INSERT INTO checklist (Task, Created) VALUES (?,?)`)
	sql.Exec(newTask, date)
}
