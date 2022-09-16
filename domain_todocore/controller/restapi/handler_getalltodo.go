package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"todoapps/domain_todocore/usecase/getalltodo"
	"todoapps/shared/infrastructure/logger"
	"todoapps/shared/infrastructure/util"
	"todoapps/shared/model/payload"
)

// getAllTodoHandler ...
func (r *Controller) getAllTodoHandler() gin.HandlerFunc {

	type request struct {
		Page int64 `form:"page,omitempty,default=0"`
		Size int64 `form:"size,omitempty,default=0"`
	}

	type response struct {
		Count int64 `json:"count"`
		Items []any `json:"items"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.Bind(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req getalltodo.InportRequest
		req.Page = jsonReq.Page
		req.Size = jsonReq.Size

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.GetAllTodoInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
