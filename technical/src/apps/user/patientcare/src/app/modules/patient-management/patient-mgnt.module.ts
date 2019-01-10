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
import { NativeScriptUIChartModule } from "nativescript-ui-chart/angular";
import { NativeScriptFormsModule } from "nativescript-angular/forms";
import { DatePipe } from "@angular/common";
import { ChartService } from "~/app/services/chart/chart.service";
import { IntakeChartComponent } from "~/app/modules/patient-management/charts/intake-chart/intake-chart.component";
import { MonitorChartComponent } from "~/app/modules/patient-management/charts/monitor-chart/monitor-chart.component";
import { MedicalDetailsComponent } from "~/app/modules/patient-management/patient-details/medical-details/medical-details.component";
import { PatientInfoComponent } from "~/app/modules/patient-management/patient-details/patient-info/patient-info.component";
import { ConfService } from "~/app/services/conf/conf.service";
import { ReportsComponent } from "~/app/modules/patient-management/reports/reports.component";
import { SectionOneComponent } from "~/app/modules/patient-management/reports/section-one/section-one.component";
import { ReactiveFormsModule } from "@angular/forms";
import { sharedModule } from "~/app/modules/shared-mudule";
import { PatientActionBarComponent } from "~/app/modules/patient-management/patient-action-bar/patient-action-bar.component";
import { Action } from "rxjs/internal/scheduler/Action";
import { ActionService } from "~/app/services/action/action.service";
import { ShowCameraImageComponent } from "~/app/modules/patient-management/reports/section-one/show-camera-image/show-camera-image.component";
import { ShowUploadedImageComponent } from "~/app/modules/patient-management/reports/section-one/show-uploaded-image/show-uploaded-image.component";
import { GetUUIDService } from "~/app/services/get-UUID.service";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        PatientMgntRoutingModule,
        HttpClientModule,
        NativeScriptUIListViewModule,
        NativeScriptUIChartModule,
        NativeScriptFormsModule,
        ReactiveFormsModule,
        sharedModule,
        NativeScriptFormsModule
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
        MedicineChartComponent,
        IntakeChartComponent,
        MonitorChartComponent,
        MedicalDetailsComponent,
        PatientInfoComponent,
        ReportsComponent, 
        SectionOneComponent,
        PatientActionBarComponent,
        ShowCameraImageComponent,
        ShowUploadedImageComponent
       
    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ],
    providers:[
        DatePipe,
        ChartService,
        ConfService,
        ActionService,
        GetUUIDService
    ]
})
export class PatientMgntModule {
    constructor(){
        console.log('PatientMgntModule muodule initiate');
    }
 }
