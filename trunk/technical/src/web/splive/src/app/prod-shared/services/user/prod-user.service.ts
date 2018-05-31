import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import {
    UserAddRequest,
    UserDataListResponse,
    UserDetailsResponse,
    UserFilterRequest,
    UserRoleidListItemResponse,
    ProductcodeRequest,
    UserMasterResponse,
    UserAddDetailsRequest,
    UserMasterUpdateRequest
} from '../../models/api/user-models';


@Injectable()
export class ProdUserService extends ListingService<UserFilterRequest, UserDataListResponse> {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }
    getDataList(dataListRequest: DataListRequest<UserFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<UserDataListResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/v1/user/list',
            dataListRequest, implicitErrorHandling);
    }

    addUser(userAddRequest: UserAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/cu/role/v1/user/add',
            userAddRequest, implicitErrorHandling);
    }

    getRoleDataList(request: ProductcodeRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<UserRoleidListItemResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/v1/urole/list',
            request, implicitErrorHandling);
    }

    getUserMasterDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<UserMasterResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/role/v1/user/info/master',
            request, implicitErrorHandling);
    }
    getUserDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<UserDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/cu/role/v1/user/info/details',
            request, implicitErrorHandling);
    }
    updateUserDetails(userAddDetailsRequest: UserAddDetailsRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/cu/role/v1/user/update/details',
            userAddDetailsRequest, implicitErrorHandling);
    }
    updateUserEdit(corporateUpadteRequest: UserMasterUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/cu/role/v1/user/update',
            corporateUpadteRequest, implicitErrorHandling);
    }
    getRoleMasterDataList(request: ProductcodeRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<UserRoleidListItemResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/v1/urole/list',
            request, implicitErrorHandling);
    }
}
