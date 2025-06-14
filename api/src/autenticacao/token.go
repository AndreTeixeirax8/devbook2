package autenticacao

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 48).Unix() //aqui adiciona quantas horas o token é valido
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("segredo"))

}
