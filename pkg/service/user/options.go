package user

type OtpOptions struct {
	Username   string
	Password   string
	PublicKey  string
	RemoteAddr string
	LoginType  string
}

type Option func(*OtpOptions)

func WithUsername(username string) Option {
	return func(args *OtpOptions) {
		args.Username = username
	}
}

func WithPassword(password string) Option {
	return func(args *OtpOptions) {
		args.Password = password
	}
}

func WithPublicKey(publicKey string) Option {
	return func(args *OtpOptions) {
		args.PublicKey = publicKey
	}
}

func WithRemoteAddr(remoteAddr string) Option {
	return func(args *OtpOptions) {
		args.RemoteAddr = remoteAddr
	}
}

func WithLoginType(loginType string) Option {
	return func(args *OtpOptions) {
		args.LoginType = loginType
	}
}
