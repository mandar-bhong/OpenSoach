import { ErrorHandler, NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptModule } from "nativescript-angular/nativescript.module";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { DatabaseService } from "./services/offline-store/database.service";
import {NativeScriptFormsModule} from "nativescript-angular/forms"
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";
import { InternetConnectionService } from "~/app/services/connectivity/internet-connection.service";
import { WorkerService } from "./services/worker.service";
import { PatientListService } from "~/app/services/patient-list/patient-list.service";
import { NetworkStatusComponent } from "~/app/network-status.component";
import { sharedModule } from "~/app/modules/shared-mudule";
import { PassDataService } from "~/app/services/pass-data-service";
import { ServerConnectivityStatusService } from "~/app/services/connectivity/server-connectivity.service";
import { NativeScriptHttpClientModule } from "nativescript-angular/http-client";
import { ServerApiInterfaceService } from "./services/server-api-interface.service";
import { AppStartupService } from "./services/app-startup.service";
import { StartupComponent } from "./startup-page/startup.component";
import { NextActionService } from "./services/action/next-action-service";
import { UserAuthService } from "./services/user-auth/user-auth-service";
import { NativeScriptAnimationsModule } from "nativescript-angular/animations";
import { NativeScriptUISideDrawerModule } from "nativescript-ui-sidedrawer/angular";

import { TraceAngularErrorHandler } from "./helpers/trace-angular-error-handler";



export function angularErrorHandlerFactory() {  
    return new  TraceAngularErrorHandler();
}


@NgModule({
    bootstrap: [
        AppComponent
    ],
    imports: [
        NativeScriptModule,
        AppRoutingModule,
        NativeScriptFormsModule,
        sharedModule,
        NativeScriptHttpClientModule,
        NativeScriptAnimationsModule,
        NativeScriptUISideDrawerModule
    ],
    declarations: [
        AppComponent,
        StartupComponent
           ],
    schemas: [
        NO_ERRORS_SCHEMA
    ],
    providers: [
        DatabaseService,
        PassDataService,
        DatabaseSchemaService,
        InternetConnectionService,
        WorkerService,
        PatientListService,
        ServerConnectivityStatusService,
        ServerApiInterfaceService,
        AppStartupService,
        NextActionService,
        UserAuthService,
        { provide: ErrorHandler, useFactory: angularErrorHandlerFactory }
    ],
    exports:[
        // NetworkStatusComponent
    ]
})
export class AppModule { }
