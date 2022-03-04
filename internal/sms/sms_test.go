package sms

import (
	"bytes"
	"encoding/json"
	"golang-training/main/internal/repository"
	"golang-training/main/internal/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddData(t *testing.T){
	validReq, _ :=  json.Marshal(&AddDataRequest{
		FirstName: "aaa",
		LastName: "tttt",
		Email: "xxxxxx",
		Age: 18,
	})
	tt := []struct{
		name 	string
		userDetailRepo repository.UserDetail
		w 		*httptest.ResponseRecorder
		r 		*http.Request
		expectResult int
	}{
		{
				name: "Status 200 : add data sucessfully", //name testcase 
				userDetailRepo: mock.UserDeatail("OK"),
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost,"http://localhost:3306",bytes.NewReader(validReq)),
				expectResult: http.StatusOK,
		},
		{
			name: "Status 500 : Internal error", //name testcase 
			userDetailRepo: mock.UserDeatail("NOK"),
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodPost,"http://localhost:3306",bytes.NewReader(validReq)),
			expectResult: http.StatusInternalServerError,
		},
		{
			name: "Status 400 : Invalid Format", //name testcase 
			w: httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodPost,"http://localhost:3306",bytes.NewReader([]byte("invalid"))),
			expectResult: http.StatusBadRequest ,

		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := service{
				Data:  []string{},
				userDetailRepo: tc.userDetailRepo,
			}
			s.AddData(tc.w, tc.r)
			assert.Equal(t, tc.expectResult, tc.w.Result().StatusCode)
		})
	}
}

func Test_ValidatData(t *testing.T) {
	validReq, _ := json.Marshal(&AddDataRequest{
		FirstName: "aaa",
		LastName: "tttt",
		Email: "xxxxxx",
		Age: 18,
	})
	tt := []struct{
		name string
		r   *http.Request
		expectResult *AddDataRequest
		expectError error
	} {
		{
			name: "request is valid", //name testcase 
			r: httptest.NewRequest(http.MethodPost,"http://localhost:3306", bytes.NewReader(validReq)),
			expectResult: &AddDataRequest{
				FirstName: "aaa",
				LastName: "tttt",
				Email: "xxxxxx",
				Age: 18,
			},
			expectError: nil,
		}, 
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r , e := validateReq(tc.r)
			if tc.expectError != nil {
				assert.NotNil(t, e)
			}
			assert.Equal(t, tc.expectResult, r)

		})
	}
}