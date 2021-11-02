package main

import (
	"VoteSth-socketgo/component/appctx"
	"VoteSth-socketgo/middleware"
	answertransport "VoteSth-socketgo/modules/answer/transport"
	questiontransport "VoteSth-socketgo/modules/question/transport"
	"VoteSth-socketgo/skio"
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
	corspolicy := cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:3000","http://localhost:3000"},
		AllowMethods: []string{"POST, OPTIONS, GET, PUT, DELETE", "PATCH"},
		AllowHeaders: []string{ "XMLHttpRequest, Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With"},
		AllowCredentials: true,
	})
	r.Use(corspolicy)
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	r.StaticFile("/demo/","asset/index.html")
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
	rtEngine := skio.NewEngine()
	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}
	//startSocketIOServer(r,appCtx)
	return r.Run()
}
//
//func startSocketIOServer(engine *gin.Engine, ctx appctx.AppContext){
//	 server := socketio.NewServer(nil)
//	 server.OnConnect("/", func(s socketio.Conn) error{
//		 fmt.Println("Connected",s.ID(),"IP:",s.RemoteHeader())
//		 return nil
//	 })
//	 server.OnEvent("/","test", func(s socketio.Conn,msg string) {
//		 fmt.Println("Server receive notice:", msg)
//		 s.Emit("test_notice","Hello")
//	 })
//	 server.OnEvent("/","getlistanswer",func(s socketio.Conn,msg string) {
//		 uid, _ := common.FromBase64(msg)
//		 db := ctx.GetMainDbConnection()
//		 store := answerstorage.NewSQLStore(db)
//		 bus := answerbus2.NewGetAnswerBus(store)
//		 data,_ := bus.GetAnswer(,int(uid.GetLocalID()))
//		 s.Emit("test_notice","Hello")
//	 })
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}