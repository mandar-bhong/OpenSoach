import { Component, OnInit, ViewChild } from '@angular/core';
import { ReportContainerModel } from '../../../models/ui/report-models';
import { ProdServicepointService } from '../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { ReportService } from '../../../services/report.service';
import { MatTableDataSource, MatPaginator } from '@angular/material';

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
    switch (this.dataModel.selecteddateoption) {
      case '0':
        this.dataModel.enddate = new Date();
        this.dataModel.startdate = new Date(this.dataModel.enddate.getFullYear(), this.dataModel.enddate.getMonth(), 1);
        break;
      case '1':
        break;
    }

    this.reportService.getReportData({
      spid: this.dataModel.selectedsp.spid,
      reportid: 1,
      startdate: this.dataModel.startdate,
      enddate: this.dataModel.enddate
    }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        const parsedHeaderData = JSON.parse(payloadResponse.data.reportheader);
        this.parsedHeaderArray = parsedHeaderData.en;
        this.parsedDataValue = payloadResponse.data.reportdata;
        this.dataSource = new MatTableDataSource<any>(this.parsedDataValue);
        console.log('datasource', this.dataSource);
        this.dataSource.paginator = this.paginator;
      }
    });
  }

  download() {

  }

}
