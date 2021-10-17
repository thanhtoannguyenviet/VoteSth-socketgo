package answertransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	answerbus "VoteSth-socketgo/modules/answer/bus"
	answerstorage "VoteSth-socketgo/modules/answer/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnswer(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase64(c.Param("id"))
		if err != nil {
			common.ErrInvalidRequest(err)
		}
		store := answerstorage.NewSQLStore(ctx.GetMainDbConnection())
		bus := answerbus.NewGetAnswerBus(store)
		data, err := bus.GetAnswer(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
