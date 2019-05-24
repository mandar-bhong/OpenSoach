import { Component, OnInit, OnDestroy } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { WorkerService } from '~/app/services/worker.service';
import { ServerConnectivityStatusService } from './services/connectivity/server-connectivity.service';

@Component({
    moduleId: module.id,
    selector: "network-status",
    template: `
    <Label  class="homeicon mdi" text="&#xf12f;" [style.color]="connectionStatus ? 'green' : 'red'"></Label>
    `,
    styles: [`   `]

})
export class NetworkStatusComponent implements OnDestroy {
    connectionStatus: boolean = false;
    connectionSubscription;
    constructor(private routerExtensions: RouterExtensions,
        private serverConnectivityStatusService: ServerConnectivityStatusService) {
        this.connectionStatus = this.serverConnectivityStatusService.ServerConnectionStatus;
        this.serverConnectivityStatusService.ServerConnectionSubject.subscribe(status => {
            this.connectionStatus = status;
        });
    }

    ngOnDestroy(): void {
        if (this.connectionSubscription)
            this.connectionSubscription.unsubscribe();
    }
}
