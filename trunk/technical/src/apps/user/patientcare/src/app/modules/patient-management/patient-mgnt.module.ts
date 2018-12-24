import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { HttpClientModule } from '@angular/common/http';

import { DetailsComponent } from "./details/details.component";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular";
import { NetworkStatusComponent } from "~/app/network-status.component";
import { FloatBtnComponent } from "~/app/float-btn.component";
import { PatientMgntRoutingModule } from "~/app/modules/patient-management/patient-mgnt-routing.module";
import { PatientMgntComponent } from "~/app/modules/patient-management/patient-mgnt.component";
import { CamerasComponent } from "~/app/modules/patient-management/cameras/cameras.component";
import { PatientDetailsComponent } from "~/app/modules/patient-management/patient-details/patient-details.component";
import { MonitorComponent } from "~/app/modules/patient-management/monitor/monitor.component";
import { ActionComponent } from "~/app/modules/patient-management/action/action.component";
import { ChartsComponent } from "~/app/modules/patient-management/charts/charts.component";
import { MedicineChartComponent } from "~/app/modules/patient-management/charts/medicine-chart/medicine-chart.component";
import { NativeScriptUIDataFormModule } from "nativescript-ui-dataform/angular";
import { NativeScriptUIChartModule } from "nativescript-ui-chart/angular";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        PatientMgntRoutingModule,
        HttpClientModule,
        NativeScriptUIListViewModule,
        NativeScriptUIDataFormModule,
        NativeScriptUIChartModule
    ],
    declarations: [
        PatientMgntComponent,
        DetailsComponent,
        FloatBtnComponent,
        CamerasComponent,
        PatientDetailsComponent,
        MonitorComponent,
        ActionComponent,
        ChartsComponent,
        MedicineChartComponent
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class PatientMgntModule { }
