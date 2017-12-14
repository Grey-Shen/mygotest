package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-oci8"
)

// system/oracle@localhost:1521&as=sysdba
// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Printf("ERROR: Please provide a DSN string in ONE argument:\n\n")
// 		fmt.Println("Shell-Conversion into DSN string:")
// 		fmt.Println("  sqlplus sys/password@tnsentry as sysdba   =>   sys/password@tnsentry?as=sysdba")
// 		fmt.Println("  sqlplus / as sysdba                       =>   sys/.@?as=sysdba")
// 		fmt.Println("instead of the tnsentry, you can also use the hostname of the IP.")
// 		os.Exit(1)
// 	}
// 	os.Setenv("NLS_LANG", "")

// 	db, err := sql.Open("oci8", os.Args[1])
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer db.Close()
// 	fmt.Println()
// 	var user string
// 	err = db.QueryRow("select user from dual").Scan(&user)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Successful 'as sysdba' connection. Current user is: %v\n", user)
// }
// select E_INSERT_mytest from user_tables;
func main() {
	const openString = "system/oracle@localhost:1521"

	// A normal simple Open to localhost would look like:
	// db, err := sql.Open("oci8", "127.0.0.1")
	// For testing, need to use additional variables
	db, err := sql.Open("oci8", openString)
	if err != nil {
		fmt.Printf("Open error is not nil: %v", err)
		return
	}
	if db == nil {
		fmt.Println("db is nil")
		return
	}

	// defer close database
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Close error is not nil:", err)
		}
	}()

	// create table
	tableName := "E_MANY_INSERT_TEST"
	// query := "create table " + tableName + " ( A INTEGER, B varchar(10) )"
	// ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	// _, err = db.ExecContext(ctx, query)
	// cancel()
	// if err != nil {
	// 	fmt.Println("ExecContext error is not nil:", err)
	// 	return
	// }

	// // prepare insert query statement
	// var stmt *sql.Stmt
	// query = "insert into " + tableName + " ( A,B ) values (:1, :2)"
	// ctx, cancel = context.WithTimeout(context.Background(), 55*time.Second)
	// stmt, err = db.PrepareContext(ctx, query)
	// cancel()
	// if err != nil {
	// 	fmt.Println("PrepareContext error is not nil:", err)
	// 	return
	// }

	// // insert 3 rows
	// for i := 0; i < 3; i++ {
	// 	ctx, cancel = context.WithTimeout(context.Background(), 55*time.Second)
	// 	_, err = stmt.ExecContext(ctx, i, fmt.Sprintf("aa%d", i))
	// 	cancel()
	// 	if err != nil {
	// 		stmt.Close()
	// 		fmt.Println("ExecContext error is not nil:", err)
	// 		return
	// 	}
	// }

	// // close insert query statement
	// err = stmt.Close()
	// if err != nil {
	// 	fmt.Println("Close error is not nil:", err)
	// 	return
	// }

	// select count/number of rows
	var rows *sql.Rows
	query := "select * from " + tableName + " where A=:1 and B=:2"
	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()
	rows, err = db.QueryContext(ctx, query, 2, "aa1")
	if err != nil {
		fmt.Println("QueryContext error is not nil:", err)
		return
	}
	if !rows.Next() {
		fmt.Println("no Next rows")
		return
	}

	var intA int64
	var strB string
	err = rows.Scan(&intA, &strB)
	if err != nil {
		fmt.Println("Scan error is not nil:", err)
		return
	}

	fmt.Println("A:", intA, "B: ", strB)

	if rows.Next() {
		fmt.Println("has Next rows")
		return
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Err error is not nil:", err)
		return
	}
	err = rows.Close()
	if err != nil {
		fmt.Println("Close error is not nil:", err)
		return
	}

	// drop table
	// query = "drop table " + tableName
	// ctx, cancel = context.WithTimeout(context.Background(), 55*time.Second)
	// _, err = db.ExecContext(ctx, query)
	// cancel()
	// if err != nil {
	// 	fmt.Println("ExecContext error is not nil:", err)
	// 	return
	// }

}

// CREATE TABLE SQUARENUM
// (
//   number int(11),
//   squareNumber int(11),
//   primary key(number)
// );

// create table classinfo(
// 	  classid number(2) primary key,
// 	 classname varchar(10)
// );
