package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "movie"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func createRecord() {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()
	// insert, err := db.Query("INSERT INTO moviedetail VALUES ( 2, 'TEST','','','','','','','','','','' )")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	data := readCSV()
	multipleInsert(db, data)
}

func multipleInsert(db *sql.DB, movieData [][]string) error {

	for _, v := range movieData {
		query := "INSERT INTO moviedetail(movietype, title,director,cast,country,dateadded,releaseyear,rating,duration,listedin,moviedescription) VALUES "
		var inserts []string
		var params []interface{}
		inserts = append(inserts, "(?, ?,?,?,?,?,?,?,?,?,?)")
		params = append(params, v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11])
		queryVals := strings.Join(inserts, "")
		query = query + queryVals
		ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelfunc()
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()
		res, err := stmt.ExecContext(ctx, params...)
		if err != nil {
			log.Printf("Error %s when inserting row into products table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d products created simulatneously", rows)
	}
	return nil
}

func readCSV() [][]string {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvLines
}
