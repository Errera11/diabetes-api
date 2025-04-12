package utils

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func SetUserIDInCtx(ctx context.Context, token string) {
	header := metadata.Pairs("sid", token)
	grpc.SendHeader(ctx, header)
}
func SetDeleteSessionFlagInCtx(ctx context.Context) {
	header := metadata.Pairs("sid-del", "true")
	grpc.SendHeader(ctx, header)
}
