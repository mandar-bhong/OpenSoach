import {
    UserLoginInfoResponse, UserAddDetailsRequest, UserDetailsResponse, UserAddRequest,
    UserAssociateProductListItemResponse,
    UserAssociateProductRequest,
    UserAssociateProductUpdateRequest
} from '../api/user-models';
import { USER_PRODUCT_MAPPING_STATE } from '../../app-common-constants';
import { CustomerAssociateProductListItemResponse } from '../../../spl/app/models/api/customer-models';
export class UserInfo {
    usrname: string;
    fname: string;
    lname: string;

    copyFrom(userLoginInfoResponse: UserLoginInfoResponse) {
        this.usrname = userLoginInfoResponse.usrname;
        this.fname = userLoginInfoResponse.fname;
        this.lname = userLoginInfoResponse.lname;
    }
}

export class UserDetailsModel {
    usrid: number;
    fname: string;
    lname: string;
    mobileno: string;
    gender: string;
    alternatecontactno: string;
    copyTo(userAddDetailsRequest: UserAddDetailsRequest) {
        userAddDetailsRequest.usrid = this.usrid;
        userAddDetailsRequest.fname = this.fname;
        userAddDetailsRequest.lname = this.lname;
        userAddDetailsRequest.mobileno = this.mobileno;
        userAddDetailsRequest.gender = this.gender;
        userAddDetailsRequest.alternatecontactno = this.alternatecontactno;
    }
    copyFrom(userDetailsResponse: UserDetailsResponse) {
        this.usrid = userDetailsResponse.usrid;
        this.fname = userDetailsResponse.fname;
        this.lname = userDetailsResponse.lname;
        this.mobileno = userDetailsResponse.mobileno;
        this.gender = userDetailsResponse.gender;
        this.alternatecontactno = userDetailsResponse.alternatecontactno;
    }
}
export class UserAddModel {
    usrname: string;
    usrpassword: string;
    usrcategory: number;
    uroleid: number;
    urolecode: string;
    usrstate: number;
    copyTo(userAddRequest: UserAddRequest) {
        userAddRequest.usrname = this.usrname;
        userAddRequest.usrcategory = this.usrcategory;
        userAddRequest.uroleid = this.uroleid;
        userAddRequest.usrstate = this.usrstate;
        userAddRequest.urolecode = this.urolecode;
    }
}

export class UserAssociateProductModel {
    usrid: number;
    cpmid: number;
    cpm: CustomerAssociateProductListItemResponse;
    usrname: string;
    uroleid: number;
    ucpmstate: USER_PRODUCT_MAPPING_STATE;
    ucpmid: number;
    prodcode: string;
    urolecode: number;
    custname: string;
    custid: number;

    copyToAddRequest(request: UserAssociateProductRequest) {
        request.userid = this.usrid;
        request.cpmid = this.cpm.cpmid;
        request.usrname = this.usrname;
        request.uroleid = this.uroleid;
        request.ucpmstate = this.ucpmstate;
    }

    copyToUpdateRequest(request: UserAssociateProductUpdateRequest) {
        request.ucpmid = this.ucpmid;
        request.ucpmstate = this.ucpmstate;
    }

    copyFrom(details: UserAssociateProductListItemResponse) {
        this.ucpmid = details.ucpmid;
        this.custname = details.custname;
        this.ucpmstate = details.ucpmstate;
        this.urolecode = details.urolecode;
        this.prodcode = details.prodcode;
    }
}
