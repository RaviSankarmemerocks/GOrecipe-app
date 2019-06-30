package main
import("database/sql"
"github.com/lib/pq"
"log")


func getConnection() *sql.DB {
    dsn:="postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"
    db,err:=sql.Open("postgres",dsn)
    if err !=nil{
    	log.Fatal(err)
    }	
    return db
}