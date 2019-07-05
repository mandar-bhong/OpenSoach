package models

func (r *ExecutionContext) GetNodeDBConnection() string {
	return r.SessionInfo.Product.NodeDbConn
}

func (r *DeviceExecutionContext) GetNodeDBConnection() string {
	return r.DeviceSessionInfo.Product.NodeDbConn
}

func (r *ExecutionContext) GetCPMID() int64 {
	return r.SessionInfo.Product.CustProdID
}

func (r *DeviceExecutionContext) GetCPMID() int64 {
	return r.DeviceSessionInfo.CpmID
}

func (r *DeviceUserExecutionContext) GetNodeDBConnection() string {
	return r.DeviceUserSessionInfo.Product.NodeDbConn
}

func (r *DeviceUserExecutionContext) GetCPMID() int64 {
	return r.DeviceUserSessionInfo.Product.CustProdID
}
