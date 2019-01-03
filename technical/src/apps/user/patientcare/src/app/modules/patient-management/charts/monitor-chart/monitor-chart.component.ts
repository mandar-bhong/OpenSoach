import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import {DatePipe} from '@angular/common';
import  { ChartModel } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";
import { Observable } from 'tns-core-modules/ui/page/page';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { ConfListViewModel, ConfModel } from '~/app/models/ui/conf-models';
import { ConfService } from '~/app/services/conf/conf.service';
import { ListPicker } from "tns-core-modules/ui/list-picker";


export class NewDataItem{
    public monitorName:string;
    public instruction:number;
    public frequency:number;
    public duration:number;
    public startDate:Date;
    public intervalHrs: number;
    public startTime:Date;
    public endTime:Date;
    public specificTime:Date;
    public specificTimes:Array<any>;
    public desc:string="";
}

export class ChartMonitorConfModel{
    public name:string;
    public foodInst:number;
    public frequency:number;
    public duration:number;
    public startDate:string;
    public intervalHrs: number;
    public startTime:string;
    public endTime:string;
    public specificTimes:Array<any>;
    public desc:string;
}


@Component({
	moduleId: module.id,
	selector: 'monitor-chart',
	templateUrl: './monitor-chart.component.html',
	styleUrls: ['./monitor-chart.component.css']
})

export class MonitorChartComponent implements OnInit {
    
    formData : NewDataItem;
    chartConfModel : ChartMonitorConfModel;
    chartDbModel : ChartModel

    public myItems: Array<SegmentedBarItem>;
    public frequencyItems: Array<SegmentedBarItem>;
    public instSelectedIndex = 0;
    public freqSelectedIndex = 0;

    private monitorConfListItems = new ObservableArray<ConfModel>();
    public monitorConf: ConfModel;

    constructor(private routerExtensions: RouterExtensions, private datePipe: DatePipe,private chartservice: ChartService, private confService:ConfService){

        this.formData = new NewDataItem();
        this.formData.specificTimes = [];
        this.chartConfModel = new ChartMonitorConfModel();
        this.chartConfModel.specificTimes = [];
        this.chartDbModel = new ChartModel();

    }

    ngOnInit(){

        this.getMonitorConfListData();

        this.myItems = [];
        this.frequencyItems = [];

        const item1 = new SegmentedBarItem();
        item1.title = "Before Meal";
        this.myItems.push(item1);
        const item2 = new SegmentedBarItem();
        item2.title = "After Meal";
        this.myItems.push(item2);

        const freqItem1 = new SegmentedBarItem();
        freqItem1.title = "Every 'X' hours";
        this.frequencyItems.push(freqItem1);
        const freqItem2= new SegmentedBarItem();
        freqItem2.title = "Specific time";
        this.frequencyItems.push(freqItem2);
        
    }

    goBackPage() {
		this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }

    onPageLoaded(args) {
        console.log("monitor form page loaded");
    }

    public onInstructionSelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.instSelectedIndex = segmetedBar.selectedIndex;
    }

    public onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;
    }

    public onSubmit(){

        this.insertData(this.formData);

    }

    public insertData(data: NewDataItem){

        if (data.frequency==0){
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime,"h:mm a");
            this.chartConfModel.endTime = this.datePipe.transform(data.endTime,"h:mm a");
            
        if (data.desc!=""){
            this.chartConfModel.desc = "Monitor every " + data.intervalHrs + " hours.\n" + data.desc ;
        }else{
            this.chartConfModel.desc = "Monitor every " + data.intervalHrs + " hours.";
        }
        }

        if (data.frequency==1){
            for(var i = 0;i<data.specificTimes.length;i++){
                this.chartConfModel.specificTimes.push(this.datePipe.transform(data.specificTimes[i],"h:mm a"));
            }
            if (data.desc!=""){
                this.chartConfModel.desc = "Monitor as per specific timings " + ".\n" + data.desc ;
            }else{
                this.chartConfModel.desc = "Monitor as per specific timings ";
            }
        }

        this.chartConfModel.name = data.monitorName;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate,"yyyy-MM-dd");
        this.chartConfModel.foodInst = data.instruction;       
        
        let confString = JSON.stringify(this.chartConfModel);

        this.chartDbModel.admissionid = 2 ;
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = "Monitor";

        this.chartservice.insertChartItem(this.chartDbModel);

        this.chartservice.getChartList()

        this.goBackPage();

    }

    public addSpecificTime(){
        this.formData.specificTimes.push(this.formData.specificTime);
    }

    public getMonitorConfListData(){
		this.confService.getMonitorConf().then(
			(val)=>{
				val.forEach(item => {
					let monitorConfListItem = new ConfModel();
                    monitorConfListItem = item;
                    monitorConfListItem.conf =  JSON.parse(item.conf);
					this.monitorConfListItems.push(monitorConfListItem);
                });
                this.monitorConf = this.monitorConfListItems.getItem(0);
			},
			(error)=>{
				console.log("confListService error:",error);
			}

		);

    }
    
    public selectedIndexChanged(args) {        
        let picker = <ListPicker>args.object;
        let picked :any;

        // console.log("picker selection: " + picker.selectedIndex);
        // console.log("picker selection value: " + (<any>picker).selectedValue);

        picked = this.monitorConf.conf.tasks[picker.selectedIndex];
        this.formData.monitorName = picked.name;

    }

}