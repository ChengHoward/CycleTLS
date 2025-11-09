package main

import (
	"fmt"
	"github.com/ChengHoward/CycleTLS/cycletls"
	"github.com/ChengHoward/CycleTLS/cycletls/imitate"
	"io"
)

func main() {
	URL := "https://tls.browserleaks.com/json"
	client := cycletls.Init()

	options := cycletls.Options{
		Timeout: 120,
		Headers: map[string]string{
			"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		},
		//Proxy: "http://127.0.0.1:1080",
	}
	imitate.Firefox(&options)
	for s, s2 := range options.Headers {
		fmt.Println(s, ": ", s2)
	}
	resp, err := client.Do(URL, options, "GET")
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	content := cycletls.DecompressBody(body, []string{resp.Headers["Content-Encoding"]}, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp.Status)
	for s, s2 := range resp.Headers {
		fmt.Println(s, ": ", s2)
	}
	fmt.Println(string(content))
}
