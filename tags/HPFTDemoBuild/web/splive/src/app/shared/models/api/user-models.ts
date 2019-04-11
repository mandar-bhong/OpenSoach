import { USER_PRODUCT_MAPPING_STATE } from '../../app-common-constants';
export class UserLoginInfoResponse {
    usrname: string;
    fname: string;
    lname: string;
}

export class UserAddRequest {
    usrname: string;
    usrpassword: string;
    usrcategory: number;
    uroleid: number;
    usrstate: number;
    urolecode: string;
}

export class UserAddDetailsRequest {
    usrid: number;
    fname: string;
    lname: string;
    mobileno: string;
    gender: string;
    alternatecontactno: string;
}
export class UserDetailsResponse {
    usrid: number;
    fname: string;
    lname: string;
    mobileno: string;
    alternatecontactno: string;
    gender: string;
    createdon: Date;
    updateon: Date;
}
export class UserDataListResponse {
    usrid: number;
    usrname: string;
    usrcategory: number;
    uroleid: number;
    urolename: string;
    usrstate: number;
    usrstatesince: number;
    fname: string;
    lname: string;
    mobileno: number;
    createdon: Date;
    updatedon: Date;
}
export class UserUpdateStateRequest {
    id: number;
    usrstate: number;
}
export class RecordChangePassRequest {
    oldpassword: string;
    newpassword: string;

}
export class UserAssociateProductListItemResponse {
    ucpmid: number;
    cpmid: number;
    custname: string;
    prodcode: string;
    urolecode: number;
    ucpmstate: USER_PRODUCT_MAPPING_STATE;
    ucpmstatesince: number;
}
export class UserAssociateProductRequest {
    userid: number;
    cpmid: number;
    usrname: string;
    uroleid: number;
    ucpmstate: USER_PRODUCT_MAPPING_STATE;
}
export class UserAssociateProductUpdateRequest {
    ucpmid: number;
    ucpmstate: USER_PRODUCT_MAPPING_STATE;
    ucpmstatesince: number;
}
export class UserRoleidListItemResponse {
    uroleid: number;
    urolecode: string;
    urolename: string;
    prodcode: string;
}
export class UserMasterUpdateRequest {
    userid: number;
    uroleid: number;
    usrstate: number;
}
export class UserMasterResponse {
    userid: number;
    usrname: string;
    usrstate: number;
    uroleid: number;
    usrcategory: number;
}
