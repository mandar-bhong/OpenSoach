import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { MonitorService } from '~/app/services/monitor/monitor.service';
import { ActionService } from '~/app/services/action/action.service';
import { TrackballCustomContentData } from 'nativescript-ui-chart';

export class MonitorUiModel {
    Comment: number;
    Value: number;
    Impact: number;
}
export class BloodUiHighModel {
    Comment: number;
    Value: number;
    Impact: number;
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



    
    dialogOpen = false;
    tempUIChart = false;
    respirationUIChart = false;
    pulseUIChart = false;

    showLabels = false;
    // pulse rate, blood pressure, respiration rate, temperature
    seriesVisiblity = [true, true, true, true];
    isLoggingIn = true;
    constructor(private monitorService: MonitorService,
        private act: ActionService) {
            // this.pulseListItems = new ObservableArray<MonitorUiModel>();
            // this.pulseRateSource.forEach(item => {
            //     this.pulseRateSeriesBinding.push(item);
            // });

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
                    temperatureListItem.Impact = 1;
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
                    // console.log('component bloodpressure', item);
                    let bloodpresHighListItem = new BloodUiHighModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.txn_date);
                    bloodpresHighListItem.Comment = getDBDate.getTime();
                    bloodpresHighListItem.Value = Number(testdata.value.high);
                    bloodpresHighListItem.Impact = 1;

                    let bloodpresLowListItem = new BloodUiHighModel();
                    bloodpresLowListItem.Comment = getDBDate.getTime();
                    bloodpresLowListItem.Value = Number(testdata.value.low);
                    bloodpresLowListItem.Impact = 1;

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
                    respirationListItem.Impact = 1;
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
                    pulseListItem.Impact = 1;
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
    onTrackBallContentRequested(args: TrackballCustomContentData) {
        let selectedItem = args.pointData;

        switch (args.seriesIndex) {
            case 0: args.content = "Pulse Rate"; break;
            case 1: args.content = String(selectedItem.Amount); break;
            case 2: args.content = "Blood Pressure (Systolic)"; break;
            case 3: args.content = String(selectedItem.Systolic); break;
            case 4: args.content = "Blood Pressure (Diastolic)"; break;
            case 5: args.content = String(selectedItem.Diastolic); break;
            case 6: args.content = "Respiratory Rate"; break;
            case 7: args.content = String(selectedItem.Amount); break;
            case 8: args.content = "Temperature"; break;
            case 9: args.content = String(selectedItem.Amount); break;
            default: args.content = " "; break;
        }
    }

    toggleLabels() {
        this.showLabels = !this.showLabels;
        this.isLoggingIn = !this.isLoggingIn;
    }

    toggleSeries(index: number) {
        this.seriesVisiblity[index] = !this.seriesVisiblity[index];
    }
}