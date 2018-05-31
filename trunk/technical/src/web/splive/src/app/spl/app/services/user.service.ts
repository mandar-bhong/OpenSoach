import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { USER_CATEGORY, USER_GENDER } from '../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import {
  UserAddDetailsRequest,
  UserAddRequest,
  UserAssociateProductListItemResponse,
  UserAssociateProductRequest,
  UserAssociateProductUpdateRequest,
  UserDataListResponse,
  UserDetailsResponse,
  UserMasterResponse,
  UserMasterUpdateRequest,
  UserRoleidListItemResponse,
} from '../../../shared/models/api/user-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import { UserFilterRequest } from '../models/api/user-models';


@Injectable()
export class UserService extends ListingService<UserFilterRequest, UserDataListResponse> {
  constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
    super();
  }

  getDataList(dataListRequest: DataListRequest<UserFilterRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<UserDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/user/list',
      dataListRequest, implicitErrorHandling);
  }

  addUser(userAddRequest: UserAddRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<RecordIDResponse>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/user/add',
      userAddRequest, implicitErrorHandling);
  }

  updateUserDetails(userAddDetailsRequest: UserAddDetailsRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/user/update/details',
      userAddDetailsRequest, implicitErrorHandling);
  }

  getUserList(dataListFilter: DataListRequest<UserFilterRequest>, implicitErrorHandling = true):
    Observable<PayloadResponse<DataListResponse<UserDataListResponse>>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/user/list',
      dataListFilter, implicitErrorHandling);
  }

  getUserDetails(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<UserDetailsResponse>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/user/info/details',
      request, implicitErrorHandling);
  }

  getCorporateRoleId(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<UserDataListResponse[]>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/urole/list',
      request, implicitErrorHandling);
  }

  associateUserToProduct(request: UserAssociateProductRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<RecordIDResponse>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/user/associate/customer',
      request, implicitErrorHandling);
  }

  updateAssociateUserToProduct(request: UserAssociateProductUpdateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/user/productassociation/update',
      request, implicitErrorHandling);
  }

  getUserProductAssociation(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<UserAssociateProductListItemResponse[]>> {
    return this.serverApiInterfaceService.getWithQueryParams(
      EnvironmentProvider.baseurl + '/api/osu/v1/user/list/productassociation',
      request, implicitErrorHandling);
  }

  getUsersCategories(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('USER_CATEGORY_', USER_CATEGORY);
  }

  getUserCategories(categories: number) {
    return 'USER_CATEGORY_' + categories;
  }

  getRoleDataList(implicitErrorHandling = true):
    Observable<PayloadResponse<UserRoleidListItemResponse[]>> {
    return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/osu/v1/urole/list',
      implicitErrorHandling);
  }

  getUserEdit(request: RecordIDRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<UserMasterResponse>> {
    return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/user/info/master',
      request, implicitErrorHandling);
  }

  updateUserEdit(corporateUpadteRequest: UserMasterUpdateRequest, implicitErrorHandling = true):
    Observable<PayloadResponse<null>> {
    return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/user/update',
      corporateUpadteRequest, implicitErrorHandling);
  }

}
