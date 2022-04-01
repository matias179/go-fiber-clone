package http

import (
	"bufio"
	"fmt"
	"github.com/matias179/go-fiber-clone/internal/config"
	"log"
	"net"
	"os"
	"regexp"
)

// New start the app.
func New() {
	config.SetEnvironmentVariables()
	fmt.Println("calling new here!")
	l, err := net.Listen("tcp", "localhost:"+config.Envs.PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Server running on localhost:" + config.Envs.PORT)
	// Close the listener when the application closes.
	defer l.Close()

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		// Print client connection address.
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		// Handle connections concurrently in a new goroutine.
		go handleConnection(c)
	}
}

// handleConnection handles logic for a single connection request.
func handleConnection(conn net.Conn) {
	// Buffer client input until a newline.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	// Log URI visited
	URIDetails := string(buffer[:len(buffer)-1])
	log.Println("Client enter to:", URIDetails)

	// primitive parse of URI from string
	regExp := regexp.MustCompile(`\/\w+`)
	URI := regExp.FindAllString(URIDetails, 1)
	for _, v := range URI {
		switch v {
		case "/hello":
			conn.Write([]byte("Hello, World 👋!"))
		case "/test":
			conn.Write([]byte("Keep trying"))
		}
	}
	conn.Close()
}
