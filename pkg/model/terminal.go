package model

type TerminalConfig struct {
	AssetListPageSize   string        `json:"TERMINAL_ASSET_LIST_PAGE_SIZE"`
	AssetListSortBy     string        `json:"TERMINAL_ASSET_LIST_SORT_BY"`
	HeaderTitle         string        `json:"TERMINAL_HEADER_TITLE"`
	PasswordAuth        bool          `json:"TERMINAL_PASSWORD_AUTH"`
	PublicKeyAuth       bool          `json:"TERMINAL_PUBLIC_KEY_AUTH"`
	ReplayStorage       ReplayConfig  `json:"TERMINAL_REPLAY_STORAGE"`
	CommandStorage      CommandConfig `json:"TERMINAL_COMMAND_STORAGE"`
	SessionKeepDuration int           `json:"TERMINAL_SESSION_KEEP_DURATION"`
	TelnetRegex         string        `json:"TERMINAL_TELNET_REGEX"`
	MaxIdleTime         int           `json:"SECURITY_MAX_IDLE_TIME"`
	HeartbeatDuration   int           `json:"TERMINAL_HEARTBEAT_INTERVAL"`
	HostKey             string        `json:"TERMINAL_HOST_KEY"`
	EnableSessionShare  bool          `json:"SECURITY_SESSION_SHARE"`
	EnableDbWeb         bool          `json:"TERMINAL_SSH_DB_ENABLED"`
	EnableRdpWeb        bool          `json:"TERMINAL_RDP_ENABLED"`
	EnableRdpNative     bool          `json:"TERMINAL_RDP_NATIVE_ENABLED"`
	EnableDbNative      bool          `json:"TERMINAL_DB_NATIVE_ENABLED"`
}

type Terminal struct {
	Name           string `json:"name"`
	Comment        string `json:"comment"`
	ServiceAccount struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		AccessKey AccessKey `json:"access_key"`
	} `json:"service_account"`
}

type HeartbeatResponse struct {
	NextHeartbeat int `json:"next_heartbeat"`
}

type TerminalTasks struct {
	Tasks       []TerminalTask `json:"tasks"`
	NextRequest int            `json:"next_request"`
}

type TerminalTask struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Args       string     `json:"args"`
	Kwargs     TaskKwargs `json:"kwargs"`
	IsFinished bool
}

const (
	TaskKillSession = "kill_session"
)

type TaskKwargs struct {
	TerminatedBy string `json:"terminated_by"`
}

type ReplayConfig struct {
	TypeName string `json:"TYPE"`

	/*
		obs oss
	*/
	Endpoint  string `json:"ENDPOINT,omitempty"`
	Bucket    string `json:"BUCKET,omitempty"`
	AccessKey string `json:"ACCESS_KEY,omitempty"`
	SecretKey string `json:"SECRET_KEY,omitempty"`

	/*
		s3、 swift cos
	*/

	Region string `json:"REGION,omitempty"`

	/*
		azure account
	*/
	AccountName    string `json:"ACCOUNT_NAME,omitempty"`
	AccountKey     string `json:"ACCOUNT_KEY,omitempty"`
	EndpointSuffix string `json:"ENDPOINT_SUFFIX,omitempty"`
	ContainerName  string `json:"CONTAINER_NAME,omitempty"`
}

type CommandConfig struct {
	TypeName string `json:"TYPE"`

	/*
		elasticsearch
	*/
	Hosts       []string `json:"HOSTS,omitempty"`
	IndexByDate bool     `json:"INDEX_BY_DATE,omitempty"`
	Index       string   `json:"INDEX,omitempty"`
	DocType     string   `json:"DOC_TYPE,omitempty"`
	Other       struct {
		IgnoreVerifyCerts bool `json:"IGNORE_VERIFY_CERTS,omitempty"`
	} `json:"OTHER"`
}
