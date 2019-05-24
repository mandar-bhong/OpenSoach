package processor

import (
	gcore "opensoach.com/core"
	"opensoach.com/core/logger"
	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	pcepproc "opensoach.com/prodcore/endpoint/processor"
)

func AuthProcessor(epmodel *gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult {
	logger.Context().WithField("EndPointModel", epmodel).LogDebug(SUB_MODULE_NAME, logger.Normal, "Starting device authrization")
	packetProcessingResult := pcepproc.AuthorizeDevice(repo.Instance().Context.Master.Cache, epmodel, authSuccessHandler)

	return packetProcessingResult
}

func authSuccessHandler(cacheCtx gcore.CacheContext, token string, chnID int) {
	chnIDvsToken[chnID] = token
	tokenvsChnID[token] = chnID

	subErr := repo.Instance().ProdTaskContext.SubmitTask(gmodels.TASK_HKT_EP_CONNECTED, token)

	if subErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while submitting task to HKT task queue.", subErr)
	}
}
