import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { MonitorService } from '~/app/services/monitor/monitor.service';
import { ActionService } from '~/app/services/action/action.service';

export class MonitorUiModel {
    Comment: number;
    Value: number;
}
@Component({
    moduleId: module.id,
    selector: 'monitor',
    templateUrl: './monitor.component.html',
    styleUrls: ['./monitor.component.css']
})

export class MonitorComponent implements OnInit {
    // categoricalSource: { Country: string, Amount: any }[] = [
    // 	{ Country: "08:00 AM", Amount: 98 },
    // 	{ Country: "12:00 AM", Amount: 98.3 },
    // 	{ Country: "04:00 PM", Amount: 99 },
    // 	{ Country: "08:00 PM", Amount: 101 },
    // 	{ Country: "12:00 PM", Amount: 97 },
    // 	{ Country: "04:00 AM", Amount: 100 }
    // ];
    tempListItems = new ObservableArray<MonitorUiModel>();
    bloodpresListItems = new ObservableArray<MonitorUiModel>();
    respirationListItems = new ObservableArray<MonitorUiModel>();
    pulseListItems = new ObservableArray<MonitorUiModel>();

    schedulardata: Schedulardata;
    // filter var
    startDateTime: any;
    endDateTime: any;

    uiStartDate: any;
    uiEndDate: any;

    majorStepUnit: any;
    majorUnit: any;
    // switch active and complited
    completeorpending: string;
    iscompleted: boolean;
    // filter buttton 
    buttonClicked: boolean = true;
    buttonCompleted: boolean = false;



    categoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 99.5 },
        { Country: "test2", Amount1: 99.5 },
        { Country: "test3", Amount1: 99.5 },
        { Country: "test4", Amount1: 99.5 },
        { Country: "test5", Amount1: 99.5 },
        { Country: "test6", Amount1: 99.5 },

    ];
    categoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 97.7 },
        { Country: "test2", Amount1: 97.7 },
        { Country: "test3", Amount1: 97.7 },
        { Country: "test4", Amount1: 97.7 },
        { Country: "test5", Amount1: 97.7 },
        { Country: "test6", Amount1: 97.7 },

    ];

    bloodSourcecategoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 140 },
        { Country: "test2", Amount1: 140 },
        { Country: "test3", Amount1: 140 },
        { Country: "test4", Amount1: 140 },
        { Country: "test5", Amount1: 140 },
        { Country: "test6", Amount1: 140 },


    ];
    bloodSourcecategoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 90 },
        { Country: "test2", Amount1: 90 },
        { Country: "test3", Amount1: 90 },
        { Country: "test4", Amount1: 90 },
        { Country: "test5", Amount1: 89 },
        { Country: "test6", Amount1: 89 },
    ];


    respirationcategoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 135 },
        { Country: "test2", Amount1: 135 },
        { Country: "test3", Amount1: 135 },
        { Country: "test4", Amount1: 135 },
        { Country: "test5", Amount1: 135 },
        { Country: "test6", Amount1: 135 },


    ];
    respirationcategoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 95 },
        { Country: "test2", Amount1: 95 },
        { Country: "test3", Amount1: 95 },
        { Country: "test4", Amount1: 95 },
        { Country: "test5", Amount1: 95 },
        { Country: "test6", Amount1: 95 },
    ];


    pulsecategoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 135 },
        { Country: "test2", Amount1: 135 },
        { Country: "test3", Amount1: 135 },
        { Country: "test4", Amount1: 135 },
        { Country: "test5", Amount1: 135 },
        { Country: "test6", Amount1: 135 },


    ];
    pulsecategoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "test1", Amount1: 95 },
        { Country: "test2", Amount1: 95 },
        { Country: "test3", Amount1: 95 },
        { Country: "test4", Amount1: 95 },
        { Country: "test5", Amount1: 95 },
        { Country: "test6", Amount1: 95 },
    ];
    constructor(private monitorService: MonitorService,
        private act: ActionService) {
    }

    ngOnInit() {

        this.bloodpressure();
        this.respiration();
        this.pulse();
        this.filter24hr();
        // this.gettestdata();
    }
    public filter24hr() {
        this.completeorpending = "24 hr";
        this.iscompleted = false;
        this.buttonCompleted = false;
        this.majorStepUnit = "Hour";
        this.majorUnit = 4;

        const newDate = new Date();
        const isodate = newDate.toISOString();
        console.log('isodate', isodate);
        const startDate = new Date(isodate);
        // console.log('startDate', startDate);
        const after24Hours = startDate.getMinutes() - 1440;
        startDate.setMinutes(after24Hours);
        const tempStart = startDate.toLocaleString();
        this.startDateTime = new Date(tempStart);
        console.log('startTimeTime', this.startDateTime);

        var curr_date = startDate.getDate();
        var curr_month = startDate.getMonth() + 1;
        var curr_year = startDate.getFullYear();
        const uiStartDate = curr_date + "/" + curr_month + "/" + curr_year;
        this.uiStartDate = uiStartDate;
        console.log('uiStartDate', this.uiStartDate);


        this.endDateTime = new Date(isodate);
        console.log('endDateTime', this.endDateTime);
        var end_curr_date = this.endDateTime.getDate();
        var end_curr_month = this.endDateTime.getMonth() + 1;
        var end_curr_year = this.endDateTime.getFullYear();
        const uiEndDate = end_curr_date + "/" + end_curr_month + "/" + end_curr_year;
        this.uiEndDate = uiEndDate;
        console.log('uiEndDate', this.uiEndDate);
  
        this.tempListItems = new ObservableArray<MonitorUiModel>();
        this.temperature();

    }
    public filter3day() {
        this.completeorpending = "3 days";
        this.iscompleted = true;
        this.majorStepUnit = "Hour";
        this.majorUnit = 8;
        this.buttonClicked = false;

        const newDate = new Date();
        const isodate = newDate.toISOString();
        // console.log('isodate', isodate); 
        const startDate = new Date(isodate);
        // console.log('startDate', startDate);
        const after24Hours = startDate.getMinutes() - 4320;
        startDate.setMinutes(after24Hours);
        const tempStart = startDate.toLocaleString();
        this.startDateTime = new Date(tempStart);
        console.log('3 day startTimeTime', this.startDateTime);

        var curr_date = startDate.getDate();
        var curr_month = startDate.getMonth() + 1;
        var curr_year = startDate.getFullYear();
        const uiStartDate = curr_date + "/" + curr_month + "/" + curr_year;
        this.uiStartDate = uiStartDate;

        this.endDateTime = new Date(isodate);
        var end_curr_date = this.endDateTime.getDate();
        var end_curr_month = this.endDateTime.getMonth() + 1;
        var end_curr_year = this.endDateTime.getFullYear();
        const uiEndDate = end_curr_date + "/" + end_curr_month + "/" + end_curr_year;
        this.uiEndDate = uiEndDate;
        console.log('3 day endDateTime', this.endDateTime);
        this.tempListItems = new ObservableArray<MonitorUiModel>();
        this.temperature();
    }
    temperature() {
        this.monitorService.getTempActionTxn().then(
            (val) => {
                val.forEach(item => {
                    let temperatureListItem = new MonitorUiModel();
                    // temperatureListItem = item;
                    const testdata = JSON.parse(item.txn_data);
                    temperatureListItem.Value = Number(testdata.value);

                    const getDBDate = new Date(item.txn_date);
                    temperatureListItem.Comment = getDBDate.getTime();
                    console.log('getDBDate', getDBDate);
                    if (getDBDate >= this.startDateTime && getDBDate <= this.endDateTime) {
                        this.tempListItems.push(temperatureListItem);                        
                    }
                   

                });
                console.log('filter data TempListItems', this.tempListItems);
            },
            (error) => {
                console.log("getChartData error:", error);
            }
        );
    }

    bloodpressure() {
        this.monitorService.getBloodPreActionTxn().then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme bloodpressure', item);
                    let bloodpresListItem = new MonitorUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    bloodpresListItem.Comment = testdata.comment;
                    bloodpresListItem.Value = Number(testdata.value);
                    this.bloodpresListItems.push(bloodpresListItem);
                    // console.log('TempListItems', this.tempListItems);
                });
                // console.log('bloodpressure outside', this.bloodpresListItems);
            },
            (error) => {
                console.log("getChartData error:", error);
            }
        );
    }


    respiration() {
        this.monitorService.getRespirationActionTxn().then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme', item);
                    let respirationListItem = new MonitorUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    respirationListItem.Comment = testdata.comment;
                    respirationListItem.Value = Number(testdata.value);
                    this.respirationListItems.push(respirationListItem);
                    // console.log('respirationListItems', this.respirationListItems);
                });
                // console.log('respirationListItems outside', this.respirationListItems);
            },
            (error) => {
                console.log("getChartData error:", error);
            }
        );
    }

    pulse() {
        this.monitorService.getPulseActionTxn().then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme', item);
                    let pulseListItem = new MonitorUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    pulseListItem.Comment = testdata.comment;
                    pulseListItem.Value = Number(testdata.value);
                    this.pulseListItems.push(pulseListItem);
                    // console.log('pulseListItems', this.pulseListItems);
                });
                // console.log('pulseListItems outside', this.pulseListItems);
            },
            (error) => {
                console.log("getChartData error:", error);
            }
        );
    }
    // gettestdata() {
    //     this.act.getActionTxnList();
    // }
}