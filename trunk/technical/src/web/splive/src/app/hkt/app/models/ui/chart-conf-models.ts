import { ServiceConfigurationRequest } from '../../../../prod-shared/models/api/service-configuration-models';
import { TaskTemplateResponse, TaskTemplateRequest, ServiceConfigurationUpdateRequest } from '../api/chart-conf-models';

export class ChartConfigurationModel {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    variableconf: VariableChartConfModel;

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
