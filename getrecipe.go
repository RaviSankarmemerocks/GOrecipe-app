package main
import 


("fmt"
"net/http"
//"io"
"html/template"
//

"database/sql"
_"github.com/lib/pq"
"log"
"errors"

)

    type content struct{
  title string
  steps string
}

  type dishes struct{
  Dishname string
      Steps string
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "root"
  dbname   = "mydbetl"
)

//default to load template in GO
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseGlob("templates/*.gohtml"))
}



 func main() {
       http.HandleFunc("/",index)
       http.HandleFunc("/sendit",display)
       http.ListenAndServe(":8080",nil)


 }


func getConnection() *sql.DB {
   // dsn:="postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"
    dsn := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db,err:=sql.Open("postgres",dsn)
    if err !=nil{
      log.Fatal(err)
    } 
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

 func index(w http.ResponseWriter, r *http.Request) {
 	tpl.ExecuteTemplate(w,"index.gohtml",nil)
 }

 func display(w http.ResponseWriter,r *http.Request) {
 	if r.Method!="POST"{
 		http.Redirect(w,r,"/",http.StatusSeeOther)
 		return
 	}
 	title1:=r.FormValue("dishname")
    recipe2:=r.FormValue("instruction")






e := content{
    title: title1,
    steps: recipe2,

  }

    
    d := dishes{
    	Dishname : title1,
    	Steps : recipe2,
    }

    err := insert(e)

  if err != nil{
    log.Fatal(err)
  }
  fmt.Println("success")
  tpl.ExecuteTemplate(w,"display.gohtml",d)

 }