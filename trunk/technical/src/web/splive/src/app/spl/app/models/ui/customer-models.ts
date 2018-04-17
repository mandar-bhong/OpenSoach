import { CUSTOMER_STATE } from '../../../../shared/app-common-constants';
import { CustomerAddRequest } from '../api/customer-models';

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
