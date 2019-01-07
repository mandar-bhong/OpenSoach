import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { Data } from '@angular/router';
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import { DatePipe } from '@angular/common';
import { Switch } from "tns-core-modules/ui/switch";
import { ChartDBModel, MedChartModel } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
    moduleId: module.id,
    selector: 'medicine-chart',
    templateUrl: './medicine-chart.component.html',
    styleUrls: ['./medicine-chart.component.css']
})

export class MedicineChartComponent implements OnInit {

    // proccess variables
    medicineForm: FormGroup;

    medNameIsValid:boolean;
    intervalHrsIsValid: boolean;
    durationIsValid: boolean;

    formData: MedChartModel;
    chartConfModel: MedChartModel;
    chartDbModel: ChartDBModel

    foodInsItems: Array<SegmentedBarItem>;
    frequencyItems: Array<SegmentedBarItem>;
    instSelectedIndex = 0;
    freqSelectedIndex = 0;

    freqMorn:boolean;
    freqAftrn:boolean;
    freqNight:boolean;
    // end of proccess variables

    constructor(private routerExtensions: RouterExtensions, private datePipe: DatePipe, private chartservice: ChartService) {

        this.formData = new MedChartModel();
        this.freqMorn = true;
        this.freqAftrn = true;
        this.freqNight = true;

        this.chartConfModel = new MedChartModel();
        this.chartDbModel = new ChartDBModel();

    }

    ngOnInit() {

        // creating form control
        this.createFormControls();

        this.foodInsItems = [];
        this.frequencyItems = [];

        // set food instruction segmented bar items
        const foodInsItem1 = new SegmentedBarItem();
        foodInsItem1.title = "Before Meal";
        this.foodInsItems.push(foodInsItem1);
        const foodInsItem2 = new SegmentedBarItem();
        foodInsItem2.title = "After Meal";
        this.foodInsItems.push(foodInsItem2);

        // set frequency segmented bar items
        const freqItem1 = new SegmentedBarItem();
        freqItem1.title = "'X'- Times a day";
        this.frequencyItems.push(freqItem1);
        const freqItem2 = new SegmentedBarItem();
        freqItem2.title = "Every 'X' hours";
        this.frequencyItems.push(freqItem2);

        // load default form data
        this.medicineForm.get('quantity').setValue(1);
        // this.medicineForm.get('startDate').setValue(new Date());

    }

    // << func for navigating previous page
    goBackPage() {
        this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    // >> func for navigating previous page

    onPageLoaded(args) {
        console.log("medicine form page loaded");
    }

    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;

        if (this.freqSelectedIndex == 1) {
            this.medicineForm.controls['intervalHrs'].setValidators([Validators.required]);
            this.medicineForm.controls['intervalHrs'].updateValueAndValidity();
        } else {
            this.medicineForm.controls['intervalHrs'].clearValidators();
            this.medicineForm.controls['intervalHrs'].updateValueAndValidity();
            this.intervalHrsIsValid = false;
        }

    }

    onInstructionSelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.instSelectedIndex = segmetedBar.selectedIndex;
    }

    // << func for submit form data
    onSubmit() {

        this.medNameIsValid = this.medicineForm.controls['name'].hasError('required');
        this.intervalHrsIsValid = this.medicineForm.controls['intervalHrs'].hasError('required');
        this.durationIsValid = this.medicineForm.controls['duration'].hasError('required');

        if (this.medicineForm.invalid) {
            console.log("validation error");
            return;
        }

        // assign form data to model
        this.formData = Object.assign({}, this.medicineForm.value);
        this.formData.freqMorn = this.freqMorn;
        this.formData.freqAftrn = this.freqAftrn;
        this.formData.freqNight = this.freqNight;

        // insert form data to sqlite db
        this.insertData(this.formData);

    }
    // >> func for submit form data

    // set frequency checkbox value
    onMorningChecked(args) {

        let Switch = <Switch>args.object;
        if (Switch.checked) {
            this.freqMorn = true;
        } else {
            this.freqMorn = false;
        }

    }

    onAfternoonChecked(args) {

        let Switch = <Switch>args.object;
        if (Switch.checked) {
            this.freqAftrn = true;
        } else {
            this.freqAftrn = false;
        }

    }

    onNightChecked(args) {

        let Switch = <Switch>args.object;
        if (Switch.checked) {
            this.freqNight = true;
        } else {
            this.freqNight = false;
        }

    }

    // << func for inserting form data to sqlite db
    insertData(data: MedChartModel) {

        //set chart conf model
        let foodIns = "";
        let desc = "";
        let count = 0;

        if (data.foodInst == 0) {
            foodIns = "before meal";
        } else {
            foodIns = "after meal";
        }

        if (data.frequency == 0) {
            if (data.freqMorn == true) {
                this.chartConfModel.freqMorn = data.freqMorn;
                count = count + 1;
                desc = desc + " Morning &"
            }else{
                this.chartConfModel.freqMorn = data.freqMorn;
            }

            if (data.freqAftrn == true) {
                this.chartConfModel.freqAftrn = data.freqAftrn;
                count = count + 1;
                desc = desc + " Afternoon &"
            }else{
                this.chartConfModel.freqAftrn = data.freqAftrn;
            }

            if (data.freqNight == true) {
                this.chartConfModel.freqNight = data.freqNight;
                count = count + 1;
                desc = desc + " Night &"
            }else{
                this.chartConfModel.freqNight = data.freqNight;
            }

            desc = desc.slice(0, -1);

            if (data.desc != null) {
                this.chartConfModel.desc = desc + foodIns + ". \n" + data.desc + ".";
            } else {
                this.chartConfModel.desc = desc + foodIns + ".";
            }
            
        } else {
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "mediumTime");

            if (data.desc != null) {
                this.chartConfModel.desc = "Every " + data.intervalHrs + " hours in a day " + foodIns + ". \n" + data.desc + ".";
            } else {
                this.chartConfModel.desc = "Every " + data.intervalHrs + " hours in a day " + foodIns + ".";
            }
        }

        this.chartConfModel.name = data.name;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate, "yyyy-MM-dd");
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.foodInst = data.foodInst;

        let confString = JSON.stringify(this.chartConfModel);

        // set db model
        this.chartDbModel.admissionid = 2;
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = "Medicine";

        // insert chart db model to sqlite db
        this.chartservice.insertChartItem(this.chartDbModel);

        // get chart data from sqlite db
        this.chartservice.getChartList();

        this.goBackPage();

    }
    // >> func for inserting form data to sqlite db


    // << func for creating form controls
    createFormControls(): void {
        this.medicineForm = new FormGroup({
            name: new FormControl('', [Validators.required]),
            quantity: new FormControl(),
            foodInst: new FormControl(),
            frequency: new FormControl(),
            intervalHrs: new FormControl(),
            startDate: new FormControl(),
            duration: new FormControl('', [Validators.required]),
            startTime: new FormControl(),
            desc: new FormControl()
        });
    }
    // >> func for creating form controls

}