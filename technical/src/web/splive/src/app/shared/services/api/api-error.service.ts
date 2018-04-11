import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { AppNotificationService } from '../../services/notification/app-notification.service';
import { PayloadResponse } from '../../models/api/payload-models';
import { Router } from '@angular/router';
import { LoginStatusService } from '../../services/login-status.service';
import { TranslatePipe } from '../../pipes/translate/translate.pipe';
import { SERVER_SYSTEM_ERROR_MAX_BOUNDARY } from '../../app-common-constants';

@Injectable()
export class ApiErrorService {
defaultErrorHandler= this.handleApiError;
    constructor(private appNotificationService: AppNotificationService,
        private router: Router,
        private loginstatusservice: LoginStatusService,
        private translatePipe: TranslatePipe) { }

    handleError<T>(url, result?: T) {
        return (error: any): Observable<T> => {
            // TODO: send the error to remote logging infrastructure
            console.error(url + ' failed', error); // log to console instead

            switch (error.status) {
                case 401:
                    // TODO: user is not autheticated, redirect to login page
                    // this.router.navigate(['login']);
                    // this.loginstatusservice.changeStatus(false);
                    // this.router.navigate(['unauthorized']);
                    break;
                case 403:
                // No access to a resource, redirect to login or not authorized page
                case 404:
                // not found, display error to user
                default:
                    // display error occured to user.
                    break;
            }

            // TODO: better job of transforming error for user consumption
            // this.log(`${operation} failed: ${error.message}`);
            // Let the app keep running by returning an empty result.
            return of(result as T);
        };
    }

    handleApiError(payloadResponse: PayloadResponse<any>) {
        if (payloadResponse && payloadResponse.issuccess) {
            return;
        }

        let errorMessage: string;
        if (payloadResponse && payloadResponse.error) {
            if (payloadResponse.error.code <= SERVER_SYSTEM_ERROR_MAX_BOUNDARY) {
                // display a common system error with error code
                errorMessage = this.translatePipe.transform('SERVER_SYSTEM_ERROR') + ' - ' + payloadResponse.error.code;
            } else {
                // retrive error messages specific to error code
                errorMessage = this.translatePipe.transform('SERVER_ERROR_' + payloadResponse.error.code);
                if (!errorMessage) {
                    // display common error message if specific error code message is not available.
                    errorMessage = this.translatePipe.transform('SERVER_ERROR') + ' - ' + payloadResponse.error.code;
                }
            }
        } else {
            // display common error message if server has not reverted with proper error.
            errorMessage = this.translatePipe.transform('SERVER_UNKNOWN_ERROR');
        }

        this.appNotificationService.error(errorMessage);
    }
}
