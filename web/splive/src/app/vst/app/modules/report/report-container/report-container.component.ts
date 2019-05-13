import { Component, OnInit, ViewChild } from '@angular/core';
import { ReportContainerModel } from '../../../models/ui/report-models';
import { ProdServicepointService } from '../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { ReportService } from '../../../services/report.service';
import { MatTableDataSource, MatPaginator, MatSort } from '@angular/material';
import { ReportRequest, ReportRequestParams } from '../../../models/api/report-models';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';

@Component({
  selector: 'app-report-container',
  templateUrl: './report-container.component.html',
  styleUrls: ['./report-container.component.css']
})
export class ReportContainerComponent implements OnInit {

  dataModel = new ReportContainerModel();
  dataSource;
  listheader;
  summaryheader = [];
  summarydata = new Array<any[]>();
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  constructor(private prodServicepointService: ProdServicepointService,
    private reportService: ReportService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe) {
    this.dataModel.selecteddateoption = '2';
    this.optionChange();
  }

  ngOnInit() {
    this.getServicepointList();
  }
  getServicepointList() {
    this.prodServicepointService.getServicepointList().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.splist = payloadResponse.data;
      }
    });
  }
  // optionReportChange() {
  //   if (this.selectedoption === '1') {

  //   } else {

  //   }
  // }
  optionChange() {
    switch (this.dataModel.selecteddateoption) {
      case '0':
        this.dataModel.enddate = new Date();
        this.dataModel.startdate = new Date(this.dataModel.enddate.getFullYear(), this.dataModel.enddate.getMonth(), 1);
        break;
      case '1':
        break;
      case '2':
        this.dataModel.startdate = new Date();
        this.dataModel.enddate = new Date();
        this.dataModel.startdate.setHours(0, 0, 0, 0);
        this.dataModel.enddate.setHours(24, 0, 0, 0);
        break;
    }
  }

  view() {
    if (this.dataModel.enddate >= this.dataModel.startdate) {
      this.reportService.getReportData(this.createReportRequest()).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          // this.summaryheader = payloadResponse.data[0].reportheader;
          // this.summarydata = payloadResponse.data[0].reportdata;
          this.listheader = payloadResponse.data[0].reportheader;
          this.dataSource = new MatTableDataSource<any>(payloadResponse.data[0].reportdata);
          console.log('datasource', this.dataSource);
          this.dataSource.paginator = this.paginator;
          this.dataSource.sort = this.sort;
        }
      });
    } else {
      this.appNotificationService.info(this.translatePipe.transform('START_DATE_MUST_BE_BEFORE_END_DATE'));
    }

  }

  download() {
    this.reportService.generateReport(this.createReportRequest()).subscribe((payloadResponse: Blob) => {
      console.log('payloadResponse', payloadResponse);
      if (payloadResponse) {
        this.reportService.saveReport(payloadResponse, 'Consolidated Report.pdf');
        // this.reportService.saveReport(payloadResponse, 'Task Report.pdf');
      }
    });
  }

  createReportRequest(): ReportRequestParams {
    const reportParams = new ReportRequestParams();
    const requestSummary = new ReportRequest();
    const requestList = new ReportRequest();
    reportParams.reportfileformat = 'pdf';

    requestSummary.queryparams = [];
    requestSummary.lang = 'en';

    requestList.queryparams = [];
    requestList.lang = 'en';

    if (this.dataModel.selectedsp) {
      requestSummary.queryparams.push(this.dataModel.selectedsp.spid);
      // requestSummary.reportcode = 'TASK_SUMMARY_SP';
      // requestList.queryparams.push(this.dataModel.selectedsp.spid);

    } else {
      // requestSummary.reportcode = 'TASK_SUMMARY_ALL';
      // requestList.reportcode = 'TASK_LIST_ALL';
      requestList.reportcode = 'CONSOLIDATED_VHL_REPORT';
    }


    switch (this.dataModel.selecteddateoption) {
      case '0':
        requestSummary.queryparams.push(new Date(this.dataModel.enddate.getFullYear(), this.dataModel.enddate.getMonth(), 1));
        requestSummary.queryparams.push(new Date());
        requestList.queryparams.push(new Date(this.dataModel.enddate.getFullYear(), this.dataModel.enddate.getMonth(), 1));
        requestList.queryparams.push(new Date());
        break;
      case '1':
        requestSummary.queryparams.push(this.dataModel.startdate);
        requestSummary.queryparams.push(new Date(this.dataModel.enddate.getFullYear(),
          this.dataModel.enddate.getMonth(), this.dataModel.enddate.getDate() + 1));
        requestList.queryparams.push(this.dataModel.startdate);
        requestList.queryparams.push(new Date(this.dataModel.enddate.getFullYear(),
          this.dataModel.enddate.getMonth(), this.dataModel.enddate.getDate() + 1));
        break;
        case '2':
        // this.dataModel.startdate = new Date();
        this.dataModel.enddate = new Date();
        this.dataModel.startdate.setHours(0, 0, 0, 0);
        this.dataModel.enddate.setHours(24, 0, 0, 0);
        requestList.queryparams.push(this.dataModel.startdate );
        requestList.queryparams.push(this.dataModel.enddate );
        console.log('this.dataModel.enddate', requestList.queryparams);
        console.log('this.dataModel.startdate', requestSummary.queryparams);
        break;
    }

    reportParams.reportreq = [];
    // reportParams.reportreq.push(requestSummary);
    reportParams.reportreq.push(requestList);
    console.log('request', reportParams);
    return reportParams;
  }

}
