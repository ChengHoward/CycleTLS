package cycletls

import (
	"bytes"
	"encoding/json"
	http "github.com/Wuhan-Dongce/fhttp"
	"github.com/Wuhan-Dongce/fhttp/http2"
	"io"
	"log"
	"net/url"
	"strings"
)

// Options sets CycleTLS client options
type Options struct {
	URL              string               `json:"url"`
	Method           string               `json:"method"`
	Headers          map[string]string    `json:"headers"`
	Body             string               `json:"body"`
	Ja3              string               `json:"ja3"`
	TLSExtensions    *TLSExtensions       `json:"-"`
	HTTP2Settings    *http2.HTTP2Settings `json:"-"`
	PHeaderOrderKeys []string             `json:"-"`
	HeaderOrderKeys  []string             `json:"-"`
	UserAgent        string               `json:"userAgent"`
	Proxy            string               `json:"proxy"`
	Cookies          []Cookie             `json:"cookies"`
	Timeout          int                  `json:"timeout"`
	DisableRedirect  bool                 `json:"disableRedirect"`
	HeaderOrder      []string             `json:"headerOrder"`
}

// rename to request+client+options
type fullRequest struct {
	req     *http.Request
	client  http.Client
	options Options
}

// Response contains Cycletls response data
type Response struct {
	Status  int
	Body    io.ReadCloser
	Headers map[string]string
	Client  http.Client
}

// JSONBody converts response body to json
func (re Response) JSONBody() map[string]interface{} {
	var data map[string]interface{}
	body, err := io.ReadAll(re.Body)
	if err != nil {
		log.Print("ReadAll failed " + err.Error())
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Print("Json Conversion failed " + err.Error())
	}
	return data
}

// CycleTLS creates full request and response
type CycleTLS struct {
}

// ready Request
func processRequest(options *Options) (result *fullRequest) {
	var browser = browser{
		JA3:           options.Ja3,
		UserAgent:     options.UserAgent,
		Cookies:       options.Cookies,
		HTTP2Settings: options.HTTP2Settings,
	}

	/*if options.Ja3 != "" && !strings.HasPrefix(options.URL, "https") {
		browser.JA3 = ""
	}*/

	client, err := newClient(
		browser,
		options.Timeout,
		options.DisableRedirect,
		options.UserAgent,
		options.Proxy,
	)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(strings.ToUpper(options.Method), options.URL, strings.NewReader(options.Body))
	if err != nil {
		log.Fatal(err)
	}

	//ordering the pseudo headers and our normal headers
	if options.PHeaderOrderKeys == nil {
		options.PHeaderOrderKeys = []string{":method", ":authority", ":scheme", ":path"}
	}

	req.Header = http.Header{
		http.HeaderOrderKey:  options.HeaderOrderKeys,
		http.PHeaderOrderKey: options.PHeaderOrderKeys,
	}
	//set our Host header
	u, err := url.Parse(options.URL)
	if err != nil {
		panic(err)
	}

	//append our normal headers
	for k, v := range options.Headers {
		if k != "Content-Length" {
			req.Header.Set(k, v)
		}
	}
	if req.Header.Get("Host") == "" {
		req.Header.Set("Host", u.Host)
	}
	req.Header.Set("user-agent", options.UserAgent)
	return &fullRequest{req: req, client: client, options: *options}

}

func dispatcher(res *fullRequest) (response Response, err error) {
	//defer res.client.CloseIdleConnections()

	resp, err := res.client.Do(res.req)
	if err != nil {
		parsedError := parseError(err)

		headers := make(map[string]string)
		// parsedError.ErrorMsg + "-> \n" + string(err.Error())
		return Response{
			parsedError.StatusCode, io.NopCloser(bytes.NewBufferString(parsedError.ErrorMsg)), headers, res.client,
		}, err

	}

	headers := make(map[string]string)

	for name, values := range resp.Header {
		if name == "Set-Cookie" {
			headers[name] = strings.Join(values, "/,/")
		} else {
			headers[name] = values[0]
		}
	}
	return Response{resp.StatusCode, resp.Body, headers, res.client}, nil

}

// Do creates a single request
func (client CycleTLS) Do(URL string, options Options, Method string) (response Response, err error) {

	options.URL = URL
	options.Method = Method

	response, err = dispatcher(processRequest(&options))
	if err != nil {
		log.Print("Request Failed: " + err.Error())
		return response, err
	}

	return response, nil
}

//TODO rename this

// Init starts the worker pool or returns a empty cycletls struct
func Init() CycleTLS {
	return CycleTLS{}
}
