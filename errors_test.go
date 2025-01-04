package errors

import (
	"encoding/json"
	"testing"
)

func TestHTTPError(t *testing.T) {
	cases := []struct {
		err           *HTTPError
		wantText      string
		wantCode      int
		wantErrorCode ErrorCode
	}{
		{
			err:           Unauthorized("example message"),
			wantText:      "example message",
			wantCode:      401,
			wantErrorCode: ErrorCodeUnauthorized,
		},
		{
			err:           InternalServerError("some message"),
			wantText:      "some message",
			wantCode:      500,
			wantErrorCode: ErrorCodeInternalServerError,
		},
	}
	for _, c := range cases {
		var (
			text       = c.err.Error()
			statusCode = c.err.StatusCode()
			errorCode  = c.err.Code
		)
		if text != c.wantText {
			t.Errorf("got text %s, want %s", text, c.wantText)
		}
		if statusCode != c.wantCode {
			t.Errorf("got status code %d, want %d", statusCode, c.wantCode)
		}
		if errorCode != c.wantErrorCode {
			t.Errorf("got error code %d, want %d", errorCode, c.wantCode)
		}
	}
}

func TestHTTPError_MarshalJSON(t *testing.T) {
	e := Unauthorized("you are not authorized to access this resource")
	out, err := json.Marshal(&e)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"code":"Unauthorized","message":"you are not authorized to access this resource"}`
	got := string(out)
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}
