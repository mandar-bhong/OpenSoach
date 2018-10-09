import { ChangeDetectorRef, Component, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatPaginator, MatSort, MatTableDataSource } from '@angular/material';
import { Router, ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';
import { JobService } from '../../../services/job.service';
import {
  JobDetailsDataListResponse, OwnerResponse, VehicleFullDetails,
  JobTxndata, JobTrnVehicleResponse, JobDetailslistResponse
} from '../../../models/api/job-models';
import { JobDetailsModel, JobTransaction } from '../../../models/ui/job-models';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';

@Component({
  selector: 'app-job-details',
  templateUrl: './job-details.component.html',
  styleUrls: ['./job-details.component.css']
})
export class JobDetailsComponent implements OnInit, OnDestroy {
  dataModel = new JobDetailsModel();
  dataModelDetails = new JobTransaction();
  selectedoption = '0';
  timline = false;
  list = true;
  dataSource;
  listdata = [];
  routeSubscription: Subscription;
  callbackUrl;
  tokendata = new JobTransaction();
  inprogressdata = new JobTransaction();
  createddata = new JobTransaction();
  deliverddata = new JobTransaction();
  completeddata = new JobTransaction();
  inprogress = [];
  inprogressarray = [];
  numbercompleted;
  numberdelivered;
  displayedColumns = ['txndate', 'fopcode', 'taskname', 'comment', 'cost'];
  sortByColumns = [{ text: 'Time', value: 'txndate' },
  { text: 'Service Personnel', value: 'fopcode' },
  { text: 'Activity', value: 'taskname' },
  { text: 'Notes', value: 'comment' },
  { text: 'Tentative Price', value: 'cost' }
  ];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  refreshTable: EventEmitter<null> = new EventEmitter();
  filteredrecords = 0;
  constructor(private jobService: JobService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) { }

  ngOnInit() {
    this.dataModel.details = new VehicleFullDetails;
    this.dataModel.details.ownerdetails = new OwnerResponse();
    this.deliverddata = new JobTransaction();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['id']) {
        this.dataModel.vehicleid = Number(params['id']);
        this.getJobsDetails();
        this.dataModel.tokenid = Number(params['id']);
        this.getJobsDetailslist();
        this.getDataList();
      }
      this.callbackUrl = params['callbackurl'];
    });
  }
  optionChange() {
    if (this.selectedoption === '1') {
      this.list = false;
      this.timline = true;
    } else {
      this.list = true;
      this.timline = false;
    }
  }
  getJobsDetails() {
    this.jobService.getJobsDetails({ recid: this.dataModel.vehicleid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data != null) {
          this.dataModel.copyFromDetails(payloadResponse.data);
        }
      }
    });
  }
  getJobsDetailslist() {
    this.jobService.getJobsDetailsList({ recid: this.dataModel.tokenid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.dataModel.copyFormList(payloadResponse.data);

          if (this.dataModel.transactions.find(a => a.status === 1) != null) {
            this.tokendata = this.dataModel.transactions.find(a => a.status === 1);
            console.log('tokendata', this.tokendata);
          }

          if (this.dataModel.transactions.find(a => a.status === 4) != null) {
            this.inprogressdata = this.dataModel.transactions.find(a => a.status === 4);
            console.log('inprogressdata', this.inprogressdata);
          }
          // this.dataModel.transactions.forEach(element => {
          //   if (element.status === 4) {
          //     this.inprogressarray.push(element);
          //     console.log('inprogressarray', this.inprogressarray);
          //   }
          // });
          if (this.dataModel.transactions.find(a => a.status === 3) != null) {
            this.createddata = this.dataModel.transactions.find(a => a.status === 3);
            console.log('createddata', this.createddata);
          }

          if (this.dataModel.transactions.find(a => a.status === 5) != null) {
            this.completeddata = this.dataModel.transactions.find(a => a.status === 5);
            console.log('completeddata', this.completeddata);
          }

          if (this.dataModel.transactions.find(a => a.status === 6) != null) {
            this.deliverddata = this.dataModel.transactions.find(a => a.status === 6);
            console.log('deliverddata testing', this.deliverddata);
          }

        }
      }
    });
  }
  getDataList() {
    this.jobService.getJobsDetailsList({ recid: this.dataModel.tokenid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        const ab = new JobDetailsModel();
        ab.copyFormList(payloadResponse.data);
        this.dataModel.transactions.forEach(element => {
          if (element.status === 4) {
            this.listdata.push(element);
            console.log('listdata', this.listdata);
            const abc = this.listdata.lastIndexOf(element);
            this.numbercompleted = abc + 4;
            const cde = this.listdata.lastIndexOf(element);
            this.numberdelivered = abc + 5;
          }
        });
        this.dataSource = new MatTableDataSource<JobTransaction>(this.listdata);
        console.log('this.dataSource', this.dataSource.data);
        this.dataSource.sort = this.sort;
        this.dataSource.paginator = this.paginator;
      }
    });
  }

  closeForm() {
    this.router.navigate([this.callbackUrl], { skipLocationChange: true });
  }
  // getTotalCost() {
  //   // calculate total price col
  //   return this.dataSource.map(t => t.price).reduce((acc, value) => acc + value, 0);
  // }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
