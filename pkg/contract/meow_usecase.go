package contract

import "github.com/kshvyryaev/cyber-meower-query-service/pkg/domain"

type MeowUsecase interface {
	Search(query string, skip int, take int) ([]domain.Meow, error)
}
