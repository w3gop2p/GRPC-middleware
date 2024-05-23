package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	order "github.com/w3gop2p/GRPC-middleware/ch8/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func getTlsCredentials() (credentials.TransportCredentials, error) {
	clientCert, clientCertErr := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if clientCertErr != nil {
		return nil, fmt.Errorf("could not load client key pair : %v", clientCertErr)
	}

	certPool := x509.NewCertPool()
	caCert, caCertErr := ioutil.ReadFile("cert/ca-cert.pem")
	if caCertErr != nil {
		return nil, fmt.Errorf("could not read Cert CA : %v", caCertErr)
	}

	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to append CA cert.")
	}

	return credentials.NewTLS(&tls.Config{
		ServerName:   "*.microservices.dev",
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}), nil
}

func main() {
	tlsCredentials, tlsCredentialsErr := getTlsCredentials()
	if tlsCredentialsErr != nil {
		log.Fatalf("failed to get tls credentials. %v", tlsCredentialsErr)
	}
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to connect order service. Err: %v", err)
	}

	defer conn.Close()

	orderClient := order.NewOrderServiceClient(conn)
	log.Println("Creating order...")
	orderResponse, errCreate := orderClient.Create(context.Background(), &order.CreateOrderRequest{
		UserId:    1,
		ProductId: 1,
		Price:     2,
	})

	if errCreate != nil {
		log.Fatalf("Failed to create order. %v", errCreate)
	} else {
		log.Printf("Order %d is created successfully.", orderResponse.OrderId)
	}
}
