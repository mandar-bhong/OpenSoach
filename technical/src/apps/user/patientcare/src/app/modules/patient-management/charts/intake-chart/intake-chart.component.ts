import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import {DatePipe} from '@angular/common';
import { Switch } from "tns-core-modules/ui/switch";
import  { ChartModel } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";


export class NewDataItem{
    public intakeName:string;
    public quantity: string;
    public frequency:number;
    public duration:number;
    public startDate:Date;
    public intervalHrs: number;
    public startTime:Date;
    public specificTime:Date;
    public specificTimes:Array<any>;
    public desc:string="";
}

export class ChartIntakeConfModel{
    public name:string;
    public quantity: string;
    public frequency:number;
    public duration:number;
    public startDate:string;
    public intervalHrs: number;
    public startTime:string;
    public specificTime:string;
    public specificTimes:Array<any>;    
    public desc:string;
}


@Component({
	moduleId: module.id,
	selector: 'intake-chart',
	templateUrl: './intake-chart.component.html',
	styleUrls: ['./intake-chart.component.css']
})

export class IntakeChartComponent implements OnInit {
    
    formData : NewDataItem;
    chartConfModel : ChartIntakeConfModel;
    chartDbModel : ChartModel

    public frequencyItems: Array<SegmentedBarItem>;
    public freqSelectedIndex = 0;

    constructor(private routerExtensions: RouterExtensions, private datePipe: DatePipe,private chartservice: ChartService){

        this.formData = new NewDataItem();
        this.formData.specificTimes = [];
        this.chartConfModel = new ChartIntakeConfModel();
        this.chartConfModel.specificTimes = [];
        this.chartDbModel = new ChartModel();

    }

    ngOnInit(){
        this.frequencyItems = [];
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
        console.log("intake form page loaded");
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
        }
        if (data.frequency==1){

            for(var i = 0;i<data.specificTimes.length;i++){
                this.chartConfModel.specificTimes.push(this.datePipe.transform(data.specificTimes[i],"h:mm a"));
            }
        }

        this.chartConfModel.name = data.intakeName;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate,"yyyy-MM-dd");

        if(data.desc!=""){
            this.chartConfModel.desc = data.quantity + "\n" + data.desc;
        }else{
            this.chartConfModel.desc = data.quantity
        }

        
        let confString = JSON.stringify(this.chartConfModel);

        this.chartDbModel.admissionid = 2 ;
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = "Intake";

        this.chartservice.insertChartItem(this.chartDbModel);

        this.chartservice.getChartList()

        this.goBackPage();

    }

    public addSpecificTime(){
        this.formData.specificTimes.push(this.formData.specificTime);
    }

}