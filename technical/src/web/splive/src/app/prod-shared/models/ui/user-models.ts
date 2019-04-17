import { USER_STATE } from '../../../shared/app-common-constants';
import {
    UserAddDetailsRequest,
    UserAddRequest,
    UserDetailsResponse,
    UserFilterRequest,
    UserMasterResponse,
    UserMasterUpdateRequest,
    UserLoginInfoResponse,
    ChangeUserPasswordRequest
} from '../api/user-models';


export class UserFilterModel {
    usrname: string;
    usrstate: USER_STATE;

    copyTo(userFilterRequest: UserFilterRequest) {
        userFilterRequest.usrname = this.usrname;
        userFilterRequest.usrstate = this.usrstate;
    }
}

export class UserAddModel {
    userid: number;
    usrname: string;
    uroleid: number;
    urolecode: string;
    prodcode: string;
    copyTo(userAddRequest: UserAddRequest) {
        userAddRequest.usrname = this.usrname;
        userAddRequest.uroleid = this.uroleid;
        userAddRequest.urolecode = this.urolecode;
    }
}

export class UserDetailsModel {
    userid: number;
    fname: string;
    lname: string;
    mobileno: string;
    gender: string;
    alternatecontactno: string;
    usrid: number;
    usrname: string;
    usrstate: number;
    uroleid: number;
    copyTo(userAddDetailsRequest: UserAddDetailsRequest) {
        userAddDetailsRequest.usrid = this.usrid;
        userAddDetailsRequest.fname = this.fname;
        userAddDetailsRequest.lname = this.lname;
        userAddDetailsRequest.mobileno = this.mobileno;
        userAddDetailsRequest.gender = this.gender;
        userAddDetailsRequest.alternatecontactno = this.alternatecontactno;
    }
    copyToUpdateRequest(userEditUpdateRequest: UserMasterUpdateRequest) {
        userEditUpdateRequest.userid = this.userid;
        userEditUpdateRequest.usrstate = this.usrstate;
        userEditUpdateRequest.uroleid = this.uroleid;
    }
    copyFrom(userDetailsResponse: UserDetailsResponse) {
        this.usrid = userDetailsResponse.usrid;
        this.fname = userDetailsResponse.fname;
        this.lname = userDetailsResponse.lname;
        this.mobileno = userDetailsResponse.mobileno;
        this.gender = userDetailsResponse.gender;
        this.alternatecontactno = userDetailsResponse.alternatecontactno;
    }
    copyFromMaster(userEditResponse: UserMasterResponse) {
        this.userid = userEditResponse.userid;
        this.usrname = userEditResponse.usrname;
        this.uroleid = userEditResponse.uroleid;
        this.usrstate = userEditResponse.usrstate;
    }
}
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

export class ConfirmPasswordModel
{
    oldpassword: string;
    newpassword: string;
    confirmpassword: string;
    copyTo(changeUserPasswordRequest: ChangeUserPasswordRequest) {
        changeUserPasswordRequest.oldpassword = this.oldpassword;
        changeUserPasswordRequest.newpassword = this.newpassword;
    }
}
