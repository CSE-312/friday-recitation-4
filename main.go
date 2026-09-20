package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
)

//TODO Fix the Code below!!

func insert_data(data string, connection net.Conn) {
	_, err := db.Exec(`
		ADD INTO greetings (message)
		VALUES ($1)
		`, data)

	if err != nil {
		message := fmt.Sprintf("Error inserting into table! Error: %v\n", err)
		handle_response(message, connection) // Send a response using the "net/http" package. Do not Touch!!
		log.Println(message)
		return
	}

	message := fmt.Sprintf("Added: %v to DB", data)

	handle_response(message, connection) // Send a response using the "net/http" package. Do not Touch!!
	if err != nil {
		log.Printf("Error writing back to socket! Error: %v\n", err)
	}
}

func get_data(connection net.Conn) {
	records, err := db.Query("SELECT  FROM greetings")
	if err != nil {
		message := fmt.Sprintf("query failed: %v", err)
		handle_response(message, connection) // Send a response using the "net/http" package. Do not Touch!!
		log.Println(message)
		return
	}

	var result strings.Builder

	for records.Next() {
		var id int
		var message string
		records.Scan(&id, &message)
		result.WriteString(message + "\n")
	}

	handle_response(result.String(), connection) // Send a response using the "net/http" package. Do not Touch!!
}

func main() {
	godotenv.Load()
	InitDatabase()

	defer CloseDB()

	tcpServer, err := net.Listen("tcp", ":8280")
	if err != nil {
		log.Printf("Error starting TCP server! Error: %v\n", err)
		return
	}

	if db == nil {
		log.Println("database not initialized")
		return
	}

	PrintStartUpMessage()

	defer tcpServer.Close()

	for {
		connection, err := tcpServer.Accept()
		if err != nil {
			log.Printf("Error accepting connection! Error: %v", err)
			continue
		}
		log.Println("Handling new client")
		go handler(connection)
	}

}

// TODO DO NOT TOUCH CODE BELOW THIS LINE!!

// Do not touch
func PrintStartUpMessage() {
	banner := `
 ________    _____    _______          ________  _______   ________  ___      ___ _______   ________     
|\_____  \  / __  \  /  ___  \        |\   ____\|\  ___ \ |\   __  \|\  \    /  /|\  ___ \ |\   __  \    
\|____|\ /_|\/_|\  \/__/|_/  /|       \ \  \___|\ \   __/|\ \  \|\  \ \  \  /  / | \   __/|\ \  \|\  \   
      \|\  \|/ \ \  \__|//  / /        \ \_____  \ \  \_|/_\ \   _  _\ \  \/  / / \ \  \_|/_\ \   _  _\  
     __\_\  \   \ \  \  /  /_/__        \|____|\  \ \  \_|\ \ \  \\  \\ \    / /   \ \  \_|\ \ \  \\  \| 
    |\_______\   \ \__\|\________\        ____\_\  \ \_______\ \__\\ _\\ \__/ /     \ \_______\ \__\\ _\ 
    \|_______|    \|__| \|_______|       |\_________\|_______|\|__|\|__|\|__|/       \|_______|\|__|\|__|
                                         \|_________|                                                    
`
	log.Println(banner)
	log.Println("listening on port 8280")
}

// Do not touch
func handler(connection net.Conn) {
	bufReader := bufio.NewReader(connection)
	req, err := http.ReadRequest(bufReader)

	if err != nil {
		log.Println("Could not read request")
		return
	}

	if req.URL.Path == "/" {
		get_data(connection)
	}

	if after, ok := strings.CutPrefix(req.URL.Path, "/Add/"); ok {
		insert_data(after, connection)
	}
	return
}

// Do not touch
func handle_response(message string, connection net.Conn) {
	resp := &http.Response{
		StatusCode: 200,
		Status:     "OK",
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	resp.Body = io.NopCloser(strings.NewReader(message))
	resp.ContentLength = int64(len(message))
	err := resp.Write(connection)
	if err != nil {
		log.Printf("Error writing back to socket! Error: %v\n", err)
	}
}
