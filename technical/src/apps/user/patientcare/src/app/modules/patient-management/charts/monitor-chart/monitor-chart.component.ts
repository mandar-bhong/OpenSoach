import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { dateProperty } from 'tns-core-modules/ui/date-picker/date-picker';
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import { DatePipe } from '@angular/common';
import { ChartDBModel, MonitorChartModel } from "~/app/models/ui/chart-models";
import { ChartService } from "~/app/services/chart/chart.service";
import { Observable } from 'tns-core-modules/ui/page/page';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { ConfListViewModel } from '~/app/models/ui/conf-models';
import { ConfService } from '~/app/services/conf/conf.service';
import { ListPicker } from "tns-core-modules/ui/list-picker";
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { PlatformHelper } from "~/app/helpers/platform-helper";

@Component({
    moduleId: module.id,
    selector: 'monitor-chart',
    templateUrl: './monitor-chart.component.html',
    styleUrls: ['./monitor-chart.component.css']
})

export class MonitorChartComponent implements OnInit {

    // proccess variables
    monitorForm: FormGroup;

    intervalHrsIsValid: boolean;
    durationIsValid: boolean;

    monitorName: string;
    specifictimes: Array<string>;

    formData: MonitorChartModel;
    chartConfModel: MonitorChartModel;
    chartDbModel: ChartDBModel

    foodInstItems: Array<SegmentedBarItem>;
    frequencyItems: Array<SegmentedBarItem>;
    foodInstSelectedIndex = 0;
    freqSelectedIndex = 0;

    monitorConfListItems = new ObservableArray<ConfListViewModel>();
    monitorConf: ConfListViewModel;
    // end of proccess variables

    constructor(
        private routerExtensions: RouterExtensions, 
        private datePipe: DatePipe, 
        private chartservice: ChartService, 
        private confService: ConfService) {

        this.formData = new MonitorChartModel();
        this.formData.specificTimes = [];
        this.specifictimes = [];
        this.chartConfModel = new MonitorChartModel();
        this.chartDbModel = new ChartDBModel();

    }

    ngOnInit() {

        // creating form control
        this.createFormControls();

        // get montior conf data for list picker
        this.getMonitorConfListData();

        this.foodInstItems = [];
        this.frequencyItems = [];

        // set food instruction segmented bar items
        const foodInstItem1 = new SegmentedBarItem();
        foodInstItem1.title = "Before Meal";
        this.foodInstItems.push(foodInstItem1);
        const foodInstItem2 = new SegmentedBarItem();
        foodInstItem2.title = "After Meal";
        this.foodInstItems.push(foodInstItem2);

        // set frequency segmented bar items
        const freqItem1 = new SegmentedBarItem();
        freqItem1.title = "Every 'X' hours";
        this.frequencyItems.push(freqItem1);
        const freqItem2 = new SegmentedBarItem();
        freqItem2.title = "Specific time";
        this.frequencyItems.push(freqItem2);

        // load default form data
        // this.monitorForm.get('startDate').setValue(new Date());

    }

    // << func for navigating previous page
    goBackPage() {
        this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    // >> func for navigating previous page

    onPageLoaded(args) {
        console.log("monitor form page loaded");
    }

    onFoodInstSelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.foodInstSelectedIndex = segmetedBar.selectedIndex;
    }

    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;

        if (this.freqSelectedIndex == 0) {
            this.monitorForm.controls['intervalHrs'].setValidators([Validators.required]);
            this.monitorForm.controls['intervalHrs'].updateValueAndValidity();
        } else {
            this.monitorForm.controls['intervalHrs'].clearValidators();
            this.monitorForm.controls['intervalHrs'].updateValueAndValidity();
            this.intervalHrsIsValid = false;
        }

    }

    // << func for submit form data
    onSubmit() {

        this.intervalHrsIsValid = this.monitorForm.controls['intervalHrs'].hasError('required');
        this.durationIsValid = this.monitorForm.controls['duration'].hasError('required');

        if (this.monitorForm.invalid) {
            console.log("validation error");
            return;
        }

        // assign form data to model
        this.formData = Object.assign({}, this.monitorForm.value);
        this.formData.name = this.monitorName;
        this.formData.specificTimes = this.specifictimes;

        // insert form data to sqlite db
        this.insertData(this.formData);

    }
    // >> func for submit form data

    // << func for inserting form data to sqlite db
    insertData(data: MonitorChartModel) {

        //set chart conf model
        if (data.frequency == 0) {
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "h:mm a");
            this.chartConfModel.endTime = this.datePipe.transform(data.endTime, "h:mm a");

            if (data.desc != null) {
                this.chartConfModel.desc = "Monitor every " + data.intervalHrs + " hours.\n" + data.desc;
            } else {
                this.chartConfModel.desc = "Monitor every " + data.intervalHrs + " hours.";
            }
        }

        if (data.frequency == 1) {
            this.chartConfModel.specificTimes = [];
            for (var i = 0; i < data.specificTimes.length; i++) {
                this.chartConfModel.specificTimes.push(this.datePipe.transform(data.specificTimes[i], "h:mm a"));
            }
            if (data.desc !=null) {
                this.chartConfModel.desc = "Monitor as per specific timings." + "\n" + data.desc;
            } else {
                this.chartConfModel.desc = "Monitor as per specific timings.";
            }
        }

        this.chartConfModel.name = data.name;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate, "yyyy-MM-dd");
        this.chartConfModel.foodInst = data.foodInst;

        let confString = JSON.stringify(this.chartConfModel);

        // set db model
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admissionid = 2;
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = "Monitor";

        // insert chart db model to sqlite db
        this.chartservice.insertChartItem(this.chartDbModel);

        // get chart data from sqlite db
        this.chartservice.getChartList()

        this.goBackPage();
        
    }
    // >> func for inserting form data to sqlite db

    // << func for specific timings
    addSpecificTime() {
        this.specifictimes.push(this.monitorForm.controls['specificTime'].value);
    }
    // >> func for specific timings

    // << func for getting monitor conf data
    getMonitorConfListData() {
        this.confService.getMonitorConf().then(
            (val) => {
                val.forEach(item => {
                    let monitorConfListItem = new ConfListViewModel();
                    monitorConfListItem.dbmodel = item;
                    monitorConfListItem.dbmodel.conf = JSON.parse(item.conf);
                    this.monitorConfListItems.push(monitorConfListItem);
                });
                this.monitorConf = this.monitorConfListItems.getItem(0);
            },
            (error) => {
                console.log("confListService error:", error);
            }

        );

    }
    // >> func for getting monitor list data

    // << func for selecting monitor name
    selectedIndexChanged(args) {
        let picker = <ListPicker>args.object;
        let picked: any;
        picked = this.monitorConf.dbmodel.conf.tasks[picker.selectedIndex];
        this.formData.name = picked.name;
        this.monitorName = picked.name;
    }
    //  >> func for selecting monitor name

    // << func for creating form controls
    createFormControls(): void {
        this.monitorForm = new FormGroup({
            foodInst: new FormControl(),
            frequency: new FormControl(),
            duration: new FormControl('', [Validators.required]),
            startDate: new FormControl(),
            intervalHrs: new FormControl('', [Validators.required]),
            startTime: new FormControl(),
            endTime: new FormControl(),
            specificTime: new FormControl(),
            desc: new FormControl()
        });
    }
    // >> func for creating form controls

}