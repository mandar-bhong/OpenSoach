import { ChangeDetectorRef, Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';

import { ServicepointListResponse } from '../../../../../../prod-shared/models/api/servicepoint-models';
import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';
import { PatientListDataModel } from '../../../../../app/models/ui/patient-models';
import { PatientDetaListResponse, PatientFilterRequest, StatusChangeRequest } from '../../../../models/api/patient-models';
import { PatientService } from '../../../../services/patient.service';


@Component({
  selector: 'app-patient-view',
  templateUrl: './patient-view.component.html',
  styleUrls: ['./patient-view.component.css']
})
export class PatientViewComponent implements OnInit, OnDestroy {
  displayedColumns = ['fname', 'patientregno', 'mobno', 'spid', 'bedno', 'status', 'action'];
  sortByColumns = [{ text: 'Patient Name', value: 'fname' },
  { text: 'Patient Registration Number', value: 'patientregno' },
  { text: 'Emergency Contact Number', value: 'mobno' },
  { text: 'Ward', value: 'spid' },
  { text: 'Bed/Room Number', value: 'bedno' },
    // {text: 'Status',value:'status'}

  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  dataSource;
  filteredrecords = 0;
  isLoadingResults = true;
  patientFilterRequest: PatientFilterRequest;
  dataListFilterChangedSubscription: Subscription;
  showEditForm = false;
  selectedPatient: PatientListDataModel;
  splist: ServicepointListResponse[] = [];

  constructor(public patientService: PatientService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.getServicepointList();
    this.patientFilterRequest = new PatientFilterRequest();
    this.patientFilterRequest.status = 1;
    this.paginator.pageSize = 10;
    this.sort.active = 'fname';
    this.sort.direction = 'asc';
    this.dataListFilterChangedSubscription = this.patientService.dataListSubject.subscribe(value => {
      this.patientFilterRequest = value;
      this.refreshTable.emit();
    });
  }

  // Accept data from ward ie. list of ward
  getServicepointList() {
    this.patientService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.splist = payloadResponse.data;
        this.setDataListing();
      }
    });

  }

  setSelectedPatient(patient: PatientListDataModel) {
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
        this.selectedPatient.status = 2;
      }
    });

  }
  setDataListing(): void {
    this.sort.sortChange.subscribe(() => this.paginator.pageIndex = 0);
    this.refreshTable.subscribe(() => this.paginator.pageIndex = 0);
    merge(this.sort.sortChange, this.paginator.page, this.refreshTable)
      .pipe(
        startWith({}),
        switchMap(() => {
          this.isLoadingResults = true;
          return this.getDataList();
        }),
        map(data => {
          this.isLoadingResults = false;
          return data;
        }),
      ).subscribe(
        payloadResponse => {
          if (payloadResponse && payloadResponse.issuccess) {
            this.filteredrecords = payloadResponse.data.filteredrecords;
            this.dataSource = payloadResponse.data.records;
            if (this.filteredrecords === 0) {
              this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];
          }
        }
      );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<PatientDetaListResponse>>> {
    const dataListRequest = new DataListRequest<PatientFilterRequest>();
    dataListRequest.filter = this.patientFilterRequest;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.patientService.getDataList(dataListRequest);
  }
  viewDetails(id: number, addid: number) {
    //setting patient id for further use
    this.patientService.patientid = id;
    this.router.navigate(['patients', 'patient_admission'], { queryParams: { id: id, addid: addid, callbackurl: 'patients' }, skipLocationChange: true });
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
  } // end 

  getSPName(value: number) {
    if (this.splist && value) {
      return this.splist.find(a => a.spid === value).spname;
    }
  }

}
