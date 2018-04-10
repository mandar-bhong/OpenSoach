import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { PayloadResponse } from '../../models/api/payload-models';
import { LoginStatusService } from '../../services/login-status.service';

@Injectable()
export class ServerApiInterfaceService {

    constructor(private http: HttpClient, private loginStatusService: LoginStatusService) { }

    post(url: string, request: any): Observable<PayloadResponse<any>> {
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
        console.log(httpOptions);
        return this.http.post<PayloadResponse<any>>
            (url, request, httpOptions);
    }

    getWithQueryParams(url: string, queryParams: any): Observable<PayloadResponse<any>> {
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
            (url, httpOptions);
    }

    get(url: string): Observable<PayloadResponse<any>> {
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
            (url, httpOptions);
    }
}
