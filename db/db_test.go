package db

import (
	"fmt"
	"os"
	"testing"
)

var (
	db *Database
)

const dbFile = "./test.sqlite3"

func TestMain(m *testing.M) {
	var exitCode = 0
	defer os.Exit(exitCode)

	err := before()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
		return
	}
	defer after()

	exitCode = m.Run()
}

func before() error {
	os.Remove(dbFile)

	var err error

	db, err = NewDatabase(dbFile)
	if err != nil {
		return err
	}

	err = db.CreateTables()
	if err != nil {
		return err
	}

	return nil
}

func after() {
	db.Close()
}
