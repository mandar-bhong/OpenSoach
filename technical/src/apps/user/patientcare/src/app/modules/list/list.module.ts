import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { HttpClientModule } from '@angular/common/http';

import { ListRoutingModule } from "./list-routing.module";
import { ListComponent } from "./list.component";
import { DetailsComponent } from "./details/details.component";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular";
import { NetworkStatusComponent } from "~/app/network-status.component";
import { FloatBtnComponent } from "~/app/float-btn.component";
import { CamerasComponent } from "~/app/modules/list/cameras/cameras.component";
import { PatientDetailsComponent } from "~/app/modules/list/patient-details/patient-details.component";
import { MonitorComponent } from "~/app/modules/list/monitor/monitor.component";
import { ActionComponent } from "~/app/modules/list/action/action.component";
import { ChartsComponent } from "~/app/modules/list/charts/charts.component";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        ListRoutingModule,
        HttpClientModule,
        NativeScriptUIListViewModule
    ],
    declarations: [
        ListComponent,
        DetailsComponent,
        NetworkStatusComponent,
        FloatBtnComponent,
        CamerasComponent,
        PatientDetailsComponent,
        MonitorComponent,
        ActionComponent,
        ChartsComponent
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class ListModule { }
