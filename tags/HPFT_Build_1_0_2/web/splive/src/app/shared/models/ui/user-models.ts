import {
    UserLoginInfoResponse, UserAddDetailsRequest, UserDetailsResponse, UserAddRequest,
    UserAssociateProductListItemResponse,
    UserAssociateProductRequest,
    UserAssociateProductUpdateRequest,
    UserMasterResponse,
    UserMasterUpdateRequest,
    ChangeUserPasswordRequest,
    ActivationChangePassword,
    ForgotPasswordRequest,
    ResetPasswordRequest
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
    userid: number;
    usrname: string;
    usrpassword: string;
    usrcategory: number;
    uroleid: number;
    urolecode: string;
    usrstate: number;
    copyTo(userAddRequest: UserAddRequest) {
        userAddRequest.usrname = this.usrname;
        userAddRequest.usrpassword = this.usrpassword;
        userAddRequest.usrcategory = this.usrcategory;
        userAddRequest.uroleid = this.uroleid;
        userAddRequest.usrstate = this.usrstate;
        userAddRequest.urolecode = this.urolecode;
    }
    copyToUpdateRequest(userEditUpdateRequest: UserMasterUpdateRequest) {
        userEditUpdateRequest.userid = this.userid;
        userEditUpdateRequest.uroleid = this.uroleid;
        userEditUpdateRequest.usrstate = this.usrstate;
    }
    copyFrom(userEditResponse: UserMasterResponse) {
        this.userid = userEditResponse.userid;
        this.usrname = userEditResponse.usrname;
        this.uroleid = userEditResponse.uroleid;
        this.usrstate = userEditResponse.usrstate;
        this.usrcategory = userEditResponse.usrcategory;
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
    urolename: string;
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

export class ConfirmPasswordModel {
    usrid: number;
    newpassword: string;
    confirmpassword: string;
    copyTo(changeUserPasswordRequest: ChangeUserPasswordRequest) {
        changeUserPasswordRequest.usrid = this.usrid;
        changeUserPasswordRequest.newpassword = this.newpassword;
    }
}
export class ActivationChangePasswordModel {
    code: string;
    copyTo(activationChangePassword: ActivationChangePassword) {
        activationChangePassword.code = this.code;
    }
}

export class ForgotPasswordModel {
    usrname: string;
    otp: string;
    newpassword: string;
    confirmpassword: string;
    copyToForgotPass(forgotPasswordRequest : ForgotPasswordRequest){
        forgotPasswordRequest.usrname = this.usrname;
    }
    copyToResetPass(resetPasswordRequest : ResetPasswordRequest){
        resetPasswordRequest.usrname = this.usrname;
        resetPasswordRequest.otp = this.otp;
        resetPasswordRequest.newpassword = this.newpassword;
    }
}
