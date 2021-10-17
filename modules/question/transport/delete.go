package questiontransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	questionbus "VoteSth-socketgo/modules/question/bus"
	questionstorage "VoteSth-socketgo/modules/question/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteQuestion(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase64(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := questionstorage.NewSQLStore(appctx.GetMainDbConnection())
		bus := questionbus.NewDeleteQuestionBus(store)

		if err := bus.DeleteQuestion(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
