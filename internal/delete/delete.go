package delete

import (
	"fmt"
	"database/sql"
	_ "github.com/glebarez/go-sqlite" 
)

func DeleteAll() {
	db, err := sql.Open("sqlite","./cmd/checklist.db")
	if err != nil {
		fmt.Println("first err")
		fmt.Println(err)
		return
	}

	defer db.Close()
	statement, err := db.Prepare(`DELETE * FROM checklist`)
	if err != nil {
		fmt.Println("something went wrong")
	}
	statement.Exec()
}
