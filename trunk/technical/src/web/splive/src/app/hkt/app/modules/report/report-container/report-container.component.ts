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
    this.dataModel.selecteddateoption = '0';
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

  optionChange() {
    switch (this.dataModel.selecteddateoption) {
      case '0':
        this.dataModel.enddate = new Date();
        this.dataModel.startdate = new Date(this.dataModel.enddate.getFullYear(), this.dataModel.enddate.getMonth(), 1);
        break;
      case '1':
        break;
    }
  }

  view() {
    if (this.dataModel.enddate >= this.dataModel.startdate) {
      this.reportService.getReportData(this.createReportRequest()).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.summaryheader = payloadResponse.data[0].reportheader;
          this.summarydata = payloadResponse.data[0].reportdata;
          this.listheader = payloadResponse.data[1].reportheader;
          this.dataSource = new MatTableDataSource<any>(payloadResponse.data[1].reportdata);
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
        this.reportService.saveReport(payloadResponse, 'Task Report.xlsx');
      }
    });
  }

  createReportRequest(): ReportRequestParams {
    const reportParams = new ReportRequestParams();
    const requestSummary = new ReportRequest();
    const requestList = new ReportRequest();

    requestSummary.queryparams = [];
    requestSummary.lang = 'en';

    requestList.queryparams = [];
    requestList.lang = 'en';

    if (this.dataModel.selectedsp) {
      requestSummary.queryparams.push(this.dataModel.selectedsp.spid);
      requestSummary.reportcode = 'TASK_SUMMARY_SP';
      requestList.queryparams.push(this.dataModel.selectedsp.spid);
      requestList.reportcode = 'TASK_LIST_SP';
    } else {
      requestSummary.reportcode = 'TASK_SUMMARY_ALL';
      requestList.reportcode = 'TASK_LIST_ALL';
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
    }

    reportParams.reportreq = [];
    reportParams.reportreq.push(requestSummary);
    reportParams.reportreq.push(requestList);
    console.log('request', reportParams);
    return reportParams;
  }

}
