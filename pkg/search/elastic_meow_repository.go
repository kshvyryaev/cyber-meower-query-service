package search

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/domain"
	"github.com/pkg/errors"
)

const _meowIndex = "meows"

type elasticSearchMeowResponse struct {
	Took int64
	Hits struct {
		Total struct {
			Value int64
		}
		Hits []*struct {
			Source domain.Meow `json:"_source"`
		}
	}
}

type ElasticMeowRepository struct {
	client *elasticsearch.Client
}

func ProvideElasticMeowRepository(client *elasticsearch.Client) *ElasticMeowRepository {
	return &ElasticMeowRepository{
		client: client,
	}
}

func (repository *ElasticMeowRepository) Search(query string, skip int, take int) ([]domain.Meow, error) {
	request := esapi.SearchRequest{
		Index:          []string{_meowIndex},
		Query:          query,
		From:           &skip,
		Size:           &take,
		DocvalueFields: []string{"body"},
		TrackTotalHits: true,
	}

	response, err := request.Do(context.Background(), repository.client)
	if err != nil {
		return nil, errors.Wrap(err, "elastic meow repository")
	}
	defer response.Body.Close()

	if response.IsError() {
		// TODO: Handle elastic error
		return nil, errors.New("elastic meow repository error")
	}

	meows, err := repository.getSerachResponse(response)
	if err != nil {
		return nil, errors.Wrap(err, "elastic meow repository")
	}

	return meows, nil
}

func (repository *ElasticMeowRepository) getSerachResponse(response *esapi.Response) ([]domain.Meow, error) {
	responseBody := elasticSearchMeowResponse{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	var meows []domain.Meow
	for _, hit := range responseBody.Hits.Hits {
		meows = append(meows, hit.Source)
	}

	return meows, nil
}

var ElasticMeowRepositorySet = wire.NewSet(
	ProvideElasticMeowRepository,
	wire.Bind(new(MeowRepository), new(*ElasticMeowRepository)),
)
