import {
    ServiceConfigurationRequest,
    ServiceConfigurationResponse,
    ServiceConfigurationUpdateRequest,
    ServicePointWithConfigurationResponse
} from '../../../../prod-shared/models/api/service-configuration-models';
import { ServicepointAssociateRequest, ServicepointListResponse } from '../../../../prod-shared/models/api/servicepoint-models';
import { ServiceInstanceTransactionResponse } from '../../../../prod-shared/models/api/service-txn-models';
import { TaskTemplateRequest, TaskTemplateResponse } from '../api/chart-conf-models';

export class ChartConfigurationModel {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    variableconf: VariableChartConfModel;
    spid: number;
    mode: number; // 0: new, 1: edit, 2: create new and associate

    copyTo(serviceConfigurationRequest: ServiceConfigurationRequest) {
        serviceConfigurationRequest.servconfid = this.servconfid;
        serviceConfigurationRequest.spcid = this.spcid;
        serviceConfigurationRequest.conftypecode = this.conftypecode;
        serviceConfigurationRequest.servconfname = this.servconfname;
        serviceConfigurationRequest.shortdesc = this.shortdesc;
        serviceConfigurationRequest.servconf = JSON.stringify(this.variableconf);
        console.log('serviceConfigurationRequest', serviceConfigurationRequest);
    }
    copyToUpdate(serviceConfigurationUpdateRequest: ServiceConfigurationUpdateRequest) {
        serviceConfigurationUpdateRequest.servconfid = this.servconfid;
        serviceConfigurationUpdateRequest.servconfname = this.servconfname;
        serviceConfigurationUpdateRequest.shortdesc = this.shortdesc;
        serviceConfigurationUpdateRequest.servconf = JSON.stringify(this.variableconf);
    }
    copyToAssociateRequest(request: ServicepointAssociateRequest) {
        request.servconfid = this.servconfid;
        request.spid = this.spid;
    }
    copyFrom(serviceConfigurationResponse: ServiceConfigurationResponse) {
        this.servconfid = serviceConfigurationResponse.servconfid;
        this.spcid = serviceConfigurationResponse.spcid;
        this.conftypecode = serviceConfigurationResponse.conftypecode;
        this.servconfname = serviceConfigurationResponse.servconfname;
        this.shortdesc = serviceConfigurationResponse.shortdesc;
        this.servconf = serviceConfigurationResponse.servconf;
        this.variableconf = new VariableChartConfModel();
        Object.assign(this.variableconf, JSON.parse(serviceConfigurationResponse.servconf));
    }
}

export class VariableChartConfModel {
    timeconf: ChartTimeConfModel;
    taskconf: ChartTaskListConfModel;
}

export class ChartTaskListConfModel {
    tasks: ChartTaskModel[];
}

export class ChartTimeConfModel {
    starttime: number;
    endtime: number;
    interval: number;
}

export class ChartTaskModel {
    taskname: string;
    spcid: number;
    shortdec: string;
    copyFrom(taskTemplateResponse: TaskTemplateResponse) {
        this.taskname = taskTemplateResponse.taskname;
        this.spcid = taskTemplateResponse.spcid;
    }
    copyTo(taskTemplateRequest: TaskTemplateRequest) {
        taskTemplateRequest.spcid = this.spcid;
    }
}

export class ChartDataViewModel {
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    variableconf: VariableChartConfModel;
    txns: ChartTransactionModel[];
    timeslots: ChartTimeSlot[];
    tasktxnslotmap: Map<string, ChartTxnSlot[]>;
    startdate: Date;
    enddate: Date;
    selecteddateoption = '0'; // 0: Today, 1: Yesterday, 2: Selected Date
    splist: ServicePointWithConfigurationResponse[];
    selectedsp: ServicePointWithConfigurationResponse;
    constructor() {
    }

    copyFromConfiguration(response: ServiceConfigurationResponse) {
        this.conftypecode = response.conftypecode;
        this.servconfname = response.servconfname;
        this.shortdesc = response.shortdesc;
        this.servconf = response.servconf;
        this.variableconf = new VariableChartConfModel();
        Object.assign(this.variableconf, JSON.parse(response.servconf));
    }

    copyFromTransactions(response: ServiceInstanceTransactionResponse[]) {
        this.txns = [];
        response.forEach(item => {
            const chartTransactionModel = new ChartTransactionModel();
            chartTransactionModel.copyFrom(item);
            this.txns.push(chartTransactionModel);
        });
    }
}

export class ChartTransactionModel {
    servintxnid: number;
    servinid: number;
    fopcode: string;
    fopname: string;
    status: number;
    txndate: Date;
    txndata: ChartTransactionDataModel;

    copyFrom(response: ServiceInstanceTransactionResponse) {
        this.servintxnid = response.servintxnid;
        this.servinid = response.servinid;
        this.fopcode = response.fopcode;
        this.fopname = response.fopname;
        this.status = response.status;
        this.txndate = new Date(response.txndate);
        this.txndata = new ChartTransactionDataModel();
        Object.assign(this.txndata, JSON.parse(response.txndata));
    }
}

export class ChartTransactionDataModel {
    taskname: string;
    slotstarttime: number;
    slotendtime: number;
    value: number;
    comment: string;
}

export class ChartTimeSlot {
    slotstarttime: number;
    slotendtime: number;
    slotdisplaytext: string;
}

export class ChartTxnSlot {
    slot: ChartTimeSlot;
    txn: ChartTransactionModel;
}

export class PatientDayWiseTxn {
    txn: ChartTransactionModel[];
    day: Date;
}
