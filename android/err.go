package android

const (
	// SignatureDecodeFailed is used when broken signature is passed
	// Signature may be malformed.
	SignatureDecodeFailed = iota + 1

	// VerificationFailed is used when verification is failed.
	// Receipt JSON may be malformed.
	VerificationFailed

	// ParseFailed is used when receipt JSON cannot be parsed.
	// Developer don't need to handle this error because
	// verification process itself is succeeded...
	ParseFailed
)

// VerifyError is an error object
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
