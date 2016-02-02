package android

const (
	ERR_SignatureDecodeFailed = iota
	ERR_VerificationFailed
	ERR_ParseFailed
)

type VerifyError struct {
	Code    int
	Message string
	Cause   error
}

func newVerifyError(code int, msg string, err error) *VerifyError {
	return &VerifyError{
		Code:    code,
		Message: msg,
		Cause:   err,
	}
}

func (o *VerifyError) Error() string {
	return o.Message
}
