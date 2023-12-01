package service

import (
	errors "Ozon/domain"
	"Ozon/pkg/logger"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/hashicorp/go-hclog"
	"math/big"
	"regexp"
)

const lengthLink = 10
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

type Repository interface {
	GetFullLink(ctx context.Context, shortLink string) (string, error)
	CreateShortLink(ctx context.Context, fullLink, shortLink string) error
}

type Service struct {
	logger hclog.Logger
	data   Repository
}

func NewService(data Repository) *Service {
	return &Service{logger: logger.GetLogger(), data: data}
}

func (s *Service) CreateShortLink(ctx context.Context, fullLink string) (string, error) {
	re := regexp.MustCompile(".+\\.+")
	if !re.MatchString(fullLink) {
		return "", errors.ErrInvalidArgument
	}
	shortLink := s.MakeShortLink(fullLink)
	err := s.data.CreateShortLink(ctx, fullLink, shortLink)
	if err != nil {
		return "", err
	}
	return shortLink, nil
}

func (s *Service) MakeShortLink(fullLink string) string {
	shortLink := make([]byte, lengthLink, lengthLink)
	v := big.Int{}
	h := sha256.New()
	h.Write([]byte(fullLink))
	hash := h.Sum(nil)
	result := v.SetBytes([]byte(hex.EncodeToString(hash)))
	bigInt := result.Uint64()
	for i := 0; i < lengthLink; i++ {
		shortLink[i] = alphabet[bigInt%62]
		bigInt /= 62
	}
	return string(shortLink)
}

func (s *Service) GetFullLink(ctx context.Context, shortLink string) (string, error) {
	re := regexp.MustCompile("^[a-zA-Z0-9_]{10}")
	if !re.MatchString(shortLink) {
		return "", errors.ErrInvalidArgument
	}
	return s.data.GetFullLink(ctx, shortLink)
}
