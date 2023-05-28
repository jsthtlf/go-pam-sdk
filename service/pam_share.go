package service

import (
	"fmt"
	"github.com/jsthtlf/go-pam-sdk/model"
)

func (s *PAMService) CreateShareRoom(sessionId string, expired int, users []string) (res model.SharingSession, err error) {
	postData := make(map[string]interface{})
	postData["session"] = sessionId
	postData["expired_time"] = expired
	postData["users"] = users
	_, err = s.authClient.Post(ShareCreateURL, postData, &res)
	return
}

func (s *PAMService) GetShareUserInfo(query string) (res []*model.MiniUser, err error) {
	params := make(map[string]string)
	params["action"] = "suggestion"
	params["search"] = query
	_, err = s.authClient.Get(UserListURL, &res, params)
	return
}

func (s *PAMService) JoinShareRoom(data SharePostData) (res model.ShareRecord, err error) {
	_, err = s.authClient.Post(ShareSessionJoinURL, data, &res)
	return
}

func (s *PAMService) FinishShareRoom(recordId string) (err error) {
	reqUrl := fmt.Sprintf(ShareSessionFinishURL, recordId)
	_, err = s.authClient.Patch(reqUrl, nil, nil)
	return
}

type SharePostData struct {
	ShareId    string `json:"sharing"`
	Code       string `json:"verify_code"`
	UserId     string `json:"joiner"`
	RemoteAddr string `json:"remote_addr"`
}
