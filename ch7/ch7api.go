package ch7

type Logic interface {
	SayHello(UserID string) (string, error)
}
