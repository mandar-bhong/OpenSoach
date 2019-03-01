import { ChangeDetectorRef, Component, OnInit, EventEmitter, ViewChild, OnDestroy } from '@angular/core';
import { MatSort, MatPaginator } from '@angular/material';
import { Subscription, merge, Observable, } from 'rxjs';
import { Router } from '@angular/router';
import { PatientService } from 'app/services/patient.service';
import { startWith, switchMap, map } from 'rxjs/operators';
import { PayloadResponse } from '../../../../../shared/models/api/payload-models';
import { DataListResponse, DataListRequest } from '../../../../../shared/models/api/data-list-models';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { PatientSearchRequestFilter, PatientSearchResponseFilter } from '../../../models/api/patient-models';
import { PatientCheckListDataModal } from '../../../models/ui/patient-models';
import { PatientAddService } from 'app/services/patient-add.service';

@Component({
  selector: 'app-patient-check-list',
  templateUrl: './patient-check-list.component.html',
  styleUrls: ['./patient-check-list.component.css']
})
export class PatientCheckListComponent implements OnInit, OnDestroy {

  displayedColumns = ['fname', 'lname', 'mobno', 'action'];
  sortByColumns = [{ text: 'First Name', value: 'fname' },
  { text: 'Last Name', value: 'lname' },
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
  dataListFilterChangedSubscription: Subscription;
  displayData: Subscription;
  showEditForm = false;
  selectedPatient: PatientCheckListDataModal;
  patientSearchRequestFilter: PatientSearchRequestFilter;


  constructor(public patientService: PatientService,
    public patientAddService: PatientAddService,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private ch: ChangeDetectorRef) { }

  ngOnInit() {
    this.paginator.pageSize = 10;
    this.sort.active = 'fname';
    this.sort.direction = 'asc';
    this.setDataListing();
    this.dataListFilterChangedSubscription = this.patientAddService.dataListSubject.subscribe((value) => {
      this.patientSearchRequestFilter = new PatientSearchRequestFilter();
      this.patientSearchRequestFilter = value;
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

  getDataList(): Observable<PayloadResponse<DataListResponse<PatientSearchResponseFilter>>> {
    const dataListRequest = new DataListRequest<PatientSearchRequestFilter>();
    dataListRequest.filter = this.patientSearchRequestFilter;
    dataListRequest.page = this.paginator.pageIndex + 1;
    dataListRequest.limit = this.paginator.pageSize;
    dataListRequest.orderby = this.sort.active;
    dataListRequest.orderdirection = this.sort.direction;
    return this.patientAddService.getDataList(dataListRequest);
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

  viewDetails(id: number, addid:number) {
    //setting patient id for further use
    this.patientService.patientid = id;
    this.patientService.admissionid = addid;
    this.router.navigate(['patients', 'add'], { queryParams: { id: id,addid: addid, callbackurl: 'patients' }, skipLocationChange: true });
  }

  ngOnDestroy(): void {
    if (this.dataListFilterChangedSubscription) {
      this.dataListFilterChangedSubscription.unsubscribe();
    }
  }

}
