package service

import "github.com/jsthtlf/go-pam-sdk/model"

func (s *PAMService) GetPublicSetting() (result model.PublicSetting, err error) {
	_, err = s.authClient.Get(PublicSettingURL, &result)
	return
}
