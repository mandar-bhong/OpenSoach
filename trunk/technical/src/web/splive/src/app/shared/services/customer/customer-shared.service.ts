import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../environment-provider';
import { CustomerAddDetailsRequest, CustomerLoginInfoResponse } from '../../models/api/customer-models';
import { PayloadResponse } from '../../models/api/payload-models';
import { ServerApiInterfaceService } from '../../services/api/server-api-interface.service';

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
}
