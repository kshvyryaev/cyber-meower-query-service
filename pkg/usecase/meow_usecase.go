package usecase

import (
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/contract"
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/domain"
	"github.com/pkg/errors"
)

type MeowUsecase struct {
	repository contract.MeowRepository
}

func ProvideMeowUsecase(repository contract.MeowRepository) *MeowUsecase {
	return &MeowUsecase{
		repository: repository,
	}
}

func (usecase *MeowUsecase) Search(query string, skip int, take int) ([]domain.Meow, error) {
	meows, err := usecase.repository.Search(query, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "meow queries")
	}

	return meows, nil
}
