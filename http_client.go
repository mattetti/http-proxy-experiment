package main

import (
	"http"
	"os"
	"log"
	"fmt"
)


type HttpClient struct {
	client   http.Client
}

// Forward a request to another domain
// By design Go doesn't let you define new methods on non-local types
func (this *HttpClient) Forward(req *http.Request, newDomain string) (resp *http.Response, err os.Error){

	log.Printf("req host: %v\n", req.Host)
	// TODO change the the domain back
	for _, cookie := range req.Cookie {
		log.Printf("req cookie: %v\n", cookie)
	}


	var base *http.URL
	/*var r *http.Response*/
	url := (newDomain + req.URL.RawPath)
	log.Printf("forwarding to: " + url)

	redirectChecker := this.client.CheckRedirect
	if redirectChecker == nil {
		redirectChecker = defaultCheckRedirect
	}
	var via []*http.Request
	
	for redirect := 0; ; redirect++ {
		var nReq = http.Request{}
		nReq.Method = req.Method
		//log.Printf("HTTP method: " + nReq.Method)
		nReq.Header = req.Header
		//log.Printf("Headers: %v\n", nReq.Header)
		
		if base == nil {
			nReq.URL, err = http.ParseURL(url)
		} else {
			nReq.URL, err = base.ParseURL(url)
		}

		if err != nil {
			break
		}

		// Set the redirection referer headers
		if len(via) > 0 {
			// Add the Referer header.
			lastReq := via[len(via)-1]
			if lastReq.URL.Scheme != "https" {
				nReq.Referer = lastReq.URL.String()
			}

			err = redirectChecker(&nReq, via)
			if err != nil {
				break
			}
		}

		// TODO support for content in the body if we aren't
		// dealing with GET/HEAD requests.

		url = nReq.URL.String()
		// Wrapped the client so I could do that ...sigh...
		// Also, #Do i a wrapper around #send
		if resp, err = this.client.Do(&nReq); err != nil {
			break
		}

		if shouldRedirect(resp.StatusCode) {
			resp.Body.Close()
			if url = resp.Header.Get("Location"); url == "" {
				err = os.ErrorString(fmt.Sprintf("%d response missing Location header", resp.StatusCode))
				break
			}
			base = req.URL
			via = append(via, &nReq)
			continue
		}
		// log.Printf("final URL: " + url)

		return
	}

	err = &http.URLError{req.Method, url, err}
	return
}


// copied over from the http package
func defaultCheckRedirect(req *http.Request, via []*http.Request) os.Error {
	if len(via) >= 10 {
		return os.ErrorString("stopped after 10 redirects")
	}
	return nil
}

// copied over from the http package
// True if the specified HTTP status code is one for which the Get utility should
// automatically redirect.
func shouldRedirect(statusCode int) bool {
	switch statusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		return true
	}
	return false
}
