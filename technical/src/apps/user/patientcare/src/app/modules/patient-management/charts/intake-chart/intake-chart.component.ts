import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { RouterExtensions } from 'nativescript-angular/router';
import { ListPicker } from 'tns-core-modules/ui/list-picker/list-picker';
import { SegmentedBarItem } from 'tns-core-modules/ui/segmented-bar';
import { ConfigCodeType, SYNC_STORE } from '~/app/app-constants';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { TimeConversion } from '~/app/helpers/time-conversion-helper';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { ChartDBModel, FrequencyValues, IntakeChartModel, PickerValues } from '~/app/models/ui/chart-models';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';

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
    intervalIsValid: boolean;
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
    serverDataStoreDataModelArray: ServerDataStoreDataModel<ScheduleDatastoreModel>[] = [];
    patientName: string;
    public frequencyType: Array<FrequencyValues> = [];
    public intakeType: Array<string> = [];
    // end of proccess variables
    frequencyList: FrequencyValues[] = [
        { name: "Every 'X' hours", value: 0 },
        { name: "Specific time", value: 1 },
        { name: "As Required", value: 2 }];
    intakeList: string[] = ["IV","Oral"];

    isSplinstructions = false;
    constructor(private routerExtensions: RouterExtensions,
        private datePipe: DatePipe,
        private params: ModalDialogParams,
        public workerService: WorkerService,
        private passDataService: PassDataService,
        private chartservice: ChartService) {
        this.formData = new IntakeChartModel();
        this.formData.specificTimes = [];
        this.specifictimes = [];
        this.chartConfModel = new IntakeChartModel();
        this.chartConfModel.specificTimes = [];
        this.chartDbModel = new ChartDBModel();
        this.frequencyType = [];
        this.intakeType = []
        for (let item of this.frequencyList) {
            this.frequencyType.push(item);
        }
        for (let item of this.intakeList) {
            this.intakeType.push(item);
        }
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
        this.params.closeCallback([]);
    }

    onPageLoaded(args) {
        console.log("intake form page loaded");
    }
    // on frequency selection changes 
    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <ListPicker>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;
        this.intakeForm.controls['interval'].clearValidators();
        this.intakeForm.controls['interval'].updateValueAndValidity();
        this.intakeForm.controls['numberofTimes'].clearValidators();
        this.intakeForm.controls['numberofTimes'].updateValueAndValidity();
        this.intakeForm.controls['splinstruction'].clearValidators();
        this.intakeForm.controls['splinstruction'].updateValueAndValidity();
        this.intervalIsValid = false;
        switch (this.freqSelectedIndex) {
            case 1:
                this.intakeForm.controls['interval'].clearValidators();
                this.intakeForm.controls['interval'].updateValueAndValidity();
                this.intakeForm.controls['numberofTimes'].clearValidators();
                this.intakeForm.controls['numberofTimes'].updateValueAndValidity();
                this.intervalIsValid = false;
                break;
            case 0:
                this.intakeForm.controls['numberofTimes'].setValidators([Validators.required]);
                this.intakeForm.controls['numberofTimes'].updateValueAndValidity();
                this.intakeForm.controls['interval'].setValidators([Validators.required]);
                this.intakeForm.controls['interval'].updateValueAndValidity();
                break;
            case 2:
                this.intakeForm.controls['splinstruction'].setValidators(Validators.required);
                this.intakeForm.controls['splinstruction'].updateValueAndValidity();
                break;
        }
    }

    // << func for submit form data
    onSubmit() {
        this.intakeNameIsValid = this.intakeForm.controls['name'].hasError('required');
        this.quantityIsValid = this.intakeForm.controls['quantity'].hasError('required');
        this.intervalIsValid = this.intakeForm.controls['interval'].hasError('required');
        this.durationIsValid = this.intakeForm.controls['duration'].hasError('required');
        this.SrtartdateIsValid = this.intakeForm.controls['startDate'].hasError('required');
        this.isNumberOfTimes = this.intakeForm.controls['numberofTimes'].hasError('required');
        this.isSplinstructions = this.intakeForm.controls['splinstruction'].hasError('required');
        if (this.intakeForm.invalid) {
            console.log("validation error");
            return;
        }
        this.formData = new IntakeChartModel();
        this.formData.specificTimes = [];
        // assign form data to model
        this.formData = Object.assign({}, this.intakeForm.value);
        this.formData.specificTimes = this.specifictimes;
        if (this.formData.frequency == 1) {
            if (this.specifictimes.length == 0) {
                this.isspecificTime = true;
                return;
            }
        }
        // creating 
        this.insertData(this.formData);

    }
    // >> func for submit form data

    // << func for inserting form data to sqlite db
    insertData(data: IntakeChartModel) {

        //set chart conf model
        if (data.frequency == 0) {    //  for  after x time interval
            this.chartConfModel.interval = data.interval * 60;
            this.chartConfModel.numberofTimes = data.numberofTimes;
            this.chartConfModel.startTime = this.datePipe.transform(data.startTime, "H.mm");
            // generate description
            let hourMinutsData = TimeConversion.timeConvert(this.chartConfModel.interval);
            let description = ` ${data.numberofTimes} times a day after every ${hourMinutsData} for ${data.duration} days.`;
            this.chartConfModel.desc = description;
        } else if (data.frequency == 1) {  //  for  at specific time
            for (var i = 0; i < data.specificTimes.length; i++) {
                this.chartConfModel.specificTimes.push(this.datePipe.transform(data.specificTimes[i], "H.mm"));
            }
            // generate description
            let desc = `At specific times for ${data.duration} days`;
            this.chartConfModel.desc = desc;
        } else if (data.frequency == 2) {   //  for  as required     
            this.chartConfModel.splinstruction = data.splinstruction;
            this.chartConfModel.desc = data.splinstruction;
        }
        this.chartConfModel.name = data.name;
        this.chartConfModel.quantity = data.quantity;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.remark = data.remark;
        this.chartConfModel.intakeType = this.intakeType[this.intakeForm.get('intakeType').value];
        this.chartConfModel.startDate = data.startDate
        let confString = JSON.stringify(this.chartConfModel);

        // set db model
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = ConfigCodeType.INTAKE;
        //  fucntion  for create actions
        this.createActions(this.chartDbModel.uuid, this.chartDbModel.admission_uuid, this.chartDbModel.conf_type_code, confString);

    }
    // >> func for inserting form data to sqlite db

    // << func for specific timings
    addSpecificTime() {
        const time = this.intakeForm.controls['specificTime'].value;
        if (time != null && time) {
            this.specifictimes.push(this.intakeForm.controls['specificTime'].value);
        }
    }

    // << func for creating form controls
    createFormControls(): void {
        this.intakeForm = new FormGroup({
            name: new FormControl('', [Validators.required]),
            quantity: new FormControl('', [Validators.required]),
            frequency: new FormControl(),
            duration: new FormControl('', [Validators.required]),
            startDate: new FormControl(),
            interval: new FormControl(),
            numberofTimes: new FormControl(),
            startTime: new FormControl(),
            specificTime: new FormControl(),
            remark: new FormControl(),
            splinstruction: new FormControl(),
            intakeType: new FormControl()
        });
    }
    // >> func for creating form controls

    // fucntion for creating intake actions
    createActions(uuid, admission_uuid, conf_type_code, conf) {
        const serverDataStoreModel = new ServerDataStoreDataModel<ScheduleDatastoreModel>();
        serverDataStoreModel.datastore = SYNC_STORE.SCHEDULE;
        serverDataStoreModel.data = new ScheduleDatastoreModel();
        serverDataStoreModel.data.uuid = uuid
        serverDataStoreModel.data.sync_pending = 1
        serverDataStoreModel.data.admission_uuid = admission_uuid;
        serverDataStoreModel.data.conf_type_code = conf_type_code;
        serverDataStoreModel.data.conf = conf;
        serverDataStoreModel.data.status = 0;
        serverDataStoreModel.data.client_updated_at = new Date().toISOString();
        this.serverDataStoreDataModelArray.push(serverDataStoreModel);
        // navigating data to schedule list page using subject
        this.params.closeCallback([serverDataStoreModel]);
    }
    // en dof fucntion
   // on inatke type selection change
   intakeTypeIndexChanged(args) {
    let picker = <ListPicker>args.object;
    let picked: any;   
}
}