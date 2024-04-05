package user

import "github.com/jsthtlf/go-pam-sdk/pkg/httplib"

type options struct {
	Username   string
	Password   string
	PublicKey  string
	RemoteAddr string
	LoginType  string
	client     *httplib.Client
}

type Option func(*options)

func WithUsername(username string) Option {
	return func(args *options) {
		args.Username = username
	}
}

func WithPasssword(password string) Option {
	return func(args *options) {
		args.Password = password
	}
}

func WithPublicKey(publicKey string) Option {
	return func(args *options) {
		args.PublicKey = publicKey
	}
}

func WithRemoteAddr(remoteAddr string) Option {
	return func(args *options) {
		args.RemoteAddr = remoteAddr
	}
}

func WithLoginType(loginType string) Option {
	return func(args *options) {
		args.LoginType = loginType
	}
}

func WithHttpClient(con *httplib.Client) Option {
	return func(args *options) {
		args.client = con
	}
}
