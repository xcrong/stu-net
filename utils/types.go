package utils

type Account struct {
	Username string
	Password string
}

type Config struct {
	account Account
	cookie  string
}
