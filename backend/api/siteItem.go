package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type listAllSiteRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=3,max=10"`
}

func (s *Server) getAllSiteItems(ctx *gin.Context) {

	var req listAllSiteRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {

		if req.PageID < 1 {

			err := errors.New("PageID must be at least 1")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		if req.PageSize < 3 || req.PageSize > 10 {

			err := errors.New("PageSize must be between 3 and 10")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

	}
	skip := (req.PageID - 1) * req.PageSize
	limit := req.PageSize
	filter := bson.M{}

	siteItems, err := s.store.GetAllSiteItem(ctx, "siteItem", filter, skip, limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"siteItems": siteItems})
}
