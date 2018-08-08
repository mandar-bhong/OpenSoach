import { ChangeDetectorRef, Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';

import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { StatusChangeRequest } from '../../../../models/api/patient-models';
import { PatientDataModel } from '../../../../models/ui/patient-models';
import { PatientService } from '../../../../services/patient.service';


@Component({
  selector: 'app-patient-view',
  templateUrl: './patient-view.component.html',
  styleUrls: ['./patient-view.component.css']
})
export class PatientViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['patientname', 'emergencycontactno', 'spname', 'bedno', 'status', 'action'];
  sortByColumns = [{ text: 'Patient Name', value: 'patientname' },
  { text: 'Emergency Contact Number', value: 'emergencycontactno' },
  { text: 'Ward', value: 'ward' },
  { text: 'Bed/Room Number', value: 'bedno' }
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

  constructor(public patientService: PatientService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private ch: ChangeDetectorRef) { }

  ngOnInit() {
    this.paginator.pageSize = 5;
    this.getDataList();
    this.patients = new PatientDataModel();
  }
  getDataList() {
    this.patientService.getDataList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        payloadResponse.data.forEach(value => {
          const a = new PatientDataModel();
          a.copyFrom(value);
          this.patient.push(a);
          console.log('this.dataSource', this.dataSource);
          this.dataSource = new MatTableDataSource<PatientDataModel>(this.patient);
          this.dataSource.sort = this.sort;
          this.dataSource.paginator = this.paginator;
        });
      } else {
        this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
      }
    });
  }

  setSelectedPatient(patient: PatientDataModel) {
    this.selectedPatient = patient;
  }

  changestatus() {
    const statusChangeRequest = new StatusChangeRequest();
    statusChangeRequest.status = 2;
    statusChangeRequest.patientid = this.selectedPatient.patientid;
    // statusChangeRequest.discharge = new Date();
    this.patientService.updateStatus(statusChangeRequest).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.appNotificationService.success();
        console.log('payloadResponse.issuccess', payloadResponse.issuccess);
        this.selectedPatient.status = 2;
      }
    });

  }
  viewDetails(id: number) {
    this.router.navigate(['patients', 'patient_chart'], { queryParams: { id: id, callbackurl: 'patients' }, skipLocationChange: true });
  }
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
