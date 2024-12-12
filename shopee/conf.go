package shopee

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"
	"time"
)

type config struct {
	proxyURL   string
	timeout    int
	sslVerify  bool
	pemCerts   []byte
	httpClient *http.Client
}

func WithProxyUrl(proxyURL string) configurer {
	return func(conf *config) {
		conf.proxyURL = proxyURL
	}
}

func WithTimeout(timeout int) configurer {
	return func(conf *config) {
		conf.timeout = timeout
	}
}

func WithSslVerify(sslVerify bool) configurer {
	return WithSslVerifyAndPemCerts(sslVerify, nil)
}

func WithSslVerifyAndPemCerts(sslVerify bool, pemCerts []byte) configurer {
	return func(conf *config) {
		conf.sslVerify = sslVerify
		conf.pemCerts = pemCerts
	}
}

func WithHttpClient(httpClient *http.Client) configurer {
	return func(conf *config) {
		conf.httpClient = httpClient
	}
}

type configurer func(conf *config)

func (conf *config) initConfig() error {
	conf.prepareConfig()

	transport := &http.Transport{
		ResponseHeaderTimeout: time.Second * time.Duration(conf.timeout),
		TLSHandshakeTimeout:   time.Second * time.Duration(conf.timeout),
		ExpectContinueTimeout: time.Second * time.Duration(conf.timeout),
	}

	if conf.proxyURL != "" {
		proxyUrl, err := url.Parse(conf.proxyURL)
		if err != nil {
			return err
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: !conf.sslVerify,
		// MinVersion:         tls.VersionTLS12,
		// MaxVersion:         tls.VersionTLS12,
	}

	if conf.sslVerify && conf.pemCerts != nil {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(conf.pemCerts)
		tlsConfig.RootCAs = pool
	}
	transport.TLSClientConfig = tlsConfig

	if conf.httpClient == nil {
		httpClient := &http.Client{
			Transport: transport,
		}
		conf.httpClient = httpClient
	}

	return nil
}

func (conf *config) prepareConfig() {
	if conf.timeout <= 0 {
		conf.timeout = DefaultTimeout
	}
}
