export class AuthRequest {
    usrname: string;
    usrpassword: string;
    prodcode: string;
}

export class AuthResponse {
    token: string;
    urolecode: string;
}
