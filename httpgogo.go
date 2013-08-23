package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// Open a channel for signal processing
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for sig := range c {
			fmt.Println("Signal received:", sig)
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}()

	// Get the command-line arguments
	argPort := flag.Int("port", 8080, "The port to run httpgogo on. Defaults to 8080.")
	argPath := flag.String("path", "./", "The path to serve. Defaults to the current directory.")
	flag.Parse()

	// Give the user some kind of feedback
	fmt.Println(fmt.Sprintf("Starting static file server at %s on port %v", *argPath, *argPort))

	// Start the server on argPort, using FileServer at argPath as the handler
	err := http.ListenAndServe(fmt.Sprintf(":%v", *argPort), http.FileServer(http.Dir(*argPath)))
	if err != nil {
		fmt.Println("Error running web server for static assets:", err)
	}
}