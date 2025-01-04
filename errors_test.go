package errors

import (
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
			wantText:      "Unauthorized: example message",
			wantCode:      401,
			wantErrorCode: ErrorCodeUnauthorized,
		},
		{
			err:           InternalServerError("some message"),
			wantText:      "Internal Server Error: some message",
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
