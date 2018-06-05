import {
    ServiceConfigurationRequest,
    ServiceConfigurationResponse,
    ServiceConfigurationUpdateRequest,
} from '../../../../prod-shared/models/api/service-configuration-models';
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
    }
    copyToUpdate(serviceConfigurationUpdateRequest: ServiceConfigurationUpdateRequest) {
        serviceConfigurationUpdateRequest.servconfid = this.servconfid;
        serviceConfigurationUpdateRequest.servconfname = this.servconfname;
        serviceConfigurationUpdateRequest.shortdesc = this.shortdesc;
        serviceConfigurationUpdateRequest.servconf = JSON.stringify(this.variableconf);
    }

    copyFrom(response: ServiceConfigurationResponse) {
        this.servconfid = response.servconfid;
        this.spcid = response.spcid;
        this.conftypecode = response.conftypecode;
        this.servconfname = response.servconfname;
        this.shortdesc = response.shortdesc;
        this.servconf = response.servconf;
        Object.assign(this.variableconf, JSON.parse(response.servconf));
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
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    variableconf: VariableChartConfModel;
    spid: number;
    txns: ChartTransactionModel[];
    timeslots: ChartTimeSlot[];
    tasktxnslotmap: Map<string, ChartTxnSlot[]>;
    startdate: Date;
    enddate: Date;
    selecteddateoption = '0'; // 0: Today, 1: Yesterday, 2: Selected Date
    constructor() {
    }

    copyFromConfiguration(response: ServiceConfigurationResponse) {
        this.servconfid = response.servconfid;
        this.spcid = response.spcid;
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
        this.status = response.status;
        this.txndate = response.txndate;
        this.txndata = new ChartTransactionDataModel();
        Object.assign(this.txndata, JSON.parse(response.txndata));
    }
}

export class ChartTransactionDataModel {
    taskname: string;
    slotstarttime: number;
    slotendtime: number;
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
