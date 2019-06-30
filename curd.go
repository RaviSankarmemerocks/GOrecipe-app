package main

import ("fmt"
"database/sql"
_"github.com/lib/pq"
"log"
"errors")

type content struct{
	title string
	steps string
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "root"
  dbname   = "mydbetl"
)



func getConnection() *sql.DB {
    //dsn:="postgres://postgres:postgres@127.0.0.1:5432/password?sslmode=disable"
    dsn := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db,err:=sql.Open("postgres",dsn)
    if err !=nil{
    	log.Fatal(err)
    }	
    fmt.Println("Successfully connected!")
    return db
}

func insert(e content) error {
	q:=`INSERT INTO mygo VALUES($1,$2);`

    db := getConnection()
    defer db.Close()
    stmt,err := db.Prepare(q)
    if err !=  nil{
    	return err
    }
    defer stmt.Close()
    r,err := stmt.Exec(e.title,e.steps)
     if err !=  nil{
    	return err
    }
    i,_ :=r.RowsAffected()
    if i!=1{
    	return errors.New("ERROR")
    }
    return nil
}


func main() {
	e := content{
		title: "fried rice",
		steps: "fry the rice",

	}
	err := insert(e)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("success")

}