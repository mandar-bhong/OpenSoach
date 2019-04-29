import { Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort } from '@angular/material';
import { Router } from '@angular/router';
import { PatientFilterRequest, PatientFilterResponse } from 'app/models/api/hospital-models';
import { PatientInfoForHospitals } from 'app/models/ui/patient-models';
import { HospitalService } from 'app/services/hospital.service';
import { PatientService } from 'app/services/patient.service';
import { merge, Observable, Subscription } from 'rxjs';
import { map, startWith, switchMap } from 'rxjs/operators';
import { DataListRequest, DataListResponse } from '../../../../../../shared/models/api/data-list-models';
import { PayloadResponse } from '../../../../../../shared/models/api/payload-models';
import { TranslatePipe } from '../../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';

@Component({
  selector: 'app-hospital-view',
  templateUrl: './hospital-view.component.html',
  styleUrls: ['./hospital-view.component.css']
})
export class HospitalViewComponent implements OnInit, OnDestroy {

  displayedColumns = ['fname', 'mobno', 'action'];
  sortByColumns = [{ text: 'Patient Name', value: 'fname' },
  { text: 'Emergency Contact Number', value: 'mobno' }
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
  patientFilterResponse: PatientFilterResponse;
  admissionid: number;
  PatientFullName: string;
  HospLname: any;
  HospFname: any;

  constructor(public hospitalService: HospitalService,
    private patientService: PatientService,
    private appNotificationService: AppNotificationService,
    private router: Router,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'fname';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.admissionid = this.hospitalService.admissionid;
    this.dataListFilterChangedSubscription = this.hospitalService.dataListSubject.subscribe(value => {
      this.patientFilterRequest = value;
      this.refreshTable.emit();
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
            let filterDataArray = [];
            filterDataArray = payloadResponse.data.records;
            console.log('main data', filterDataArray);
            const filterItem = filterDataArray.filter(a => a.admissionid != null);
            this.dataSource = filterItem;
            console.log('filter data ', this.dataSource);
            if (this.filteredrecords === 0) {
              this.appNotificationService.info(this.translatePipe.transform('INFO_NO_RECORDS_FOUND'));
            }
          } else {
            this.dataSource = [];
          }
        }
      );
  }

  getDataList(): Observable<PayloadResponse<DataListResponse<PatientFilterResponse>>> {
    const dataListRequest = new DataListRequest<PatientFilterRequest>();
    // dataListRequest.filter = this.patientFilterRequest;
    dataListRequest.filter = new PatientFilterRequest();
    dataListRequest.filter.cpmid = this.hospitalService.cpmid;
    // dataListRequest.filter.admissionid = this.admissionid;
    // console.log("dataListRequest.filter.admissionid", dataListRequest.filter.admissionid);
    console.log("this.hospitalService.cpmid", this.hospitalService.cpmid);
    console.log("dataListRequest.filter.cpmid", dataListRequest.filter.cpmid);
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    console.log("dataListRequest", dataListRequest);
    return this.hospitalService.getDataList(dataListRequest);
  }

  viewDetails(admissionid: number, fname, lname) {
    this.HospFname = fname;
    this.HospLname = lname;

    // setting patient info in service for pathology report componet use
    const patinetInfo = new PatientInfoForHospitals()
    patinetInfo.isvisible = true;
    patinetInfo.patintname = this.HospFname + ' ' + this.HospLname;
    this.patientService.patinetInfo = patinetInfo;
    console.log("this.patientService.patinetInfo",this.patientService.patinetInfo);
  // end
  
    this.hospitalService.admissionid = admissionid;
    this.patientService.admissionid = admissionid;
    this.router.navigate(['patients', 'pathology_report'], { queryParams: { id: admissionid, callbackurl: 'hospitals' }, skipLocationChange: true });
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
