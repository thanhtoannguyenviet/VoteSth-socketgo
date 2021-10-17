package questiontransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	questionbus "VoteSth-socketgo/modules/question/bus"
	questionstorage "VoteSth-socketgo/modules/question/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase64(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := questionstorage.NewSQLStore(appCtx.GetMainDbConnection())
		bus := questionbus.NewGetQuestionBus(store)

		data, err := bus.GetQuestion(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		data.Mask()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
