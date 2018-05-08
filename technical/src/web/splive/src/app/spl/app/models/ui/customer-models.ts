import { CUSTOMER_PRODUCT_MAPPING_STATE, CUSTOMER_STATE } from '../../../../shared/app-common-constants';
import {
    CustomerAddRequest,
    CustomerAssociateProductListItemResponse,
    CustomerAssociateProductRequest,
    CustomerAssociateProductUpdateRequest,
    CustomerFilterRequest,
} from '../api/customer-models';

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

export class CustomerAssociateProductModel {
    custid: number;
    prodid: number;
    dbiid: number;
    cpmstate: CUSTOMER_PRODUCT_MAPPING_STATE;
    cpmid: number;

    copyToAddRequest(request: CustomerAssociateProductRequest) {
        request.custid = this.custid;
        request.prodid = this.prodid;
        request.dbiid = this.dbiid;
        request.cpmstate = this.cpmstate;
    }

    copyToUpdateRequest(request: CustomerAssociateProductUpdateRequest) {
        request.cpmid = this.cpmid;
        request.cpmstate = this.cpmstate;
    }

    copyFrom(details: CustomerAssociateProductListItemResponse) {
        this.cpmid = details.cpmid;
        this.cpmstate = details.cpmstate;
        this.dbiid = details.dbiid;
        this.prodid = details.prodid;
    }
}
