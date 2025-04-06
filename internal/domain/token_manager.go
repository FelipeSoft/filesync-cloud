package domain

type TokenManager interface {
	AssignRSA(sub string) (string, error)
	VerifyRSA(token string) (any, error)
}
