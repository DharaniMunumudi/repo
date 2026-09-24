package main

import (
	"fmt"
	"log"
	"mcp-secure-platform/gateway/auth"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "MCP SECURE PLATFORM GATEWAY")

	})
	http.Handle("/", auth.Middleware(handler))
	log.Println("Gateway running on port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
