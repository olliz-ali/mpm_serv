package databaseinter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBinter struct {
	con *sql.DB
}

func NewDBinter(user string, password string, database string) *DBinter {
	con, err := sql.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/"+database)

	if err != nil {
		panic(err)
	}
	con.SetConnMaxLifetime(time.Minute * 4)
	con.SetMaxOpenConns(10)
	con.SetMaxIdleConns(10)
	dbcon := DBinter{con}
	return &dbcon

}

func (db *DBinter) InsertJSON(json_data string, tabel_name string) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(json_data), &data)
	if err != nil {
		log.Fatal(err)
	}
	insert_string, leftovers := db.insertJSONtoExestring(data, tabel_name)
	tx, err := db.con.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insert_string)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	var lastInsertID int64
	err = db.con.QueryRow("SELECT LAST_INSERT_ID();").Scan(&lastInsertID)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range leftovers {
		if value, ok := v.(map[string]interface{}); ok {
			db.insertFOREIGEN(int(lastInsertID), value, k)
		}
	}

}

func (db *DBinter) insertJSONtoExestring(jsonData map[string]interface{}, tabel_name string) (string, map[string]interface{}) {
	insert_str := "INSERT INTO " + tabel_name + " ("
	value_str := "VALUES ("
	leftovers := make(map[string]interface{})
	for k, v := range jsonData {
		//check if value is not a string
		if nestedMap, ok := v.(map[string]interface{}); ok {
			leftovers[k] = nestedMap
		} else {
			insert_str += k + ","
			switch value := v.(type) {
			case int:
				value_str += "'" + fmt.Sprintf("%d", value) + "',"
			case float32, float64:
				value_str += "'" + fmt.Sprintf("%.2f", value) + "',"
			case string:
				value_str += "'" + value + "',"
			}
		}

	}
	insert_str = insert_str[:len(insert_str)-1] + ") "
	value_str = value_str[:len(value_str)-1] + "); "
	return insert_str + value_str, leftovers
}

func (db *DBinter) insertFOREIGEN(foreigne_id int, jsonvalue map[string]interface{}, foreigen_tabel string) {
	jsonvalue["owner_id"] = foreigne_id
	exe_string, leftovers := db.insertJSONtoExestring(jsonvalue, foreigen_tabel)
	tx, err := db.con.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(exe_string)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	var lastInsertID int64
	err = db.con.QueryRow("SELECT LAST_INSERT_ID();").Scan(&lastInsertID)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range leftovers {
		if value, ok := v.(map[string]interface{}); ok {
			db.insertFOREIGEN(int(lastInsertID), value, k)
		}

	}

}
