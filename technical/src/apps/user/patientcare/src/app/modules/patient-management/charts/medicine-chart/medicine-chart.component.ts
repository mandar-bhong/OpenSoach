import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { Data } from '@angular/router';
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import {DatePipe} from '@angular/common';
import { Switch } from "tns-core-modules/ui/switch";
import  { ChartModel } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";

export class NewDataItem{
    public medicineName:string;
    public quantity:number = 1;
    public startDate:Date = new Date();
    public duration:number;
    public frequency:number;
    public freqMorning:boolean;
    public freqAfternoon:boolean;  
    public freqNight:boolean;  
    public instruction:number;
    public intervalHrs: number;
    public desc:string="";
    public starttime:Date;
}

export class ChartMedConfModel{
    public name:string;
    public quantity:number;
    public startDate:string;
    public duration:number;
    public frequency:number;
    public freqMorning:boolean;
    public freqAfternoon:boolean;  
    public freqNight:boolean;  
    public foodInst:number;
    public intervalHrs: number;
    public desc:string;
    public starttime:string;
}


@Component({
	moduleId: module.id,
	selector: 'medicine-chart',
	templateUrl: './medicine-chart.component.html',
	styleUrls: ['./medicine-chart.component.css']
})

export class MedicineChartComponent implements OnInit {

    formData : NewDataItem;
    chartConfModel : ChartMedConfModel;
    chartDbModel : ChartModel

    public myItems: Array<SegmentedBarItem>;
    public frequencyItems: Array<SegmentedBarItem>;
    public instSelectedIndex = 0;
    public freqSelectedIndex = 0;

    constructor(private routerExtensions: RouterExtensions, private datePipe: DatePipe,private chartservice: ChartService) { 

        this.formData = new NewDataItem();
        this.formData.freqAfternoon = true;
        this.formData.freqMorning = true;
        this.formData.freqNight = true;

        this.chartConfModel = new ChartMedConfModel();
        this.chartDbModel = new ChartModel();

     }

    ngOnInit() {         

        this.myItems = [];
        this.frequencyItems = [];

        const item1 = new SegmentedBarItem();
        item1.title = "Before Meal";
        this.myItems.push(item1);
        const item2 = new SegmentedBarItem();
        item2.title = "After Meal";
        this.myItems.push(item2);

        const freqItem1 = new SegmentedBarItem();
        freqItem1.title = "'X'- Times a day";
        this.frequencyItems.push(freqItem1);
        const freqItem2 = new SegmentedBarItem();
        freqItem2.title = "Every 'X' hours";
        this.frequencyItems.push(freqItem2);

    }
    
    goBackPage() {
		this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    
    onPageLoaded(args) {
        console.log("medicine form page loaded");
    }

    // public onTap() {
        
    //     this.myCommitDataFormComp.dataForm.commitAll();

    //     const data = this.myCommitDataFormComp.dataForm.source;

    //     console.log("data",data);

    //     alert(
    //         {
    //             title: "Medicine Details",
    //             message: JSON.stringify(data),
    //             okButtonText: "OK"
    //         });
    // }

    // public onPropertyCommit(args: DataFormEventData) {
    //     let dataForm = <RadDataForm>args.object;
    // }

    public onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;
    }

    public onInstructionSelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.instSelectedIndex = segmetedBar.selectedIndex;
    }

    public onSubmit(){

        this.insertData(this.formData);

    }

    public onMorningChecked(args) {

        let Switch = <Switch>args.object;
        if (Switch.checked){
            this.formData.freqMorning = true;
        }else{
            this.formData.freqMorning = false;
        }

    }

    public onAfternoonChecked(args) {

        let Switch = <Switch>args.object;
        if (Switch.checked){
            this.formData.freqAfternoon = true;
        }else{
            this.formData.freqAfternoon = false;
        }

    }

    public onNightChecked(args) {

        let Switch = <Switch>args.object;
        if (Switch.checked){
            this.formData.freqNight = true;
        }else{
            this.formData.freqNight = false;
        }

    }

    public insertData(data: NewDataItem){

        let foodIns = "";
        let desc = "";
        let count = 0;

        if(data.instruction == 0){
            foodIns = "before meal";
        }else{
            foodIns = "after meal";
        }

        if(data.frequency == 0){
            if( data.freqMorning  == true){
                this.chartConfModel.freqMorning = data.freqMorning;                
                count = count + 1;
                desc = desc + " Morning &"
            }
            if( data.freqAfternoon  == true){
                this.chartConfModel.freqAfternoon = data.freqAfternoon;
                count = count + 1;
                desc = desc + " Afternoon &"
            }
            if( data.freqNight  == true){
                this.chartConfModel.freqNight = data.freqNight;
                count = count + 1;
                desc = desc + " Night &"
            }
        }else{
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.starttime = this.datePipe.transform(data.starttime,"mediumTime");
        }

        desc = desc.slice(0,-1);
        console.log("desc",desc);

        if (data.frequency == 0){
            if (data.desc!=""){
                this.chartConfModel.desc = desc + foodIns + ". \n" + data.desc + ".";
            }else{
                this.chartConfModel.desc = desc + foodIns + ".";
            } 
        }else{
            if(data.desc!=""){
                this.chartConfModel.desc = "Every " + data.intervalHrs + " hours in a day " + foodIns + ". \n" + data.desc + ".";
            }else{
                this.chartConfModel.desc = "Every " + data.intervalHrs + " hours in a day " + foodIns + ".";
            }
                        
        }

        console.log("this.chartConfModel.desc",this.chartConfModel.desc);

        this.chartConfModel.name = data.medicineName;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate,"yyyy-MM-dd");
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.foodInst = data.instruction; 

        let confString = JSON.stringify(this.chartConfModel);

        this.chartDbModel.admissionid = 2 ;
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = "Medicine";

        this.chartservice.insertChartItem(this.chartDbModel);

        this.chartservice.getChartList();

        this.goBackPage();

    }

}