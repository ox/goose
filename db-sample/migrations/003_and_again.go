package main

import (
	"database/sql"
	"fmt"
)

func Up_3(txn *sql.Tx) {
	fmt.Println("Hello from migration 3 Up!")
}

func Down_3(txn *sql.Tx) {
	fmt.Println("Hello from migration 3 Down!")
}
