import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { USER_PRODUCT_MAPPING_STATE } from '../../app-common-constants';
import { EnvironmentProvider } from '../../environment-provider';
import { PayloadResponse } from '../../models/api/payload-models';
import { EnumDataSourceItem } from '../../models/ui/enum-datasource-item';
import { UserLoginInfoResponse, UserAddDetailsRequest } from '../../models/api/user-models';
import { ServerApiInterfaceService } from '../../services/api/server-api-interface.service';
import { EnumNumberDatasource } from '../../utility/enum-number-datasource';
@Injectable()
export class UserSharedService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    getLoginInfo(implicitErrorHandling = true): Observable<PayloadResponse<UserLoginInfoResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/v1/user/info/login', implicitErrorHandling);
    }

    updateUserDetails(userAddDetailsRequest: UserAddDetailsRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/cu/v1/user/update/details',
            userAddDetailsRequest, implicitErrorHandling);
    }

    getUcpmStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('USER_PRODUCT_MAPPING_STATE_', USER_PRODUCT_MAPPING_STATE);
    }

    getUcpmState(state: number) {
        return 'USER_PRODUCT_MAPPING_STATE_' + state;
    }
}
