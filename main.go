package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
)

var recordCount int = 1000000

func main() {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", "myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase")
	if err != nil {
		panic(fmt.Errorf("failed to open database connection: %w", err))
	}
	defer db.Close()

	// err = insertUuidsChar(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = insertUuidsBinary(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = insertUuidsBinarySwapped(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = insertUlids(db)
	// if err != nil {
	// 	panic(err)
	// }

	// DO ONLY THIS ONE ON NEXT RUN
	err = insertUlidsBinary(db)
	if err != nil {
		panic(err)
	}

	// err = insertKsuids(db)
	// if err != nil {
	// 	panic(err)
	// }

	// err = insertKsuidsBinary(db)
	// if err != nil {
	// 	panic(err)
	// }

}

func insertUuidsChar(db *sql.DB) error {
	_, err := db.Exec("delete from uuid_v1_char")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO uuid_v1_char (id, name) VALUES (UUID(), ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}

func insertUuidsBinary(db *sql.DB) error {
	_, err := db.Exec("delete from uuid_v1_binary")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO uuid_v1_binary (id, name) VALUES (UUID_TO_BIN(UUID()), ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}

func insertUuidsBinarySwapped(db *sql.DB) error {
	_, err := db.Exec("delete from uuid_v1_binary_swapped")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO uuid_v1_binary_swapped (id, name) VALUES (UUID_TO_BIN(UUID(), 1), ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}

func insertUlids(db *sql.DB) error {
	_, err := db.Exec("delete from ulids_char")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO ulids_char (id, name) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(ulid.Make().String(), fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}

func insertUlidsBinary(db *sql.DB) error {
	_, err := db.Exec("delete from ulids_binary")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO ulids_binary (id, name) VALUES (BINARY(?), ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(ulid.Make().String(), fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}

func insertKsuids(db *sql.DB) error {
	_, err := db.Exec("delete from ksuids_char")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO ksuids_char (id, name) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(ksuid.New().String(), fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}

func insertKsuidsBinary(db *sql.DB) error {
	_, err := db.Exec("delete from ksuids_binary")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO ksuids_binary (id, name) VALUES (BINARY(?), ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < recordCount; i++ {
		_, err := stmt.Exec(ksuid.New().String(), fmt.Sprintf("Record %d", i))
		if err != nil {
			return fmt.Errorf("failed to insert record: %w", err)
		}
	}

	return nil
}
