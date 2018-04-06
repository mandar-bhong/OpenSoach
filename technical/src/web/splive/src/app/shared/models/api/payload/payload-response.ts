import { PayloadError } from './payload-error';
export class PayloadResponse<T> {
    issuccess: boolean;
    data: T;
    error: PayloadError;
  }
