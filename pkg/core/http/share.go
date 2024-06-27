package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) CreateShareRoom(sessionId string, expired int, users []string) (res model.SharingSession, err error) {
	postData := make(map[string]interface{})
	postData["session"] = sessionId
	postData["expired_time"] = expired
	postData["users"] = users
	_, err = p.post(UrlShareCreate, postData, &res)
	return
}

func (p *httpProvider) JoinShareRoom(data model.SharePostData) (res model.ShareRecord, err error) {
	_, err = p.post(UrlShareSessionJoin, data, &res)
	return
}

func (p *httpProvider) FinishShareRoom(recordId string) (err error) {
	reqUrl := fmt.Sprintf(UrlShareSessionFinish, recordId)
	_, err = p.patch(reqUrl, nil, nil)
	return
}
