package service_test

import (
	"github.com/Twitter/src/domain"
	"github.com/Twitter/src/service"
	"testing"
)

//Los tests siempre arrancan con la palabra Test
func TestUserIsSaved(t *testing.T) {

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)

	// Validation
	if usuario == nil {
		t.Error("Expected error")
		return
	}

}

func TestUserWhichExceeding40CharactersIsNotAdd(t *testing.T) {

	// Initialization
	usuarioManager := service.NewUsuarioManager()

	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel312323213123123213121223232423423534534564654645654"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)

	// Operation
	var err error
	err = usuarioManager.AddUser(usuario)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "name exceeds 40 characters" {
		t.Error("Expected error is name exceeds 40 characters")
	}

}

func TestReturns2users(t *testing.T) {

	// Initialization
	usuarioManager := service.NewUsuarioManager()

	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	nombre = "Lucas"
	mail = "lucas@hotmail.com"
	nick = "Lu"
	clave = "5678"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	// Validation
	usuarios := usuarioManager.GetUsuarios()

	if len(usuarios) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(usuarios))
		return
	}

}

func TestVerifyUser(t *testing.T) {

	// Initialization
	usuarioManager := service.NewUsuarioManager()

	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	nombre = "Lucas"
	mail = "lucas@hotmail.com"
	nick = "Lu"
	clave = "5678"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	if usuarioManager.VerifyUser(usuario.Nombre)==false {

		t.Errorf("Expected user")
		return
	}

}