import { CustomerAddDetailsRequest, CustomerDetailsResponse } from '../api/customer-models';
import { CustomerLoginInfoResponse } from '../api/customer-models';

export class CustomerDetailsModel {
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

    copyTo(customerAddDetailsRequest: CustomerAddDetailsRequest) {
        customerAddDetailsRequest.custid = this.custid;
        customerAddDetailsRequest.poc1name = this.poc1name;
        customerAddDetailsRequest.poc1emailid = this.poc1emailid;
        customerAddDetailsRequest.poc1mobileno = this.poc1mobileno;
        customerAddDetailsRequest.poc2name = this.poc2name;
        customerAddDetailsRequest.poc2emailid = this.poc2emailid;
        customerAddDetailsRequest.poc2mobileno = this.poc2mobileno;
        customerAddDetailsRequest.address = this.address;
        customerAddDetailsRequest.addressstate = this.addressstate;
        customerAddDetailsRequest.addresscity = this.addresscity;
        customerAddDetailsRequest.addresspincode = this.addresspincode;
    }

    copyFrom(customerDetailsResponse: CustomerDetailsResponse) {
        this.custid = customerDetailsResponse.custid;
        this.poc1name = customerDetailsResponse.poc1name;
        this.poc1emailid = customerDetailsResponse.poc1emailid;
        this.poc1mobileno = customerDetailsResponse.poc1mobileno;
        this.poc2name = customerDetailsResponse.poc2name;
        this.poc2emailid = customerDetailsResponse.poc2emailid;
        this.poc2mobileno = customerDetailsResponse.poc2mobileno;
        this.address = customerDetailsResponse.address;
        this.addressstate = customerDetailsResponse.addressstate;
        this.addresscity = customerDetailsResponse.addresscity;
        this.addresspincode = customerDetailsResponse.addresspincode;
    }

}


export class CustomerInfo {
    corpname: string;
    custname: string;

    copyFrom(customerLoginInfoResponse: CustomerLoginInfoResponse) {
        this.corpname = customerLoginInfoResponse.corpname;
        this.custname = customerLoginInfoResponse.custname;
    }
}



