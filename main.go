package main

import (
	"VoteSth-socketgo/component/appctx"
	"VoteSth-socketgo/middleware"
	answertransport "VoteSth-socketgo/modules/answer/transport"
	questiontransport "VoteSth-socketgo/modules/question/transport"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	if err := runServer(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}
func runServer(db *gorm.DB, secretKey string) error {
	appCtx := appctx.NewAppContext(db, secretKey)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	v1 := r.Group("/v1")
	question := v1.Group("/question")
	{
		question.POST("", questiontransport.CreateQuestion(appCtx))
		question.GET("/:id", questiontransport.GetQuestion(appCtx))
		question.GET("", questiontransport.ListQuestion(appCtx))
		question.DELETE("/:id", questiontransport.DeleteQuestion(appCtx))
	}
	answer := v1.Group("/answer")
	{
		answer.POST("", answertransport.CreateQuestion(appCtx))
		answer.GET("/:id", answertransport.GetAnswer(appCtx))
		answer.GET("", answertransport.ListAnswer(appCtx))
		answer.PATCH("/:id", answertransport.VoteAnswer(appCtx))
		answer.DELETE("/:id", answertransport.DeleteAnswer(appCtx))
	}
	return r.Run()
}
