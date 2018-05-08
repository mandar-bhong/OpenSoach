import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { USER_CATEGORY, USER_STATE, USER_GENDER } from '../../../shared/app-common-constants';
import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { UserAddDetailsRequest, UserDetailsResponse, UserDataListResponse } from '../../../shared/models/api/user-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { EnumDataSourceItem } from '../../../shared/models/ui/enum-datasource-item';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import { EnumNumberDatasource } from '../../../shared/utility/enum-number-datasource';
import { UserFilterRequest } from '../models/api/user-models';
import { UserAddRequest } from '../../../shared/models/api/user-models';


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

  getUserStates(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('USER_STATE_', USER_STATE);
  }

  getUserState(state: number) {
    return 'USER_STATE_' + state;
  }

  getUsersCategories(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('USER_CATEGORY_', USER_CATEGORY);
  }

  getUserCategories(categories: number) {
    return 'USER_CATEGORY_' + categories;
  }

  getUsersGender(): EnumDataSourceItem<number>[] {
    return EnumNumberDatasource.getDataSource('USER_GENDER_', USER_GENDER);
  }

  getUserGenders(genders: number) {
    return 'USER_GENDER_' + genders;
  }

}
