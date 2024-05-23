package middleware

import (
	"context"
	"fmt"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
)

func CircuitBreakerClientInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		result, cbErr := cb.Execute(func() (interface{}, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return nil, err
			}
			return reply, nil
		})
		if cbErr != nil {
			return cbErr
		}
		// Process the result if needed
		processResult(result)
		return nil
	}
}
func processResult(result interface{}) {
	fmt.Println("result is processed", result)
}
