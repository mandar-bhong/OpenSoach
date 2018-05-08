import { UserLoginInfoResponse, UserAddDetailsRequest, UserDetailsResponse, UserAddRequest } from '../api/user-models';
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
    usrstate: number;
    copyTo(userAddRequest: UserAddRequest) {
        userAddRequest.usrname = this.usrname;
        userAddRequest.usrcategory = this.usrcategory;
        userAddRequest.uroleid = this.uroleid;
        userAddRequest.usrstate = this.usrstate;
    }
}
