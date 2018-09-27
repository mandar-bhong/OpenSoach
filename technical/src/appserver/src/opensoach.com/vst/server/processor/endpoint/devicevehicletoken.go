package endpoint

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
	pcmgr "opensoach.com/prodcore/server/manager"
	"opensoach.com/vst/constants"
	hktmodels "opensoach.com/vst/models"
	lconst "opensoach.com/vst/server/constants"
	"opensoach.com/vst/server/dbaccess"
	lhelper "opensoach.com/vst/server/helper"
	lmodels "opensoach.com/vst/server/models"
	repo "opensoach.com/vst/server/repository"
)

func ProcessVehicleTokenData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketVehicleTokenData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		false, 0, nil)

	packetVehicleTokenInsertData := *devicePacket.Payload.(*lmodels.PacketVehicleTokenData)

	var vehicleID int64

	err, vehicleToken := generateVehicleToken(ctx.InstanceDBConn)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while generating vehicle token.", convErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	dbErr, vehicleData := dbaccess.EPGetVehicleId(ctx.InstanceDBConn, packetVehicleTokenInsertData.VehicleNo)

	if dbErr != nil {
		logger.Context().WithField("Token", ctx.Token).
			WithField("Vehicle No.", packetVehicleTokenInsertData.VehicleNo).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting vehicle id.", dbErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	dbTxErr, tx := dbaccess.GetDBTransaction(ctx.InstanceDBConn)

	if dbTxErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Transaction Error.", convErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	vehicleDataItem := *vehicleData
	if len(vehicleDataItem) < 1 {

		dbVehicleInsertRowModel := hktmodels.DBVehicleInsertRowModel{}
		dbVehicleInsertRowModel.CpmId = ctx.TokenInfo.CpmID
		dbVehicleInsertRowModel.VehicleNo = packetVehicleTokenInsertData.VehicleNo
		dbVehicleInsertRowModel.Details = "{}"

		dbErr, insertedID := dbaccess.EPInsertVehicleData(tx, dbVehicleInsertRowModel)

		if dbErr != nil {

			txErr := tx.Rollback()

			if txErr != nil {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
			}

			logger.Context().WithField("Token", ctx.Token).
				WithField("VehicleData", dbVehicleInsertRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving vehicle data.", dbErr)

			packetProcessingResult.IsSuccess = false
			packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
			return
		}

		vehicleID = insertedID

	} else {
		vehicleID = vehicleDataItem[0].VehicleId
	}

	tokenMappingDetailsModel := hktmodels.TokenMappingDetailsModel{}
	tokenMappingDetailsModel.TokenConfigId = 0
	tokenMappingDetailsModel.JobCreationId = 0
	tokenMappingDetailsModel.JobExeId = 0

	isSucess, mappingDetailsJsonString := ghelper.ConvertToJSON(tokenMappingDetailsModel)
	if isSucess == false {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Failed to covert to json")
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	dbTokenInsertRowModel := hktmodels.DBTokenInsertRowModel{}
	dbTokenInsertRowModel.VhlId = vehicleID
	dbTokenInsertRowModel.GeneratedOn = ghelper.GetCurrentTime()
	dbTokenInsertRowModel.Token = vehicleToken
	dbTokenInsertRowModel.State = constants.DB_VEHICLE_TOKEN_STATE_OPEN
	dbTokenInsertRowModel.MappingDetails = mappingDetailsJsonString

	dbErr, tokenInsertedId := dbaccess.EPInsertVstTokenData(tx, dbTokenInsertRowModel)

	if dbErr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().WithField("Token", ctx.Token).
			WithField("VehicleTokenData", dbTokenInsertRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving vehicle data.", dbErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	dbEPSPVhlTokenDataModel := hktmodels.DBEPSPVhlTokenDataModel{}
	dbEPSPVhlTokenDataModel.TokenId = tokenInsertedId
	dbEPSPVhlTokenDataModel.Token = dbTokenInsertRowModel.Token
	dbEPSPVhlTokenDataModel.VhlId = dbTokenInsertRowModel.VhlId
	dbEPSPVhlTokenDataModel.VehicleNo = packetVehicleTokenInsertData.VehicleNo
	dbEPSPVhlTokenDataModel.State = dbTokenInsertRowModel.State
	dbEPSPVhlTokenDataModel.GeneratedOn = dbTokenInsertRowModel.GeneratedOn

	commandAck = lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, dbEPSPVhlTokenDataModel)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true

	//get online devices
	getErr, tokenlistjsonstring := repo.Instance().ProdTaskContext.ProcessTask(pcconst.TASK_GET_ONLINE_DEVICES, "")
	if getErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task.", getErr)
		return
	}

	issuccess, deviceDataList := pcmgr.GetOnlineDevices(repo.Instance().Context.Master.Cache, tokenlistjsonstring, ctx.TokenInfo.CpmID)
	if issuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting online devices.", nil)
		return
	}

	//get device service points
	dbDeviceServicePointDataModelList := []hktmodels.DBDeviceServicePointDataModel{}

	for _, deviceData := range deviceDataList {

		dbErr, devspdata := dbaccess.TaskGetServicePointByDevId(ctx.InstanceDBConn, deviceData.DevID)

		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting device service points.", dbErr)
			return
		}

		dbDeviceServicePointDataModelList = append(dbDeviceServicePointDataModelList, devspdata)

	}

	DBDeviceVhlTokenModelList := []hktmodels.DBDeviceVhlTokenModel{}

	for i := 0; i < len(dbDeviceServicePointDataModelList); i++ {
		dbDeviceVhlTokenModel := hktmodels.DBDeviceVhlTokenModel{}
		dbDeviceVhlTokenModel.SpId = dbDeviceServicePointDataModelList[i].SpId
		dbDeviceVhlTokenModel.DeviceId = dbDeviceServicePointDataModelList[i].DevId
		dbDeviceVhlTokenModel.TokenId = tokenInsertedId
		dbDeviceVhlTokenModel.Token = dbTokenInsertRowModel.Token
		dbDeviceVhlTokenModel.VhlId = dbTokenInsertRowModel.VhlId
		dbDeviceVhlTokenModel.VehicleNo = packetVehicleTokenInsertData.VehicleNo
		dbDeviceVhlTokenModel.State = dbTokenInsertRowModel.State
		dbDeviceVhlTokenModel.GeneratedOn = dbTokenInsertRowModel.GeneratedOn
		if ctx.TokenInfo.DevID != dbDeviceVhlTokenModel.DeviceId {
			DBDeviceVhlTokenModelList = append(DBDeviceVhlTokenModelList, dbDeviceVhlTokenModel)
		}
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, dbDeviceVhlTokenData := range DBDeviceVhlTokenModelList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, dbDeviceVhlTokenData.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "failed to get device tokens", nil)
			continue
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = dbDeviceVhlTokenData
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_VHL_TOKEN_ADDED

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	ProcessDeviceVhlToken(epTaskSendPacketDataList)

}

func generateVehicleToken(DBConn string) (error, int64) {

	var vhlToken int64

	dbErr, vhlTokenData := dbaccess.EPGetLastVhlTokenRecord(DBConn)

	if dbErr != nil {
		return dbErr, 0
	}

	vhlTokendataItem := *vhlTokenData

	currentDate := ghelper.GetCurrentTime().Format("2006-01-02")
	if currentDate == vhlTokendataItem[0].CreatedOn.Format("2006-01-02") {
		vhlToken = vhlTokendataItem[0].Token + 1
	} else {
		vhlToken = 1
	}

	return nil, vhlToken

}

func ProcessVehicleDetailsData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketVehicleDetailsData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetVehicleDetailsUpdateData := *devicePacket.Payload.(*lmodels.PacketVehicleDetailsData)

	dbVehicleDetailsUpdateModel := hktmodels.DBVehicleDetailsUpdateModel{}
	dbVehicleDetailsUpdateModel.VehicleNo = packetVehicleDetailsUpdateData.VehicleNo
	dbVehicleDetailsUpdateModel.Details = packetVehicleDetailsUpdateData.Details

	dbErr, _ := dbaccess.EPUpdateVehicleDetailsData(ctx.InstanceDBConn, dbVehicleDetailsUpdateModel)

	if dbErr != nil {
		logger.Context().WithField("Token", ctx.Token).
			WithField("VehicleDetailsData", dbVehicleDetailsUpdateModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while updating vehicle details data.", dbErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, nil)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true
}

func ProcessTokenGenerationData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	ProcessJobData(ctx, packetProcessingResult, constants.PROCESS_TYPE_TOKEN_GENERATION)

}

func ProcessJobCreationData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	ProcessJobData(ctx, packetProcessingResult, constants.PROCESS_TYPE_JOB_CREATION)

}

func ProcessJobExeData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	ProcessJobData(ctx, packetProcessingResult, constants.PROCESS_TYPE_JOB_EXE)

}

func ProcessGenerateTokenClaimData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	ProcessJobClaimData(ctx, packetProcessingResult, constants.PROCESS_TYPE_TOKEN_GENERATION_CLAIM)

}

func ProcessJobExeClaimData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	ProcessJobClaimData(ctx, packetProcessingResult, constants.PROCESS_TYPE_JOB_EXE_CLAIM)

}

func ProcessJobDeliveredData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	ProcessJobClaimData(ctx, packetProcessingResult, constants.PROCESS_TYPE_JOB_DELIVERED)

}

func ProcessJobData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult, processType int) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketVhlTokenTxnData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	dbTxErr, tx := dbaccess.GetDBTransaction(ctx.InstanceDBConn)

	if dbTxErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Transaction Error.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetVhlTokenTxnData := *devicePacket.Payload.(*lmodels.PacketVhlTokenTxnData)

	dbErr, TokenData := dbaccess.EPGetTokenMappingDetailsData(ctx.InstanceDBConn, packetVhlTokenTxnData.TokenId)

	if dbErr != nil {
		logger.Context().WithField("Token", ctx.Token).
			WithField("Token Id.", packetVhlTokenTxnData.TokenId).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting vhl token data.", dbErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	dbServiceInstanceDataRowModel := hktmodels.DBServiceInstanceTxDataRowModel{}
	dbServiceInstanceDataRowModel.CpmId = ctx.TokenInfo.CpmID
	dbServiceInstanceDataRowModel.ServiceInstanceID = packetVhlTokenTxnData.ServiceInstanceID
	dbServiceInstanceDataRowModel.TransactionData = packetVhlTokenTxnData.TxnData
	dbServiceInstanceDataRowModel.TransactionDate = packetVhlTokenTxnData.TxnDate
	dbServiceInstanceDataRowModel.FOPCode = packetVhlTokenTxnData.FOPCode
	dbServiceInstanceDataRowModel.Status = packetVhlTokenTxnData.Status

	dbErr, insertedId := dbaccess.EPInsertServiceInstanceTxnData(tx, dbServiceInstanceDataRowModel)

	if dbErr != nil {
		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().WithField("Token", ctx.Token).
			WithField("InstanceData", dbServiceInstanceDataRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving service instance txn data.", dbErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	tokenDataItem := *TokenData

	tokenMappingDetailsModel := hktmodels.TokenMappingDetailsModel{}

	isSuccess := ghelper.ConvertFromJSONString(tokenDataItem[0].MappingDetails, &tokenMappingDetailsModel)
	if isSuccess == false {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Failed to convert json")
		packetProcessingResult.IsSuccess = false
		return
	}

	dbTokenMappingDetailsUpdateModel := hktmodels.DBTokenMappingDetailsUpdateModel{}
	dbTokenMappingDetailsUpdateModel.TokenId = packetVhlTokenTxnData.TokenId

	switch processType {
	case constants.PROCESS_TYPE_TOKEN_GENERATION: // generate token
		tokenMappingDetailsModel.TokenConfigId = insertedId
		dbTokenMappingDetailsUpdateModel.State = constants.DB_VEHICLE_TOKEN_STATE_GENERATE_TOKEN
	case constants.PROCESS_TYPE_JOB_CREATION: // job create
		tokenMappingDetailsModel.JobCreationId = insertedId
		dbTokenMappingDetailsUpdateModel.State = constants.DB_VEHICLE_TOKEN_STATE_JOB_CREATION
	case constants.PROCESS_TYPE_JOB_EXE: // job execution
		tokenMappingDetailsModel.JobExeId = insertedId
		dbTokenMappingDetailsUpdateModel.State = constants.DB_VEHICLE_TOKEN_STATE_JOB_EXE
	}

	isSuccess, tokenMappingDetailsString := ghelper.ConvertToJSON(tokenMappingDetailsModel)
	if isSuccess == false {
		logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Failed to convert to json")
		packetProcessingResult.IsSuccess = false
		return
	}

	dbTokenMappingDetailsUpdateModel.MappingDetails = tokenMappingDetailsString

	dbErr, _ = dbaccess.EPUpdateTokenMappingDetailsData(tx, dbTokenMappingDetailsUpdateModel)

	if dbErr != nil {
		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().WithField("Token", ctx.Token).
			WithField("Update token mapping details", dbTokenMappingDetailsUpdateModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while updating token mapping details.", dbErr)

		packetProcessingResult.IsSuccess = false
		return
	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, nil)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true

}

func ProcessJobClaimData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult, processType int) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketVhlTokenTxnData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		false, 0, nil)

	dbTxErr, tx := dbaccess.GetDBTransaction(ctx.InstanceDBConn)

	if dbTxErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Transaction Error.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetVhlTokenTxnData := *devicePacket.Payload.(*lmodels.PacketVhlTokenTxnData)

	dbServiceInstanceDataRowModel := hktmodels.DBServiceInstanceTxDataRowModel{}
	dbServiceInstanceDataRowModel.CpmId = ctx.TokenInfo.CpmID
	dbServiceInstanceDataRowModel.ServiceInstanceID = packetVhlTokenTxnData.ServiceInstanceID
	dbServiceInstanceDataRowModel.TransactionData = packetVhlTokenTxnData.TxnData
	dbServiceInstanceDataRowModel.TransactionDate = packetVhlTokenTxnData.TxnDate
	dbServiceInstanceDataRowModel.FOPCode = packetVhlTokenTxnData.FOPCode
	dbServiceInstanceDataRowModel.Status = packetVhlTokenTxnData.Status

	dbErr, _ := dbaccess.EPInsertServiceInstanceTxnData(tx, dbServiceInstanceDataRowModel)

	if dbErr != nil {
		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().WithField("Token", ctx.Token).
			WithField("InstanceData", dbServiceInstanceDataRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving service instance txn data.", dbErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	dbTokenStateUpdateModel := hktmodels.DBTokenStateUpdateModel{}
	dbTokenStateUpdateModel.TokenId = packetVhlTokenTxnData.TokenId

	switch processType {
	case constants.PROCESS_TYPE_TOKEN_GENERATION_CLAIM: // generate token claim
		dbTokenStateUpdateModel.State = constants.DB_VEHICLE_TOKEN_STATE_GENERATE_TOKEN_CLAIM
	case constants.PROCESS_TYPE_JOB_EXE_CLAIM: // job exec claim
		dbTokenStateUpdateModel.State = constants.DB_VEHICLE_TOKEN_STATE_JOB_EXE_CLAIM
	case constants.PROCESS_TYPE_JOB_DELIVERED: //vehicle delivered
		dbTokenStateUpdateModel.State = constants.DB_VEHICLE_TOKEN_STATE_JOB_DELIVERED
	}

	dberr, _ := dbaccess.EPUpdateTokenStateData(tx, dbTokenStateUpdateModel)

	if dberr != nil {

		txErr := tx.Rollback()

		if txErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to rollback transaction", txErr)
		}

		logger.Context().WithField("Token", ctx.Token).
			WithField("InstanceData", dbTokenStateUpdateModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while updating vhl token status.", dberr)
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		packetProcessingResult.IsSuccess = false
		return
	}

	txErr := tx.Commit()

	if txErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to commit transaction", txErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	dbErr, tokenData := dbaccess.EPGetTokenDataById(ctx.InstanceDBConn, packetVhlTokenTxnData.TokenId)

	if dbErr != nil {
		logger.Context().WithField("Token", ctx.Token).
			WithField("Token Id.", packetVhlTokenTxnData.TokenId).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting vhl token data.", dbErr)
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		packetProcessingResult.IsSuccess = false
		return
	}

	tokendata := *tokenData

	dbEPSPVhlTokenDataModel := hktmodels.DBEPSPVhlTokenDataModel{}
	dbEPSPVhlTokenDataModel.TokenId = tokendata[0].TokenId
	dbEPSPVhlTokenDataModel.Token = tokendata[0].Token
	dbEPSPVhlTokenDataModel.VhlId = tokendata[0].VhlId
	dbEPSPVhlTokenDataModel.VehicleNo = tokendata[0].VehicleNo
	dbEPSPVhlTokenDataModel.State = tokendata[0].State
	dbEPSPVhlTokenDataModel.GeneratedOn = tokendata[0].GeneratedOn

	commandAck = lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, dbEPSPVhlTokenDataModel)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true

	//send notification to all devices
	//get online devices
	getErr, tokenlistjsonstring := repo.Instance().ProdTaskContext.ProcessTask(pcconst.TASK_GET_ONLINE_DEVICES, "")
	if getErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task.", getErr)
		return
	}

	issuccess, deviceDataList := pcmgr.GetOnlineDevices(repo.Instance().Context.Master.Cache, tokenlistjsonstring, ctx.TokenInfo.CpmID)
	if issuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task.", nil)
		return
	}

	//get device service points
	dbDeviceServicePointDataModelList := []hktmodels.DBDeviceServicePointDataModel{}

	for _, deviceData := range deviceDataList {

		dbErr, devspdata := dbaccess.TaskGetServicePointByDevId(ctx.InstanceDBConn, deviceData.DevID)

		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching field operator by fopid.", dbErr)
			return
		}

		dbDeviceServicePointDataModelList = append(dbDeviceServicePointDataModelList, devspdata)

	}

	DBDeviceVhlTokenModelList := []hktmodels.DBDeviceVhlTokenModel{}

	for i := 0; i < len(dbDeviceServicePointDataModelList); i++ {
		dbDeviceVhlTokenModel := hktmodels.DBDeviceVhlTokenModel{}
		dbDeviceVhlTokenModel.SpId = dbDeviceServicePointDataModelList[i].SpId
		dbDeviceVhlTokenModel.DeviceId = dbDeviceServicePointDataModelList[i].DevId
		dbDeviceVhlTokenModel.TokenId = tokendata[0].TokenId
		dbDeviceVhlTokenModel.Token = tokendata[0].Token
		dbDeviceVhlTokenModel.VhlId = tokendata[0].VhlId
		dbDeviceVhlTokenModel.VehicleNo = tokendata[0].VehicleNo
		dbDeviceVhlTokenModel.State = tokendata[0].State
		dbDeviceVhlTokenModel.GeneratedOn = tokendata[0].GeneratedOn
		if ctx.TokenInfo.DevID != dbDeviceVhlTokenModel.DeviceId {
			DBDeviceVhlTokenModelList = append(DBDeviceVhlTokenModelList, dbDeviceVhlTokenModel)
		}
	}

	epTaskSendPacketDataList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, dbDeviceVhlTokenData := range DBDeviceVhlTokenModelList {
		deviceTokenKey := fmt.Sprintf("%s%d", pcconst.CACHE_DEVICE_TOKEN_MAPPING_KEY_PREFIX, dbDeviceVhlTokenData.DeviceId)
		fmt.Println(deviceTokenKey)

		isTokenGetSucc, deviceToken := repo.Instance().Context.Master.Cache.Get(deviceTokenKey)

		if isTokenGetSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device tokens", nil)
			continue
		}

		epTaskSendPacketDataModel := pcmodels.EPTaskSendPacketDataModel{}
		epTaskSendPacketDataModel.Token = deviceToken
		epTaskSendPacketDataModel.Data = dbDeviceVhlTokenData
		epTaskSendPacketDataModel.TaskType = constants.TASK_TYPE_VHL_TOKEN_ADDED

		epTaskSendPacketDataList = append(epTaskSendPacketDataList, epTaskSendPacketDataModel)
	}

	ProcessDeviceVhlToken(epTaskSendPacketDataList)

}

func ProcessDeviceTokenList(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketVehicleTokenData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		false, 0, nil)

	dbErr, vhlTokenDataList := dbaccess.EPGetTokenList(ctx.InstanceDBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get vhl token data list", dbErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	packetProcessingResult.IsSuccess = true

	commandAck = lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CONFIG_DEVICE_TOKEN_LIST,
		devicePacket.Header.SeqID,
		true, 0, vhlTokenDataList)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

}

func ProcessDeviceJobExeConfigList(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketTokenData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		false, 0, nil)

	packetTokenData := *devicePacket.Payload.(*lmodels.PacketTokenData)

	tokensConfigList := []hktmodels.DBTokenConfigModel{}

	for i := 0; i < len(packetTokenData.TokenId); i++ {

		dbTokenConfigModel := hktmodels.DBTokenConfigModel{}
		dbTokenConfigModel.TokenId = packetTokenData.TokenId[i]

		dbErr, tokenConfigList := dbaccess.EPGetConfigListByTokenId(ctx.InstanceDBConn, packetTokenData.TokenId[i])
		if dbErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get vhl token config list", dbErr)
			packetProcessingResult.IsSuccess = false
			return
		}

		dbTokenConfigModel.TokenConfig = tokenConfigList

		tokensConfigList = append(tokensConfigList, dbTokenConfigModel)
	}

	packetProcessingResult.IsSuccess = true

	commandAck = lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CONFIG_DEVICE_TOKEN_CONFIG_LIST,
		devicePacket.Header.SeqID,
		true, 0, tokensConfigList)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

}

func ProcessDeviceVehicleDetails(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketVehicleTokenData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		false, 0, nil)

	packetVehicleTokenData := *devicePacket.Payload.(*lmodels.PacketVehicleTokenData)

	dbErr, vhlDetailsData := dbaccess.EPGetVehicleDetailsDataByVhlNo(ctx.InstanceDBConn, packetVehicleTokenData.VehicleNo)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get vhl details data ", dbErr)
		packetProcessingResult.IsSuccess = false
		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)
		return
	}

	vhlDetails := *vhlDetailsData

	packetProcessingResult.IsSuccess = true

	commandAck = lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CONFIG_DEVICE_VHL_DETAILS,
		devicePacket.Header.SeqID,
		true, 0, vhlDetails[0])

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

}

//process device vhltoken packet
func ProcessDeviceVhlToken(epTaskSendPacketDataModelList []pcmodels.EPTaskSendPacketDataModel) {

	var tokenSPServices map[string]map[int64][]hktmodels.DBDeviceVhlTokenModel
	tokenSPServices = make(map[string]map[int64][]hktmodels.DBDeviceVhlTokenModel)

	for _, epTask := range epTaskSendPacketDataModelList {

		_, hasToken := tokenSPServices[epTask.Token]

		if hasToken == false {
			tokenSPServices[epTask.Token] = map[int64][]hktmodels.DBDeviceVhlTokenModel{}
		}

		dbDeviceVhlTokenModel := epTask.Data.(hktmodels.DBDeviceVhlTokenModel)

		_, hasSP := tokenSPServices[epTask.Token][dbDeviceVhlTokenModel.SpId]

		if hasSP == false {
			tokenSPServices[epTask.Token][dbDeviceVhlTokenModel.SpId] = []hktmodels.DBDeviceVhlTokenModel{dbDeviceVhlTokenModel}
		} else {
			tokenSPServices[epTask.Token][dbDeviceVhlTokenModel.SpId] = append(tokenSPServices[epTask.Token][dbDeviceVhlTokenModel.SpId], dbDeviceVhlTokenModel)
		}
	}

	epPackets := []pcmodels.TaskEPPacketSendDataModel{}

	for token, tokenSPItem := range tokenSPServices {

		for spid, spTokenItem := range tokenSPItem {

			isJsonSucc, jsonData := createVhlTokenPacket(spid, spTokenItem)

			if isJsonSucc == false {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert to json data", nil)
				continue
			}

			packetSendDataModel := pcmodels.TaskEPPacketSendDataModel{}
			packetSendDataModel.Token = token
			packetSendDataModel.Packet = jsonData

			epPackets = append(epPackets, packetSendDataModel)
		}
	}

	SendPacketToEP(epPackets)
}

func createVhlTokenPacket(spid int64, vhlTokenModels []hktmodels.DBDeviceVhlTokenModel) (bool, string) {

	vhltokeninfo := &gmodels.DevicePacket{}
	vhltokeninfo.Header = gmodels.DeviceHeaderData{}
	vhltokeninfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	vhltokeninfo.Header.CommandID = lconst.DEVICE_CMD_CONFIG_SERVICE_POINTS_TOKEN
	vhltokeninfo.Header.SPID = spid

	var vhlTokenList []hktmodels.DBEPSPVhlTokenDataModel

	for _, vhlTokenModel := range vhlTokenModels {
		dbEPSPVhlTokenDataModel := hktmodels.DBEPSPVhlTokenDataModel{}

		dbEPSPVhlTokenDataModel.TokenId = vhlTokenModel.TokenId
		dbEPSPVhlTokenDataModel.Token = vhlTokenModel.Token
		dbEPSPVhlTokenDataModel.VhlId = vhlTokenModel.VhlId
		dbEPSPVhlTokenDataModel.VehicleNo = vhlTokenModel.VehicleNo
		dbEPSPVhlTokenDataModel.State = vhlTokenModel.State
		dbEPSPVhlTokenDataModel.GeneratedOn = vhlTokenModel.GeneratedOn

		vhlTokenList = append(vhlTokenList, dbEPSPVhlTokenDataModel)
	}

	vhltokeninfo.Payload = vhlTokenList

	isJsonSucc, jsonData := ghelper.ConvertToJSON(vhltokeninfo)

	return isJsonSucc, jsonData

}
