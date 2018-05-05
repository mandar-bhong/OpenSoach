import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { CUSTOMER_PRODUCT_MAPPING_STATE } from '../../app-common-constants';
import { EnvironmentProvider } from '../../environment-provider';
import { CustomerAddDetailsRequest, CustomerLoginInfoResponse } from '../../models/api/customer-models';
import { PayloadResponse } from '../../models/api/payload-models';
import { EnumDataSourceItem } from '../../models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../services/api/server-api-interface.service';
import { EnumNumberDatasource } from '../../utility/enum-number-datasource';

@Injectable()
export class CustomerSharedService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    getCusomerLoginInfo(implicitErrorHandling = true): Observable<PayloadResponse<CustomerLoginInfoResponse>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/cu/v1/customer/info/login', implicitErrorHandling);
    }

    updateCustomerDetails(customerAddDetailsRequest: CustomerAddDetailsRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/cu/v1/customer/update/details',
            customerAddDetailsRequest, implicitErrorHandling);
    }

    getCpmStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('CUSTOMER_PRODUCT_MAPPING_STATE_', CUSTOMER_PRODUCT_MAPPING_STATE);
    }

    getCpmState(state: number) {
        return 'CUSTOMER_PRODUCT_MAPPING_STATE_' + state;
    }
}
