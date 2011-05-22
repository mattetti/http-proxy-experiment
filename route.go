package main

import(
	"log"
	"http"
	"io"
	"regexp"
	//"os"
)

type Route struct {
  uri string
  endpoint string
	domain string
	client HttpClient
}


// Implementing the Handler interface
func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request){
    // http.Redirect(w, req, r.endpoint, 302)

    // Proxy the request
     // should use a connection pool and reuse connections
			//if err != nil { log.Fatal(err) }

    pResponse, err := r.client.Forward(req, r.endpoint)
		/*log.Printf("forwarded response: %v\n", pResponse)*/
		if err != nil { log.Fatal(err) }
		if pResponse == nil{
			log.Printf("Empty response")
		} else {
			//pResponse, err := client.Do(pReq)
			defer pResponse.Body.Close()
			if err != nil { log.Fatal(err) }

			// Convert the cookies
			// TODO: convert to the domain saw by the client
			for _, cookie := range pResponse.SetCookie {
				rexp, err := regexp.Compile("domain=." + r.domain)
				if err != nil { log.Fatal(err) }
				newCookie :=  rexp.ReplaceAllString(cookie.Raw, "domain=127.0.0.1")
				w.Header().Add("Set-Cookie", newCookie)
				log.Printf("Modified Cookie: " + newCookie)
			}

			log.Printf("Proxied url: " + r.endpoint + req.URL.RawPath + "\n\n")
			

			// push back to the client
			cw := &ChunkedWriter{w}
			_, err = io.Copy(cw, pResponse.Body)
			if err == nil { err = cw.Close() }
		}
}


func (r *Route) newDispatcher(){
  log.Printf("uri: " + r.uri + " endpoint: " + r.endpoint)
  http.Handle(r.uri, r)
}


/*
// Could have also been implemented like that.
// Where we dynamically create a dynamic handler and use that.

// in main
for _, value := range routes {
    value.newDispatcher()
}

// then here

func (r *Route) newDispatcher(){
  log.Printf("uri: " + r.uri + " endpoint: " + r.endpoint)
  http.Handle(r.uri, r)
}

func (r *Route) newDispatcher(){
  log.Printf("uri: " + r.uri + " endpoint: " + r.endpoint)
  
  dynHandler := func(w http.ResponseWriter, req *http.Request){
    // Reimplementing http://golang.org/pkg/http/#ReverseProxy
    // http.Redirect(w, req, r.endpoint, 302)

    // Proxy the request
    client := http.Client{}
    pResponse, url, err := client.Get(r.endpoint + req.URL.RawPath)
    if err != nil { log.Fatal(err) }

    // copy the headers
    for key, value := range pResponse.Header {
      log.Printf(key + " : " + value[0])
      if key == "Cookie" {
        w.Header().Set("Set-Cookie", value[0])
      } else {
        w.Header().Set(key, value[0])
      }
    }

    log.Printf("Proxied url: " + url)
    defer pResponse.Body.Close()

    cw := &CustomChunkedWriter{w}
    _, err = io.Copy(cw, pResponse.Body)
    if err == nil { err = cw.Close() }
  }
  
  http.HandleFunc(r.uri, dynHandler)
}
*/


