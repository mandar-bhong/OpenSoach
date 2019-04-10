import { CUSTOMER_PRODUCT_MAPPING_STATE, CUSTOMER_STATE } from '../../../../shared/app-common-constants';

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
    custstate: CUSTOMER_STATE;
    corpname: string;
    poc1name: string;
    poc1emailid: string;
    poc1mobileno: string;
    createdon: Date;
    updateon: Date;
}

export class CustomerAssociateProductRequest {
    custid: number;
    prodid: number;
    dbiid: number;
    cpmstate: CUSTOMER_PRODUCT_MAPPING_STATE;
}

export class CustomerAssociateProductListItemResponse {
    cpmid: number;
    prodid: number;
    prodcode: string;
    dbiid: number;
    dbiname: string;
    cpmstate: CUSTOMER_PRODUCT_MAPPING_STATE;
}

export class CustomerAssociateProductUpdateRequest {
    cpmid: number;
    cpmstate: CUSTOMER_PRODUCT_MAPPING_STATE;
}

export class CustomerListItemResponse {
    custid: number;
    custname: string;
    createdon: Date;
    updatedon: Date;
}
export class CustomerRoleidListItemResponse {
    uroleid: number;
    urolecode: string;
    urolename: string;
    prodcode: string;
}
export class CustomerRoleListRequest {
    prodcode: string;
}
export class CustomerServiceAssociateListResponse {
    prodcode: string;
    spcount: number;
    cpmid: number;
}
export class CustomerServiceAssociateUpdateRequest {
    updatecount: number;
    cpmid: number;
}
