package culi

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dataSourceName = "root:',.flacless@tcp(127.0.0.1:3306)/callogit?charset=utf8"
var database Database
var err error

type Database struct {
	db *sql.DB
}

func init() {
	database.db, err = sql.Open("mysql", dataSourceName)
	err = database.db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database init()")
	}
}

func (db Database) begin() (tx *sql.Tx) {
	tx, e := db.db.Begin()
	if e != nil {
		log.Println(e)
		return nil
	}
	return tx
}

func (db Database) prepare(q string) (stmt *sql.Stmt) {
	stmt, e := db.db.Prepare(q)
	if e != nil {
		log.Println(e)
		return nil
	}
	return stmt
}

func (db Database) query(q string, args ...interface{}) (rows *sql.Rows) {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}

func taskQuery(sql string, args ...interface{}) error {
	log.Print("inside task query")
	SQL := database.prepare(sql)
	tx := database.begin()
	_, err = tx.Stmt(SQL).Exec(args...)
	if err != nil {
		log.Println("taskQuery: ", err)
		tx.Rollback()
	} else {
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("Commit successful")
	}
	return err
}

func Close() {
	database.db.Close()
}
