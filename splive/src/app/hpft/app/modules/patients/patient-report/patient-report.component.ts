import {  ChangeDetectorRef, Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';

import { ServicepointConfigureListResponse } from '../../../../../prod-shared/models/api/service-configuration-models';
import { ServicepointListResponse } from '../../../../../prod-shared/models/api/servicepoint-models';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { MedicalDetailAddRequest, PatientDataAddRequest, PatientDetailAddRequest,PersonDataAddRequest } from '../../../models/api/patient-data-models';
import { PatientDataModel } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { PatientAdmissionComponent } from '../patient-admission/patient-admission.component';
import { StatusChangeRequest } from '../../../models/api/patient-data-models';

@Component({
  selector: 'app-patient-report',
  templateUrl: './patient-report.component.html',
  styleUrls: ['./patient-report.component.css']
})
export class PatientReportComponent implements OnInit, OnDestroy {
  displayedColumns = ['reportname', 'persondate','personview','persondelete','personupload'];
  sortByColumns = [{ text: 'Report Name', value: 'reportname' },
  { text: 'Date', value: 'persondate' },
  { text: 'View', value: 'personview' },
  { text: 'Delete', value: 'persondelete' },
  { text: 'File Upload', value: 'personupload' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  patients = new PatientDataModel();
  patient = [];
  dataSource;
  stat;
  filteredrecords = 0;
  isLoadingResults = true;
  dataListFilterChangedSubscription: Subscription;
  showEditForm = false;
  selectedPatient: PatientDataModel;

  constructor(
    public patientService: PatientService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private ch: ChangeDetectorRef
  ) {

   }

  ngOnInit() {
    this.paginator.pageSize = 10;
    // this.getDataList();
    this.patients = new PatientDataModel();
  }

  setSelectedPatient(patient: PatientDataModel) {
    this.selectedPatient = patient;
  }
  viewDetails(id: number) {
    this.router.navigate(['patients', 'patient_chart'], { queryParams: { id: id, callbackurl: 'patients' }, skipLocationChange: true });
  }
  // closeForm() {
  //   this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  // }
  sortByChanged() {
    this.sort.sortChange.next(this.sort);
  }

  sortDirectionAsc() {
    this.sort.direction = 'asc';
    this.sort.sortChange.next(this.sort);
  }

  sortDirectionDesc() {
    this.sort.direction = 'desc';
    this.sort.sortChange.next(this.sort);
  }
  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }
}
