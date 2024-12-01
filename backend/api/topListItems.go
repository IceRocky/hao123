package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type topItemListRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=3,max=10"`
}

func (s *Server) getAllTopListItems(ctx *gin.Context) {
	var req topItemListRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {

		if req.PageID < 1 {

			err := errors.New("PageID must be at least 1")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		if req.PageSize < 5 || req.PageSize > 10 {

			err := errors.New("PageSize must be between 3 and 10")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

	}

	skip := (req.PageID - 1) * req.PageSize
	limit := req.PageSize
	filter := bson.M{}

	topListItems, err := s.store.GetAllTopListItem(ctx, "topListItems", filter, skip, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"topListItems": topListItems})
}
