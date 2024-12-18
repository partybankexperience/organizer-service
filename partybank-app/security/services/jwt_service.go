package services

type JwtService interface {
	IsNonExpired(token string) (bool, error)
}

type PartyBankJwtService struct {
}

func NewJwtService() JwtService {
	return &PartyBankJwtService{}
}

func (partybankJwtService *PartyBankJwtService) IsNonExpired(token string) (bool, error) {
	return false, nil
}
