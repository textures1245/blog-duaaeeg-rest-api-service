package auth

type AuthRepository interface {
	SignUsersAccessToken(req *UsersPassport) (string, error)
}
