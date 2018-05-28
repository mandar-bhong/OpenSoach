package processor

import (
	

	pcmodels "opensoach.com/prodcore/models"
	"opensoach.com/hkt/server/processor/endpoint"
)

func ProcessEndPointTask(epmodelList []pcmodels.EPTaskSendPacketDataModel) {


	serviceConfigTaskList := []pcmodels.EPTaskSendPacketDataModel{}

	for _, epmodel := range epmodelList {
		switch epmodel.TaskType {
		case "ServiceConfig":
			serviceConfigTaskList = append(serviceConfigTaskList,epmodel )
		}

	}

	
if len(serviceConfigTaskList) > 0{
	endpoint.ProcessDeviceServiceConfig(serviceConfigTaskList)
}
}
