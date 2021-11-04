package query

import (
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/search"
	"github.com/pkg/errors"
)

type MeowQueries struct {
	meowRepository search.MeowRepository
}

func ProvideMeowQueries(meowRepository search.MeowRepository) *MeowQueries {
	return &MeowQueries{
		meowRepository: meowRepository,
	}
}

func (queries *MeowQueries) Search(request *SearchRequest) ([]MeowResponse, error) {
	meows, err := queries.meowRepository.Search(request.Query, request.Skip, request.Take)
	if err != nil {
		return nil, errors.Wrap(err, "meow queries")
	}

	response := make([]MeowResponse, 0, len(meows))
	for _, meow := range meows {
		response = append(response, MeowResponse{
			ID:        meow.ID,
			Body:      meow.Body,
			CreatedOn: meow.CreatedOn,
		})
	}

	return response, nil
}
