import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { ConfigCodeType, GRACE_PERIOD } from '~/app/app-constants';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { TimeConversion } from '~/app/helpers/time-conversion-helper';
import { AftrnFreqInfo, ChartDBModel, FrequencyValues, MedChartModel, MornFreqInfo, NightFreqInfo } from '~/app/models/ui/chart-models';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';

@Component({
	moduleId: module.id,
	selector: 'medicine-actions',
	templateUrl: './medicine-actions.component.html',
	styleUrls: ['./medicine-actions.component.css']
})

export class MedicineActionsComponent implements OnInit {
	// proccess variables
	medicineForm: FormGroup;
	medNameIsValid: boolean;
	// formData: MedChartModel;
	chartConfModel: MedChartModel;
	chartDbModel: ChartDBModel;
	public medicineType: Array<string> = [];
	// end of proccess variables
	frequencyList: FrequencyValues[] = [
		{ name: "'X'- Times a day", value: 0 },
		{ name: "Every 'X' hours", value: 1 },
		// { name: "As Required", value: 2 }
	];
	isInstruction = false;
	isSplinstructions = false;
	isDosage = false;
	isXTimesDay = false;
	constructor(
		private params: ModalDialogParams,
		private passDataService: PassDataService,
		private chartservice: ChartService) {
			console.log('add medicine actions executed');
		this.chartConfModel = new MedChartModel();
	}
	ngOnInit() {
		
		this.createFormControls();
		this.getMedicineType();
	}

	// << func for navigating previous page
	goBackPage() {
		this.params.closeCallback([]);	
	}
	onSubmit() {
		this.medNameIsValid = this.medicineForm.controls['name'].hasError('required');
		// this.isInstruction = this.medicineForm.controls['desc'].hasError('required');
		this.isSplinstructions = this.medicineForm.controls['splinstructions'].hasError('required');
		this.isDosage = this.medicineForm.controls['quantity'].hasError('required');
		//  validation for checing grace peroid of schedule generation.
		const strDate = new Date();
		console.log('strDate', strDate);
		strDate.setMinutes(strDate.getMinutes() + GRACE_PERIOD);
		const currentDate = new Date();
		console.log('currentDate', currentDate);
		console.log('strDate', strDate);
		//  validation check  if form is valid 
		if (this.medicineForm.invalid) {
			console.log("validation error", this.medicineForm.controls['duration'].hasError('pattern'));
			return;
		}
		// assign form data to model
		const formData = new MedChartModel();
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
		// this.chartConfModel.startDate =data.startDate
		this.chartConfModel.duration = data.duration;
		this.chartConfModel.frequency = data.frequency;
		this.chartConfModel.foodInst = data.foodInst;
		this.chartConfModel.remark = data.remark;

		let confString = JSON.stringify(this.chartConfModel);
		this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
		this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
		this.chartDbModel.conf = confString;
		this.chartDbModel.start_date = data.startDate.toISOString();
		this.chartDbModel.conf_type_code = ConfigCodeType.MEDICINE

	}
	// >> func for inserting form data to sqlite db

	// << func for creating form controls
	createFormControls(): void {
		this.medicineForm = new FormGroup({
			name: new FormControl('', [Validators.required]),
			remark: new FormControl(),
			medicineType: new FormControl(),
			splinstructions: new FormControl(),
		});
	}

	// fucntion for getting  medicine type form database
	public getMedicineType() {
		this.chartservice.getAllData('medicineType').then(
			(success) => {
				if (success.length > 0) {
					const medicineType = JSON.parse(success[0].conf);
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
