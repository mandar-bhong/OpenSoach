package models

import (
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
)

func (r *DevicePacketProccessExecution) GetCPMID() int64 {

	var cpmID int64

	switch r.DeviceContext.(type) {
	case *gmodels.DeviceTokenModel:
		cpmID = r.DeviceContext.(*gmodels.DeviceTokenModel).CpmID
		break
	case *gmodels.DeviceUserSessionInfo:
		cpmID = r.DeviceContext.(*gmodels.DeviceUserSessionInfo).Product.CustProdID
		break
	}
	return cpmID
}

func (r *DevicePacketProccessExecution) GetDeviceID() (bool, int64) {

	var deviceId int64

	switch r.DeviceContext.(type) {
	case *gmodels.DeviceTokenModel:
		deviceId = r.DeviceContext.(*gmodels.DeviceTokenModel).DevID
		return true, deviceId
	}
	return false, deviceId
}

func (r *DevicePacketProccessExecution) GetDeviceUserID() (bool, int64) {

	var deviceUserID int64

	switch r.DeviceContext.(type) {
	case *gmodels.DeviceUserSessionInfo:
		deviceUserID = r.DeviceContext.(*gmodels.DeviceUserSessionInfo).UserID
		return true, deviceUserID
	}
	return false, deviceUserID
}

func (r *DevicePacketProccessExecution) GetDeviceContextType() int {

	switch r.DeviceContext.(type) {
	case *gmodels.DeviceTokenModel:
		return pcconst.DEVICE_TYPE_SHARED_DEVICE
	case *gmodels.DeviceUserSessionInfo:
		return pcconst.DEVICE_TYPE_USER_DEVICE
	}

	return pcconst.DEVICE_TYPE_NONE
}
