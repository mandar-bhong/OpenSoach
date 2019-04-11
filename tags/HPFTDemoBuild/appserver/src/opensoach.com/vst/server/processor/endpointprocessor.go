package processor

import (
	pcmodels "opensoach.com/prodcore/models"
	"opensoach.com/vst/constants"
	"opensoach.com/vst/server/processor/endpoint"
)

func ProcessEndPointTask(epmodelList []pcmodels.EPTaskSendPacketDataModel) {

	serviceConfigTaskList := []pcmodels.EPTaskSendPacketDataModel{}
	fieldOperatorTaskList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, epmodel := range epmodelList {
		switch epmodel.TaskType {
		case constants.TASK_TYPE_SERV_CONF:
			serviceConfigTaskList = append(serviceConfigTaskList, epmodel)
		case constants.TASK_TYPE_FIELD_OPERATOR_ASSOCIATED:
			fieldOperatorTaskList = append(fieldOperatorTaskList, epmodel)
		case constants.TASK_TYPE_FIELD_OPERATOR_DEASSOCIATED:
			fieldOperatorTaskList = append(fieldOperatorTaskList, epmodel)
		case constants.TASK_TYPE_FIELD_OPERATOR_ADDED:
			fieldOperatorTaskList = append(fieldOperatorTaskList, epmodel)
		}

	}

	if len(serviceConfigTaskList) > 0 {
		endpoint.ProcessDeviceServiceConfig(serviceConfigTaskList)
	}

	if len(fieldOperatorTaskList) > 0 {
		endpoint.ProcessDeviceFieldOperator(fieldOperatorTaskList)
	}

}
