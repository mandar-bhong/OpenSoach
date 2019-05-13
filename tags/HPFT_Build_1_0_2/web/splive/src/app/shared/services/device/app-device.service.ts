import { Injectable } from '@angular/core';

import { DEVICE_STATE } from '../../app-common-constants';
import { EnumDataSourceItem } from '../../models/ui/enum-datasource-item';
import { EnumNumberDatasource } from '../../utility/enum-number-datasource';

@Injectable()
export class AppDeviceService {
    constructor() {
    }

    getDeviceStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('DEVICE_STATE_', DEVICE_STATE);
    }

    getDeviceState(state: number) {
        return 'DEVICE_STATE_' + state;
    }
}
