package test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"shrading/helper"
	"shrading/routes"
	"testing"
)

const UnregistSubjectMethodName = "UnregistSubject"

func TestUnregistSubjectActivity(t *testing.T) {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	routes.RegisterAPI(app)

	testTable := []struct {
		name string
		// The HTTP method to use in our call
		method string
		// The URL path that is being requested
		path string
		// The expected response status code
		statusCode int
		// The expected response body, as string
		body string
		// The request body to sent with the request
		requestBody map[string]interface{}
		// The name of the AppHandlerFake method that we want to spy on
		handlerMethodName string
		// The headers that are being set for the request
		requestHeaders map[string]string
		// The response headers we want to test on
		headers map[string]string
	}{
		{
			name:           `POST endpoint to unregist subject`, // test thieu truong
			method:         http.MethodPost,
			path:           `/unregist_subject`,
			statusCode:     4000,
			body:           `Invalid payload`,
			requestHeaders: map[string]string{`Content-Type`: `application/text`},
			headers:        map[string]string{`Content-Type`: `text/plain; charset=utf-8`},
		},
		{
			name:       `POST endpoint to unregist subject`, // test thieu truong
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4000,
			body:       `Invalid payload`,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN213",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			requestHeaders:    map[string]string{`Content-Type`: `application/text`},
			headers:           map[string]string{`Content-Type`: `text/plain; charset=utf-8`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test truong hop fail trong elk
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4001,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN213",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test voi hoc sinh xyz da dang ki mon nay r => bao loi
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4002,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test mon hoc khong ton tai
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4003,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "test_123",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test full slot mon hoc
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4004,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN405",
				"id_mon":   58,
				"ma_mon":   "BAS1152",
				"nhom_lop": "28",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test invalid slot so cho > sy so
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4005,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test lock record phai tat commnet trong ham handler
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4006,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test delete database rs
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4013,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test insert elasticsearch rs
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4014,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test update tkb database
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4009,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test rollback fail
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4010,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test rollback es fail
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4011,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to unregist subject`, //test commit fail
			method:     http.MethodPost,
			path:       `/unregist_subject`,
			statusCode: 4012,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: UnregistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			// Create and send request
			rbody, _ := json.Marshal(tc.requestBody)
			request := httptest.NewRequest(tc.method, tc.path, bytes.NewReader(rbody))
			request.Header.Add(`Content-Type`, `application/json`)

			// Request Headers
			for k, v := range tc.requestHeaders {
				request.Header.Add(k, v)
			}

			response, _ := app.Test(request)
			body, _ := ioutil.ReadAll(response.Body)
			var resp helper.Response
			json.Unmarshal(body, &resp)

			// Status Code
			statusCode := resp.Error.ErrorCode
			if statusCode != tc.statusCode {
				t.Errorf("StatusCode was incorrect, got: %d, want: %d.", statusCode, tc.statusCode)
			}

		})
	}
}
