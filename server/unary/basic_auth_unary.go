package unary

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"strings"

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
	fmt.Printf("Validating the Authorization(%v) in headers...\n", authorization)
	if len(authorization) <= 0 {
		return false, nil
	}
	// parse clientId and clientSecret from Authorization in headers
	id, secret, err := ParseBasicAuthorization(authorization[0])
	if err != nil {
		return false, err
	}
	fmt.Printf("[CREDENTIALS] Validated id: %s, secret: %s \n", id, secret)

	return true, nil
}

func ParseBasicAuthorization(authorzation string) (id string, secret string, err error) {
	authorizationSplitedArr := strings.Split(authorzation, " ")
	if len(authorizationSplitedArr) < 2 {
		return "", "", errors.New("Invalid authorization.")
	}
	if authType := authorizationSplitedArr[0]; !(authType == "Basic") {
		return "", "", errors.New("Invalid type of authorization: " + authType)
	}

	authClaim := authorizationSplitedArr[1]
	authorizationDecodedBytes, derr := b64.URLEncoding.DecodeString(authClaim)
	if derr != nil {
		return "", "", derr
	}

	authorizationDecodedStringArr := strings.Split(string(authorizationDecodedBytes), ":")
	if len(authorizationDecodedStringArr) < 2 {
		return "", "", errors.New("Invalid authorization.")
	}

	id = authorizationDecodedStringArr[0]
	secret = authorizationDecodedStringArr[1]
	return id, secret, nil
}
