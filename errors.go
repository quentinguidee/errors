package errors

import (
	"fmt"
	"net/http"
)

type ErrorCode int

// Codes from https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses
const (
	ErrorCodeBadRequest                  ErrorCode = 400
	ErrorCodeUnauthorized                ErrorCode = 401
	ErrorCodePaymentRequired             ErrorCode = 402
	ErrorCodeForbidden                   ErrorCode = 403
	ErrorCodeNotFound                    ErrorCode = 404
	ErrorCodeMethodNotAllow              ErrorCode = 405
	ErrorCodeNotAcceptable               ErrorCode = 406
	ErrorCodeProxyAuthRequired           ErrorCode = 407
	ErrorCodeRequestTimeout              ErrorCode = 408
	ErrorCodeConflict                    ErrorCode = 409
	ErrorCodeGone                        ErrorCode = 410
	ErrorCodeLengthRequired              ErrorCode = 411
	ErrorCodePreconditionFailed          ErrorCode = 412
	ErrorCodeContentTooLarge             ErrorCode = 413
	ErrorCodeURITooLong                  ErrorCode = 414
	ErrorCodeUnsupportedMediaType        ErrorCode = 415
	ErrorCodeRangeNotSatisfiable         ErrorCode = 416
	ErrorCodeExpectationFailed           ErrorCode = 417
	ErrorCodeTeapot                      ErrorCode = 418
	ErrorCodeMisdirectedRequest          ErrorCode = 421
	ErrorCodeUnprocessableContent        ErrorCode = 422
	ErrorCodeLocked                      ErrorCode = 423
	ErrorCodeFailedDependency            ErrorCode = 424
	ErrorCodeTooEarly                    ErrorCode = 425
	ErrorCodeUpgradeRequired             ErrorCode = 426
	ErrorCodePreconditionRequired        ErrorCode = 428
	ErrorCodeTooManyRequests             ErrorCode = 429
	ErrorCodeRequestHeaderFieldsTooLarge ErrorCode = 431
	ErrorCodeUnavailableForLegalReasons  ErrorCode = 451

	ErrorCodeInternalServerError           ErrorCode = 500
	ErrorCodeNotImplemented                ErrorCode = 501
	ErrorCodeBadGateway                    ErrorCode = 502
	ErrorCodeServiceUnavailable            ErrorCode = 503
	ErrorCodeGatewayTimeout                ErrorCode = 504
	ErrorCodeHTTPVersionNotSupported       ErrorCode = 505
	ErrorCodeVariantAlsoNegotiates         ErrorCode = 506
	ErrorCodeInsufficientStorage           ErrorCode = 507
	ErrorCodeLoopDetected                  ErrorCode = 508
	ErrorCodeNotExtended                   ErrorCode = 510
	ErrorCodeNetworkAuthenticationRequired ErrorCode = 511
)

func (e ErrorCode) String() string {
	return http.StatusText(int(e))
}

type HTTPError struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

func NewHTTPError(code ErrorCode, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: msg,
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code.String(), e.Message)
}

func (e *HTTPError) StatusCode() int {
	return int(e.Code)
}

// 400

func BadRequest(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeBadRequest, msg)
}
func Unauthorized(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeUnauthorized, msg)
}
func PaymentRequired(msg string) *HTTPError {
	return NewHTTPError(ErrorCodePaymentRequired, msg)
}
func Forbidden(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeForbidden, msg)
}
func NotFound(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeNotFound, msg)
}
func MethodNotAllow(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeMethodNotAllow, msg)
}
func NotAcceptable(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeNotAcceptable, msg)
}
func ProxyAuthRequired(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeProxyAuthRequired, msg)
}
func RequestTimeout(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeRequestTimeout, msg)
}
func Conflict(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeConflict, msg)
}
func Gone(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeGone, msg)
}
func LengthRequired(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeLengthRequired, msg)
}
func PreconditionFailed(msg string) *HTTPError {
	return NewHTTPError(ErrorCodePreconditionFailed, msg)
}
func ContentTooLarge(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeContentTooLarge, msg)
}
func URITooLong(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeURITooLong, msg)
}
func UnsupportedMediaType(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeUnsupportedMediaType, msg)
}
func RangeNotSatisfiable(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeRangeNotSatisfiable, msg)
}
func ExpectationFailed(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeExpectationFailed, msg)
}
func Teapot(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeTeapot, msg)
}
func MisdirectedRequest(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeMisdirectedRequest, msg)
}
func UnprocessableContent(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeUnprocessableContent, msg)
}
func Locked(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeLocked, msg)
}
func FailedDependency(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeFailedDependency, msg)
}
func TooEarly(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeTooEarly, msg)
}
func UpgradeRequired(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeUpgradeRequired, msg)
}
func PreconditionRequired(msg string) *HTTPError {
	return NewHTTPError(ErrorCodePreconditionRequired, msg)
}
func TooManyRequests(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeTooManyRequests, msg)
}
func RequestHeaderFieldsTooLarge(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeRequestHeaderFieldsTooLarge, msg)
}
func UnavailableForLegalReasons(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeUnavailableForLegalReasons, msg)
}

// 500

func InternalServerError(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeInternalServerError, msg)
}
func NotImplemented(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeNotImplemented, msg)
}
func BadGateway(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeBadGateway, msg)
}
func ServiceUnavailable(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeServiceUnavailable, msg)
}
func GatewayTimeout(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeGatewayTimeout, msg)
}
func HTTPVersionNotSupported(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeHTTPVersionNotSupported, msg)
}
func VariantAlsoNegotiates(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeVariantAlsoNegotiates, msg)
}
func InsufficientStorage(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeInsufficientStorage, msg)
}
func LoopDetected(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeLoopDetected, msg)
}
func NotExtended(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeNotExtended, msg)
}
func NetworkAuthenticationRequired(msg string) *HTTPError {
	return NewHTTPError(ErrorCodeNetworkAuthenticationRequired, msg)
}
