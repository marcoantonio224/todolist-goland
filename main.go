package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pb "todolist/protobuf/todo"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
	// Configure database
	db := configureDB()
	defer db.Close()
	// Configure grpc dial connection
	conn, err := grpc.Dial("grpc_server:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodolistClient(conn)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		res, err := client.GetFirstItem(c, &pb.FirstItemRequest{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.GetTodo()),
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not run gin server")
	}
}
