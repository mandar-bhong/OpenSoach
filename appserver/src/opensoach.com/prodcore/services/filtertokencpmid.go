package services

import (
	"errors"

	"opensoach.com/core/logger"
	pchelper "opensoach.com/prodcore/helper"
)

//FilterCPMIDService This service gets devices token by cpmid
type FilterTokenCPMIDService struct {
	*ServiceContext
	NextHandler IHandler
}

func (r *FilterTokenCPMIDService) Handle(serctx *ServiceContext) error {

	var newTokenList []string

	isSuccess, sourceDeviceTokenModel, _ := pchelper.CacheGetDeviceInfo(r.Repo.Context.Master.Cache, r.ServiceConfig.SourceToken)
	if isSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device token model from cache.", nil)
		return errors.New("Failed to get device token model from cache.")
	}

	for _, token := range r.ServiceRuntime.Tokens {
		isSuccess, deviceTokenModel, _ := pchelper.CacheGetDeviceInfo(r.Repo.Context.Master.Cache, token)
		if isSuccess == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device token model from cache.", nil)
			return errors.New("Failed to get device token model from cache.")
		}

		if deviceTokenModel.CpmID == sourceDeviceTokenModel.CpmID {
			newTokenList = append(newTokenList, token)
		}

	}

	r.ServiceRuntime.Tokens = []string{}
	r.ServiceRuntime.Tokens = newTokenList

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(serctx)
		return err
	}

	return nil

}
