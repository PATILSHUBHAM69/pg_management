package main

import (
	"database/sql"
	"fmt"
	"log"

	controller "github.com/shahrammohd3010/Property_Management/controller"

	"github.com/gin-gonic/gin"
)

func db_creation() {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS pms")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Successfully Created")
}

func main() {
	db_creation()
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pms")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	Create, err := db.Query("CREATE TABLE IF NOT EXISTS Owner_Details (ID INT not null AUTO_INCREMENT primary key,First_Name varchar(30),Last_Name VARCHAR(30),Email VARCHAR(30),PhoneNumber VARCHAR(10), PG_Type VARCHAR(30), PropertyID INT not null, City VARCHAR(50), Pin_Code INT, State VARCHAR(50), Location VARCHAR(100));")
	if err != nil {
		panic(err.Error())
	}
	defer Create.Close()
	fmt.Println("Successfully Created")
	r := gin.Default()
	r.POST("/post", controller.Add_owner())
	r.GET("/get", controller.get_owner_byID())
	r.GET("/get2", controller.get_Allowner())
	r.PUT("/put", controller.update_owner())
	r.DELETE("/delete", controller.delete_owner())
	r.Run("localhost:8080")
}
