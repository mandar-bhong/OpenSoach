import { Component, OnInit } from '@angular/core';
import { JobFilterModel } from '../../../../models/ui/job-models';
import { EnumDataSourceItem } from '../../../../../../shared/models/ui/enum-datasource-item';
import { JobFiltrRequest } from '../../../../models/api/job-models';
import { JobService } from '../../../../services/job.service';

@Component({
  selector: 'app-job-search',
  templateUrl: './job-search.component.html',
  styleUrls: ['./job-search.component.css']
})
export class JobSearchComponent implements OnInit {
  dataModel = new JobFilterModel();
  isExpanded = false;
  jobStates: EnumDataSourceItem<number>[];
  selectedDate = new Date();
  selecteddateoption = '1';
  constructor(private jobService: JobService) { }

  ngOnInit() {
    this.jobStates = this.jobService.getJobStates();
  }
  search() {
    this.isExpanded = false;
    const jobFiltrRequest = new JobFiltrRequest();
    this.dataModel.copyTo(jobFiltrRequest);
    this.jobService.dataListSubjectTrigger(jobFiltrRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }
  optionChange() {
    switch (this.selecteddateoption) {
      case '0':
        this.selectedDate = new Date();
        this.dataModel.startdate = new Date(this.selectedDate);
        this.dataModel.enddate = new Date(this.selectedDate);
        this.dataModel.startdate.setHours(0, 0, 0, 0);
        this.dataModel.enddate.setHours(24, 0, 0, 0);
        console.log('this.dataModel.startdate', this.dataModel.startdate);
        console.log('this.dataModel.enddate', this.dataModel.enddate);
        break;
      case '2':
        break;
      case '1':
        const currentDate = new Date();
        this.dataModel.enddate = new Date(Date.UTC(
        currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), currentDate.getHours(), currentDate.getMinutes()));
        console.log('this.dataModel.enddate', this.dataModel.enddate);
        const ticks = Date.UTC(this.dataModel.enddate.getUTCFullYear(), this.dataModel.enddate.getUTCMonth() - 23, 1);
        this.dataModel.startdate = new Date(ticks);
        console.log('this.dataModel.startdate', this.dataModel.startdate);
        break;
    }
  }
  dateChanged(value: any) {
    this.selectedDate = value;
    this.dataModel.startdate = new Date(this.selectedDate);
    this.dataModel.enddate = new Date(this.selectedDate);
    this.dataModel.startdate.setHours(0, 0, 0, 0);
    this.dataModel.enddate.setHours(24, 0, 0, 0);
    console.log('this.dataModel.startdate', this.dataModel.startdate);
    console.log('this.dataModel.enddate', this.dataModel.enddate);
  }
}
