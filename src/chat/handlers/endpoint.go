package handlers

import (
	"context"
	"math"
	retrieveUseCase "prismapp-test/src/chat/usecase/retrieve"
	sendUseCase "prismapp-test/src/chat/usecase/send"

	"github.com/go-kit/kit/endpoint"
)

// MakeSendNewMessageEndpoint specific endpoint for SignUp Validate
func MakeSendNewMessageEndpoint(sendService sendUseCase.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendNewMessageRequest)
		err := sendService.SendNewMessage(req.Message)
		if err != nil {
			return nil, err
		}

		return StandardResponse{
			HTTPStatusCode: 200,
			Success:        true,
			StateCode:      "SEND_SUCCESS",
			Message:        "Send message successfully",
			Data:           nil,
		}, nil
	}
}

func MakeRetrievePreviousMessageEndpoint(retrieveService retrieveUseCase.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RetrievePreviousMessageRequest)

		response, count, err := retrieveService.RetrievePreviousMessage(req.Page, req.Limit)
		if err != nil {
			return nil, err
		}

		d := float64(count) / float64(req.Limit)
		totalPages := int(math.Ceil(d))

		return StandardResponse{
			HTTPStatusCode: 200,
			Success:        true,
			StateCode:      "RETRIEVE_SUCCESS",
			Message:        "Retrieve message successfully",
			Data:           response,
			Pagination: &Pagination{
				TotalCount:  count,
				TotalPages:  totalPages,
				CurrentPage: req.Page,
				Limit:       req.Limit,
			},
		}, nil
	}
}
