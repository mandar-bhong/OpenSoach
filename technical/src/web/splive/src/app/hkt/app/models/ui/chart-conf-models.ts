import {
    ServiceConfigurationRequest,
    ServiceConfigurationResponse,
    ServiceConfigurationUpdateRequest,
} from '../../../../prod-shared/models/api/service-configuration-models';
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

