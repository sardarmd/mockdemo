package domain

type User struct {
	Id   string
	Name string
	City string
}

//go:generate mockgen -destination=./mocks/domain/mockAccountRepository.go  -package=mock github.com/sardarmd/mockgendemo/domain UserRepo
type UserRepo interface {
	AddUser(user User)
	FindUserByPosition(index int64)
}
