package service

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) GetUserNodeAssets(userID, nodeID string,
	params model.PaginationParam) (resp model.PaginationResponse, err error) {
	Url := fmt.Sprintf(UserPermsNodeAssetsListURL, userID, nodeID)
	return s.getPaginationResult(Url, params)
}

func (s *PAMService) GetUserNodes(userId string) (nodes model.NodeList, err error) {
	Url := fmt.Sprintf(UserPermsNodesListURL, userId)
	_, err = s.authClient.Get(Url, &nodes)
	return
}

func (s *PAMService) RefreshUserNodes(userId string) (nodes model.NodeList, err error) {
	params := map[string]string{
		"rebuild_tree": "1",
	}
	Url := fmt.Sprintf(UserPermsNodesListURL, userId)
	_, err = s.authClient.Get(Url, &nodes, params)
	return
}
