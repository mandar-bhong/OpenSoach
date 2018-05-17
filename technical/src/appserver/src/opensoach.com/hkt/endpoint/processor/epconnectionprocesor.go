package processor

func EPOnConnectProcessExecutor(chnid int) {
	//TODO: Relation should be maintain that if device connected and does not send auth command
	// for longer duration then device should be disconnected
}

func EPOnDisConnectProcessExecutor(chnid int) {

	//TODO From the channel id find the device information and send it with device id

	//repo.Instance().ProdTaskContext.SubmitTask(gmodels.TASK_HKT_EP_DISCONNECTED, "jsonMsg")

}
