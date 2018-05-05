import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ProductListItemResponse } from '../models/api/product-models';

@Injectable()
export class ProductService {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }

    getDataList(implicitErrorHandling = true):
        Observable<PayloadResponse<ProductListItemResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/osu/v1/product/list',
            implicitErrorHandling);
    }
}
