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
	log     hclog.Logger
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{log: logger.GetLogger(), service: service}
}

func (h *Handler) CreateShortLink(ctx context.Context, lr *protos.LinkRequest) (*protos.LinkResponse, error) {
	h.log.Info("[GET SHORT LINK] Handle GetShortLink", "link", lr.GetLink())
	if lr.Link == "" {
		err := status.Newf(
			codes.InvalidArgument,
			"Link can not be empty",
		)
		err, wde := err.WithDetails(lr)
		if wde != nil {
			return nil, wde
		}
		return nil, err.Err()
	}
	sh, err := h.service.CreateShortLink(ctx, lr.Link)
	if err != nil {
		if err == errors.ErrInvalidArgument {
			er := status.Newf(
				codes.InvalidArgument,
				"Is not a link",
			)
			er, wde := er.WithDetails(lr)
			if wde != nil {
				return nil, wde
			}
			return nil, er.Err()
		}
		h.log.Error("[ERROR] Trouble with adding data into DataBase", err.Error())
		return nil, err
	}
	return &protos.LinkResponse{Link: sh}, nil
}

func (h *Handler) GetFullLink(ctx context.Context, lr *protos.LinkRequest) (*protos.LinkResponse, error) {
	h.log.Info("[GET FULL LINK] Handle GetFullLink", "link", lr.GetLink())

	if len(lr.Link) != 10 {
		err := status.Newf(
			codes.InvalidArgument,
			"The link must be 10 characters long",
		)
		err, wde := err.WithDetails(lr)
		if wde != nil {
			return nil, wde
		}
		return nil, err.Err()
	}

	full, err := h.service.GetFullLink(ctx, lr.Link)
	if err != nil {
		if err == errors.ErrNoRecordFound {
			er := status.Newf(
				codes.NotFound,
				"Full link for this short link not found",
			)
			er, wde := er.WithDetails(lr)
			if wde != nil {
				return nil, wde
			}
			h.log.Error("[ERROR] For short link we do not have a full link")
			return nil, er.Err()
		} else if err == errors.ErrInvalidArgument {
			er := status.Newf(
				codes.InvalidArgument,
				"Invalid short link value",
			)
			er, wde := er.WithDetails(lr)
			if wde != nil {
				return nil, wde
			}
			h.log.Error("[ERROR] Invalid short link")
			return nil, er.Err()
		}
		h.log.Error("[ERROR] Trouble with taking data from DataBase", err.Error())
		return nil, err
	}
	return &protos.LinkResponse{Link: full}, nil
}
