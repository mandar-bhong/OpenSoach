import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordAddResponse } from '../../../shared/models/api/common-models';
import { CustomerAddDetailsRequest } from '../../../shared/models/api/customer-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { CustomerAddRequest, CustomerFilterRequest, CustomerDataListingItemResponse } from '../models/api/customer-models';

@Injectable()
export class CustomerService {

    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }
    addCustomer(customerAddRequest: CustomerAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordAddResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/add',
            customerAddRequest, implicitErrorHandling);
    }

    updateCustomerDetails(customerAddDetailsRequest: CustomerAddDetailsRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/update/details',
            customerAddDetailsRequest, implicitErrorHandling);
    }

    getCustomerList(customerFilterRequest: CustomerFilterRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerDataListingItemResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/customer/list',
        customerFilterRequest, implicitErrorHandling);
    }
}
