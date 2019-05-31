import { Component, OnInit } from '@angular/core';

import { FeedbackSummaryModel } from '../../../models/ui/dashboard-models';
import { DashboardService } from '../../../services/dashboard.service';

@Component({
  selector: 'app-feedback-summary',
  templateUrl: './feedback-summary.component.html',
  styleUrls: ['./feedback-summary.component.css',
    '../default-dashboard/default-dashboard.component.css']
})
export class FeedbackSummaryComponent implements OnInit {

  feedbacksummary = new FeedbackSummaryModel();

  selectedoption = '0';
  constructor(private dashboardService: DashboardService) { }

  ngOnInit() {
    this.getFeedbackSummaryTillDate();
  }

  getRatingStar(rating: number, averagerating: number) {
    if (rating <= averagerating) {
      return 'fa fa-star ratingStar filled';
    } else {
      if (rating - averagerating <= .5) {
        return 'fa fa-star-half-o ratingStar filled';
      } else {
        return 'fa fa-star ratingStar notfilled';
      }
    }
  }

  optionChange() {
    if (this.selectedoption === '1') {
      this.getFeedbackSummaryThisMonth();
    } else {
      this.getFeedbackSummaryTillDate();
    }
  }

  getFeedbackSummaryTillDate() {
    this.dashboardService.getFeedbackSummary().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.feedbacksummary.copyFrom(payloadResponse.data);
      }
    });
  }

  getFeedbackSummaryThisMonth() {
    const dt = new Date();
    const firstDayofMonth = new Date(dt.getFullYear(), dt.getMonth(), 1);

    this.dashboardService.getFeedbackSummary(
      { spid: undefined, startdate: firstDayofMonth, enddate: dt }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.feedbacksummary.copyFrom(payloadResponse.data);
        }
      });
  }
}
