import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";

import { DetailsComponent } from "./details/details.component";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular";
import { NetworkStatusComponent } from "~/app/network-status.component";
import { FloatBtnComponent } from "~/app/float-btn.component";
import { PatientMgntRoutingModule } from "./patient-mgnt-routing.module";
import { PatientMgntComponent } from "./patient-mgnt.component";
import { CamerasComponent } from "./cameras/cameras.component";
import { MonitorComponent } from "./monitor/monitor.component";
import { ActionComponent } from "./action/action.component";
import { ChartsComponent } from "./charts/charts.component";
import { MedicineChartComponent } from "./charts/medicine-chart/medicine-chart.component";
import { NativeScriptUIChartModule } from "nativescript-ui-chart/angular";
import { NativeScriptFormsModule } from "nativescript-angular/forms";
import { DatePipe } from "@angular/common";
import { ChartService } from "~/app/services/chart/chart.service";
import { IntakeChartComponent } from "./charts/intake-chart/intake-chart.component";
import { MonitorChartComponent } from "./charts/monitor-chart/monitor-chart.component";
import { ReportsComponent } from "./reports/reports.component";
import { SectionOneComponent } from "./reports/section-one/section-one.component";
import { ReactiveFormsModule } from "@angular/forms";
import { sharedModule } from "~/app/modules/shared-mudule";
import { PatientActionBarComponent } from "./patient-action-bar/patient-action-bar.component";
import { Action } from "rxjs/internal/scheduler/Action";
import { ActionService } from "~/app/services/action/action.service";
import { ShowCameraImageComponent } from "./reports/section-one/show-camera-image/show-camera-image.component";
import { ShowUploadedImageComponent } from "./reports/section-one/show-uploaded-image/show-uploaded-image.component";
import { NativeScriptUIAutoCompleteTextViewModule } from "nativescript-ui-autocomplete/angular";
import { MonitorService } from "~/app/services/monitor/monitor.service";
import { UserAuthComponent } from "./user_auth/user_auth.component";
import { DoctorOrdersComponent } from "./doctor-orders/doctor-orders.component";
import { ActionFabComponent } from "./action-fab/action-fab.component";
import { SchedularFabComponent } from "./schedular-fab/schedular-fab.component";
import { TreatmentReportsComponent } from "./reports/treatment-reports/treatment-reports.component";
import { OutputChartComponent } from "./charts/output-chart/output-chart.component";
import { ReportsService } from "~/app/services/reports/reports-service";
import { AppNotificationService } from "~/app/services/app-notification-service";
import { MedicineActionsComponent } from "./action/medicine-actions/medicine-actions.component";
import { ImageModalComponent } from "./image-modal/image-modal.component";

@NgModule({
    imports: [
        NativeScriptCommonModule,
        PatientMgntRoutingModule,
        NativeScriptUIListViewModule,
        NativeScriptUIChartModule,
        NativeScriptFormsModule,
        ReactiveFormsModule,
        sharedModule,
        NativeScriptFormsModule,
        NativeScriptUIAutoCompleteTextViewModule
    ],
    declarations: [
        PatientMgntComponent,
        DetailsComponent,
        FloatBtnComponent,
        CamerasComponent,
        MonitorComponent,
        ActionComponent,
        ChartsComponent,
        MedicineChartComponent,
        IntakeChartComponent,
        MonitorChartComponent,
        ReportsComponent,
        SectionOneComponent,
        PatientActionBarComponent,
        ShowCameraImageComponent,
        ShowUploadedImageComponent,
        UserAuthComponent,
        DoctorOrdersComponent,
        ActionFabComponent,
        SchedularFabComponent,
        TreatmentReportsComponent,
        OutputChartComponent,
        MedicineActionsComponent,
        ImageModalComponent

    ],
    schemas: [
        NO_ERRORS_SCHEMA
    ],
    providers: [
        DatePipe,
        ChartService,
        ActionService,
        MonitorService,
        ReportsService,
        AppNotificationService
    ],
    entryComponents: [
        DoctorOrdersComponent,
        ActionFabComponent,
        SchedularFabComponent,
        IntakeChartComponent,
        MedicineChartComponent,
        MonitorChartComponent,
        OutputChartComponent,
        MedicineActionsComponent,
        ImageModalComponent
    ]
})
export class PatientMgntModule {
    constructor() {
        console.log('PatientMgntModule muodule initiate');
    }
}
