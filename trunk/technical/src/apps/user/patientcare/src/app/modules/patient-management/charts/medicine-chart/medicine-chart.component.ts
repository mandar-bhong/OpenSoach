import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { RouterExtensions } from 'nativescript-angular/router';
import { ListPicker } from 'tns-core-modules/ui/list-picker/list-picker';
import { SegmentedBar, SegmentedBarItem } from 'tns-core-modules/ui/segmented-bar';
import { Switch } from 'tns-core-modules/ui/switch';
import { ConfigCodeType, SYNC_STORE } from '~/app/app-constants';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { TimeConversion } from '~/app/helpers/time-conversion-helper';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import {
    AftrnFreqInfo,
    ChartDBModel,
    FrequencyValues,
    MedChartModel,
    MornFreqInfo,
    NightFreqInfo,
} from '~/app/models/ui/chart-models';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';

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
    intervalIsValid: boolean;
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
    patientName: string;
    public medicineType: Array<string> = [];
    public frequencyType: Array<FrequencyValues> = [];
    public picked: string;
    // end of proccess variables
    frequencyList: FrequencyValues[] = [
        { name: "'X'- Times a day", value: 0 },
        { name: "Every 'X' hours", value: 1 },
        { name: "As Required", value: 2 }];
    isInstruction = false;
    isSplinstructions = false;
    isDosage = false;
    isXTimesDay = false;
    constructor(private routerExtensions: RouterExtensions,
        private datePipe: DatePipe,
        public workerService: WorkerService,
        private params: ModalDialogParams,
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
        this.frequencyType = [];      
        for (let item of this.frequencyList) {
            this.frequencyType.push(item);
        }
    }

    ngOnInit() {
        // this.patientName = 'Raj Ghadage';
        // creating form control
        this.createFormControls();
        this.getMedicineType();
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
        this.medicineForm.get('startDate').setValue(new Date());
        this.medicineForm.get('startTime').setValue(new Date());
        this.medicineForm.get('mornQuantity').setValue('1');
        this.medicineForm.get('aftrnQuantity').setValue('1');
        this.medicineForm.get('nightQuantity').setValue('1');
    }

    // << func for navigating previous page
    goBackPage() {
        this.params.closeCallback([]);
        //  this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
    }
    // >> func for navigating previous page

    onPageLoaded(args) {

    }

    onFrequencySelectedIndexChange(args) {
        let segmetedBar = <ListPicker>args.object;
        this.freqSelectedIndex = segmetedBar.selectedIndex;
        this.medicineForm.controls['interval'].clearValidators();
        this.medicineForm.controls['interval'].updateValueAndValidity();
        this.medicineForm.controls['numberofTimes'].clearValidators();
        this.medicineForm.controls['numberofTimes'].updateValueAndValidity();
        this.medicineForm.controls['splinstructions'].clearValidators();
        this.medicineForm.controls['splinstructions'].updateValueAndValidity();
        this.medicineForm.controls['quantity'].clearValidators();
        this.medicineForm.controls['quantity'].updateValueAndValidity();
        this.intervalIsValid = false;
        switch (this.freqSelectedIndex) {
            case 1:
                this.medicineForm.controls['numberofTimes'].setValidators(Validators.required);
                this.medicineForm.controls['numberofTimes'].updateValueAndValidity();
                this.medicineForm.controls['interval'].setValidators([Validators.required]);
                this.medicineForm.controls['interval'].updateValueAndValidity();
                this.medicineForm.controls['quantity'].setValidators([Validators.required]);
                this.medicineForm.controls['quantity'].updateValueAndValidity();
                break;
            case 0:

                break;
            case 2:
                this.medicineForm.controls['splinstructions'].setValidators(Validators.required);
                this.medicineForm.controls['splinstructions'].updateValueAndValidity();
                break;
        }
    }

    onInstructionSelectedIndexChange(args) {
        let segmetedBar = <SegmentedBar>args.object;
        this.instSelectedIndex = segmetedBar.selectedIndex;
    }

    // << func for submit form data
    onSubmit() {

        this.medNameIsValid = this.medicineForm.controls['name'].hasError('required');
        this.intervalIsValid = this.medicineForm.controls['interval'].hasError('required');
        this.durationIsValid = this.medicineForm.controls['duration'].hasError('required');
        this.numberOfTimesValid = this.medicineForm.controls['numberofTimes'].hasError('required');
        this.startDateValid = this.medicineForm.controls['startDate'].hasError('required');
        // this.isInstruction = this.medicineForm.controls['desc'].hasError('required');
        this.isSplinstructions = this.medicineForm.controls['splinstructions'].hasError('required');
        this.isDosage = this.medicineForm.controls['quantity'].hasError('required');
        // frequency check
        switch (this.medicineForm.get('frequency').value) {
            case 0:
                if (!(this.freqMorn || this.freqAftrn || this.freqNight)) {
                    this.isXTimesDay = true;
                    return;
                }
                break;
            default:
                break;
        }
        //  validation check  if form is valid 
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
        formData.interval = this.medicineForm.get('interval').value;
        formData.numberofTimes = this.medicineForm.get('numberofTimes').value;
        formData.startDate = this.medicineForm.get('startDate').value;
        formData.duration = this.medicineForm.get('duration').value;
        formData.startTime = this.medicineForm.get('startTime').value;
        formData.remark = this.medicineForm.get('remark').value;
        formData.splinstruction = this.medicineForm.get('splinstructions').value;
        this.insertData(formData);
    }
    // >> func for submit form data

    createActions(uuid, admission_uuid, conf_type_code, conf) {
        const initModel = new ServerDataProcessorMessageModel();
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
        this.params.closeCallback([serverDataStoreModel]);
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
        this.chartConfModel = new MedChartModel();
        this.chartConfModel.mornFreqInfo = new MornFreqInfo();
        this.chartConfModel.aftrnFreqInfo = new AftrnFreqInfo();
        this.chartConfModel.nightFreqInfo = new NightFreqInfo();
        this.chartDbModel = new ChartDBModel();
        //set chart conf model
        let foodIns = "";
        let desc = "";
        let count = 0;
        if (data.foodInst == 0) {
            foodIns = "before meal";
        } else {
            foodIns = "after meal";
        }
        //  for x times in day 
        if (data.frequency == 0) {
            if (data.mornFreqInfo.freqMorn == true) {
                this.chartConfModel.mornFreqInfo.freqMorn = data.mornFreqInfo.freqMorn;
                this.chartConfModel.mornFreqInfo.mornFreqQuantity = data.mornFreqInfo.mornFreqQuantity;
                count = count + 1;
                // desc = desc + " Morning &"
            } else {
                this.chartConfModel.mornFreqInfo.freqMorn = data.mornFreqInfo.freqMorn;
                this.chartConfModel.mornFreqInfo.mornFreqQuantity = 0;
            }

            if (data.aftrnFreqInfo.freqAftrn == true) {
                this.chartConfModel.aftrnFreqInfo.freqAftrn = data.aftrnFreqInfo.freqAftrn;
                this.chartConfModel.aftrnFreqInfo.aftrnFreqQuantity = data.aftrnFreqInfo.aftrnFreqQuantity;
                count = count + 1;
                // desc = desc + " Afternoon &"
            } else {
                this.chartConfModel.aftrnFreqInfo.freqAftrn = data.aftrnFreqInfo.freqAftrn;
                this.chartConfModel.aftrnFreqInfo.aftrnFreqQuantity = 0;
            }

            if (data.nightFreqInfo.freqNight == true) {
                this.chartConfModel.nightFreqInfo.freqNight = data.nightFreqInfo.freqNight;
                this.chartConfModel.nightFreqInfo.nightFreqQuantity = data.nightFreqInfo.nightFreqQuantity;
                count = count + 1;
                // desc = desc + " Night &"
            } else {
                this.chartConfModel.nightFreqInfo.freqNight = data.nightFreqInfo.freqNight;
                this.chartConfModel.nightFreqInfo.nightFreqQuantity = 0;
            }

            // code for generating scheduel description.
            switch (count) {
                case 1:
                    let descText = '';
                    if (data.mornFreqInfo.freqMorn) {
                        descText = "morning";
                    } else if (data.aftrnFreqInfo.freqAftrn) {
                        descText = "afternoon";
                    }
                    else if (data.nightFreqInfo.freqNight) {
                        descText = "night";
                    }
                    desc = `Every ${descText} ${foodIns} for ${data.duration} days`;
                    break;
                case 2:
                    let descTextData = '';
                    if (data.mornFreqInfo.freqMorn && data.aftrnFreqInfo.freqAftrn) {
                        descTextData = "morning & afternoon"
                    }
                    else if (data.mornFreqInfo.freqMorn && data.nightFreqInfo.freqNight) {
                        descTextData = "morning & night";
                    } else if (data.aftrnFreqInfo.freqAftrn && data.nightFreqInfo.freqNight) {
                        descTextData = "afternoon & night";
                    }
                    desc = `Every ${descTextData} ${foodIns} for ${data.duration} days`;
                    break;
                case 3:
                    desc = `3 times a day ${foodIns} for ${data.duration} days`;
                    break;
                default:
                    break;
            }
            this.chartConfModel.desc = desc;

        } else if (data.frequency == 1) {  // for every xtimes in a day

            this.chartConfModel.quantity = data.quantity;
            this.chartConfModel.numberofTimes = data.numberofTimes;
            if (data.interval != null) {
                this.chartConfModel.interval = data.interval * 60;
            }
            this.chartConfModel.startTime =TimeConversion.getStartTime(this.datePipe.transform(data.startTime, "H.mm"));

            let description = '';
            let hourMinutsData = TimeConversion.timeConvert(this.chartConfModel.interval);
            description = ` ${data.numberofTimes} times a day after every ${hourMinutsData} for ${data.duration} days.`;
            this.chartConfModel.desc = description;
        } else if (data.frequency == 2) { //  for as required
            this.chartConfModel.splinstruction = data.splinstruction;
            this.chartConfModel.desc = data.splinstruction;
        }

        this.chartConfModel.name = data.name;
        this.chartConfModel.medicinetype = this.medicineType[this.medicineForm.get('medicineType').value];
        this.chartConfModel.startDate =data.startDate
        this.chartConfModel.duration = data.duration;
        this.chartConfModel.frequency = data.frequency;
        this.chartConfModel.foodInst = data.foodInst;
        this.chartConfModel.remark = data.remark;

        let confString = JSON.stringify(this.chartConfModel);
        this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
        this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
        this.chartDbModel.conf = confString;
        this.chartDbModel.conf_type_code = ConfigCodeType.MEDICINE
        this.createActions(this.chartDbModel.uuid, this.chartDbModel.admission_uuid, this.chartDbModel.conf_type_code, confString)
    }
    // >> func for inserting form data to sqlite db


    // << func for creating form controls
    createFormControls(): void {

        this.medicineForm = new FormGroup({
            name: new FormControl('', [Validators.required]),
            quantity: new FormControl(),
            foodInst: new FormControl(),
            frequency: new FormControl(),
            interval: new FormControl(),
            numberofTimes: new FormControl(),
            startDate: new FormControl(),
            duration: new FormControl('', [Validators.required]),
            startTime: new FormControl(),
            remark: new FormControl(),
            mornQuantity: new FormControl(),
            aftrnQuantity: new FormControl(),
            nightQuantity: new FormControl(),
            medicineType: new FormControl(),
            splinstructions: new FormControl(),
        });
    }
    // >> func for creating form controls
    // << func for selecting monitor name
    selectedIndexChanged(args) {
        let picker = <ListPicker>args.object;
        let picked: any;
        // picked = this.monitorConf.dbmodel.conf.tasks[picker.selectedIndex];
        // this.formData.name = picked.name;
        // this.monitorName = picked.name;
    }
// fucntion for getting  medicine type form database
public getMedicineType() {
    this.chartservice.getAllData('medicineType').then(
        (success) => {	          
            if (success.length > 0) {
             const   medicineType = JSON.parse(success[0].conf);				
                this.medicineType = [];
                for (let item of medicineType) {
                    this.medicineType.push(item);
                }
            }
        },
        (error) => {
            console.log("getChartData error:", error);
        }
    );
}

}
