package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "ApiKey fsdifgew8rywegfwe", want: "fsdifgew8rywegfwe"},
		{input: "ApiKey svns12312-123-124124-gfsbi", want: "svns12312-123-124124-gfsbi"},
		{input: "ApiKey 1111111-2222222-33333-444-5-5", want: "1111111-2222222-33333-444-5-5"},
	}

	for _, tc := range tests {
		mockRequest := httptest.NewRequest("GET", "https://dummyURL", nil)
		mockRequest.Header.Add("Authorization", tc.input)
		headers := mockRequest.Header
		got, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("Got unexpected error: %s", err)
		}
		if got != tc.want {
			t.Errorf("Got: %s, Expected: %s", got, tc.want)
		}
	}
}

//func GetAPIKey(headers http.Header) (string, error)
