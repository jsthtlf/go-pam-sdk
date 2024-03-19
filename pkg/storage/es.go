package storage

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch"
	"github.com/jsthtlf/go-pam-sdk/model"
)

type ESCommandStorage struct {
	Hosts   []string
	Index   string
	DocType string

	InsecureSkipVerify bool
}

func (es ESCommandStorage) BulkSave(commands []*model.Command) error {
	var buf bytes.Buffer
	transport := http.DefaultTransport.(*http.Transport).Clone()
	tlsClientConfig := &tls.Config{InsecureSkipVerify: es.InsecureSkipVerify}
	transport.TLSClientConfig = tlsClientConfig
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: es.Hosts,
		Transport: transport,
	})
	if err != nil {
		return err
	}
	for _, item := range commands {
		meta := []byte(fmt.Sprintf(`{ "index" : { } }%s`, "\n"))
		data, err := json.Marshal(item)
		if err != nil {
			return err
		}
		data = append(data, "\n"...)
		buf.Write(meta)
		buf.Write(data)
	}

	response, err := esClient.Bulk(bytes.NewReader(buf.Bytes()),
		esClient.Bulk.WithIndex(es.Index), esClient.Bulk.WithDocumentType(es.DocType))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var (
		blk *bulkResponse
		raw map[string]interface{}
	)
	if response.IsError() {
		if err = json.NewDecoder(response.Body).Decode(&raw); err != nil {
			return fmt.Errorf("es failure to parse error response body: %s", err)
		}

		return fmt.Errorf("es return [%d] %s: %s",
			response.StatusCode,
			raw["error"].(map[string]interface{})["type"],
			raw["error"].(map[string]interface{})["reason"],
		)
	}

	if err = json.NewDecoder(response.Body).Decode(&blk); err != nil {
		return fmt.Errorf("es failure to parse response body: %s", err)
	}

	var (
		numErrors  int
		numIndexed int
		errorsWrap error
	)

	errorsWrap = errors.New("save commands failed")

	for _, d := range blk.Items {
		if d.Index.Status > 201 {
			numErrors++
			errorsWrap = errors.Join(errorsWrap,
				fmt.Errorf("ES failure to save: [%d]: %s: %s: %s: %s",
					d.Index.Status,
					d.Index.Error.Type,
					d.Index.Error.Reason,
					d.Index.Error.Cause.Type,
					d.Index.Error.Cause.Reason,
				))
		} else {
			numIndexed++
		}
	}

	if numErrors > 0 {
		return errorsWrap
	}

	return nil
}

func (es ESCommandStorage) TypeName() string {
	return StorageTypeES
}

// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html#bulk-api-response-body
type bulkResponse struct {
	Errors bool `json:"errors"`
	Items  []struct {
		Index struct {
			ID     string `json:"_id"`
			Result string `json:"result"`
			Status int    `json:"status"`
			Error  struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
				Cause  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
				} `json:"caused_by"`
			} `json:"error"`
		} `json:"index"`
	} `json:"items"`
}
