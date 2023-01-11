/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/adaptive-scale/logchain/internal"
	"github.com/adaptive-scale/logchain/pkg/logchain"
	"github.com/adaptive-scale/logchain/pkg/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// startCmd represents the start command

var (
	port string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a logchain service",
	Long:  `//TODO`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.AutomaticEnv()

		config := internal.Configuration{
			DatabaseType: viper.GetString("DATABASE_TYPE"),
			DSN:          viper.GetString("DSN"),
			Port:         viper.GetInt("PORT"),
			Protected:    viper.GetBool("PROTECTED"),
			CertLocation: viper.GetString("CERT_LOCATION"),
			KeyLocation:  viper.GetString("KEY_LOCATION"),
		}

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var s *grpc.Server

		if config.Protected {

			tlsCredentials, err := loadTLSCredentials(config.CertLocation, config.KeyLocation)
			if err != nil {
				log.Fatal("cannot load TLS credentials: ", err)
			}
			s = grpc.NewServer(grpc.Creds(tlsCredentials))
		} else {
			s = grpc.NewServer()
		}

		logchain.RegisterLogChainServer(s, service.New(config))
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func loadTLSCredentials(cert string, key string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
