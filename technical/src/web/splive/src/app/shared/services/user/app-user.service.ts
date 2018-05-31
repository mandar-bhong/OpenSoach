import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { USER_GENDER, USER_PRODUCT_MAPPING_STATE, USER_STATE } from '../../app-common-constants';
import { EnvironmentProvider } from '../../environment-provider';
import { PayloadResponse } from '../../models/api/payload-models';
import { UserLoginInfoResponse } from '../../models/api/user-models';
import { EnumDataSourceItem } from '../../models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../services/api/server-api-interface.service';
import { EnumNumberDatasource } from '../../utility/enum-number-datasource';

@Injectable()
export class AppUserService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    getLoginInfo(implicitErrorHandling = true): Observable<PayloadResponse<UserLoginInfoResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/v1/user/info/login', implicitErrorHandling);
    }

    getUcpmStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('USER_PRODUCT_MAPPING_STATE_', USER_PRODUCT_MAPPING_STATE);
    }

    getUcpmState(state: number) {
        return 'USER_PRODUCT_MAPPING_STATE_' + state;
    }
    getUserStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('USER_STATE_', USER_STATE);
    }

    getUserState(state: number) {
        return 'USER_STATE_' + state;
    }
    getUsersGender(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('USER_GENDER_', USER_GENDER);
    }

    getUserGenders(genders: number) {
        return 'USER_GENDER_' + genders;
    }
}
