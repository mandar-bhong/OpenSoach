package processor

import (
	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/endpoint/repository"
	gmodels "opensoach.com/models"
)

func EPOnConnectProcessExecutor(chnid int) {
	//TODO: Relation should be maintain that if device connected and does not send auth command
	// for longer duration then device should be disconnected
}

func EPOnDisConnectProcessExecutor(chnid int) {

	token, hasChnID := chnIDvsToken[chnid]

	if hasChnID {
		delete(chnIDvsToken, chnid)
	}

	_, hasToken := tokenvsChnID[token]

	if hasToken {
		delete(tokenvsChnID, token)
	}

	subErr := repo.Instance().ProdTaskContext.SubmitTask(gmodels.TASK_HKT_EP_DISCONNECTED, token)

	if subErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Unable to submit endpoint disconnect task", subErr)
	}
}
