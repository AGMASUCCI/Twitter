package service

import (
	"errors"
	"github.com/Twitter/src/domain"
)

var usr *domain.Usuario

type UsuarioManager struct {
	usrs []*domain.Usuario
}

var um UsuarioManager

func (um *UsuarioManager)AddUser(u *domain.Usuario) (error) {

	if u.Nombre == "" {
		return errors.New("nombre is required")
	}

	if len(u.Nombre) > 40 {
		return errors.New("name exceeds 40 characters")
	}

	usr = u
	um.usrs = append(um.usrs, usr)

	return  nil
}


func (um *UsuarioManager)GetUsuario() *domain.Usuario {

	return usr;
}

func (um *UsuarioManager)GetUsuarios() []*domain.Usuario{

	return um.usrs;
}

func (um *UsuarioManager)InitializeService(){

	um.usrs = make([]*domain.Usuario, 0)

}

func (um *UsuarioManager)VerifyUser(usr string) bool {

	for _, value := range um.usrs {
		if value.Nombre == usr {
			return true
		}
	}

	return false
}

func NewUsuarioManager() *UsuarioManager{

	usuarioManager := UsuarioManager{make([]*domain.Usuario, 0)}
	return &usuarioManager

}

