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
