package model

type PublicSetting struct {
	LoginTitle string `json:"LOGIN_TITLE"`
	LogoURLS   struct {
		LogOut  string `json:"logo_logout"`
		Index   string `json:"logo_index"`
		Image   string `json:"login_image"`
		Favicon string `json:"favicon"`
	} `json:"LOGO_URLS"`
	EnableWatermark    bool `json:"SECURITY_WATERMARK_ENABLED"`
	EnableSessionShare bool `json:"SECURITY_SESSION_SHARE"`
}
