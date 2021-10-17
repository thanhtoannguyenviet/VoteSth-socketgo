package answertransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	answerbus "VoteSth-socketgo/modules/answer/bus"
	answerstorage "VoteSth-socketgo/modules/answer/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAnswer(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fullfill()
		store := answerstorage.NewSQLStore(ctx.GetMainDbConnection())
		bus := answerbus.NewListAnswerBus(store)
		data, err := bus.ListAnswer(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask()
			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
