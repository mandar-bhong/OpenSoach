package models

import "bytes"

type IDBConnection interface {
	GetNodeDBConnection() string
	GetCPMID() int64
}

type ExecutionContext struct {
	SessionInfo             UserSessionInfo
	SelectedCustomerProduct int64
	RequestJSON             string
	SessionToken            string
}

type DeviceExecutionContext struct {
	DeviceSessionInfo       DeviceTokenModel
	SelectedCustomerProduct int64
	RequestJSON             string
	SessionToken            string
}

type DeviceUserExecutionContext struct {
	DeviceUserSessionInfo   DeviceUserSessionInfo
	SelectedCustomerProduct int64
	RequestJSON             string
	SessionToken            string
}

type ServerListingResultModel struct {
	RecordCount int
	RecordList  interface{}
}

type PacketProcessingTaskModel struct {
	Token         string `json:"token"`
	EPTaskListner string `json:"eptsklistner"`
	ChannelID     int    `json:"chnid"`
	Message       []byte `json:"msg"`
}

type PacketProcessingTaskResult struct {
	IsSuccess  bool            `json:"issuccess"`
	StatusCode int             `json:"status"`
	AckPayload []*DevicePacket `json:"ackpayload"`
	ChannelID  int             `json:"chnid"`
}

type ExcelData struct {
	SheetName  string
	StartCell  string
	IsVertical bool
	Headers    []string
	Data       [][]string
	StartDate  string
	EndDate    string
}

type PdfDataItem struct {
	IsSummary bool
	Headers   []string
	Data      [][]string
	ColsWidth []float64
	ColsAlign []string
}

type PdfDataModel struct {
	ReportName string
	StartDate  string
	EndDate    string
	PdfData    []PdfDataItem
}

type HTMLPDFDataModel struct {
	TemplatePath  string
	TemplateData  interface{}
	PDFBuffer     *bytes.Buffer
	PDFOutputPath string
	HeaderPath    string
}
