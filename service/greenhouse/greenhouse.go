package greenhouse

import (
	"context"
	greenhouseapi "intelligent-greenhouse-service/api/web/greenhouse"
	domain_greenhouse "intelligent-greenhouse-service/domain/greenhouse"
)

type GreenhouseService struct {
	uc *domain_greenhouse.GreenhouseDomain
}

func NewUserService(uc *domain_greenhouse.GreenhouseDomain) *GreenhouseService {
	return &GreenhouseService{uc: uc}
}

func (g GreenhouseService) CreateGreenhouse(ctx context.Context, info *greenhouseapi.CreateGreenhouseInfo) (*greenhouseapi.CreateGreenhouseReply, error) {
	//TODO implement me
	panic("implement me")
}

func (g GreenhouseService) GetGreenhouseInfoByUserId(ctx context.Context, request *greenhouseapi.GetGreenhouseInfoByUserIdRequest) (*greenhouseapi.GetGreenhouseInfoByUserIdReply, error) {
	//TODO implement me
	panic("implement me")
}
