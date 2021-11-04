package search

import (
	"github.com/kshvyryaev/cyber-meower-query-service/pkg/domain"
)

type MeowRepository interface {
	Search(query string, skip int, take int) ([]domain.Meow, error)
}
