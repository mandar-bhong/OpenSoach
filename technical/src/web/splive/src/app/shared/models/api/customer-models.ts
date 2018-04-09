import { CUSTOMER_STATE } from '../../app-common-constants';

export class CustomerLoginInfoResponse {
    corpname: string;
    custname: string;
}

export class CustomerAddRequest {
    custname: string;
    custstate: CUSTOMER_STATE;
}

export class CustomerAddDetailsRequest {
    custid: number;
    corpid: number;
    poc1name: string;
    poc1emailid: string;
    poc1mobileno: string;
    poc2name: string;
    poc2emailid: string;
    poc2mobileno: string;
    address: string;
    addressstate: string;
    addresscity: string;
    addresspincode: string;
}
