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

const RegistSubjectMethodName = "RegistSubject"
const unregistSubjectMethodName = "UnregistSubject"

func TestRegistSubjectActivity(t *testing.T) {

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
			name:           `POST endpoint to regist subject`, // test thieu truong
			method:         http.MethodPost,
			path:           `/regist_subject`,
			statusCode:     4000,
			body:           `Invalid payload`,
			requestHeaders: map[string]string{`Content-Type`: `application/text`},
			headers:        map[string]string{`Content-Type`: `text/plain; charset=utf-8`},
		},
		{
			name:       `POST endpoint to regist subject`, //test truong hop fail trong elk
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4001,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN213",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test voi hoc sinh xyz da dang ki mon nay r => bao loi
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4002,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test mon hoc khong ton tai
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4003,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "test_123",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test full slot mon hoc
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4004,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN405",
				"id_mon":   58,
				"ma_mon":   "BAS1152",
				"nhom_lop": "28",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test invalid slot so cho > sy so
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4005,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test lock record phai tat commnet trong ham handler
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4006,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test insert database rs
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4007,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test insert elasticsearch rs
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4008,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   1464,
				"ma_mon":   "INT1336_CLC",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test update tkb database
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4009,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test rollback fail
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4010,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test rollback es fail
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4011,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
			headers:           map[string]string{`Content-Type`: `application/json`},
		},
		{
			name:       `POST endpoint to regist subject`, //test commit fail
			method:     http.MethodPost,
			path:       `/regist_subject`,
			statusCode: 4012,
			requestBody: map[string]interface{}{
				"ma_sv":    "B18DCCN341",
				"id_mon":   127,
				"ma_mon":   "INT14105",
				"nhom_lop": "1",
			},
			handlerMethodName: RegistSubjectMethodName,
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

// Just duplicated code for demo purposes
//func TestAppNoMock(t *testing.T) {
//
//	testTable := []struct {
//		name                  string
//		method                string
//		path                  string
//		statusCode            int
//		body                  string
//		requestBody           map[string]interface{}
//		handlerMethodName     string
//		handlerToBeCalledWith []interface{}
//		requestHeaders        map[string]string
//		headers               map[string]string
//	}{
//		{
//			name:                  `GET endpoint to get a sum`,
//			method:                http.MethodGet,
//			path:                  `/sum?x=5&y=2`,
//			statusCode:            200,
//			requestBody:           nil,
//			body:                  `{"value":7}`,
//			handlerMethodName:     registSubjectMethodName,
//			handlerToBeCalledWith: []interface{}{5, 2},
//			headers:               map[string]string{`Content-Type`: `application/json`},
//		},
//	}
//
//	for _, tc := range testTable {
//		t.Run(tc.name, func(t *testing.T) {
//
//			// Create and send request
//			rbody, _ := json.Marshal(tc.requestBody)
//			request := httptest.NewRequest(tc.method, tc.path, bytes.NewReader(rbody))
//			request.Header.Add(`Content-Type`, `application/json`)
//
//			// Request Headers
//			for k, v := range tc.requestHeaders {
//				request.Header.Add(k, v)
//			}
//
//			response, _ := APP.Test(request)
//
//			// Status Code
//			statusCode := response.StatusCode
//			if statusCode != tc.statusCode {
//				t.Errorf("StatusCode was incorrect, got: %d, want: %d.", statusCode, tc.statusCode)
//			}
//
//			// Headers
//			for k, want := range tc.headers {
//				headerValue := response.Header.Get(k)
//				if headerValue != want {
//					t.Errorf("Response header '%s' was incorrect, got: '%s', want: '%s'", k, headerValue, want)
//				}
//			}
//
//			// Response Body
//			body, _ := ioutil.ReadAll(response.Body)
//			actual := string(body)
//			if actual != tc.body {
//				t.Errorf("Body was incorrect, got: %v, want: %v", actual, tc.body)
//			}
//
//		})
//	}
//
//}
