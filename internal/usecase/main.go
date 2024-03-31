package usecase

type UseCases struct {
	Authorization IAuthorization
	Jwt           IJwt
	Shelters      IShelters
	Files         IFiles
}
