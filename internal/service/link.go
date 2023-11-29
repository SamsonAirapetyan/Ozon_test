package service

import (
	"Ozon/pkg/logger"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"math/big"
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
	shortLink := MakeShortLink(fullLink)
	err := s.data.CreateShortLink(ctx, fullLink, shortLink)
	if err != nil {
		return "", err
	}
	return shortLink, nil
}

func MakeShortLink(fullLink string) string {
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
	fmt.Println("sucesfully made short link ", string(shortLink))
	return string(shortLink)
}

func (s *Service) GetFullLink(ctx context.Context, shortLink string) (string, error) {
	fullLink, err := s.data.GetFullLink(ctx, shortLink)
	if err != nil {
		return "", err
	}
	return fullLink, nil
}
