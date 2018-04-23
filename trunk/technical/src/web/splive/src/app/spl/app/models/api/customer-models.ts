import { CUSTOMER_STATE } from '../../../../shared/app-common-constants';

export class CustomerAddRequest {
    custname: string;
    corpid: number;
    custstate: CUSTOMER_STATE;
}

export class CustomerFilterRequest {
    custname: string;
    custstate: CUSTOMER_STATE;
    corpid: number;
}

export class CustomerDataListingItemResponse {
    custid: number;
    corpid: number;
    custname: string;
    corpname: string;
    poc1name: string;
    poc1emailid: string;
    poc1mobileno: string;
    createdon: Date;
    updateon: Date;
}
