package processor

import (
	"opensoach.com/hpft/constants"
	"opensoach.com/hpft/server/processor/endpoint"
	pcmodels "opensoach.com/prodcore/models"
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
