package unary

import (
    "fmt"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func BasicAuthUnary(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    // retrieve metadata from context
    md, ok := metadata.FromContext(ctx)
    _ = ok

    // validate 'Authorization' metadata
    isValid, err := ValidateBasicCredentials(md["authorization"])
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "falied to validate Authorization.")
    }

    if isValid {
        // proceed the request
        return handler(ctx, req)
    } else {
        return nil, grpc.Errorf(codes.Unauthenticated, "authentication required.")
    }
}

func ValidateBasicCredentials(authorization []string) (isValid bool, err error) {
    fmt.Println("Validating the Authorization in headers...")
    fmt.Println(authorization)
    return true, nil
}
