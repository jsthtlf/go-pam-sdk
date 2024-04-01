package httplib

import (
	"net/http"
	"net/url"
	"sync"
)

type simpleCookieJar struct {
	mu   sync.Mutex
	data map[string]string
}

func (c *simpleCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := range cookies {
		name := cookies[i].Name
		value := cookies[i].Value
		c.data[name] = value
	}
}

func (c *simpleCookieJar) Cookies(u *url.URL) []*http.Cookie {
	c.mu.Lock()
	defer c.mu.Unlock()
	cookies := make([]*http.Cookie, 0, len(c.data))
	for k, v := range c.data {
		cookie := http.Cookie{Value: v, Name: k}
		cookies = append(cookies, &cookie)
	}
	return cookies
}
