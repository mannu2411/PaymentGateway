package server

import (
	"github.com/razorpay/razorpay-go"
	"os"
)

type Server struct {
	Client *razorpay.Client
}

func InitServices() *Server {
	return &Server{
		Client: razorpay.NewClient(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY")),
	}
}
