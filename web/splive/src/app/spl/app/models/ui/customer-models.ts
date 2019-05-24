import { CUSTOMER_PRODUCT_MAPPING_STATE, CUSTOMER_STATE } from '../../../../shared/app-common-constants';
import { CustomerMasterResponse, CustomerMasterUpdateRequest } from '../../../../shared/models/api/customer-models';
import {
    CustomerAddRequest,
    CustomerAssociateProductListItemResponse,
    CustomerAssociateProductRequest,
    CustomerAssociateProductUpdateRequest,
    CustomerFilterRequest,
    CustomerServiceAssociateUpdateRequest,
    CustomerServiceAssociateListResponse,
} from '../api/customer-models';

export class CustomerAddModel {
    custid: number;
    custname: string;
    corpid: number;
    custstate: CUSTOMER_STATE;

    copyTo(customerAddRequest: CustomerAddRequest) {
        customerAddRequest.custname = this.custname;
        customerAddRequest.corpid = this.corpid;
        customerAddRequest.custstate = this.custstate;
    }
    copyToUpdateRequest(customerMasterUpdateRequest: CustomerMasterUpdateRequest) {
        customerMasterUpdateRequest.custid = this.custid;
        customerMasterUpdateRequest.custname = this.custname;
        customerMasterUpdateRequest.custstate = this.custstate;
    }
    copyFrom(customerMasterResponse: CustomerMasterResponse) {
        this.custid = customerMasterResponse.custid;
        this.corpid = customerMasterResponse.corpid;
        this.custname = customerMasterResponse.custname;
        this.custstate = customerMasterResponse.custstate;

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

export class CustomerSeviceAssociateProductModel {
    custid: number;
    prodcode: string;
    updatecount: number;
    spcount: number;
    cpmid: number;
    copyToUpdateRequest(request: CustomerServiceAssociateUpdateRequest) {
        request.cpmid = this.cpmid;
        request.updatecount = this.updatecount;
    }

    copyFrom(details: CustomerServiceAssociateListResponse) {
        this.prodcode = details.prodcode;
        this.spcount = details.spcount;
        this.cpmid = details.cpmid;

    }
}
