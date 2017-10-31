package db

import (
	
	"fmt"
	"github.com/lucashenriqueteixeira/list_go/lib/util"	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Open() (*sql.DB){
	host := os.Getenv("HOST")
	password := os.Getenv("PASSWORD")
	user := os.Getenv("USER")
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	access := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",user,password,host,port,name)

	db, err := sql.Open("mysql",access)
	util.VerifyErr(err)

	//check if connection is active
	err = db.Ping()
	util.VerifyErr(err)

	return db
}