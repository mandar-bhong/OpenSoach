import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PatientsRoutingModule } from './patients-routing.module';
import { PatientListComponent } from './patient-list/patient-list.component';
import { PatientAddComponent } from './patient-add/patient-add.component';
import { PatientViewComponent } from './patient-list/patient-view/patient-view.component';
import { MaterialModules } from '../../../../shared/modules/material/material-modules';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ProdCommonModule } from '../../../../prod-shared/prod-common.module';
import { AppCommonModule } from '../../../../shared/app-common.module';
import { PatientChartComponent } from './patient-chart/patient-chart.component';
import { PatientDayWiseComponent } from './patient-day-wise/patient-day-wise.component';

import { PatientsPersonalDetailComponent } from './patients-personal-detail/patients-personal-detail.component';
import { PatientPersonAccompaniesComponent } from './patient-person-accompanies/patient-person-accompanies.component';

import { PatientSearchComponent } from './patient-list/patient-search/patient-search.component';
import { PatientCheckListComponent } from './patient-check-list/patient-check-list.component';
import { PatientCheckSearchComponent } from './patient-check-search/patient-check-search.component';
import { PatientAdmissionComponent } from './patient-admission/patient-admission.component';
import { PatientDetailsComponent } from './patient-details/patient-details.component';
import { PatientMedicalComponent } from './patient-medical/patient-medical.component';
import { PatientReportComponent } from './patient-report/patient-report.component';
import { AmazingTimePickerModule, AmazingTimePickerService } from 'amazing-time-picker';
// import { Mul } from './patient-medical/app-multiple-comment-input/app-multiple-comment-input.component';
import { MultipleCommentInputComponent } from './patient-medical/multiple-comment-input/multiple-comment-input.component';
import { MedicalPersonalHistoryComponent } from './patient-medical/medical-personal-history/medical-personal-history.component';
import { MatRadioModule } from '@angular/material/radio';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MedicalContainerComponent } from './medical-container/medical-container.component';
import { ScheduleDetailsComponent } from './schedule-details/schedule-details.component';
import { DoctorOrdersComponent } from './doctor-orders/doctor-orders.component';
import { TransactionDetailsComponent } from './transaction-details/transaction-details.component';
import { ViewMedicineTransactionComponent } from './transaction-details/view-medicine-transaction/view-medicine-transaction.component';
import { ViewIntakeTransactionComponent } from './transaction-details/view-intake-transaction/view-intake-transaction.component';
import { ViewMonitorTransactionComponent } from './transaction-details/view-monitor-transaction/view-monitor-transaction.component';
import { ViewOutputTransactionComponent } from './transaction-details/view-output-transaction/view-output-transaction.component';


@NgModule({
  imports: [
    CommonModule,
    PatientsRoutingModule,
    MaterialModules,
    ProdCommonModule,
    FormsModule,
    ReactiveFormsModule,
    AppCommonModule,
    AmazingTimePickerModule,
    MatRadioModule,
    MatCheckboxModule
  ],
  declarations: [PatientListComponent,
    PatientAddComponent,
    PatientViewComponent,
    PatientChartComponent,
    PatientDayWiseComponent,
    PatientDetailsComponent,
    PatientsPersonalDetailComponent,
    PatientPersonAccompaniesComponent,
    PatientAdmissionComponent,
    PatientMedicalComponent,
    PatientReportComponent,
    PatientSearchComponent,
    PatientCheckListComponent,
    PatientCheckSearchComponent,
    MedicalContainerComponent,
    ScheduleDetailsComponent,
    DoctorOrdersComponent,
    TransactionDetailsComponent,
    ViewMedicineTransactionComponent,
    ViewIntakeTransactionComponent,
    ViewMonitorTransactionComponent,
    MultipleCommentInputComponent,
    MedicalPersonalHistoryComponent,
    ViewOutputTransactionComponent],
  providers: [
    AmazingTimePickerService,
    MatRadioModule,
    MatCheckboxModule
  ]
})
export class PatientsModule { }
