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
	greenhouse, err := g.uc.CreateGreenhouse(ctx, info.Size, info.Name, info.Pos)
	if err != nil {
		return nil, err
	}

	return &greenhouseapi.CreateGreenhouseReply{GreenhouseId: greenhouse.ID}, nil
}

func (g GreenhouseService) GetGreenhouseInfoByUserId(ctx context.Context, request *greenhouseapi.GetGreenhouseInfoByUserIdRequest) (*greenhouseapi.GetGreenhouseInfoByUserIdReply, error) {
	list, err := g.uc.GetGreenhouseListByUserId(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	var greenHouseList []*greenhouseapi.GreenhouseInfo
	for _, g := range list {
		greenHouseList = append(greenHouseList, &greenhouseapi.GreenhouseInfo{
			GreenhouseId: g.ID,
			Pos:          g.Pos,
			Size:         g.Size,
		})
	}

	return &greenhouseapi.GetGreenhouseInfoByUserIdReply{List: greenHouseList}, nil
}
