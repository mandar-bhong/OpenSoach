package api

import (
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/dbaccess"
	repo "opensoach.com/hkt/server/repository"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceSPAssociated(jsonmsg string) error {

	return nil
}

func ProcessSerConfigOnSP(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskResultModel) {

	taskSerConfigAddedOnSPModel := ctx.TaskData.(*hktmodels.TaskSerConfigAddedOnSPModel)

	dbaccess.GetInstanceDBConn(repo.Instance().Context.Master.DBConn, taskSerConfigAddedOnSPModel.CpmId)

	apiTaskResultModel := &pcmodels.APITaskResultModel{}
	return nil, apiTaskResultModel
}
