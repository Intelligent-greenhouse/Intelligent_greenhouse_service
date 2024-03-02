package device

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "intelligent-greenhouse-service/api/mq/device"
	domain_device "intelligent-greenhouse-service/domain/device"
)

type DeviceService struct {
	uc *domain_device.DeviceDomain
}

func NewDeviceService(uc *domain_device.DeviceDomain) *DeviceService {
	return &DeviceService{uc: uc}
}

func (d DeviceService) UpdateDeviceInfo(ctx context.Context, request *v1.UpdateDeviceInfoRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
