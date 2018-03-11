package handlers

import (
	"context"
	"net/http"
	retrieveUseCase "prismapp-test/src/chat/usecase/retrieve"
	sendUseCase "prismapp-test/src/chat/usecase/send"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(
	sendService sendUseCase.Service,
	retrieveService retrieveUseCase.Service,
	logger kitlog.Logger,
) http.Handler {

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()

	r.Handle("/chats", httptransport.NewServer(
		MakeSendNewMessageEndpoint(sendService),
		DecodeSendNewMessageRequest,
		EncodeResponse,
		opts...,
	)).Methods("POST")

	r.Handle("/chats", httptransport.NewServer(
		MakeRetrievePreviousMessageEndpoint(retrieveService),
		DecodeRetrievePreviousMessageRequest,
		EncodeResponse,
		opts...,
	)).Methods("GET")

	return r
}

// DecodeSendRequest decoding register request from client
func DecodeSendNewMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	sendRequest := SendNewMessageRequest{
		Message: r.FormValue("message"),
	}

	return sendRequest, nil
}

func DecodeRetrievePreviousMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	retrieveRequest := RetrievePreviousMessageRequest{
		Page:  1,
		Limit: 20,
	}

	return retrieveRequest, nil
}
