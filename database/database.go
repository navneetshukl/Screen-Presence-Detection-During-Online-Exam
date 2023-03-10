package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

//var conn *sql.DB
//var err error
func DB_Connect() (*sql.DB,error){

	conn, err:= sql.Open("pgx", "host=localhost port=5432 dbname=UserImageDB user=postgres password=password")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to connect: %v\n", err))
		return nil,err
	}

	//defer conn.Close()

	log.Println("Connected to database")

	err=conn.Ping()
	if err!=nil{
		log.Fatal("Cannot Ping the database")
		return nil,err
	}
	log.Println("pinged database")

	return conn,nil

}
func InsertRow(name,email,invitation_code string){
	conn,err:=DB_Connect()
	defer conn.Close()
	if err!=nil{
		panic(err)
	}
	query:=`insert into usertable(name,email,invitation_code) values($1,$2,$3)`

	_,err =conn.Exec(query,name,email,invitation_code)

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Inserted a row")
}

func InsertCameraData(timestamp,image,email string){

	conn,err:=DB_Connect()
	defer conn.Close()
	if err!=nil{
		panic(err)
	}
	query:=`insert into imagetable(timestamp,image,email) values($1,$2,$3)`

	_,err =conn.Exec(query,timestamp,image,email)

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Inserted a row")
}
func Getdata(){
	conn,err:=DB_Connect()
	defer conn.Close()
	if err!=nil{
		panic(err)
	}
	query := `
        SELECT *
        FROM usertable
        JOIN imagetable ON usertable.email = imagetable.email;
    `

    // Execute the query and retrieve the result set
    rows, err := conn.Query(query)
    if err != nil {
        panic(err)
    }
    defer rows.Close()
for rows.Next() {
	var col1 int
	var col2 string
	var col3 string
	var col4 string
	var col5 int
	var col6 string
	var col7 string
	var col8 string
	// add more columns as needed
	err = rows.Scan(&col1, &col2,&col3,&col4,&col5,&col6,&col7,&col8)
	if err != nil {
		panic(err)
	}
	fmt.Println(col1, col2, col3,col4,col5,col6,col7,col8)
}
if err = rows.Err(); err != nil {
	panic(err)
}
}