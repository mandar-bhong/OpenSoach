import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { CUSTOMER_STATE } from '../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import {
    CustomerAddDetailsRequest,
    CustomerDetailsResponse,
    CustomerMasterResponse,
    CustomerMasterUpdateRequest,
} from '../../../shared/models/api/customer-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import {
    CustomerAddRequest,
    CustomerAssociateProductListItemResponse,
    CustomerAssociateProductRequest,
    CustomerAssociateProductUpdateRequest,
    CustomerDataListingItemResponse,
    CustomerFilterRequest,
    CustomerListItemResponse,
    CustomerRoleidListItemResponse,
    CustomerRoleListRequest,
    CustomerServiceAssociateListResponse,
    CustomerServiceAssociateUpdateRequest
} from '../models/api/customer-models';

@Injectable()
export class CustomerService extends ListingService<CustomerFilterRequest, CustomerDataListingItemResponse> {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }

    getDataList(dataListRequest: DataListRequest<CustomerFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<CustomerDataListingItemResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/customer/list',
            dataListRequest, implicitErrorHandling);
    }

    addCustomer(customerAddRequest: CustomerAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/add',
            customerAddRequest, implicitErrorHandling);
    }

    updateCustomerDetails(customerAddDetailsRequest: CustomerAddDetailsRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/update/details',
            customerAddDetailsRequest, implicitErrorHandling);
    }

    getCustomerList(dataListFilter: DataListRequest<CustomerFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<CustomerDataListingItemResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/customer/list',
            dataListFilter, implicitErrorHandling);
    }

    getCustomerDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/customer/info/details',
            request, implicitErrorHandling);
    }

    getCustomerStates(): EnumDataSourceItem<number>[] {
        return EnumNumberDatasource.getDataSource('CUSTOMER_STATE_', CUSTOMER_STATE);
    }

    getCustomerState(state: number) {
        return 'CUSTOMER_STATE_' + state;
    }

    associateCustomerToProduct(request: CustomerAssociateProductRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/associate/product',
            request, implicitErrorHandling);
    }

    updateAssociateCustomerToProduct(request: CustomerAssociateProductUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/productassociation/update',
            request, implicitErrorHandling);
    }

    getCustomerProductAssociation(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerAssociateProductListItemResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(
            EnvironmentProvider.baseurl + '/api/osu/v1/customer/list/productassociation',
            request, implicitErrorHandling);
    }
    getCustomerNameList(implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerListItemResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/osu/v1/customer/list/short',
            implicitErrorHandling);
    }
    getCustRoleDataList(request: CustomerRoleListRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerRoleidListItemResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/v1/urole/list',
            request, implicitErrorHandling);
    }

    getCustomerMaster(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerMasterResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/customer/info/master',
            request, implicitErrorHandling);
    }
    updateCustomerMaster(customerMasterUpdateRequest: CustomerMasterUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/update',
        customerMasterUpdateRequest, implicitErrorHandling);
    }
    getCustomerServiceProductAssociation(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CustomerServiceAssociateListResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(
            EnvironmentProvider.baseurl + '/api/osu/v1/customer/association/servicepoint',
            request, implicitErrorHandling);
    }
    updateCustomerServicePointProduct(request: CustomerServiceAssociateUpdateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/customer/association/servicepoint/update',
        request, implicitErrorHandling);
}
}
