package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)
	sendGetReq()
	fmt.Println("\nSend Post:")
	sendPostReq()
	fmt.Println("\n\nPost not encoded:")
	sendPostReqNotEncoded()
	fmt.Println("\nCreate simple POST:")
	createSimplePostReq()
	fmt.Println("\nUsing conventional funcs:")
	usingConvenFuncs()
	fmt.Println("\nUsing cookies:")
	usingCookies()
	fmt.Println("\nUse separate clients and cookies:")
	separateClinets()
	fmt.Println("\nUsing redirects:")
	usingRedirects()
	fmt.Println("\nCreate multipart forms:")
	createMultipartForms()
}

func createMultipartForms() {
	var buffer bytes.Buffer
	formWriter := multipart.NewWriter(&buffer)
	fieldWriter, err := formWriter.CreateFormField("name")
	if err == nil {
		io.WriteString(fieldWriter, "Alice")
	}
	fieldWriter, err = formWriter.CreateFormField("city")
	if err == nil {
		io.WriteString(fieldWriter, "New York")
	}
	fileWriter, err := formWriter.CreateFormFile("codeFile", "printer.go")
	if err == nil {
		fileData, err := os.ReadFile("./printer.go")
		if err == nil {
			fileWriter.Write(fileData)
		}
	}
	formWriter.Close()
	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:5000/form", &buffer)
	req.Header["Content-Type"] = []string{formWriter.FormDataContentType()}
	if err == nil {
		var response *http.Response
		response, err = http.DefaultClient.Do(req)
		if err == nil {
			io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func usingRedirects() {
	http.DefaultClient.CheckRedirect = func(req *http.Request,
		previous []*http.Request) error {
		if len(previous) == 3 {
			url, _ := url.Parse("http://localhost:5000/html")
			req.URL = url
		}
		return nil
	}
	req, err := http.NewRequest(http.MethodGet,
		"http://localhost:5000/redirect1", nil)
	if err == nil {
		var response *http.Response
		response, err = http.DefaultClient.Do(req)
		if err == nil {
			io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func separateClinets() {
	// If multiple Client values are required but cookies should be shared, then a single
	jar, err := cookiejar.New(nil)
	clients := make([]http.Client, 3)
	for index, client := range clients {
		//jar, err := cookiejar.New(nil)
		if err == nil {
			client.Jar = jar
		}
		for i := 0; i < 3; i++ {
			req, err := http.NewRequest(http.MethodGet,
				"http://localhost:5000/cookie", nil)
			if err == nil {
				response, err := client.Do(req)
				if err == nil && response.StatusCode == http.StatusOK {
					fmt.Fprintf(os.Stdout, "Client %v: ", index)
					io.Copy(os.Stdout, response.Body)
					defer response.Body.Close()
				} else {
					Printfln("Request Error: %v", err.Error())
				}
			} else {
				Printfln("Request Init Error: %v", err.Error())
			}
		}
	}
}

func usingCookies() {
	jar, err := cookiejar.New(nil)
	if err == nil {
		http.DefaultClient.Jar = jar
	}
	for i := 0; i < 3; i++ {
		req, err := http.NewRequest(http.MethodGet,
			"http://localhost:5000/cookie", nil)
		if err == nil {
			response, err := http.DefaultClient.Do(req)
			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Request Init Error: %v", err.Error())
		}
	}
}

func usingConvenFuncs() {
	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/echo",
			io.NopCloser(strings.NewReader(builder.String())))
		if err == nil {
			req.Header["Content-Type"] = []string{"application/json"}
			response, err := http.DefaultClient.Do(req)
			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Request Init Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}
}

func createSimplePostReq() {
	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		reqURL, err := url.Parse("http://localhost:5000/echo")
		if err == nil {
			req := http.Request{
				Method: http.MethodPost,
				URL:    reqURL,
				Header: map[string][]string{
					"Content-Type": {"application.json"},
				},
				Body: io.NopCloser(strings.NewReader(builder.String())),
			}
			response, err := http.DefaultClient.Do(&req)
			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Parse Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}
}

func sendPostReqNotEncoded() {
	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		response, err := http.Post("http://localhost:5000/echo",
			"application/json",
			strings.NewReader(builder.String()))
		if err == nil && response.StatusCode == http.StatusOK {
			io.Copy(os.Stdout, response.Body)
			defer response.Body.Close()
		} else {
			Printfln("Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func sendPostReq() {
	formData := map[string][]string{
		"name":     {"Kayak "},
		"category": {"Watersports"},
		"price":    {"279"},
	}
	response, err := http.PostForm("http://localhost:5000/echo", formData)
	if err == nil && response.StatusCode == http.StatusOK {
		io.Copy(os.Stdout, response.Body)
		defer response.Body.Close()
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}
}

func sendGetReq() {
	response, err := http.Get("http://localhost:5000/json")
	if err == nil && response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		data := []Product{}
		err = json.NewDecoder(response.Body).Decode(&data)
		if err == nil {
			for _, p := range data {
				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
			}
		} else {
			Printfln("Decode error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}
}
