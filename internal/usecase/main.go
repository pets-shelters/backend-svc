package usecase

type UseCases struct {
	Authorization IAuthorization
	Jwt           IJwt
	Shelters      IShelters
	Files         IFiles
	Employees     IEmployees
	Locations     ILocations
	Animals       IAnimals
	Adopters      IAdopters
	Tasks         ITasks
	Walkings      IWalkings
}
