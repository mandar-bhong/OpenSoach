export class AuthRequest {
    username: string;
    password: string;
    prodcode: string;
}

export class AuthResponse {
    token: string;
    urolecode: string;
}

export class ValidateAuthTokenRequest {
    token: string;
}
