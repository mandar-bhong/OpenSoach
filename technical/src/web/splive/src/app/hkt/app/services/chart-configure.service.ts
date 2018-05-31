import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { AppDataStoreService } from '../../../shared/services/app-data-store/app-data-store-service';
import { APP_DATA_STORE_KEYS } from '../app-constants';
import {
    CategoriesShortDataResponse, TaskTemplateRequest,
    TaskTemplateResponse, ChartsDetailsResponse
} from '../models/api/chart-conf-models';
import {
    ChartConfigurationModel,
    ChartTaskListConfModel,
    ChartTaskModel,
    ChartTimeConfModel,
    VariableChartConfModel,
} from '../models/ui/chart-conf-models';

@Injectable()
export class ChartConfigureService {
    constructor(private appDataStoreService: AppDataStoreService,
        private serverApiInterfaceService: ServerApiInterfaceService
    ) { }

    getDataModel(): ChartConfigurationModel {
        return this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.CHART_CONFIG).
            getObject<ChartConfigurationModel>(APP_DATA_STORE_KEYS.CHART_CONFIG, ChartConfigurationModel);
    }

    setDataModel(dataModel: ChartConfigurationModel) {
        console.log('setDataModel', dataModel);
        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.CHART_CONFIG).
            setObject<ChartConfigurationModel>(APP_DATA_STORE_KEYS.CHART_CONFIG, dataModel);
    }

    createDataModel() {
        const dataModel = new ChartConfigurationModel();
        dataModel.variableconf = new VariableChartConfModel();
        dataModel.variableconf.timeconf = new ChartTimeConfModel();
        dataModel.variableconf.taskconf = new ChartTaskListConfModel();
        dataModel.variableconf.taskconf.tasks = new Array<ChartTaskModel>();
        this.setDataModel(dataModel);
        return dataModel;
    }

    removeDataModel() {
        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.CHART_CONFIG).
            removeObject(APP_DATA_STORE_KEYS.CHART_CONFIG);
    }
    getCategoriesShortDataList(implicitErrorHandling = true):
        Observable<PayloadResponse<CategoriesShortDataResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.appbaseurl + '/api/v1/servicepoint/category/list/short',
            implicitErrorHandling);
    }
    addTask(deviceAddRequest: TaskTemplateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.appbaseurl + '/api/v1/task/add',
            deviceAddRequest, implicitErrorHandling);
    }
    getTaskDataList(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<TaskTemplateResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/task/list',
            request, implicitErrorHandling);
    }
    getChartConfigureDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<ChartsDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/config/info',
            request, implicitErrorHandling);
    }
}
