import { USER_CATEGORY } from '../../app-common-constants';
export class AuthRequest {
    username: string;
    password: string;
    prodcode: string;
}

export class AuthResponse {
    token: string;
    urolecode: string;
    usrcategory: USER_CATEGORY;
}

export class ValidateAuthTokenRequest {
    token: string;
}
