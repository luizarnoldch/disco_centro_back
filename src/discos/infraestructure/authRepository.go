package infraestructure

type AuthRepository interface {
	IsAuthorized(string, string, map[string]string)bool
}