package verificationcode

type VerificationCode interface {
	Valid() bool
	Expire()
	GetCode() string
}

type VerificationStore interface {
	Generate() (VerificationCode, error)
	Delete(code string)
	Check(code string) bool
}
