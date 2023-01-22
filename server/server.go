package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var (
	clients = make(map[string]net.Conn)
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running on port", port)

	defer listener.Close()

	for {
		id := uuid.New().String()

		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		clients[id] = con
		go handleConnection(con, id)
	}
}

func handleConnection(con net.Conn, id string) {
	defer con.Close()

	clientReader := bufio.NewReader(con)
	con.Write([]byte("Your id is " + id + "\n"))
	log.Println("The client: " + id + " sucessfully connected to the server")

	for {
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			// clientRequest := strings.TrimSpace(clientRequest)
			clientRequest := strings.Split(clientRequest, " ")
			if clientRequest[0] != id {
				con.Write([]byte("Wrong ID, please re sent your ID \n"))
				break
			}

			if clientRequest[1] == "LIST" {
				log.Println("client requested server to list all connected clients")
				listClients(con)
			}

			if clientRequest[1] == "QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			} else {
				log.Println("The client " + id + " sent: ")
				log.Println(clientRequest[2])
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		// Responding to the client request
		if _, err = con.Write([]byte("\nRecieve the Request\n")); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}

func listClients(conn net.Conn) {
	var clientList []string
	for id := range clients {
		clientList = append(clientList, id)
	}
	b, _ := json.Marshal(clientList)

	conn.Write(b)
}
