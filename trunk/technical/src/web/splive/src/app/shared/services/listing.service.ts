import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';

import { DataListRequest, DataListResponse } from '../models/api/data-list-models';
import { PayloadResponse } from '../models/api/payload-models';

@Injectable()
export abstract class ListingService<Request, Response> {
    dataListSubject = new Subject<Request>();
    dataListSubjectTrigger(value: Request) {
        this.dataListSubject.next(value);
    }

    abstract getDataList(dataListRequest: DataListRequest<Request>, implicitErrorHandling):
        Observable<PayloadResponse<DataListResponse<Response>>>;
}
