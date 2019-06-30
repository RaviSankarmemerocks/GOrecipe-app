package main
 import ("errors")

type content struct{
	title string
	steps string
}


func insert(e content) error {
	q:=`INSERRT INTO curd VALUES($1,$2);`

    db := getConnection()
    defer db.close()
    stmt,err := db.Prepare(q)
    if err !=  nill{
    	return err
    }
    defer stmt.Close()
    r,err := stmt.Exec(e.Title,e.Steps)
     if err !=  nill{
    	return err
    }
    i,_ :=r.RowsAffected()
    if i!=1{
    	return error.New("ERROR")
    }
    return nil
}