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
import { SERVER_WORKER_MSG_TYPE, SYNC_STORE, ConfigCodeType } from '~/app/app-constants';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { WorkerService } from '~/app/services/worker.service';
import { PassDataService } from '~/app/services/pass-data-service';

@Component({
    moduleId: module.id,
    selector: 'intake-chart',
    templateUrl: './intake-chart.component.html',
    styleUrls: ['./intake-chart.component.css']
})

export class IntakeChartComponent implements OnInit {

    // proccess variables
    intakeForm: FormGroup;

    intakeNameIsValid: boolean;
    quantityIsValid: boolean;
    intervalHrsIsValid: boolean;
    durationIsValid: boolean;

    specifictimes: Array<string>;

    formData: IntakeChartModel;
    chartConfModel: IntakeChartModel;
    chartDbModel: ChartDBModel

    frequencyItems: Array<SegmentedBarItem>;
    freqSelectedIndex = 0;
    SrtartdateIsValid: boolean;
    isNumberOfTimes: boolean;
    isspecificTime: boolean;
    // end of proccess variables

    constructor(private routerExtensions: RouterExtensions,
        private datePipe: DatePipe,
        public workerService: WorkerService,
        private passDataService: PassDataService,
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
        this.intakeForm.get('startDate').setValue(new Date());
        this.intakeForm.get('startTime').setValue(new Date());
    }

    // << func for navigating previous page
    goBackPage() {
        this.routerExtensions.back();
        //  this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    // >> func for navigating previous page

    onPageLoaded(args) {
        console.log("intake form page loaded");
    }

    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;

        if (this.freqSelectedIndex == 0) {
            this.intakeForm.controls['numberofTimes'].setValidators([Validators.required]);
            this.intakeForm.controls['numberofTimes'].updateValueAndValidity();
            this.intakeForm.controls['intervalHrs'].setValidators([Validators.required]);
            this.intakeForm.controls['intervalHrs'].updateValueAndValidity();
        } else {
            this.intakeForm.controls['intervalHrs'].clearValidators();
            this.intakeForm.controls['intervalHrs'].updateValueAndValidity();
            this.intakeForm.controls['numberofTimes'].clearValidators();
            this.intakeForm.controls['numberofTimes'].updateValueAndValidity();
            this.intervalHrsIsValid = false;
        }

    }

    // << func for submit form data
    onSubmit() {

        this.intakeNameIsValid = this.intakeForm.controls['name'].hasError('required');
        this.quantityIsValid = this.intakeForm.controls['quantity'].hasError('required');
        this.intervalHrsIsValid = this.intakeForm.controls['intervalHrs'].hasError('required');
        this.durationIsValid = this.intakeForm.controls['duration'].hasError('required');
        this.SrtartdateIsValid = this.intakeForm.controls['startDate'].hasError('required');
        this.isNumberOfTimes = this.intakeForm.controls['numberofTimes'].hasError('required');
        if (this.intakeForm.invalid) {
            console.log("validation error");
            return;
        }
        // assign form data to model
        this.formData = Object.assign({}, this.intakeForm.value);
        this.formData.specificTimes = this.specifictimes;
        if (this.formData.frequency == 1) {
            if (this.specifictimes.length == 0) {
                this.isspecificTime = true;
                return;
            }
        }
        // insert form data to sqlite db
        this.insertData(this.formData);

    }
    // >> func for submit form data

    // << func for inserting form data to sqlite db
    insertData(data: IntakeChartModel) {

        //set chart conf model
        if (data.frequency == 0) {
            this.chartConfModel.intervalHrs = data.intervalHrs;
            this.chartConfModel.numberofTimes = data.numberofTimes;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "H.mm");
        }
        if (data.frequency == 1) {
            for (var i = 0; i < data.specificTimes.length; i++) {
                this.chartConfModel.specificTimes.push(this.datePipe.transform(data.specificTimes[i], "H.mm"));
            }
        }

        this.chartConfModel.name = data.name;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        const currentTime = this.datePipe.transform(Date.now(), "H:mm");
        console.log("currentTime", currentTime);
        this.chartConfModel.startDate = this.datePipe.transform(data.startDate, "yyyy-MM-dd") + " " + currentTime;

        if (data.desc != null) {
            this.chartConfModel.desc = data.quantity + "\n" + data.desc;
        } else {
            this.chartConfModel.desc = data.quantity
        }


        let confString = JSON.stringify(this.chartConfModel);

        // set db model
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = ConfigCodeType.INTAKE;

        // insert chart db model to sqlite db       
        this.createActions(this.chartDbModel.uuid, this.chartDbModel.admission_uuid, this.chartDbModel.conf_type_code, confString);
        // get chart data from sqlite db
        this.chartservice.getChartList()

        this.goBackPage();

    }
    // >> func for inserting form data to sqlite db

    // << func for specific timings
    addSpecificTime() {
        console.log('addSpecificTime Taped');
        const time = this.intakeForm.controls['specificTime'].value;
        if (time != null && time) {
            this.specifictimes.push(this.intakeForm.controls['specificTime'].value);
        }
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
            numberofTimes: new FormControl(),
            startTime: new FormControl(),
            specificTime: new FormControl(),
            desc: new FormControl()
        });
    }
    // >> func for creating form controls

    // fucntion for creating intake actions
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
    // en dof fucntion

}