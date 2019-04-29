import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AppGlobalContext } from '../app-global-context';

export class ServerAPIErrorModel {
    handled: boolean;
    error: any;

    constructor(handled: boolean, error: any) {
        this.handled = handled;
        this.error = error;
    }
}

@Injectable()
export class ServerApiInterfaceService {
    constructor(private http: HttpClient) {
    }

    get<T>(url: string, queryParams?: any): Promise<T> {
        return new Promise((resolve, reject) => {
            try {
                let httpOptions = {
                    headers: new HttpHeaders(
                        {
                            'Content-Type': 'application/json'
                        }),
                    params: new HttpParams()
                };
                if (AppGlobalContext.Token) {
                    httpOptions = {
                        headers: new HttpHeaders(
                            {
                                'Content-Type': 'application/json',
                                'Authorization': AppGlobalContext.Token
                            }
                        ),
                        params: new HttpParams()
                    };
                }

                if (queryParams) {
                    httpOptions.params = new HttpParams().set('params', JSON.stringify(queryParams))
                }

                this.http.get<any>
                    (url, httpOptions).subscribe(payloadResponse => {
                        console.log('response received', payloadResponse);
                        if (payloadResponse) {
                            if (payloadResponse.issuccess)
                                resolve(payloadResponse.data);
                            else
                                reject(new ServerAPIErrorModel(false, payloadResponse.error));
                        }
                        else {
                            reject(new ServerAPIErrorModel(false, null));
                        }
                    }, error => {
                        reject(this.handleError(url, error));
                    });
            } catch (error) {
                reject(new ServerAPIErrorModel(false, null));
            }
        });
    }

    post<T>(url: string, requestBody?: any, queryParams?: any, ): Promise<T> {
        return new Promise((resolve, reject) => {
            try {
                let httpOptions = {
                    headers: new HttpHeaders(
                        {
                            'Content-Type': 'application/json'
                        }),
                    params: new HttpParams()
                };
                if (AppGlobalContext.Token) {
                    httpOptions = {
                        headers: new HttpHeaders(
                            {
                                'Content-Type': 'application/json',
                                'Authorization': AppGlobalContext.Token
                            }
                        ),
                        params: new HttpParams()
                    };
                }

                if (queryParams) {
                    httpOptions.params = new HttpParams().set('params', JSON.stringify(queryParams))
                }

                this.http.post<any>
                    (url, requestBody, httpOptions).subscribe(payloadResponse => {                     
                        if (payloadResponse) {
                            if (payloadResponse.issuccess)
                                resolve(payloadResponse.data);
                            else
                                reject(new ServerAPIErrorModel(false, payloadResponse.error));
                        }
                        else {
                            reject(new ServerAPIErrorModel(false, null));
                        }
                    }, error => {
                        reject(this.handleError(url, error));
                    });
            } catch (error) {
                reject(new ServerAPIErrorModel(false, null));
            }
        });
    }

    handleError(url, error) {

        console.error(url + ' failed', error);

        const errorModel = new ServerAPIErrorModel(false, error.message);
        errorModel.error = error.message;
        switch (error.status) {
            case 401:
                errorModel.handled = true;
                break;
            case 403:
                errorModel.handled = true;
                break;
            case 404:
                errorModel.handled = true;
            default:

                break;
        }

        return errorModel;
    }
}