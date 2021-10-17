package answertransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	answerbus "VoteSth-socketgo/modules/answer/bus"
	answermodel "VoteSth-socketgo/modules/answer/model"
	answerstorage "VoteSth-socketgo/modules/answer/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateQuestion(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data answermodel.Answer
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := answerstorage.NewSQLStore(appctx.GetMainDbConnection())
		bus := answerbus.NewCreateAnswerBus(store)

		if err := bus.CreateAnswer(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
