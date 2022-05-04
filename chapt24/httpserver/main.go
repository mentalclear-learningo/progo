package main

import (
	"io"
	"net/http"
	"strings"
)

type StringHandler struct {
	message string
}

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
// 	request *http.Request) {
// 	// printReqest(request)
// 	//workAroundFavicon(writer, request)
// 	Printfln("Request for %v", request.URL.Path)
// 	switch request.URL.Path {
// 	case "/favicon.ico":
// 		http.NotFound(writer, request)
// 	case "/message":
// 		io.WriteString(writer, sh.message)
// 	default:
// 		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
// 	}
// 	io.WriteString(writer, sh.message)
// }
func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func HTTPSRedirect(writer http.ResponseWriter,
	request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer",
			"certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

// func main() {
// 	//err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World\n"})
// 	http.Handle("/message", StringHandler{"Hello, World"})
// 	http.Handle("/favicon.ico", http.NotFoundHandler())
// 	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
// 	err := http.ListenAndServe(":5000", nil)
// 	if err != nil {
// 		Printfln("Error: %v", err.Error())
// 	}
// }

func workAroundFavicon(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/favicon.ico" {
		Printfln("Request for icon detected - returning 404")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	Printfln("Request for %v", request.URL.Path)
}

func printReqest(req *http.Request) {
	Printfln("Method: %v", req.Method)
	Printfln("URL: %v", req.URL)
	Printfln("HTTP Version: %v", req.Proto)
	Printfln("Host: %v", req.Host)
	for name, val := range req.Header {
		Printfln("Header: %v, Value: %v", name, val)
	}
	Printfln("---")
}
