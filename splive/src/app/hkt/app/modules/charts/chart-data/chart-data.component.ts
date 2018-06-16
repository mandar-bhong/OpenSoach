import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';

import { SpServiceConfService } from '../../../../../prod-shared/services/spservice/sp-service-conf.service';
import { SpServiceTxnService } from '../../../../../prod-shared/services/spservice/sp-service-txn.service';
import {
  ChartDataViewModel,
  ChartTimeSlot,
  ChartTransactionModel,
  ChartTxnSlot,
} from '../../../models/ui/chart-conf-models';

@Component({
  selector: 'app-chart-data',
  templateUrl: './chart-data.component.html',
  styleUrls: ['./chart-data.component.css']
})
export class ChartDataComponent implements OnInit, OnDestroy {

  dataModel = new ChartDataViewModel();
  routeSubscription: Subscription;
  isDataLoaded = false;
  optionYesterday = 1;
  optionToday = 0;
  optionCustomDate = 2;

  selectedDate = new Date();

  constructor(private route: ActivatedRoute,
    private spServiceConfService: SpServiceConfService,
    private spServiceTxnService: SpServiceTxnService) { }

  ngOnInit() {
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      this.dataModel.spid = Number(params['spid']);
      this.dataModel.servconfid = Number(params['servconfid']);
      this.getConfiguration();
    });
  }

  getConfiguration() {
    this.spServiceConfService.getServiceConf({ recid: this.dataModel.servconfid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.copyFromConfiguration(payloadResponse.data);
        this.createChartSlots();
        this.getChartTransactions();
      }
    });
  }

  getChartTransactions() {
    this.dataModel.startdate = new Date(this.selectedDate.setHours(0, 0, 0, 0));
    this.dataModel.enddate = new Date(this.selectedDate.setHours(24, 0, 0, 0));
    this.spServiceTxnService.getServiceTransactions({
      spid: this.dataModel.spid,
      startdate: this.dataModel.startdate,
      enddate: this.dataModel.enddate
    }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.copyFromTransactions(payloadResponse.data);
        this.createChartTxns();
        this.setTxnsInChartSlots();
        this.isDataLoaded = true;
      }
    });

    console.log('datamodel', this.dataModel);
  }
  createChartSlots() {
    this.dataModel.timeslots = [];
    let i = this.dataModel.variableconf.timeconf.starttime;
    while (i < this.dataModel.variableconf.timeconf.endtime) {
      const d = new Date();
      d.setHours(0, i);
      const slot = new ChartTimeSlot();
      slot.slotdisplaytext = d.toLocaleTimeString('en-us',
        { hour12: true, hour: '2-digit', minute: '2-digit' });
      slot.slotstarttime = i;
      i = i + this.dataModel.variableconf.timeconf.interval;
      slot.slotendtime = i;
      this.dataModel.timeslots.push(slot);
    }
  }

  createChartTxns() {
    this.dataModel.tasktxnslotmap = new Map<string, ChartTxnSlot[]>();
    this.dataModel.variableconf.taskconf.tasks.forEach(task => {
      const chartTxnSlots: ChartTxnSlot[] = [];

      this.dataModel.timeslots.forEach(slot => {
        const chartTxnSlot = new ChartTxnSlot();
        chartTxnSlot.slot = slot;
        chartTxnSlots.push(chartTxnSlot);
      });
      this.dataModel.tasktxnslotmap.set(task.taskname, chartTxnSlots);
    });
  }

  setTxnsInChartSlots() {
    this.dataModel.txns.forEach(txn => {
      const chartTxnSlots = this.dataModel.tasktxnslotmap.get(txn.txndata.taskname);
      if (chartTxnSlots && chartTxnSlots.length > 0) {
        const chartTxnSlot = chartTxnSlots.find(txnSlot => txnSlot.slot.slotstarttime === txn.txndata.slotstarttime
          && txnSlot.slot.slotendtime === txn.txndata.slotendtime);
        if (chartTxnSlot) {
          // TODO: set the operator name here
          txn.fopname = txn.fopcode;
          chartTxnSlot.txn = txn;
        } else {
          // TODO: need to show in exceptions that slot no more exist in chart
        }
      } else {
        // TODO: need to show in exceptions that task no more exist in chart
      }

    });
  }

  optionChange() {
    switch (this.dataModel.selecteddateoption) {
      case '0':
        this.selectedDate = new Date();
        this.getChartTransactions();
        break;
      case '1':
        this.selectedDate = new Date(new Date().setDate(new Date().getDate() - 1));
        this.getChartTransactions();
        break;
    }
  }

  dateChanged(value: any) {
    this.selectedDate = value;
    this.getChartTransactions();
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }

  setStatusCss(txn: ChartTransactionModel) {
    switch (txn.status) {
      case 1:
        return 'onTimeTask';
      case 2:
        return 'delayedTask';
    }
  }
}
