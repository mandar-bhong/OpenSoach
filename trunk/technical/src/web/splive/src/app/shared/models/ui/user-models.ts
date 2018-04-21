import { UserLoginInfoResponse } from '../api/user-models';
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
