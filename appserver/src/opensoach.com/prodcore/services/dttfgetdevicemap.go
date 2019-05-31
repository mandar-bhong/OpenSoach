package services

import (
	"errors"

	"opensoach.com/core/logger"
	pchelper "opensoach.com/prodcore/helper"
)

//DttfGetDevMapService This service gets device mappings
type DttfGetDevMapService struct {
	*ServiceContext
	NextHandler IHandler
}

func (r *DttfGetDevMapService) Handle(serctx *ServiceContext) error {

	r.ServiceRuntime.DeviceLocationMap = make(map[int64][]int64)

	for _, token := range r.ServiceRuntime.Tokens {

		isSuccess, deviceTokenModel, _ := pchelper.CacheGetDeviceInfo(r.Repo.Context.Master.Cache, token)
		if isSuccess == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device token model from cache.", nil)
			return errors.New("Failed to get device token model from cache.")
		}

		_, ok := r.ServiceRuntime.DeviceLocationMap[deviceTokenModel.DevID]
		if !ok {
			r.ServiceRuntime.DeviceLocationMap[deviceTokenModel.DevID] = append(r.ServiceRuntime.DeviceLocationMap[deviceTokenModel.DevID], 0)
		}
	}

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(serctx)
		return err
	}

	return nil

}
