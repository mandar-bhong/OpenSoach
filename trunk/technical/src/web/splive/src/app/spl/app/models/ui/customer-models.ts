import { CUSTOMER_STATE } from '../../../../shared/app-common-constants';
import { CustomerAddRequest, CustomerFilterRequest } from '../api/customer-models';

export class CustomerAddModel {
    custname: string;
    corpid: number;
    custstate: CUSTOMER_STATE;

    copyTo(customerAddRequest: CustomerAddRequest) {
        customerAddRequest.custname = this.custname;
        customerAddRequest.corpid = this.corpid;
        customerAddRequest.custstate = this.custstate;
    }
}

export class CustomerFilterModel {
    custname: string;
    custstate: CUSTOMER_STATE;
    corpid: number;

    copyTo(customerFilterRequest: CustomerFilterRequest) {
        customerFilterRequest.custname = this.custname;
        customerFilterRequest.custstate = this.custstate;
        customerFilterRequest.corpid = this.corpid;
    }
}
