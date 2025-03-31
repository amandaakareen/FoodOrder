package domain

type User struct {
	name      string
	password  string
	cpf       string
	email     string
	telephone string
}

type UserRepository interface {
	FindByCpfPassword(cpf string, password string)
}
