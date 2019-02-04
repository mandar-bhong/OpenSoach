import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { MonitorService } from '~/app/services/monitor/monitor.service';

export class MonitorUiModel {
    Comment: string;
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
    constructor(private monitorService: MonitorService) {

    }

    ngOnInit() {
        this.temperature();
        this.bloodpressure();
        this.respiration();
        this.pulse();

    }
    temperature() {
        this.monitorService.getTempActionTxn().then(
            (val) => {
                val.forEach(item => {
                    let temperatureListItem = new MonitorUiModel();
                    // temperatureListItem = item;
                    const testdata = JSON.parse(item.txn_data);
                    temperatureListItem.Comment = testdata.comment;
                    temperatureListItem.Value = Number(testdata.value);
                    // console.log('temperatureListItem.comment', temperatureListItem.comment);
                    // console.log('temperatureListItem.value', temperatureListItem.value);
                    this.tempListItems.push(temperatureListItem);
                    // console.log('TempListItems', this.tempListItems);
                });
                // console.log('TempListItems outside', this.tempListItems);
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
}