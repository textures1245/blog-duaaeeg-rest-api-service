package user

type UserService interface {
	OnFetchUserByUUID(userUuid string) (*UserResDat, error)
	OnUpdateUserProfile(userUuid string, req *UserProfileDataRequest) (*UserProfileRes, error)
}
