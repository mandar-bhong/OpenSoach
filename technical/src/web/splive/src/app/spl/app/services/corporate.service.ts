import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { RecordIDRequest, RecordIDResponse } from '../../../shared/models/api/common-models';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import {
    CorpDetailsResponse,
    CorporateAddRequest,
    CorporateDataListingItemResponse,
    CorporateFilterRequest,
    CorporateShortDataResponse,
    CorporateUpdateRequest,
} from '../models/api/corporate-models';

@Injectable()
export class CorporateService extends ListingService<CorporateFilterRequest, CorporateDataListingItemResponse> {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) {
        super();
    }

    getDataList(dataListRequest: DataListRequest<CorporateFilterRequest>, implicitErrorHandling = true):
        Observable<PayloadResponse<DataListResponse<CorporateDataListingItemResponse>>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/corporate/list',
            dataListRequest, implicitErrorHandling);
    }

    getCorporateShortDataList(implicitErrorHandling = true):
        Observable<PayloadResponse<CorporateShortDataResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/osu/v1/corporate/list/short',
            implicitErrorHandling);
    }
    addCorporate(customerAddRequest: CorporateAddRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<RecordIDResponse>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/corporate/add',
            customerAddRequest, implicitErrorHandling);
    }

    getCorporateDetails(request: RecordIDRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<CorpDetailsResponse>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.baseurl + '/api/osu/v1/corporate/info/master',
            request, implicitErrorHandling);
    }

    updateCorporateDetails(corporateUpadteRequest: CorporateUpdateRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<null>> {
        return this.serverApiInterfaceService.post(EnvironmentProvider.baseurl + '/api/osu/v1/corporate/update',
            corporateUpadteRequest, implicitErrorHandling);
    }
}
