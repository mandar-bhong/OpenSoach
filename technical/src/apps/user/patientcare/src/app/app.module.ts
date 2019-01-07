import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptModule } from "nativescript-angular/nativescript.module";
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { DatabaseService } from "./services/offline-store/database.service";
import {NativeScriptFormsModule} from "nativescript-angular/forms"
import { DatabaseSchemaService } from "./services/offline-store/database-schema.service";
import { InternetConnectionService } from "~/app/services/internet-status/internet-connection.service";
import { WorkerService } from "./services/worker.service";
import { PatientListService } from "~/app/services/patient-list/patient-list.service";
import { NetworkStatusComponent } from "~/app/network-status.component";
import { sharedModule } from "~/app/modules/shared-mudule";
import { PassDataService } from "~/app/services/pass-data-service";

@NgModule({
    bootstrap: [
        AppComponent
    ],
    imports: [
        NativeScriptModule,
        AppRoutingModule,
        NativeScriptFormsModule,
        HttpClientModule,
        sharedModule
    ],
    declarations: [
        AppComponent
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
        PatientListService
    ],
    exports:[
        // NetworkStatusComponent
    ]
})
export class AppModule { }
