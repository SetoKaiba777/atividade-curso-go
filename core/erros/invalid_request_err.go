package erros

type InvalidRequestErr struct {
}

func NewInvalidRequestErr() error {
	return InvalidRequestErr{}
}

func (c InvalidRequestErr) Error() string {
	return "requisicao inválida!"
}
