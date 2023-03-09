package controller

import (
	"github.com/shahrammohd3010/Property_Management/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func db_creation() {
// 	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS pms")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Database Successfully Created")
// }

func get_owner_byID() gin.HandlerFunc{
	retrun func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pms")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var get_pro models.owner
		err = c.BindJSON(&get_pro)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM Owner_Details WHERE ID=%d", get_pro.ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var Id_ uint
			var F_Name string
			var L_Name string
			var Email string
			var Phonenumber string
			var PG_Type string
			var PropertyID uint
			var City string
			var PinCode int
			var State string
			var Location string
			err = results.Scan(&Id_, &F_Name, &L_Name, &Email, &Phonenumber, &PG_Type, &PropertyID, &City, &PinCode, &State, &Location)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d %s %s %s %s %s %d %s %d %s %s \n", Id_, F_Name, L_Name, Email, Phonenumber, PG_Type, PropertyID, City, PinCode, State, Location)
			c.IndentedJSON(http.StatusOK, output)	
		}
	}
}

func get_Allowner(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pms")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM Owner_Details")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	var output interface{}
	for results.Next() {
		var Id_ uint
		var F_Name string
		var L_Name string
		var Email string
		var Phonenumber string
		var PG_Type string
		var PropertyID uint
		var City string
		var PinCode int
		var State string
		var Location string
		err = results.Scan(&Id_, &F_Name, &L_Name, &Email, &Phonenumber, &PG_Type, &PropertyID, &City, &PinCode, &State, &Location)
		if err != nil {
			panic(err.Error())
		}
		output = fmt.Sprintf("%d %s %s %s %s %s %d %s %d %s %s \n", Id_, F_Name, L_Name, Email, Phonenumber, PG_Type, PropertyID, City, PinCode, State, Location)
		c.IndentedJSON(http.StatusOK, output)
	}
}

func Add_owner(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pms")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var add_newpro models.owner
	err = c.BindJSON(&add_newpro)
	if err != nil {
		return
	}

	query_data := fmt.Sprintf(`INSERT INTO Owner_Details VALUES("%d", "%s", "%s", "%s", "%s", "%s", "%d", "%s", "%d", "%s", "%s")`, add_newpro.ID, add_newpro.First_Name, add_newpro.Last_Name, add_newpro.Email, add_newpro.PhoneNumber, add_newpro.PG_Type, add_newpro.PropertyID, add_newpro.City, add_newpro.Pin_Code, add_newpro.State, add_newpro.Location)

	insert, err := db.Query(query_data)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yes, values added!")
}

func update_owner(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pms")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connected")
	defer db.Close()
	var data models.owner
	err = c.BindJSON(&data)
	if err != nil {
		return
	}
	query := fmt.Sprintf("UPDATE Owner_details SET First_Name='%s', Last_Name='%s', Email='%s', PhoneNumber='%s', PropertyID='%d', PG_Type='%s', City='%s',Pin_Code='%d', State='%s', Location='%s' WHERE ID=%d", data.First_Name, data.Last_Name, data.Email, data.PhoneNumber, data.PropertyID, data.PG_Type, data.City, data.Pin_Code, data.State, data.Location, data.ID)
	results, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer results.Close()
}

func delete_owner(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pms")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connected")
	defer db.Close()
	var get_pro models.owner
	err = c.BindJSON(&get_pro)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM Owner_Details WHERE ID=%d", get_pro.ID)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
}
