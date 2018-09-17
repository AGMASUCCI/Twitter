package domain

type Usuario struct{
	Nombre , Mail, Nick, Clave string
}

func NewUsuario(nombre , mail, nick, clave string) *Usuario {
	return &Usuario{nombre,mail, nick, clave}
}

