package questiontransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	questionbus "VoteSth-socketgo/modules/question/bus"
	questionmodel "VoteSth-socketgo/modules/question/model"
	questionstorage "VoteSth-socketgo/modules/question/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateQuestion(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data questionmodel.Question
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := questionstorage.NewSQLStore(appctx.GetMainDbConnection())

		bus := questionbus.NewCreateQuestionBus(store)

		if err := bus.CreateQuestion(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
