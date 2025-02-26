package utils

import "context"

type errorContextKey struct{}

func SetErrorInContext(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorContextKey{}, err)
}

func GetErrorFromContext(ctx context.Context) error {
	if err, ok := ctx.Value(errorContextKey{}).(error); ok {
		return err
	}
	return nil
}
