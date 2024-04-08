package http

import (
	"fmt"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"github.com/jsthtlf/go-pam-sdk/pkg/utils"
)

func (p *httpProvider) CreateSession(sess model.Session) error {
	_, err := p.authClient.Post(UrlSessionList, sess, nil)
	return err
}

func (p *httpProvider) sessionPatch(sid string, data interface{}) error {
	Url := fmt.Sprintf(UrlSessionDetail, sid)
	_, err := p.authClient.Patch(Url, data, nil)
	return err
}

func (p *httpProvider) SessionSuccess(sid string) error {
	data := map[string]bool{
		"is_success": true,
	}
	return p.sessionPatch(sid, data)
}

func (p *httpProvider) SessionFailed(sid string, err error) error {
	data := map[string]bool{
		"is_success": false,
	}
	return p.sessionPatch(sid, data)
}

func (p *httpProvider) SessionDisconnect(sid string) error {
	return p.SessionFinished(sid, time.Now())
}

func (p *httpProvider) SessionFinished(sid string, t time.Time) error {
	data := map[string]interface{}{
		"is_finished": true,
		"date_end":    utils.NewUTCTime(t),
	}
	return p.sessionPatch(sid, data)
}

func (p *httpProvider) GetSessionById(sid string) (data model.Session, err error) {
	reqURL := fmt.Sprintf(UrlSessionDetail, sid)
	_, err = p.authClient.Get(reqURL, &data)
	return
}

func (p *httpProvider) CreateSessionTicketRelation(sid, ticketId string) (err error) {
	data := map[string]string{
		"session": sid,
		"ticket":  ticketId,
	}
	_, err = p.authClient.Post(UrlSessionTicketRelation, data, nil)
	return
}

func (p *httpProvider) UploadReplay(sid, gZipFile string) error {
	version := model.ParseReplayVersion(gZipFile, model.Version3) // TODO перенести
	var res map[string]interface{}
	Url := fmt.Sprintf(UrlSessionReplay, sid)
	fields := make(map[string]string)
	fields["version"] = string(version)
	return p.authClient.PostFileWithFields(Url, gZipFile, fields, &res)
}

func (p *httpProvider) FinishReply(sid string) error {
	data := map[string]bool{"has_replay": true}
	return p.sessionPatch(sid, data)
}
