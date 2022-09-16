package restapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"todoapps/domain_todocore/usecase/runtodocreate"
	"todoapps/shared/infrastructure/logger"
	"todoapps/shared/infrastructure/util"
	"todoapps/shared/model/payload"
)

// runTodoCreateHandler ...
func (r *Controller) runTodoCreateHandler() gin.HandlerFunc {

	type request struct {
		Message string `json:"message"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runtodocreate.InportRequest
		req.Message = jsonReq.Message
		req.Now = time.Now()

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunTodoCreateInport.Execute(ctx, req)
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
