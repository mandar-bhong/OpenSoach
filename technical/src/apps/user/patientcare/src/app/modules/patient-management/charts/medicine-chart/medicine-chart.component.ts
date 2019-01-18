import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { Data } from '@angular/router';
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import { DatePipe } from '@angular/common';
import { Switch } from "tns-core-modules/ui/switch";
import { ChartDBModel, MedChartModel, MornFreqInfo, AftrnFreqInfo, NightFreqInfo } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";
import { FormGroup, FormControl, Validators, FormBuilder } from '@angular/forms';
import { PlatformHelper } from "~/app/helpers/platform-helper";

@Component({
    moduleId: module.id,
    selector: 'medicine-chart',
    templateUrl: './medicine-chart.component.html',
    styleUrls: ['./medicine-chart.component.css']
})

export class MedicineChartComponent implements OnInit {

    // proccess variables
    medicineForm: FormGroup;

    medNameIsValid: boolean;
    intervalHrsIsValid: boolean;
    durationIsValid: boolean;

    // formData: MedChartModel;
    chartConfModel: MedChartModel;
    chartDbModel: ChartDBModel

    foodInsItems: Array<SegmentedBarItem>;
    frequencyItems: Array<SegmentedBarItem>;
    instSelectedIndex = 0;
    freqSelectedIndex = 0;

    freqMorn: boolean;
    freqAftrn: boolean;
    freqNight: boolean;
    // end of proccess variables

    constructor(private routerExtensions: RouterExtensions,
        private datePipe: DatePipe,
        private chartservice: ChartService) {

        this.freqMorn = true;
        this.freqAftrn = true;
        this.freqNight = true;

        this.chartConfModel = new MedChartModel();
        this.chartConfModel.mornFreqInfo = new MornFreqInfo();
        this.chartConfModel.aftrnFreqInfo = new AftrnFreqInfo();
        this.chartConfModel.nightFreqInfo = new NightFreqInfo();
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
        // this.medicineForm.get('mornQuantity').setValue(1);
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
        const formData = new MedChartModel();
        formData.mornFreqInfo = new MornFreqInfo();
        formData.aftrnFreqInfo = new AftrnFreqInfo();
        formData.nightFreqInfo = new NightFreqInfo();

        formData.name = this.medicineForm.get('name').value;
        formData.quantity = this.medicineForm.get('quantity').value;
        formData.foodInst = this.medicineForm.get('foodInst').value;
        formData.frequency = this.medicineForm.get('frequency').value;
        formData.mornFreqInfo.freqMorn = this.freqMorn;
        formData.mornFreqInfo.mornFreqQuantity = this.medicineForm.get('mornQuantity').value;
        formData.aftrnFreqInfo.freqAftrn = this.freqAftrn;
        formData.aftrnFreqInfo.aftrnFreqQuantity = this.medicineForm.get('aftrnQuantity').value;
        formData.nightFreqInfo.freqNight = this.freqNight;
        formData.nightFreqInfo.nightFreqQuantity = this.medicineForm.get('nightQuantity').value;
        formData.nightFreqInfo.nightFreqQuantity = this.medicineForm.get('nightQuantity').value;
        formData.intervalHrs = this.medicineForm.get('intervalHrs').value;
        formData.startDate = this.medicineForm.get('startDate').value;
        formData.duration = this.medicineForm.get('duration').value;
        formData.startTime = this.medicineForm.get('startTime').value;
        formData.desc = this.medicineForm.get('desc').value;

        console.log("formData",formData);

        // insert form data to sqlite db
        this.insertData(formData);

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
            if (data.mornFreqInfo.freqMorn == true) {
                this.chartConfModel.mornFreqInfo.freqMorn = data.mornFreqInfo.freqMorn;
                this.chartConfModel.mornFreqInfo.mornFreqQuantity = data.mornFreqInfo.mornFreqQuantity;
                count = count + 1;
                desc = desc + " Morning &"
            } else {
                this.chartConfModel.mornFreqInfo.freqMorn = data.mornFreqInfo.freqMorn;
                this.chartConfModel.mornFreqInfo.mornFreqQuantity = 0;
            }

            if (data.aftrnFreqInfo.freqAftrn == true) {
                this.chartConfModel.aftrnFreqInfo.freqAftrn = data.aftrnFreqInfo.freqAftrn;
                this.chartConfModel.aftrnFreqInfo.aftrnFreqQuantity = data.aftrnFreqInfo.aftrnFreqQuantity;
                count = count + 1;
                desc = desc + " Afternoon &"
            } else {
                this.chartConfModel.aftrnFreqInfo.freqAftrn = data.aftrnFreqInfo.freqAftrn;
                this.chartConfModel.aftrnFreqInfo.aftrnFreqQuantity = 0;
            }

            if (data.nightFreqInfo.freqNight == true) {
                this.chartConfModel.nightFreqInfo.freqNight = data.nightFreqInfo.freqNight;
                this.chartConfModel.nightFreqInfo.nightFreqQuantity = data.nightFreqInfo.nightFreqQuantity;
                count = count + 1;
                desc = desc + " Night &"
            } else {
                this.chartConfModel.nightFreqInfo.freqNight = data.nightFreqInfo.freqNight;
                this.chartConfModel.nightFreqInfo.nightFreqQuantity = 0;
            }

            desc = desc.slice(0, -1);

            if (data.desc != null) {
                this.chartConfModel.desc = desc + foodIns + ". \n" + data.desc + ".";
            } else {
                this.chartConfModel.desc = desc + foodIns + ".";
            }

        } else {
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "H:mm");

            if (data.desc != null) {
                this.chartConfModel.desc = "Every " + data.intervalHrs + " hours in a day " + foodIns + ". \n" + data.desc + ".";
            } else {
                this.chartConfModel.desc = "Every " + data.intervalHrs + " hours in a day " + foodIns + ".";
            }
        }

        const currentTime = this.datePipe.transform(Date.now(), "H:mm");
        console.log("currentTime",currentTime);

        this.chartConfModel.name = data.name;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate, "yyyy-MM-dd") + " " + currentTime;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.foodInst = data.foodInst;

        let confString = JSON.stringify(this.chartConfModel);

        // set db model
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = "PA001";
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
            desc: new FormControl(),
            mornQuantity: new FormControl(),
            aftrnQuantity: new FormControl(),
            nightQuantity: new FormControl()
        });
    }
    // >> func for creating form controls

}