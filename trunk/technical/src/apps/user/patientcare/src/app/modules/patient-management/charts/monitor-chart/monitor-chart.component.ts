import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { RouterExtensions } from 'nativescript-angular/router';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { ListPicker } from 'tns-core-modules/ui/list-picker';
import { SegmentedBar, SegmentedBarItem } from 'tns-core-modules/ui/segmented-bar';
import { ConfigCodeType, SYNC_STORE } from '~/app/app-constants';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { TimeConversion } from '~/app/helpers/time-conversion-helper';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { ChartDBModel, ChartListViewModel, FrequencyValues, MonitorChartModel } from '~/app/models/ui/chart-models';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';

@Component({
    moduleId: module.id,
    selector: 'monitor-chart',
    templateUrl: './monitor-chart.component.html',
    styleUrls: ['./monitor-chart.component.css']
})

export class MonitorChartComponent implements OnInit {

    // proccess variables
    monitorForm: FormGroup;
    intervalIsValid: boolean;
    durationIsValid: boolean;
    monitorName: string;
    specifictimes: Array<string>;
    formData: MonitorChartModel;
    chartConfModel: MonitorChartModel;
    chartDbModel: ChartDBModel;
    foodInstItems: Array<SegmentedBarItem>;
    frequencyItems: Array<SegmentedBarItem>;
    foodInstSelectedIndex = 0;
    freqSelectedIndex = 0;
    monitorConfListItems = new ObservableArray<ChartListViewModel>();
    monitorConf: ChartListViewModel;
    isspecificTime: boolean;
    isNumberOfTimes: boolean;
    // end of proccess variables
    public frequencyType: Array<FrequencyValues> = [];
    frequencyList: FrequencyValues[] = [
        { name: "Every 'X' hours", value: 0 },
        { name: "Specific time", value: 1 }]

    // end of proccess variables

    constructor(
        private routerExtensions: RouterExtensions,
        private passDataService: PassDataService,
        private datePipe: DatePipe,
        private params: ModalDialogParams,
        public workerService: WorkerService,
        private chartService: ChartService) {

        this.formData = new MonitorChartModel();
        this.formData.specificTimes = [];
        this.specifictimes = [];
        this.chartConfModel = new MonitorChartModel();
        this.chartDbModel = new ChartDBModel();
        this.frequencyType = [];
        for (let item of this.frequencyList) {
            this.frequencyType.push(item);
        }

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
        this.monitorForm.get('startDate').setValue(new Date());
        this.monitorForm.get('startTime').setValue(new Date());
        // load default form data
        // this.monitorForm.get('startDate').setValue(new Date());
    }

    // << func for navigating previous page
    goBackPage() {
        this.params.closeCallback([]);
        // this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
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
        let segmetedBar = <ListPicker>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;
        if (this.freqSelectedIndex == 0) {
            this.monitorForm.controls['interval'].setValidators([Validators.required]);
            this.monitorForm.controls['interval'].updateValueAndValidity();
            this.monitorForm.controls['numberofTimes'].setValidators([Validators.required]);
            this.monitorForm.controls['numberofTimes'].updateValueAndValidity();
        } else {
            this.monitorForm.controls['interval'].clearValidators();
            this.monitorForm.controls['interval'].updateValueAndValidity();
            this.monitorForm.controls['numberofTimes'].clearValidators();
            this.monitorForm.controls['numberofTimes'].updateValueAndValidity();
            this.intervalIsValid = false;
        }

    }

    // << func for submit form data
    onSubmit() {

        this.intervalIsValid = this.monitorForm.controls['interval'].hasError('required');
        this.durationIsValid = this.monitorForm.controls['duration'].hasError('required');
        this.isNumberOfTimes = this.monitorForm.controls['numberofTimes'].hasError('required');
        if (this.monitorForm.invalid) {
            console.log("validation error");
            return;
        };
        this.formData = new MonitorChartModel();
        this.formData.specificTimes = [];
        // assign form data to model
        this.formData = Object.assign({}, this.monitorForm.value);
        this.formData.name = this.monitorName;
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
    insertData(data: MonitorChartModel) {

        //set chart conf model
        if (data.frequency == 0) {
            this.chartConfModel.interval = data.interval * 60;
            this.chartConfModel.numberofTimes = data.numberofTimes;
            this.chartConfModel.startTime =TimeConversion.getStartTime(this.datePipe.transform(data.startTime, "H.mm"));
            // this.chartConfModel.endTime = this.datePipe.transform(data.endTime, "H.mm");
            // generate description
            let hourMinutsData = TimeConversion.timeConvert(this.chartConfModel.interval);
            let description = ` ${data.numberofTimes} times a day after every ${hourMinutsData} for ${data.duration} days.`;
            this.chartConfModel.desc = description;
        } else if (data.frequency == 1) {
            this.chartConfModel.specificTimes = [];
            for (var i = 0; i < data.specificTimes.length; i++) {
                this.chartConfModel.specificTimes.push(TimeConversion.getStartTime(this.datePipe.transform(data.specificTimes[i], "H.mm")));
            }
            // generate description
            let desc = `At specific times for ${data.duration} days`;
            this.chartConfModel.desc = desc;
        }
        this.chartConfModel.name = data.name;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.remark = data.remark;
        // this.chartConfModel.startDate = data.startDate;
        // this.chartConfModel.foodInst = data.foodInst;

        let confString = JSON.stringify(this.chartConfModel);
        // set db model
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
        this.chartDbModel.conf = confString;
        this.chartDbModel.start_date = data.startDate.toISOString();
        this.chartDbModel.conf_type_code = ConfigCodeType.MONITOR;
        console.log(this.chartDbModel);
        //cehcking existing monitor schedule
        //  to do 
        this.createActions(this.chartDbModel, confString);
        // get chart data from sqlite db
        // this.chartService.getChartList();
        //   this.goBackPage();
    }
    // >> func for inserting form data to sqlite db

    // << func for specific timings
    addSpecificTime() {
        const timeValue = this.monitorForm.controls['specificTime'].value;
        if (timeValue && timeValue != null) {
            this.specifictimes.push(this.monitorForm.controls['specificTime'].value);
        }
    }
    // >> func for specific timings

    // << func for getting monitor conf data
    getMonitorConfListData() {
        this.chartService.getMonitorConf().then(
            (val) => {
                val.forEach(item => {
                    let monitorConfListItem = new ChartListViewModel();
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
            // foodInst: new FormControl(),
            frequency: new FormControl(),
            duration: new FormControl('', [Validators.required]),
            startDate: new FormControl(),
            interval: new FormControl('', [Validators.required]),
            numberofTimes: new FormControl(),
            startTime: new FormControl(),
            // endTime: new FormControl(),
            specificTime: new FormControl(),
            remark: new FormControl()
        });
    }
    // >> func for creating form controls
    // fucntion for creating intake actions
    createActions(monitormodel, conf) {
        const initModel = new ServerDataProcessorMessageModel();
        const serverDataStoreModel = new ServerDataStoreDataModel<ScheduleDatastoreModel>();
        serverDataStoreModel.datastore = SYNC_STORE.SCHEDULE;
        serverDataStoreModel.data = new ScheduleDatastoreModel();
        serverDataStoreModel.data = monitormodel;
        serverDataStoreModel.data.sync_pending = 1
        serverDataStoreModel.data.client_updated_at = new Date().toISOString();
        serverDataStoreModel.data.conf = conf;
        serverDataStoreModel.data.status = 0;
        this.params.closeCallback([serverDataStoreModel]);

    }
    // en dof fucntion


}// end of class