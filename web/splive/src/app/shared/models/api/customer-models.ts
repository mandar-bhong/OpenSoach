import { CUSTOMER_STATE } from '../../app-common-constants';

export class CustomerLoginInfoResponse {
    corpname: string;
    custname: string;
}

export class CustomerDetailsResponse {
    custid: number;
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
    createdon: Date;
    updateon: Date;
}

export class CustomerMasterResponse {
    custid: number;
    corpid: number;
    custname: string;
    custstate: CUSTOMER_STATE;
    custstatesince: Date;
    createdon: Date;
    updateon: Date;
}

export class CustomerMasterUpdateRequest {
    custid: number;
    custname: string;
    custstate: number;
}

export class CorporateDetailsResponse {
    corpid: number;
    corpname: string;
    corpmobileno: string;
    corpemailid: string;
    corplandlineno: string;
    createdon: Date;
    updateon: Date;
}

export class CustomerAddDetailsRequest {
    custid: number;
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

