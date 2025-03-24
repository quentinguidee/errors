package errors

import (
	"encoding/json"
	"errors"
	"net/http"
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
	t.Run("WithoutCode", func(t *testing.T) {
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
	})

	t.Run("WithCode", func(t *testing.T) {
		e := UnauthorizedNamed("ERR_UNAUTHORIZED", "you are not authorized to access this resource")
		out, err := json.Marshal(&e)
		if err != nil {
			t.Fatal(err)
		}
		want := `{"code":"Unauthorized","name":"ERR_UNAUTHORIZED","message":"you are not authorized to access this resource"}`
		got := string(out)
		if want != got {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestIsStatus(t *testing.T) {
	cases := []struct {
		name       string
		err        error
		statusCode int
		want       bool
	}{
		{
			name:       "IsNotFound",
			err:        NotFound("not found"),
			statusCode: http.StatusNotFound,
			want:       true,
		},
		{
			name:       "IsUnauthorized",
			err:        Unauthorized("you are not authorized to access this resource"),
			statusCode: http.StatusUnauthorized,
			want:       true,
		},
		{
			name:       "IsNotUnauthorized",
			err:        NotFound("not found"),
			statusCode: http.StatusUnauthorized,
			want:       false,
		},
		{
			name:       "PlainError",
			err:        errors.New("test"),
			statusCode: http.StatusOK,
			want:       false,
		},
		{
			name:       "Nil",
			err:        nil,
			statusCode: http.StatusOK,
			want:       false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := IsStatus(c.statusCode, c.err)
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
