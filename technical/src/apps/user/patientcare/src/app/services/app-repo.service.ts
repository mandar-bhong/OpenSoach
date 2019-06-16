import { Injectable } from "@angular/core";

@Injectable({ providedIn: 'root' })
export class AppRepoService {

    public static Instance = new AppRepoService();

    public DBWorkerRequestMapper: Map<number, any> = new Map<number, any>();
    public DBWorkerRequestCounter = 0;

    constructor() {

    }

}