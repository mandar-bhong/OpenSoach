import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { MonitorService } from '~/app/services/monitor/monitor.service';
import { ActionService } from '~/app/services/action/action.service';

export class MonitorUiModel {
    Comment: number;
    Value: number;
}
export class BloodUiHighModel {
    Comment: number;
    Value: number;
}
export class BloodUiLowModel {
    Comment: number;
    Value: BloodMonitorModel;
}
export class BloodMonitorModel {
    high: string;
    low: string
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
    bloodpresHighListItems = new ObservableArray<BloodUiHighModel>();
    bloodpresLowListItems = new ObservableArray<BloodUiHighModel>();
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
    dialogOpen = false;
    tempUIChart = false;
    respirationUIChart = false;
    pulseUIChart = false;
    constructor(private monitorService: MonitorService,
        private act: ActionService) {
            
    }

    ngOnInit() {

        this.bloodpressure();
        this.respiration();
        this.pulse();
        this.temperature();
        // this.filter24hr();
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
        this.majorStepUnit = "Day";
        this.monitorService.getTempActionTxn().then(
            (val) => {
                val.forEach(item => {
                    let temperatureListItem = new MonitorUiModel();
                    // temperatureListItem = item;
                    const testdata = JSON.parse(item.txn_data);
                    temperatureListItem.Value = Number(testdata.value);

                    const getDBDate = new Date(item.txn_date);
                    // const asc_date =
                    temperatureListItem.Comment = getDBDate.getTime();
                    // console.log('getDBDate', getDBDate);
                    // filter data condition 24 hr and last 3 days
                    // if (getDBDate >= this.startDateTime && getDBDate <= this.endDateTime) {

                    // }
                    this.tempListItems.push(temperatureListItem);
                    // this.tempListItems.filter(a => a.Va)

                });
                // console.log('filter data TempListItems', this.tempListItems);
            },
            (error) => {
                console.log("getChartData error:", error);
            }
        );
    }

    bloodpressure() {
        this.bloodpresHighListItems = new ObservableArray<BloodUiHighModel>();
        this.bloodpresLowListItems = new ObservableArray<BloodUiHighModel>();
        this.majorStepUnit = "Day";
        this.monitorService.getBloodPreActionTxn().then(
            (val) => {
                val.forEach(item => {
                    console.log('component bloodpressure', item);
                    let bloodpresHighListItem = new BloodUiHighModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.txn_date);
                    bloodpresHighListItem.Comment = getDBDate.getTime();               
                    bloodpresHighListItem.Value = Number(testdata.value.high);
                    console.log('testdata', testdata);
                    console.log('value high',bloodpresHighListItem.Value);

                    let bloodpresLowListItem = new BloodUiHighModel();
                    bloodpresLowListItem.Comment = getDBDate.getTime();  
                    bloodpresLowListItem.Value = Number(testdata.value.low);
                    console.log('value low',bloodpresHighListItem.Value);

                    this.bloodpresHighListItems.push(bloodpresHighListItem);
                    this.bloodpresLowListItems.push(bloodpresLowListItem);
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
        this.majorStepUnit = "Day";
        this.monitorService.getRespirationActionTxn().then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme', item);
                    let respirationListItem = new MonitorUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.txn_date);
                    respirationListItem.Comment = getDBDate.getTime();
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
        this.majorStepUnit = "Day";
        this.monitorService.getPulseActionTxn().then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme', item);
                    let pulseListItem = new MonitorUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.txn_date);
                    pulseListItem.Comment = getDBDate.getTime();
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
    showDialog() {
        this.dialogOpen = true;
        this.tempUIChart = true;
    }

    closeDialog() {
        this.dialogOpen = false;
        this.tempUIChart = false;
    }
    showrespDialog() {
        this.dialogOpen = true;
        this.tempUIChart = false;
        this.respirationUIChart = true;
    }

    closerespDialog() {
        this.dialogOpen = false;
        this.tempUIChart = false;
        this.respirationUIChart = false;
    }
    showPuleDialog() {
        this.dialogOpen = true;
        this.tempUIChart = false;
        this.respirationUIChart = false;
        this.pulseUIChart = true;
    }

    closePuleDialog() {
        this.dialogOpen = false;
        this.tempUIChart = false;
        this.respirationUIChart = false;
        this.pulseUIChart = false;
    }
}