package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/tmpgat/pb"
	"github.com/tmpgat/pb/pbconnect"
)

type AuthService struct {
}

var _ pbconnect.AuthServiceHandler = &AuthService{}

func (s *AuthService) Login(
	_ context.Context,
	_ *connect.Request[pb.User],
) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func main() {
	svc := &AuthService{}

	mux := http.NewServeMux()
	mux.Handle(pbconnect.NewAuthServiceHandler(svc))

	c := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			log.Printf("origin %#v\n", origin)
			return true
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	h := c.Handler(mux)

	handler := h2c.NewHandler(h, &http2.Server{})
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
