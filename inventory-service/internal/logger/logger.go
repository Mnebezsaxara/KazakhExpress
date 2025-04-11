package logger

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LoggingInterceptor создает gRPC interceptor для логирования запросов
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		// Логируем входящий запрос
		InfoLogger.Printf("Request: method=%s, request=%v", info.FullMethod, req)

		// Вызываем обработчик
		resp, err := handler(ctx, req)

		// Логируем результат
		duration := time.Since(start)
		if err != nil {
			st, _ := status.FromError(err)
			ErrorLogger.Printf("Response: method=%s, duration=%s, error=%v, code=%s", 
				info.FullMethod, duration, err, st.Code())
		} else {
			InfoLogger.Printf("Response: method=%s, duration=%s, response=%v", 
				info.FullMethod, duration, resp)
		}

		return resp, err
	}
}

// LogError логирует ошибку с дополнительным контекстом
func LogError(method string, err error) error {
	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Internal, err.Error())
	}
	ErrorLogger.Printf("Method: %s, Error: %v, Code: %s", method, err, st.Code())
	return st.Err()
}

// LogInfo логирует информационное сообщение
func LogInfo(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
} 