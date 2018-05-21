import { Injectable } from '@angular/core';

import { AppDataStoreService } from '../../../shared/services/app-data-store/app-data-store-service';
import { APP_DATA_STORE_KEYS } from '../app-constants';
import {
    ChartConfigurationModel,
    ChartTaskListConfModel,
    ChartTaskModel,
    ChartTimeConfModel,
    VariableChartConfModel,
} from '../models/ui/chart-conf-models';

@Injectable()
export class ChartConfigureService {
    constructor(private appDataStoreService: AppDataStoreService) { }

    getDataModel(): ChartConfigurationModel {
        return this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.CHART_CONFIG).
            getObject<ChartConfigurationModel>(APP_DATA_STORE_KEYS.CHART_CONFIG);
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
    }

    removeDataModel() {
        this.appDataStoreService.getDataStore(APP_DATA_STORE_KEYS.CHART_CONFIG).
            removeObject(APP_DATA_STORE_KEYS.CHART_CONFIG);
    }
}
