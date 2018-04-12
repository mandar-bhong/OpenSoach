import { CUSTOMER_STATE } from '../../../../shared/app-common-constants';

export class CustomerAddRequest {
    custname: string;
    corpid: number;
    custstate: CUSTOMER_STATE;
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

export class CustomerFilterRequest {
    custname: string;
    custstate: CUSTOMER_STATE;
}

export class CustomerDataListingModel {
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
