import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import { DatePipe } from '@angular/common';
import { Switch } from "tns-core-modules/ui/switch";
import { ChartDBModel, IntakeChartModel } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { PlatformHelper } from "~/app/helpers/platform-helper";

@Component({
    moduleId: module.id,
    selector: 'intake-chart',
    templateUrl: './intake-chart.component.html',
    styleUrls: ['./intake-chart.component.css']
})

export class IntakeChartComponent implements OnInit {

    // proccess variables
    intakeForm: FormGroup;

    intakeNameIsValid:boolean;
    quantityIsValid:boolean;
    intervalHrsIsValid: boolean;
    durationIsValid: boolean;

    specifictimes: Array<string>;

    formData: IntakeChartModel;
    chartConfModel: IntakeChartModel;
    chartDbModel: ChartDBModel

    frequencyItems: Array<SegmentedBarItem>;
    freqSelectedIndex = 0;
    // end of proccess variables

    constructor(private routerExtensions: RouterExtensions, 
        private datePipe: DatePipe, 
        private chartservice: ChartService) {

        this.formData = new IntakeChartModel();
        this.formData.specificTimes = [];
        this.specifictimes = [];
        this.chartConfModel = new IntakeChartModel();
        this.chartConfModel.specificTimes = [];
        this.chartDbModel = new ChartDBModel();

    }

    ngOnInit() {

        // creating form control
        this.createFormControls();

        this.frequencyItems = [];
        const freqItem1 = new SegmentedBarItem();
        freqItem1.title = "Every 'X' hours";
        this.frequencyItems.push(freqItem1);
        const freqItem2 = new SegmentedBarItem();
        freqItem2.title = "Specific time";
        this.frequencyItems.push(freqItem2);

    }

    // << func for navigating previous page
    goBackPage() {
        this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    // >> func for navigating previous page

    onPageLoaded(args) {
        console.log("intake form page loaded");
    }

    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;

        if (this.freqSelectedIndex == 0) {
            this.intakeForm.controls['intervalHrs'].setValidators([Validators.required]);
            this.intakeForm.controls['intervalHrs'].updateValueAndValidity();
        } else {
            this.intakeForm.controls['intervalHrs'].clearValidators();
            this.intakeForm.controls['intervalHrs'].updateValueAndValidity();
            this.intervalHrsIsValid = false;
        }

    }

    // << func for submit form data
    onSubmit() {

        this.intakeNameIsValid = this.intakeForm.controls['name'].hasError('required');
        this.quantityIsValid = this.intakeForm.controls['quantity'].hasError('required');
        this.intervalHrsIsValid = this.intakeForm.controls['intervalHrs'].hasError('required');
        this.durationIsValid = this.intakeForm.controls['duration'].hasError('required');

        if (this.intakeForm.invalid) {
            console.log("validation error");
            return;
        }

        // assign form data to model
        this.formData = Object.assign({}, this.intakeForm.value);
        this.formData.specificTimes = this.specifictimes;

        // insert form data to sqlite db
        this.insertData(this.formData);

    }
    // >> func for submit form data

    // << func for inserting form data to sqlite db
    insertData(data: IntakeChartModel) {

        //set chart conf model
        if (data.frequency == 0) {
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "h:mm a");
        }
        if (data.frequency == 1) {
            for (var i = 0; i < data.specificTimes.length; i++) {
                this.chartConfModel.specificTimes.push(this.datePipe.transform(data.specificTimes[i], "h:mm a"));
            }
        }

        this.chartConfModel.name = data.name;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate, "yyyy-MM-dd");

        if (data.desc != null) {
            this.chartConfModel.desc = data.quantity + "\n" + data.desc;
        } else {
            this.chartConfModel.desc = data.quantity
        }


        let confString = JSON.stringify(this.chartConfModel);

        // set db model
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = "PA001";
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = "Intake";

        // insert chart db model to sqlite db
        this.chartservice.insertChartItem(this.chartDbModel);

        // get chart data from sqlite db
        this.chartservice.getChartList()

        this.goBackPage();

    }
    // >> func for inserting form data to sqlite db

    // << func for specific timings
    addSpecificTime() {
        this.specifictimes.push(this.intakeForm.controls['specificTime'].value);
    }
    // >> func for specific timings

    // << func for creating form controls
    createFormControls(): void {
        this.intakeForm = new FormGroup({
            name: new FormControl('', [Validators.required]),
            quantity: new FormControl('', [Validators.required]),
            frequency: new FormControl(),
            duration: new FormControl('', [Validators.required]),
            startDate: new FormControl(),
            intervalHrs: new FormControl(),
            startTime: new FormControl(),
            specificTime: new FormControl(),
            desc: new FormControl()
        });
    }
    // >> func for creating form controls

}