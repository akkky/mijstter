package main

import (
	"fmt"
	"mijstter/db"
	"os"
)

var (
	subCommand   string
	databasePath string
)

func parseCommand() error {
	if len(os.Args) < 3 {
		return fmt.Errorf("パラメーターが不正です。")
	}

	subCommand = os.Args[1]
	databasePath = os.Args[2]

	return nil
}

func create() error {
	db, err := db.NewDatabase(databasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.CreateTables()
	if err != nil {
		return err
	}

	return nil
}

func drop() error {
	db, err := db.NewDatabase(databasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	// User
	err = db.DropTables()
	if err != nil {
		return err
	}

	return nil
}

func clear() error {
	db, err := db.NewDatabase(databasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

func _main() (int, error) {
	err := parseCommand()
	if err != nil {
		return 1, err
	}

	switch subCommand {
	case "create":
		err = create()
		if err != nil {
			return 11, err
		}
	case "drop":
		err = drop()
		if err != nil {
			return 12, err
		}
	case "clear":
		err = clear()
		if err != nil {
			return 13, err
		}
	default:
		return 2, fmt.Errorf("コマンドが存在しません。")
	}

	return 0, nil
}

func main() {
	if status, err := _main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(status)
	}
}
