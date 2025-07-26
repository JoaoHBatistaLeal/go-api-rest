package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidateNotaAvaliacao(avaliacao models.Avaliacao) error {
	if avaliacao.Nota < 0 || avaliacao.Nota > 5 {
		return errors.New("a nota deve estar entre 1 e 5")
	}
	return nil
}
