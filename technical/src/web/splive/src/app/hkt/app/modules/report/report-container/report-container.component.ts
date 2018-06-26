import { Component, OnInit, ViewChild } from '@angular/core';
import { ReportContainerModel } from '../../../models/ui/report-models';
import { ProdServicepointService } from '../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { ReportService } from '../../../services/report.service';
import { MatTableDataSource, MatPaginator,MatSort } from '@angular/material';
import { ReportRequest } from '../../../models/api/report-models';

@Component({
  selector: 'app-report-container',
  templateUrl: './report-container.component.html',
  styleUrls: ['./report-container.component.css']
})
export class ReportContainerComponent implements OnInit {

  dataModel = new ReportContainerModel();
  dataSource;
  parsedDataValue: string[][];
  parsedHeaderArray = [];
  @ViewChild(MatPaginator)
  paginator: MatPaginator;
  @ViewChild(MatSort)
  sort: MatSort;
  constructor(private prodServicepointService: ProdServicepointService,
    private reportService: ReportService) {
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
    this.reportService.getReportData(this.createReportRequest()).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        const parsedHeaderData = JSON.parse(payloadResponse.data.list.listheader);
        this.parsedHeaderArray = parsedHeaderData.en;
       // this.parsedDataValue = payloadResponse.data.reportdata;
        this.dataSource = new MatTableDataSource<any>(payloadResponse.data.list.listdata);
        console.log('datasource', this.dataSource);
        this.dataSource.paginator = this.paginator;
        this.dataSource.sort = this.sort;
      }
    });
  }

  download() {

  }

  createReportRequest(): ReportRequest {
    const request = new ReportRequest();
    request.queryparams = [];
    request.reportcode = 'TASKLIST';
    request.lang = 'en';

    if (this.dataModel.selectedsp) {
      request.queryparams.push(this.dataModel.selectedsp.spid);
    } else {
      request.queryparams.push();
    }


    // switch (this.dataModel.selecteddateoption) {
    //   case '0':
    //     request.queryparams.push(new Date());
    //     request.queryparams.push(new Date(this.dataModel.enddate.getFullYear(), this.dataModel.enddate.getMonth(), 1));
    //     break;
    //   case '1':
    //     request.queryparams.push(this.dataModel.startdate);
    //     request.queryparams.push(new Date(this.dataModel.enddate.getFullYear(),
    //       this.dataModel.enddate.getMonth(), this.dataModel.enddate.getDate() + 1));
    //     break;
    // }

    console.log('request', request);
    return request;
  }

}
