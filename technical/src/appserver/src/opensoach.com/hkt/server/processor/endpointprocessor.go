package processor

import (
	"opensoach.com/hkt/constants"
	"opensoach.com/hkt/server/processor/endpoint"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessEndPointTask(epmodelList []pcmodels.EPTaskSendPacketDataModel) {

	serviceConfigTaskList := []pcmodels.EPTaskSendPacketDataModel{}
	fieldOperatorTaskList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, epmodel := range epmodelList {
		switch epmodel.TaskType {
		case constants.TASK_TYPE_SERV_CONF:
			serviceConfigTaskList = append(serviceConfigTaskList, epmodel)
		case constants.TASK_TYPE_FIELD_OPERATOR:
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
