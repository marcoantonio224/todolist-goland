package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Todo struct {
	Id        int `gorm:"primary_key"`
	Todo      string
	Completed bool
}

func configureDB() *gorm.DB {
	godotenv.Load(".env")
	path := fmt.Sprintf("%s:%s@tcp(database:3306)/todolist?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"))

	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}

	return db
}

func getFirstTodo(db gorm.DB) Todo {
	var todo Todo
	db.First(&todo)
	return todo
}

func main() {
	db := configureDB()
	defer db.Close()
	firstTodo := getFirstTodo(*db)

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("First todo item: %s", firstTodo.Todo))
	})

	router.Run()
}
