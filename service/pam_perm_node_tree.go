package service

import (
	"fmt"
	"github.com/jsthtlf/go-pam-sdk/model"
)

func (s *PAMService) GetNodeTreeByUserAndNodeKey(userID, nodeKey string) (nodeTrees model.NodeTreeList, err error) {
	payload := map[string]string{}
	if nodeKey != "" {
		payload["key"] = nodeKey
	}
	apiURL := fmt.Sprintf(UserPermsNodeTreeWithAssetURL, userID)
	_, err = s.authClient.Get(apiURL, &nodeTrees, payload)
	return
}
