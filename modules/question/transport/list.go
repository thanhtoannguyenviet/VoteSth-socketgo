package questiontransport

import (
	"VoteSth-socketgo/common"
	"VoteSth-socketgo/component/appctx"
	questionbus "VoteSth-socketgo/modules/question/bus"
	questionstorage "VoteSth-socketgo/modules/question/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fullfill()
		store := questionstorage.NewSQLStore(appCtx.GetMainDbConnection())
		bus := questionbus.NewListQuestionBus(store)

		rs, err := bus.ListQuestion(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}
		for i := range rs {
			rs[i].Mask()
			if i == len(rs)-1 {
				paging.NextCursor = rs[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(rs, paging, nil))
	}
}
