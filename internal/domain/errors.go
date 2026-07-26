package domain

import "errors"

var (
	ErrTitleRequired  = errors.New("O título do livro é obrigatório")
	ErrAuthorRequired = errors.New("O autor é obrigatório")
	ErrInvalidPages   = errors.New("O número de páginas deve ser maior que zero")
	ErrBookNotFound   = errors.New("Livro não Encontrado")
)
