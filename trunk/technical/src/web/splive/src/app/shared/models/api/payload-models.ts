export class PayloadResponse<T> {
    issuccess: boolean;
    data: T;
    error: PayloadError;
}

export class PayloadError {
    code: number;
}
