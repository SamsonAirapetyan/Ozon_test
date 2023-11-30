package repository

import (
	errors "Ozon/domain"
	"Ozon/pkg/logger"
	"context"
	"github.com/hashicorp/go-hclog"
	"sync"
)

type MemoryRepository struct {
	logger hclog.Logger
	base   map[string]string
	mu     *sync.Mutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		logger: logger.GetLogger(),
		base:   make(map[string]string),
		mu:     &sync.Mutex{},
	}
}

// Есть вариант сюда добавить ошибку, мол есть уже такая ссылка
// Ограничить  размер кеша, а то долго прога работать не будет при огромном количестве входных данных
func (r *MemoryRepository) CreateShortLink(ctx context.Context, fullLink, shortLink string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.base[shortLink]; !ok {
		r.base[shortLink] = fullLink
	}
	return nil
}

func (r *MemoryRepository) GetFullLink(ctx context.Context, shortLink string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	fullLink, ok := r.base[shortLink]
	if !ok {
		return "", errors.ErrNoRecordFound
	}
	return fullLink, nil
}
