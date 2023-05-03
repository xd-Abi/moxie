package main

import (
	"context"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/xd-Abi/moxie/pkg/config"
	"github.com/xd-Abi/moxie/pkg/logging"
	"github.com/xd-Abi/moxie/pkg/network"
	"github.com/xd-Abi/moxie/pkg/proto/hello"
)

var (
	log = logging.New()
)

type HelloServiceServer struct {
	SendGridApiKey string
	SendGridSender string
	hello.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) SendWelcomeEmail(ctx context.Context, request *hello.WelcomeEmailRequest) (*hello.WelcomeEmailResponse, error) {
	from := mail.NewEmail("Moxie", s.SendGridSender)
	subject := "Welcome to moxie 😎"
	to := mail.NewEmail(request.Username, request.Email)
	plainTextContent := fmt.Sprintf("Hi %s", request.Username)
	htmlContent := fmt.Sprintf("Hi <strong>%s</strong>", request.Username)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.SendGridApiKey)
	_, err := client.Send(message)

	if err != nil {
		log.Error("%v", err)
		// @TODO: Return error
	}

	return &hello.WelcomeEmailResponse{
		Success: true,
	}, nil
}

func main() {
	config.LoadEnvVariables(log)

	app := network.NewMicroServiceServer(config.GetUint("HELLO_PORT"), log)
	hello.RegisterHelloServiceServer(app.InternalServer, &HelloServiceServer{
		SendGridApiKey: config.GetString("HELLO_SENDGRID_API_KEY"),
		SendGridSender: config.GetString("HELLO_SENDGRID_SENDER"),
	})
	app.Start()
}
