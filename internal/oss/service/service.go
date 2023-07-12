package service

type Service interface {
	Echo(string) string
}

type service struct{}

func (s service) Echo(str string) string {
	return str
}

func NewService() Service {
	return service{}
}
