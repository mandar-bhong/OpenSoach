import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';

import { EnvironmentProvider } from '../../../shared/environment-provider';
import { PayloadResponse } from '../../../shared/models/api/payload-models';
import { ServerApiInterfaceService } from '../../../shared/services/api/server-api-interface.service';
import { ServiceInstanceTransactionResponse, ServiceTxnRequest } from '../../models/api/service-txn-models';
import { PatientTxnRequest } from '../../../hpft/app/models/api/patient-data-models';

@Injectable()
export class SpServiceTxnService {
    constructor(
        private serverApiInterfaceService: ServerApiInterfaceService) {
    }

    getServiceTransactions(request: ServiceTxnRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<ServiceInstanceTransactionResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/transaction/list',
            request, implicitErrorHandling);
    }
    getPatientTransactions(request: PatientTxnRequest, implicitErrorHandling = true):
        Observable<PayloadResponse<ServiceInstanceTransactionResponse[]>> {
        return this.serverApiInterfaceService.getWithQueryParams(EnvironmentProvider.appbaseurl + '/api/v1/service/transaction/list',
            request, implicitErrorHandling);
    }
}

