import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { DataListRequest, DataListResponse } from '../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ListingService } from '../../../shared/services/listing.service';
import {
    CorporateDataListingItemResponse,
    CorporateFilterRequest,
    CorporateShortDataResponse,
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
}
