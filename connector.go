package main

import (
    "fmt"
    //packages we should include
    "time"
    "database/sql"
    //mysql driver
    //The reason for using underscore is that we will include this package but not use it here
    _ "github.com/go-sql-driver/mysql"

)

func main() {
    start := time.Now()
    //my mysql user name is root and my passaword 1234 
    //my database name is deneme2
    //my localhost port number is 3306
   
    
    db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/deneme2")

    if err != nil {
        panic(err.Error())
    }
    //close my database
    defer db.Close()

    //here we are trying to add database file
    db.Query("INSERT INTO otomobil_fiyatlari (id,marka,model,donanim,motor,yakit,vites,fiyat,websitesi) VALUES ('300','ZEO','A3 Sedan','1.0 TFSI 116 hp Sport Line S tronic PI','1600','Benzin','Otomatik','987999','https://www.zoe.com.tr')")
    fmt.Println("Connected to database !")
    duration := time.Since(start)
    fmt.Println(duration.Nanoseconds())
 
}