package main

import (
	"VoteSth-socketgo/component/appctx"
	"VoteSth-socketgo/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	secretKey := os.Getenv("SecretKey")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := runServer(db,secretKey); err != nil{
		log.Fatalln(err)
	}
}
func runServer(db *gorm.DB,secretKey string) error {
	appCtx := appctx.NewAppContext(db,secretKey)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	return r.Run()
}

