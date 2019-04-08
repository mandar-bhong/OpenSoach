import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { MonitorService } from '~/app/services/monitor/monitor.service';
import { ActionService } from '~/app/services/action/action.service';
import { TrackballCustomContentData } from 'nativescript-ui-chart';
import { PassDataService } from '~/app/services/pass-data-service';

export class MonitorChartUiModel {
    timeStamp: number;
    Amount: number;
    Systolic: number;
    Diastolic: number;
    Impact: number;
}
@Component({
    moduleId: module.id,
    selector: 'monitor',
    templateUrl: './monitor.component.html',
    styleUrls: ['./monitor.component.css']
})

export class MonitorComponent implements OnInit {

    tempListItems = new ObservableArray<MonitorChartUiModel>();
    bloodpresListItems = new ObservableArray<MonitorChartUiModel>();
    respirationListItems = new ObservableArray<MonitorChartUiModel>();
    pulseListItems = new ObservableArray<MonitorChartUiModel>();

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
        private passdataservice:PassDataService,
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

        this.tempListItems = new ObservableArray<MonitorChartUiModel>();
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
        this.tempListItems = new ObservableArray<MonitorChartUiModel>();
        this.temperature();
    }
    temperature() {
        this.majorStepUnit = "Day";
        this.monitorService.getTempActionTxn(this.passdataservice.getAdmissionID()).then(
            (val) => {
                val.forEach(item => {
                    let temperatureListItem = new MonitorChartUiModel();
                    // temperatureListItem = item;
                    const testdata = JSON.parse(item.txn_data);
                    temperatureListItem.Amount = Number(testdata.value);

                    const getDBDate = new Date(item.scheduled_time);
                    // const asc_date =
                    temperatureListItem.timeStamp = getDBDate.getTime();
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
        this.bloodpresListItems = new ObservableArray<MonitorChartUiModel>();
         this.majorStepUnit = "Day";
        this.monitorService.getBloodPreActionTxn(this.passdataservice.getAdmissionID()).then(
            (val) => {               
                val.forEach(item => {
                     console.log('component bloodpressure', item);
                    let bloodpresHighListItem = new MonitorChartUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.scheduled_time);
                    bloodpresHighListItem.timeStamp = getDBDate.getTime();
                    bloodpresHighListItem.Systolic = Number(testdata.value.high);
                    bloodpresHighListItem.Impact = 1;
                    bloodpresHighListItem.timeStamp = getDBDate.getTime();
                    bloodpresHighListItem.Diastolic = Number(testdata.value.low);
                    bloodpresHighListItem.Impact = 1;

                    this.bloodpresListItems.push(bloodpresHighListItem);
                    // console.log('TempListItems', this.tempListItems);
                });
            },
            (error) => {
                console.log("getChartData error:", error);
            }
        );
    }


    respiration() {
        this.majorStepUnit = "Day";
        this.monitorService.getRespirationActionTxn(this.passdataservice.getAdmissionID()).then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme', item);
                    let respirationListItem = new MonitorChartUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.scheduled_time);
                    respirationListItem.timeStamp = getDBDate.getTime();
                    respirationListItem.Amount = Number(testdata.value);
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
        this.monitorService.getPulseActionTxn(this.passdataservice.getAdmissionID()).then(
            (val) => {
                val.forEach(item => {
                    // console.log('item homme', item);
                    let pulseListItem = new MonitorChartUiModel();
                    const testdata = JSON.parse(item.txn_data);
                    const getDBDate = new Date(item.scheduled_time);
                    pulseListItem.timeStamp = getDBDate.getTime();
                    pulseListItem.Amount = Number(testdata.value);
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