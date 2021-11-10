package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/contract"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http/request"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/controller/http/response"
)

type MeowController struct {
	usecase contract.MeowUsecase
}

func ProvideMeowController(usecase contract.MeowUsecase) *MeowController {
	return &MeowController{
		usecase: usecase,
	}
}

func (controller *MeowController) Route(router *gin.Engine) {
	meow := router.Group("/meow")
	{
		meow.GET("/search", controller.Search)
	}
}

func (controller *MeowController) Search(context *gin.Context) {
	request := request.SearchRequest{
		Query: "",
		Skip:  0,
		Take:  20,
	}

	if err := context.Bind(&request); err != nil {
		context.Error(err)
		return
	}

	meows, err := controller.usecase.Search(request.Query, request.Skip, request.Take)
	if err != nil {
		context.Error(err)
		return
	}

	responses := make([]response.MeowResponse, 0, len(meows))
	for _, meow := range meows {
		responses = append(responses, response.MeowResponse{
			ID:        meow.ID,
			Body:      meow.Body,
			CreatedOn: meow.CreatedOn,
		})
	}

	context.JSON(http.StatusOK, responses)
}
