package services

import (
	"errors"

	"opensoach.com/core/logger"
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
)

//FilterCPMIDService This service gets devices token by cpmid
type FilterCPMIDService struct {
	*ServiceContext
	NextHandler IHandler
}

func (r *FilterCPMIDService) Handle(serctx *ServiceContext) error {

	var newTokenList []string

	var deviceCpmID int64

	for _, token := range r.ServiceRuntime.Tokens {

		isSuccess, deviceType, deviceTokenModel, userTokenModel, _ := pchelper.CacheGetDeviceInfoData(r.Repo.Context.Master.Cache, token)
		if isSuccess == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device token model from cache.", nil)
			return errors.New("Failed to get device token model from cache.")
		}

		if deviceType == pcconst.DEVICE_TYPE_SHARED_DEVICE {
			deviceCpmID = deviceTokenModel.CpmID
		} else if deviceType == pcconst.DEVICE_TYPE_USER_DEVICE {
			deviceCpmID = userTokenModel.Product.CustProdID
		}

		if deviceCpmID == r.ServiceConfig.CPMID {
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
