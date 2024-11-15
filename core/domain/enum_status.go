package domain

type Status int

const (
	PENDENTE Status = iota
	PROCESSADO
	ENVIADO
	CONCLUIDO
)


func (s Status) Next() Status {
	if s < CONCLUIDO {
		return s+1
	}
	return s
}

func (s Status) Previous() Status {
	if s > PENDENTE {
		return s-1
	}
	return s
}