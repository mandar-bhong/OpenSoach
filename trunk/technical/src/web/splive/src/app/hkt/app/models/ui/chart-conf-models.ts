import { ServiceConfigurationRequest } from '../../../../prod-shared/models/api/service-configuration-models';
import { TaskTemplateModel } from '../api/chart-conf-models';

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

    copyFrom(taskTemplateModel: TaskTemplateModel) {
        this.taskname = taskTemplateModel.taskname;
    }
}
