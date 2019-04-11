import { USER_STATE, USER_CATEGORY } from '../../../shared/app-common-constants';

export class UserFilterRequest {
    usrname: string;
    usrstate: USER_STATE;
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
export class UserAddRequest {
    usrname: string;
    usrpassword: string;
    uroleid: number;
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
export class UserMasterResponse {
    userid: number;
    usrname: string;
    usrstate: number;
    uroleid: number;
    usrstatesince: number;
}
export class UserRoleidListItemResponse {
    uroleid: number;
    urolecode: string;
    urolename: string;
    prodcode: string;
}
export class ProductcodeRequest {
    prodcode: string;
}
export class UserMasterUpdateRequest {
    userid: number;
    uroleid: number;
    usrstate: number;
}
export class UserLoginInfoResponse {
    usrname: string;
    fname: string;
    lname: string;
}
