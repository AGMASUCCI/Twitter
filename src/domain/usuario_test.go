package domain_test

import (
	"github.com/Twitter/src/domain"
	"testing"
)

func TestValidateUserInfo(t *testing.T) {

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)

	// Validation
	if usuario.Nombre != nombre {
		t.Error("Nombre incorrecto")
	}

	if usuario.Mail != mail {
		t.Error("Mail incorrecto")
	}

	if usuario.Nick != nick {
		t.Error("Nick incorrecto")
	}

	if usuario.Clave != clave {
		t.Error("Clave incorrecto")
	}

}
