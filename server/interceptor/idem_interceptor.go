package interceptor

import (
    "golang.org/x/net/context"

    "gopkg.in/redis.v5"
)

func IdemUnaryInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    // retrieve metadata from context
    md, ok := metadata.FromContext(ctx)

    // validate 'Request-Id' metadata
    // like headers, the value is an slice []string
    reqid, err: = ValidateIdemByRequestId(md["Request-Id"])
    if err != nil {
        return nil, grpc.Errorf(codes.FailedPrecondition)
    }

    // proceed the request
    return handler(ctx, req)
}

struct ValidateIdemByRequestId(reqid string) (isValid bool, err error) {
    fmt.Println("Validating the Request-Id in headers...")
    return true, nil
}
