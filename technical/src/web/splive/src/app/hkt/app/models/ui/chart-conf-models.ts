import { ServiceConfigurationRequest } from '../../../../shared/models/api/service-configuration-models';

export class ChartConfigurationModel {
    servconfid: number;
    spcid: number;
    conftypecode: string;
    servconfname: string;
    shortdesc: string;
    servconf: string;
    variableserviceconfig: VariableServiceConfig;

    copyTo(serviceConfigurationRequest: ServiceConfigurationRequest) {
        serviceConfigurationRequest.servconfid = this.servconfid;
        serviceConfigurationRequest.spcid = this.spcid;
        serviceConfigurationRequest.conftypecode = this.conftypecode;
        serviceConfigurationRequest.servconfname = this.servconfname;
        serviceConfigurationRequest.shortdesc = this.shortdesc;
        serviceConfigurationRequest.servconf = JSON.stringify(this.variableserviceconfig);
        console.log('serviceConfigurationRequest', serviceConfigurationRequest);
    }
}

export class VariableServiceConfig {
    timeconfig: TimeConfigurationModel;
    taskconfig: TaskConfigurationModel;
}

export class TimeConfigurationModel {
    starttime: number;
    endtime: number;
    interval: number;
}

export class TaskConfigurationModel {
    taskname: string;
    taskorder: string;
}
