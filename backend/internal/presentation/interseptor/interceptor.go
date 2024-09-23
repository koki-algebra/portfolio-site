package interseptor

import "connectrpc.com/connect"

func NewCommonInterceptors() []connect.Interceptor {
	return []connect.Interceptor{
		UnexpectedErrorHandler(),
	}
}
