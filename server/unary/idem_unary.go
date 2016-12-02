package unary

import (
    "fmt"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func IdemUnary(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    // retrieve metadata from context
    md, ok := metadata.FromContext(ctx)
    _ = ok

    // validate 'Request-Id' metadata
    isValid, err := ValidateIdemByRequestId(md["request-id"])
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "failed to validate Request-Id.")
    }

    if isValid {
        // proceed the request
        return handler(ctx, req)
    } else {
        return nil, grpc.Errorf(codes.FailedPrecondition, "duplicated request cannot be accepted.")
    }
}

func ValidateIdemByRequestId(requestId []string) (isValid bool, err error) {
    fmt.Println("Validating the Request-Id in headers...")
    fmt.Println(requestId)
    return true, nil
}
