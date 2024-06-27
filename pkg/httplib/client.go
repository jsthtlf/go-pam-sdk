package httplib

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const minTimeout = time.Second * 30

type AuthSign interface {
	Sign(req *http.Request) error
}

type Client struct {
	baseUrl  string
	cookies  map[string]string
	headers  map[string]string
	http     *http.Client
	authSign AuthSign
}

func NewClient(baseUrl string, timeout time.Duration) (*Client, error) {
	_, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if timeout < minTimeout {
		timeout = minTimeout
	}

	return &Client{
		baseUrl: baseUrl,
		cookies: make(map[string]string),
		headers: make(map[string]string),
		http: &http.Client{
			Timeout: timeout,
			Jar: &simpleCookieJar{
				data: map[string]string{},
			},
		},
	}, nil
}

func (c *Client) Clone() Client {
	return Client{
		baseUrl: c.baseUrl,
		cookies: make(map[string]string),
		headers: make(map[string]string),
		http: &http.Client{
			Timeout: c.http.Timeout,
			Jar: &simpleCookieJar{
				data: map[string]string{},
			},
		},
	}
}

func (c *Client) SetCookie(key string, value string) {
	c.cookies[key] = value
}

func (c *Client) SetHeader(key, value string) {
	c.headers[key] = value
}

func (c *Client) SetAuthSign(auth AuthSign) {
	c.authSign = auth
}

func (c *Client) setReqAuthHeader(r *http.Request) error {
	for k, v := range c.cookies {
		co := http.Cookie{Name: k, Value: v}
		r.AddCookie(&co)
	}

	if c.authSign != nil {
		return c.authSign.Sign(r)
	}

	return nil
}

func (c *Client) setReqHeaders(req *http.Request) error {
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.setReqAuthHeader(req)
}

func (c *Client) parseQueryUrl(reqUrl string, params []map[string]string) string {
	if len(params) < 1 {
		return reqUrl
	}
	query := url.Values{}
	for _, item := range params {
		for k, v := range item {
			query.Add(k, v)
		}
	}

	if strings.Contains(reqUrl, "?") {
		reqUrl += "&" + query.Encode()
	} else {
		reqUrl += "?" + query.Encode()
	}

	return reqUrl
}

func (c *Client) parseUrl(reqUrl string, params []map[string]string) string {
	reqUrl = c.parseQueryUrl(reqUrl, params)
	if c.baseUrl != "" {
		reqUrl = strings.TrimSuffix(c.baseUrl, "/") + reqUrl
	}

	return reqUrl
}

func (c *Client) newRequest(method, reqUrl string, data interface{}, params []map[string]string) (*http.Request, error) {
	reqUrl = c.parseUrl(reqUrl, params)
	dataRaw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(dataRaw)
	req, err := http.NewRequest(method, reqUrl, reader)
	if err != nil {
		return req, err
	}
	err = c.setReqHeaders(req)

	return req, err
}

func (c *Client) Do(method, reqUrl string, data, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	req, err := c.newRequest(method, reqUrl, data, params)
	if err != nil {
		return
	}
	resp, err = c.http.Do(req)
	if err != nil {
		return
	}

	return resp, c.handleResp(resp, res)
}

func (c *Client) Get(reqUrl string, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	return c.Do("GET", reqUrl, nil, res, params...)
}

func (c *Client) Post(reqUrl string, data interface{}, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	return c.Do("POST", reqUrl, data, res, params...)
}

func (c *Client) Delete(reqUrl string, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	return c.Do("DELETE", reqUrl, nil, res, params...)
}

func (c *Client) Put(reqUrl string, data interface{}, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	return c.Do("PUT", reqUrl, data, res, params...)
}

func (c *Client) Patch(reqUrl string, data interface{}, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	return c.Do("PATCH", reqUrl, data, res, params...)
}

func (c *Client) PostFileWithFields(reqUrl string, gFile string, fields map[string]string, res interface{}) error {
	fd, err := os.Open(gFile)
	if err != nil {
		return err
	}
	bufferFd := bufio.NewReader(fd)
	defer fd.Close()
	fi, err := fd.Stat()
	if err != nil {
		return err
	}
	var size = fi.Size()
	startPartBuf := bytes.NewBufferString("")
	partWriter := multipart.NewWriter(startPartBuf)
	for name, value := range fields {
		_ = partWriter.WriteField(name, value)
	}
	_, _ = partWriter.CreateFormFile("file", fi.Name())
	boundary := partWriter.Boundary()
	endString := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	endPartBuf := bytes.NewBufferString(endString)
	bodyReader := io.MultiReader(startPartBuf, bufferFd, endPartBuf)
	contentLen := int64(startPartBuf.Len()) + size + int64(endPartBuf.Len())
	reqUrl = c.parseUrl(reqUrl, nil)
	req, err := http.NewRequest(http.MethodPost, reqUrl, bodyReader)
	if err != nil {
		return err
	}
	if err = c.setReqHeaders(req); err != nil {
		return err
	}
	req.ContentLength = contentLen
	req.Header.Set("Content-Type", partWriter.FormDataContentType())

	client := http.Client{
		Jar: c.http.Jar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	return c.handleResp(resp, res)
}

var (
	CodeTerminalAlreadyExist = "terminal_already_exist"
	CodeObjectNotFound       = "object_does_not_exist"
	CodeAuthFailed           = "authentication_failed"
)

type ErrResponseType struct {
	Method  string
	UrlPath string

	Detail string `json:"detail"`
	Code   string `json:"code"`
}

func (e ErrResponseType) Error() string {
	return fmt.Sprintf("%s %s failed: %s (Code: %s)", e.Method, e.UrlPath, e.Detail, e.Code)
}

func (c *Client) handleResp(resp *http.Response, res interface{}) (err error) {
	defer resp.Body.Close()

	if buf, ok := res.(*bytes.Buffer); ok {
		_, err = buf.ReadFrom(resp.Body)
		return err
	}

	req := resp.Request

	if resp.StatusCode >= http.StatusBadRequest {
		exception := &ErrResponseType{
			Method:  req.Method,
			UrlPath: req.URL.Path,
		}
		decodeErr := json.NewDecoder(resp.Body).Decode(exception)
		if decodeErr == nil {
			return exception
		}

		return fmt.Errorf("%s %s failed: %s", req.Method, req.URL.Path, resp.Status)
	}

	if res != nil {
		switch {
		case strings.Contains(resp.Header.Get("Content-Type"), "application/json"):
			err = json.NewDecoder(resp.Body).Decode(res)
			if err != nil {
				return fmt.Errorf("%s %s failed, json unmarshal failed: %s", req.Method, req.URL.Path, err)
			}
		}
	}

	return nil
}
