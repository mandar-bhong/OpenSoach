package processor

import (
	"fmt"

	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	pcepproc "opensoach.com/prodcore/endpoint/processor"
)

func AuthProcessor(epmodel *gmodels.EndPointToServerTaskModel) *gmodels.PacketProcessingResult {

	packetProcessingResult := &gmodels.PacketProcessingResult{}
	//chnIDAuthData

	fmt.Printf("Repo ctx : %#v \n", repo.Instance().Context)
	pcepproc.AuthorizeDevice(repo.Instance().Context.Master.Cache, chnIDAuthData, epmodel)

	return packetProcessingResult
}
