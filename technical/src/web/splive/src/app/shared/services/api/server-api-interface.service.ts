import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { catchError, tap } from 'rxjs/operators';
import { PayloadResponse } from '../../models/api/payload-models';
import { LoginStatusService } from '../../services/login-status.service';
import { ApiErrorService } from '../../services/api/api-error.service';

@Injectable()
export class ServerApiInterfaceService {

    constructor(private http: HttpClient,
        private loginStatusService: LoginStatusService,
        private apiErrorService: ApiErrorService) { }

    post(url: string, request: any, implicitErrorHandling = true): Observable<PayloadResponse<any>> {
        let httpOptions = {
            headers: new HttpHeaders(
                {
                    'Content-Type': 'application/json'
                }),
            withCredentials: true,
        };
        if (this.loginStatusService.isLoggedIn) {
            httpOptions = {
                headers: new HttpHeaders(
                    {
                        'Content-Type': 'application/json',
                        'Authorization': this.loginStatusService.authToken
                    }
                ),
                withCredentials: true,
            };
        }
        return this.http.post<PayloadResponse<any>>
            (url, request, httpOptions).pipe(
            tap(
                payloadResponse => {
                    if (implicitErrorHandling) {
                        this.apiErrorService.handleApiError(payloadResponse);
                    }
                }
            ),
            catchError(this.apiErrorService.handleError<PayloadResponse<any>>(url)));
    }

    getWithQueryParams(url: string, queryParams: any, implicitErrorHandling = true): Observable<PayloadResponse<any>> {
        const httpParams = new HttpParams();
        httpParams.set('params', JSON.stringify(queryParams));
        let httpOptions = {
            headers: new HttpHeaders(
                {
                    'Content-Type': 'application/json'
                }),
            withCredentials: true,
            params: httpParams
        };
        if (this.loginStatusService.isLoggedIn) {
            httpOptions = {
                headers: new HttpHeaders(
                    {
                        'Content-Type': 'application/json',
                        'Authorization': this.loginStatusService.authToken
                    }
                ),
                withCredentials: true,
                params: httpParams
            };
        }
        return this.http.get<PayloadResponse<any>>
            (url, httpOptions).pipe(
            tap(
                payloadResponse => {
                    if (implicitErrorHandling) {
                        this.apiErrorService.handleApiError(payloadResponse);
                    }
                }
            ),
            catchError(this.apiErrorService.handleError<PayloadResponse<any>>(url)));
    }

    get(url: string, implicitErrorHandling = true): Observable<PayloadResponse<any>> {
        let httpOptions = {
            headers: new HttpHeaders(
                {
                    'Content-Type': 'application/json'
                }),
            withCredentials: true
        };
        if (this.loginStatusService.isLoggedIn) {
            httpOptions = {
                headers: new HttpHeaders(
                    {
                        'Content-Type': 'application/json',
                        'Authorization': this.loginStatusService.authToken
                    }
                ),
                withCredentials: true
            };
        }
        return this.http.get<PayloadResponse<any>>
            (url, httpOptions).pipe(
            tap(
                payloadResponse => {
                    if (implicitErrorHandling) {
                        this.apiErrorService.handleApiError(payloadResponse);
                    }
                }
            ),
            catchError(this.apiErrorService.handleError<PayloadResponse<any>>(url)));
    }
}
