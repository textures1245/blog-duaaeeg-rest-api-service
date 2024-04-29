package auth

type AuthUsecase interface {
	Login(req *UsersCredentials) (*UsersLoginRes, error)
	Register(req *UsersCredentials) (*UsersLoginRes, error)
}
