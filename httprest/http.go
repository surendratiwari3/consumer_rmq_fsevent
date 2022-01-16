package httprest

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type RestInterface interface {
	Get(data map[string]interface{}, urls string) (int, []byte, error)
	Post(data map[string]interface{}, urls string) (int, []byte, error)
	GetOrCreatConnection() *resty.Client
}

type RestHandler struct {
	Conn *resty.Client
}

func NewHttpRestHandler() RestInterface {
	return &RestHandler{}
}

func (h *RestHandler) GetOrCreatConnection() *resty.Client {
	if h.Conn != nil {
		return h.Conn
	}
	reusableConnection := resty.New()
	transport := createTransport(nil)
	reusableConnection.SetTransport(transport)
	return reusableConnection
}

func createTransport(localAddr net.Addr) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	if localAddr != nil {
		dialer.LocalAddr = localAddr
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   20,
		MaxConnsPerHost:       100,
	}
}

func (h *RestHandler) Get(data map[string]interface{}, urls string) (int, []byte, error) {
	var restClient = h.GetOrCreatConnection()

	u, err := url.Parse(urls)
	if err != nil {
		log.Println("error while parsing url - returning 400 ", err)
		return 400, nil, err
	}
	if u.User != nil {
		if pwd, ok := u.User.Password(); ok {
			restClient.SetBasicAuth(u.User.Username(), pwd)
			baseUrlStr := fmt.Sprint(u.Scheme, "://", u.Host, u.Path)
			base, err := url.Parse(baseUrlStr)
			if err != nil {
				log.Println("error while parsing auth scheme of url - returning 400 ", err)
				return 400, nil, err
			}
			urls = fmt.Sprint(base.ResolveReference(u))
		}
	}

	if data == nil {
		resp, err := restClient.R().
			EnableTrace().
			SetHeader("Accept", "application/json").
			Get(urls)
		if resp == nil || err != nil {
			log.Println("response is nil or error while query to url ", err)
			return 400, nil, err
		}
		return resp.StatusCode(), resp.Body(), err
	}
	reqData := make(map[string]string)
	for key, value := range data {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		reqData[strKey] = strValue
	}
	resp, err := restClient.R().
		SetQueryParams(reqData).
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get(urls)

	// Close the connection to reuse it
	if resp == nil {
		log.Println("response is nil or error while query to url ", err)
		return 400, nil, err
	}
	return resp.StatusCode(), resp.Body(), err
}

func (h *RestHandler) Post(data map[string]interface{}, urls string) (int, []byte, error) {
	u, err := url.Parse(urls)
	if err != nil {
		log.Println("error while parsing url - returning 400 ", err)
		return 400, nil, err
	}

	var restClient = h.GetOrCreatConnection()

	if u.User != nil {
		if pwd, ok := u.User.Password(); ok {
			restClient.SetBasicAuth(u.User.Username(), pwd)
			baseUrlStr := fmt.Sprint(u.Scheme, "://", u.Host, u.Path)
			base, err := url.Parse(baseUrlStr)
			if err != nil {
				log.Println("error while parsing auth scheme of url - returning 400 ", err)
				return 400, nil, err
			}
			urls = fmt.Sprint(base.ResolveReference(u))
		}
	}

	if data == nil {
		resp, err := restClient.R().
			EnableTrace().
			SetHeader("Accept", "application/json").
			Post(urls)
		if resp == nil {
			return 400, nil, err
		}
		return resp.StatusCode(), resp.Body(), err
	}
	resp, err := restClient.R().
		EnableTrace().
		SetBody(data).
		SetHeader("Accept", "application/json").
		Post(urls)
	if resp == nil {
		return 400, nil, err
	}
	return resp.StatusCode(), resp.Body(), err
}
