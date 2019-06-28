import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { NativeScriptCommonModule } from "nativescript-angular/common";
import { PatientDetailsRoutingModule } from "./patient-details-routing.module";
import { NativeScriptUIListViewModule } from "nativescript-ui-listview/angular/listview-directives";
import { NativeScriptUIChartModule } from "nativescript-ui-chart/angular/chart-directives";
import { NativeScriptFormsModule } from "nativescript-angular/forms";
import { ReactiveFormsModule } from "@angular/forms";
// import { sharedModule } from "../shared-mudule";
import { NativeScriptUIAutoCompleteTextViewModule } from "nativescript-ui-autocomplete/angular/autocomplete-directives";
import { DatePipe } from "@angular/common";
import { ChartService } from "~/app/services/chart/chart.service";
import { ActionService } from "~/app/services/action/action.service";
import { MonitorService } from "~/app/services/monitor/monitor.service";
import { ReportsService } from "~/app/services/reports/reports-service";
import { AppNotificationService } from "~/app/services/app-notification-service";
import { AdmissionDetailsComponent } from "./admission-details/admission-details.component";
import { PersonAccompanyingDetailsComponent } from "./person-accompanying-details/person-accompanying-details.component";
import { ReasonForAdmissionDetailsComponent } from "./reason-for-admission-details/reason-for-admission-details.component";
import { HistoryOfPresentIllnessComponent } from "./history-of-present-Illness/history-of-present-Illness.component";
import { PastHistoryAboutHealthComponent } from "./past-history-about-health/past-history-about-health.component";
import { TreatmentBeforeAdmissionComponent } from "./treatment-before-admission/treatment-before-admission.component";
import { InvestigationBeforeAdmissionComponent } from "./investigation-before-admission/investigation-before-admission.component";
import { FamilyHistoryComponent } from "./family-history/family-history.component";
import { PersonalHistoryComponent } from "./personal-history/personal-history.component";
import { AllergiesComponent } from "./allergies/allergies.component";
import { MedicalDetailsComponent } from "./medical-details/medical-details.component";
import { PatientInfoComponent } from "./patient-info/patient-info.component";
import { PatientDetailsComponent } from "./patient-details.component";



@NgModule({
    imports: [
        NativeScriptCommonModule,
        PatientDetailsRoutingModule,
        NativeScriptUIListViewModule,
        NativeScriptUIChartModule,
        NativeScriptFormsModule,
        ReactiveFormsModule,
        NativeScriptFormsModule,
        NativeScriptUIAutoCompleteTextViewModule
    ],
    declarations: [
        AdmissionDetailsComponent,
        AllergiesComponent,
        FamilyHistoryComponent,        
        HistoryOfPresentIllnessComponent,        
        InvestigationBeforeAdmissionComponent,
        MedicalDetailsComponent,
        PastHistoryAboutHealthComponent,
        PatientInfoComponent,
        PersonAccompanyingDetailsComponent,
        PersonalHistoryComponent,
        ReasonForAdmissionDetailsComponent,
        PatientDetailsComponent,
        TreatmentBeforeAdmissionComponent,
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
    ]
})
export class PatientDetailsModule {
    constructor() {
        console.log('PatientDetails    Module muodule initiate');
    }
}
