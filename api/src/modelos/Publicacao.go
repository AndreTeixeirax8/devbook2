package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64    `json:"id, omitempty"` // ID da publicação
	Titulo    string    `json:"titulo, omitempty"`
	Conteudo  string    `json:"conteudo, omitempty"`
	AutorID   uint64    `json:"autorId, omitempty"`
	AutorNick string    `json:"autorNick, omitempty"` // Nick do autor da publicação
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm, omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil

}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o título da publicação não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo da publicação não pode estar em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
