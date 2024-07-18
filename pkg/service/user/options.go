package user

type otpOptions struct {
	Username   string
	Password   string
	PublicKey  string
	RemoteAddr string
	LoginType  string
}

type Option func(*otpOptions)

func WithUsername(username string) Option {
	return func(args *otpOptions) {
		args.Username = username
	}
}

func WithPassword(password string) Option {
	return func(args *otpOptions) {
		args.Password = password
	}
}

func WithPublicKey(publicKey string) Option {
	return func(args *otpOptions) {
		args.PublicKey = publicKey
	}
}

func WithRemoteAddr(remoteAddr string) Option {
	return func(args *otpOptions) {
		args.RemoteAddr = remoteAddr
	}
}

func WithLoginType(loginType string) Option {
	return func(args *otpOptions) {
		args.LoginType = loginType
	}
}
