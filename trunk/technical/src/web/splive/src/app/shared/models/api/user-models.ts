export class UserLoginInfoResponse {
    fname: string;
    lname: string;
}

export class UserAddRequest {
    usrname: string;
    usrpassword: string;
    usrcategory: number;
	uroleidfk: number;
	usrstate: number;
}

export class UserAddDetailsRequest {
    usridfk: number;
    fname: string;
    lname: string;
    mobileno: string;
    alternatecontactno: string;
}

export class UserUpdateStateRequest {
	id: number;
	usrstate: number;
}

export class RecordChangePassRequest {
	oldpassword: string;
	newpassword: string;
}