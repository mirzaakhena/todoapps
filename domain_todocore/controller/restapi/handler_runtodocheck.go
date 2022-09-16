package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"todoapps/domain_todocore/usecase/runtodocheck"
	"todoapps/shared/infrastructure/logger"
	"todoapps/shared/infrastructure/util"
	"todoapps/shared/model/payload"
)

// runTodoCheckHandler ...
func (r *Controller) runTodoCheckHandler() gin.HandlerFunc {

	type request struct {
		//TodoID string `json:"todo_id"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		//var jsonReq request
		//if err := c.BindJSON(&jsonReq); err != nil {
		//	r.Log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		//	return
		//}

		var req runtodocheck.InportRequest
		//req.TodoID = jsonReq.TodoID
		req.TodoID = c.Param("todo_id")

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunTodoCheckInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
