export class DataDBModel {
	patientid: number;
	admissionid: number;
	cpmid: number;
	usrid: number;
	patientregno: string;
	bedno: string;
	fname: string;
	lname: string;
	spid: number;
	spname: string;
	monitored: number;
	upmmid: number;
	upmmidpatientid: number;
	upmmidspid: number;
}

export class UiViewModel {
	dbmodel: DataDBModel;
	checked: any;
	isDisabled: boolean;
}

export class ApiRequestModel {
	filter: FilterRequest;
	page: number;
	limit: number;
	orderby: string;
	orderdirection: string;

}
export class FilterRequest {
	fname: string;
	bedno: string;
	lname: string;
	spname: string;
	mobno: string;
	patientregno: string;
	status: number
}
export class MonitoredRequest {
	usrid: number;
	spid: number;
	patientid: number;
}