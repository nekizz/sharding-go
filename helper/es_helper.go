package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"shrading/connection"
	"strings"
)

func QueryES(index string, query map[string]interface{}) (ESResultRegistSubject, int, error) {
	var rs ESResultRegistSubject
	//var listSourceUser []ESSourceUser
	var totalRecord int

	connectES, errES := connection.ConnectElastic()
	if errES != nil {
		fmt.Println("connect ES error")
		return ESResultRegistSubject{}, 0, errES
	}
	bodySendEs, _ := json.Marshal(query)
	res, err := connectES.Search(
		connectES.Search.WithContext(context.Background()),
		connectES.Search.WithIndex(index),
		connectES.Search.WithBody(strings.NewReader(string(bodySendEs))),
		connectES.Search.WithTrackTotalHits(true),
		connectES.Search.WithPretty())

	if err != nil {
		fmt.Println("get user from ES error", err)
		return ESResultRegistSubject{}, 0, err
	}

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&rs); err != nil {
		fmt.Println("decode from ES error", err)
		return ESResultRegistSubject{}, 0, err
	}

	totalRecord = rs.Hits.Total.Value

	return rs, totalRecord, nil
}

func InsertToElastic(docStruct interface{}, index string, documentID string, documentType string) (interface{}, error) {
	doc, err := json.Marshal(docStruct)

	if err != nil {
		return nil, err
	}

	if connection.ESLivechatConnection == nil {
		return nil, errors.New("connect ESLivechatConnection failed")
	}

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(string(doc)),
		Refresh: "true",
	}

	if len(documentID) > 0 {
		req.DocumentID = documentID
	}

	if len(documentType) > 0 {
		req.DocumentType = documentType
	}

	res, err := req.Do(context.Background(), connection.ESLivechatConnection)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			return nil, errors.New(e["error"].(map[string]interface{})["reason"].(string))
		}
	}
	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	} else {
		return r["result"], nil
	}
}

func DeleteByQueryES(index string, query map[string]interface{}) (map[string]interface{}, error) {
	var mapResp map[string]interface{}
	connectES := connection.ConnectLivechatElastic()
	bodySendEs, _ := json.Marshal(query)
	res, err := connectES.DeleteByQuery([]string{index}, bytes.NewReader(bodySendEs))

	if err != nil {
		fmt.Println("delete by query ES error", err)
		return mapResp, err
	} else {
		defer res.Body.Close()
		mapResp := make(map[string]interface{})
		if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
			return mapResp, nil
		}
	}
	return mapResp, nil
}
