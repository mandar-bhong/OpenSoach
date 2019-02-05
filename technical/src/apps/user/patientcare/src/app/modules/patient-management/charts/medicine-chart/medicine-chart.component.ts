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
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { SERVER_WORKER_MSG_TYPE, SYNC_STORE, ConfigCodeType } from '~/app/app-constants';
import { WorkerService } from '~/app/services/worker.service';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { PassDataService } from '~/app/services/pass-data-service';

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
    numberOfTimesValid: boolean;
    startDateValid: boolean;
    // end of proccess variables

    constructor(private routerExtensions: RouterExtensions,
        private datePipe: DatePipe,
        public workerService: WorkerService,
        private passDataService: PassDataService,
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

        // remove this code block once u done testing of 
        //  this.createActions();
        this.medicineForm.get('startDate').setValue(new Date());
        this.medicineForm.get('startTime').setValue(new Date());
    }

    // << func for navigating previous page
    goBackPage() {
        this.routerExtensions.back();
        //  this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    // >> func for navigating previous page

    onPageLoaded(args) {
        console.log("medicine form page loaded");
    }

    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;

        if (this.freqSelectedIndex == 1) {
            this.medicineForm.controls['numberofTimes'].setValidators(Validators.required);
            this.medicineForm.controls['numberofTimes'].updateValueAndValidity();
            this.medicineForm.controls['intervalHrs'].setValidators([Validators.required]);
            this.medicineForm.controls['intervalHrs'].updateValueAndValidity();
        } else {
            this.medicineForm.controls['intervalHrs'].clearValidators();
            this.medicineForm.controls['intervalHrs'].updateValueAndValidity();
            this.medicineForm.controls['numberofTimes'].clearValidators();
            this.medicineForm.controls['numberofTimes'].updateValueAndValidity();
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
        this.numberOfTimesValid = this.medicineForm.controls['numberofTimes'].hasError('required');
        this.startDateValid = this.medicineForm.controls['startDate'].hasError('required');
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
        formData.numberofTimes = this.medicineForm.get('numberofTimes').value;
        formData.startDate = this.medicineForm.get('startDate').value;
        formData.duration = this.medicineForm.get('duration').value;
        formData.startTime = this.medicineForm.get('startTime').value;
        formData.desc = this.medicineForm.get('desc').value;
        // replace this with user entered value of how many times take 
        //  console.log("formData", formData);

        // insert form data to sqlite db
        this.insertData(formData);
    }
    // >> func for submit form data

    createActions(uuid, admission_uuid, conf_type_code, conf) {
        const initModel = new ServerDataProcessorMessageModel();
        const serverDataStoreModel = new ServerDataStoreDataModel<ScheduleDatastoreModel>();
        serverDataStoreModel.datastore = SYNC_STORE.SCHEDULE;
        serverDataStoreModel.data = new ScheduleDatastoreModel();
        // serverDataStoreModel.data.uuid = '11'
        // serverDataStoreModel.data.sync_pending = 1
        // serverDataStoreModel.data.admission_uuid = "11";
        // serverDataStoreModel.data.conf_type_code = 'Medicine';
        //  serverDataStoreModel.data.conf = '{"mornFreqInfo":{"freqMorn":true},"aftrnFreqInfo":{"freqAftrn":true},"nightFreqInfo":{"freqNight":true},"desc":" Morning & Afternoon & Night before meal Test.","name":"Cipla","quantity":11,"startDate":"2019-01-23T08:30:00.438Z","duration":3,"frequency":1,"startTime":"20.30","intervalHrs":180,"foodInst":1,"endTime":"12.30","numberofTimes":3,"specificTimes":[11.3,12.3]}';

        // serverDataStoreModel.data.conf = JSON.stringify(formData);
        serverDataStoreModel.data.uuid = uuid
        serverDataStoreModel.data.sync_pending = 1
        serverDataStoreModel.data.admission_uuid = admission_uuid;
        serverDataStoreModel.data.conf_type_code = conf_type_code;
        serverDataStoreModel.data.conf = conf;
        console.log('created data', serverDataStoreModel.data)
        initModel.data = [serverDataStoreModel];
        initModel.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;
        this.workerService.ServerDataProcessorWorker.postMessage(initModel);
    }

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
            this.chartConfModel.numberofTimes = data.numberofTimes;
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "H.mm");
            console.log(' this.chartConfModel.startTime', this.chartConfModel.startTime);
            if (data.desc != null) {
                this.chartConfModel.desc = "Every " + data.intervalHrs + " minutes in a day " + foodIns + ". \n" + data.desc + ".";
            } else {
                this.chartConfModel.desc = "Every " + data.intervalHrs + " minutes in a day " + foodIns + ".";
            }
        }
        const currentTime = this.datePipe.transform(Date.now(), "H:mm");
        console.log("currentTime", currentTime);

        this.chartConfModel.name = data.name;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate, "yyyy-MM-dd") + " " + currentTime;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.foodInst = data.foodInst;

        let confString = JSON.stringify(this.chartConfModel);
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = ConfigCodeType.MEDICINE
        this.createActions(this.chartDbModel.uuid, this.chartDbModel.admission_uuid, this.chartDbModel.conf_type_code, confString)
        // get chart data from sqlite db
         // this.chartservice.getChartList();
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
            numberofTimes: new FormControl(),
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