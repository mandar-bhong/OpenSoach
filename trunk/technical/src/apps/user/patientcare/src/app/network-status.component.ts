import { Component, OnInit, OnDestroy } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";


import { InternetConnectionService } from '~/app/services/internet-status/internet-connection.service';

@Component({
    moduleId: module.id,
    selector: "network-status",
    template: `
    <Label  class="homeicon mdi" text="&#xf12f;" [style.color]="connectionStatus ? 'green' : 'red'"></Label>
    `,
    styles: [`   ` ]

})
export class NetworkStatusComponent implements OnInit, OnDestroy  {
    connectionStatus: boolean = true;
    connectiontext:string;
    connection$;
    constructor(private routerExtensions: RouterExtensions,
        private _internetConnection: InternetConnectionService) {
    }
    ngOnInit() {
		this.connection$ = this._internetConnection.connectionStatus$.subscribe(data => {
            this.connectionStatus = data.valueOf();
        });
    }

    ngOnDestroy(): void {
        if (this.connection$)
            this.connection$.unsubscribe();
    }
}
