package server

import (
	"golang.org/x/net/context"

	pb "github.com/trumanw/cloud-auth-go/pb"
)

type clientCredentialsServer struct{}

func newClientCredentialsServer() pb.CilentCredentialsServiceServer {
	return new(clientCredentialsServer)
}

func (s *clientCredentialsServer) CreateClientCredentials(ctx context.Context, clientCredentials *pb.ClientCredentialsCreatedEvent) (*pb.Token, error) {
	token := &pb.Token{
		AccessToken:  "EEwJ6tF9x5WCIZDYzyZGaz6Khbw7raYRIBV_WxVvgmsG",
		TokenType:    "Bearer",
		ExpiresIn:    28800,
		RefreshToken: "8xLOxBtZp8",
		Scope:        "https://api.paypal.com/v1/payments/.*%20https://api.paypal.com/v1/vault/credit-card%20https://api.paypal.com/v1/vault/credit-card/.*",
		State:        "DUMP",
	}
	return token, nil
}
