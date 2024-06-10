package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connect MySQL")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/schedule-maker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO user (account, password) VALUES('test', 'test')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	router := gin.Default()
	router.GET("/returnJSON", returnJSON)
	router.GET("/returnJSON2", returnJSON2)
	router.Run(":8080")
}

func returnJSON(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func returnJSON2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "from returnJSON2",
	})
}
