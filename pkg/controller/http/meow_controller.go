package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/query"
)

type MeowController struct {
	meowQueries *query.MeowQueries
}

func ProvideMeowController(meowQueries *query.MeowQueries) *MeowController {
	return &MeowController{
		meowQueries: meowQueries,
	}
}

func (controller *MeowController) Route(router *gin.Engine) {
	meow := router.Group("/meow")
	{
		meow.GET("/search", controller.Search)
	}
}

func (controller *MeowController) Search(context *gin.Context) {
	request := query.SearchRequest{
		Query: "",
		Skip:  0,
		Take:  20,
	}

	if err := context.Bind(&request); err != nil {
		context.Error(err)
		return
	}

	response, err := controller.meowQueries.Search(&request)
	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, response)
}
