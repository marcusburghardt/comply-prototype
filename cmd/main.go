package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/marcusburghardt/comply-prototype/proto"
)

const ServerSocketFile = "server_socket.sock"

func getSocketPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	socketPath := homeDir + "/" + ServerSocketFile
	return socketPath, nil
}

var rootCmd = &cobra.Command{
	Use:   "comply-prototype",
	Short: "comply-prototype CLI command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide an action")
			os.Exit(1)
		}

		action := args[0]

		// connect
		socket, err := getSocketPath()
		if err != nil {
			log.Fatalf("Failed to determine the server socket path: %v", err)
		}
		conn, err := grpc.Dial(
			"unix://"+socket,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("Failed to connect to plugin gRPC server: %v", err)
		}
		defer conn.Close()

		client := proto.NewScanServiceClient(conn)

		// call
		resp, err := client.Execute(context.Background(), &proto.ScanRequest{Action: action})
		if err != nil {
			log.Fatalf("Error calling plugin: %v", err)
		}

		// process response
		if resp.GetReturnCode() == 0 {
			fmt.Println("Command executed successfully")
		} else {
			fmt.Println("Command failed")
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
