import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';

import { ServicePointWithConfigurationResponse } from '../../../../../prod-shared/models/api/service-configuration-models';
import { ProdServicepointService } from '../../../../../prod-shared/services/servicepoint/prod-servicepoint.service';
import { SpServiceConfService } from '../../../../../prod-shared/services/spservice/sp-service-conf.service';
import { SpServiceTxnService } from '../../../../../prod-shared/services/spservice/sp-service-txn.service';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import {
  ChartDataViewModel,
  ChartTimeSlot,
  ChartTransactionModel,
  ChartTxnSlot,
} from '../../../models/ui/chart-conf-models';
import { FloatingMenu, FloatingMenuItem } from '../../../../../shared/models/ui/floating-menu';
import { DEFAULT_PAGE_MENU } from '../../../../../shared/app-common-constants';
import { FloatingButtonMenuService } from '../../../../../shared/services/floating-button-menu.service';

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
  spidparam: number;

  selectedDate = new Date();

  constructor(private route: ActivatedRoute,
    private spServiceConfService: SpServiceConfService,
    private spServiceTxnService: SpServiceTxnService,
    public prodServicepointService: ProdServicepointService,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
    private floatingButtonMenuService: FloatingButtonMenuService) { }

  ngOnInit() {
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['spid']) {
        this.spidparam = Number(params['spid']);
      }
    });
    // this.dataModel.splist = ServicePointWithConfigurationResponse[];
    this.getServicepointList();
    this.setFloatingMenu();
  }

  getServicepointList() {
    this.spServiceConfService.getServicePointsWithConfigurations().subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        this.dataModel.splist = payloadResponse.data;

        if (this.spidparam) {
          this.dataModel.selectedsp = this.dataModel.splist.find(sp => sp.spid === this.spidparam);
        } else if (this.dataModel.splist.length > 0) {
          this.dataModel.selectedsp = this.dataModel.splist[0];
        }

        this.getConfiguration();
      }
    });
  }

  getConfiguration() {
    this.isDataLoaded = false;
    if (!this.dataModel.selectedsp) {
      return;
    }

    if (this.dataModel.selectedsp.servconfid > 0) {
      this.spServiceConfService.getServiceConf({ recid: this.dataModel.selectedsp.servconfid }).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.dataModel.copyFromConfiguration(payloadResponse.data);
          this.createChartSlots();
          this.getChartTransactions();
        }
      });
    } else {
      this.appNotificationService.info(this.translatePipe.transform('CHART_DATA_NO_CHART_CONFIGURED'));
    }
  }

  getChartTransactions() {
    this.dataModel.startdate = new Date(this.selectedDate);
    this.dataModel.enddate = new Date(this.selectedDate);
    this.dataModel.startdate.setHours(0, 0, 0, 0);
    this.dataModel.enddate.setHours(24, 0, 0, 0);
    this.spServiceTxnService.getServiceTransactions({
      spid: this.dataModel.selectedsp.spid,
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

  selectedSpChange(sp: ServicePointWithConfigurationResponse) {
    this.dataModel.selectedsp = sp;
    this.getConfiguration();
  }

  setFloatingMenu() {
    const floatingMenu = new FloatingMenu();
    floatingMenu.menuInstanceKey = DEFAULT_PAGE_MENU;
    floatingMenu.items = new Array<FloatingMenuItem>();
    const item = new FloatingMenuItem();
    item.icon = 'view_list';
    item.title = 'Chart Templates';
    item.navigate = true;
    item.url = 'charts/templatelist';
    item.data = { queryParams: { callbackurl: 'servicepoints' }, skipLocationChange: true };
    floatingMenu.items.push(item);
    this.floatingButtonMenuService.setFloatingMenu(floatingMenu);
  }
}
