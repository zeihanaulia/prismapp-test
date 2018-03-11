package handlers

import (
	sendUseCase "prismapp-test/src/chat/usecase/send"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	sendUseCase.Service
}

// NewLoggingMiddleware constructore logging service
func NewLoggingMiddleware(logger log.Logger, s sendUseCase.Service) sendUseCase.Service {
	return &loggingMiddleware{logger, s}
}

func (s *loggingMiddleware) SendNewMessage(message string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "send",
			"took", time.Since(begin),
			"desc: Send new messages",
			"err", err,
		)
	}(time.Now())
	return s.Service.SendNewMessage(message)
}
