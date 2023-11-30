package handlers

import (
	errors "Ozon/domain"
	"Ozon/pkg/logger"
	protos "Ozon/protos/links"
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	MakeShortLink(fullLink string) string
	CreateShortLink(ctx context.Context, fullLink string) (string, error)
	GetFullLink(ctx context.Context, shortLink string) (string, error)
}

type Handler struct {
	logger  hclog.Logger
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{logger: logger.GetLogger(), service: service}
}

func (h *Handler) CreateShortLink(ctx context.Context, lr *protos.LinkRequest) (*protos.LinkResponse, error) {
	h.logger.Info("[GET SHORT LINK] Handle GetShortLink", "link", lr.GetLink())
	if lr.Link == "" {
		err := status.Newf(
			codes.InvalidArgument,
			"Link can not be empty",
		)
		return nil, err.Err()
	}
	sh, err := h.service.CreateShortLink(ctx, lr.Link)
	if err != nil {
		h.logger.Error("[ERROR] Trouble with adding data into DataBase", err.Error())
		return nil, err
	}
	return &protos.LinkResponse{Link: sh}, nil
}

func (h *Handler) GetFullLink(ctx context.Context, lr *protos.LinkRequest) (*protos.LinkResponse, error) {
	h.logger.Info("[GET FULL LINK] Handle GetFullLink", "link", lr.GetLink())
	if lr.Link == "" {
		err := status.Newf(
			codes.InvalidArgument,
			"Link can not be empty",
		)
		return nil, err.Err()
	}

	full, err := h.service.GetFullLink(ctx, lr.Link)
	if err != nil {
		if err == errors.ErrNoRecordFound {
			err := status.Newf(
				codes.NotFound,
				"Full link for this short link not found",
			)
			h.logger.Error("[ERROR] For this short link we do not have a full link")
			return nil, err.Err()
		}
		h.logger.Error("[ERROR] Trouble with taking data from DataBase", err.Error())
		return nil, err
	}
	return &protos.LinkResponse{Link: full}, nil
}
