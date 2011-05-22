package main

import (
  "http"
  "log"
)

// TODO: the uri removal from namespaced domains doesn't work yet
func main(){
  // Define the routes
  routes := []*Route{ &Route{ uri: "/", 
															endpoint: "http://google.com",
															domain: "google.com",
														  client: HttpClient{client: http.Client{}} }, 
                      &Route{ uri: "/yahoo",
															endpoint: "http://yahoo.com",
															domain: "yahoo.com",
															client: HttpClient{client: http.Client{}}} }
  
  // Set the route handlers
  for _, route := range routes {
		log.Printf("Handling base route: " + route.uri)
    http.Handle(route.uri, route)
  }

	// Start the server
	log.Printf("Starting server on http://127.0.0.1:10980")
  err := http.ListenAndServe(":10980", nil)
	if err != nil { log.Fatal(err) }
}
