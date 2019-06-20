import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { DbInstanceListItemResponse } from '../models/api/db-instance-models';

@Injectable()
export class DBInstanceService {
    constructor(private serverApiInterfaceService: ServerApiInterfaceService) { }

    getDataList(implicitErrorHandling = true):
        Observable<PayloadResponse<DbInstanceListItemResponse[]>> {
        return this.serverApiInterfaceService.get(EnvironmentProvider.baseurl + '/api/osu/v1/dbinstance/list',
            implicitErrorHandling);
    }
}
