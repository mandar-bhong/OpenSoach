import { DatePipe } from '@angular/common';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';
import { ConfigCodeType, SYNC_STORE, ActionStatus, MAXIMUM_SCHEDULE_DURATION } from '~/app/app-constants';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { ChartDBModel, OutputChartModel } from '~/app/models/ui/chart-models';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';
import { VALIDATION_REQUIRED_FIELD } from '~/app/common-constants';
import { OsSelectionListComponent, SELECTION_TYPE } from '~/app/os-selection-list.component';

@Component({
	moduleId: module.id,
	selector: 'output-chart',
	templateUrl: './output-chart.component.html',
	styleUrls: ['./output-chart.component.css']
})

export class OutputChartComponent implements OnInit {

	// proccess variables
	outputForm: FormGroup;

	durationIsValid: boolean;
	chartDbModel: ChartDBModel;
	chartConfModel: OutputChartModel;
	formData: OutputChartModel;
	// end of proccess variables
	public outputType: Array<string> = [];
	outputList: string[] = [];
	pattern = '^[0-9]*$';
	// end of proccess variables
	VALIDATION_REQUIRED_FIELD = VALIDATION_REQUIRED_FIELD;

	@ViewChild("outputSelectionControl",{static:true}) outputSelectionCtl: OsSelectionListComponent;
	SELECTION_TYPE = SELECTION_TYPE;

	constructor(
		private passDataService: PassDataService,
		private datePipe: DatePipe,
		private params: ModalDialogParams,
		public workerService: WorkerService,
		private chartService: ChartService) {

		this.chartDbModel = new ChartDBModel();
		this.chartConfModel = new OutputChartModel();
	}

	public displayText = (item: any) => {
        return item;
    }

	ngOnInit() {

		// creating form control
		this.createFormControls();
		this.getOutputType();
		// get montior conf data for list picker
		this.outputForm.get('startDate').setValue(new Date());
	}

	selectionOutputData() {
        this.outputSelectionCtl.Items = this.outputType;
        if (this.outputSelectionCtl.Items.length > 0) {
            this.outputSelectionCtl.SelectedItems.push(this.outputSelectionCtl.Items[0]);
        }
        this.outputSelectionCtl.Init();
    }

	// << func for navigating previous page
	goBackPage() {
		this.params.closeCallback([]);
		// this.routerExtensions.navigate(['patientmgnt', 'details'], { clearHistory: true });
	}
	// >> func for navigating previous page

	// << func for submit form data
	onSubmit() {
		this.formData = new OutputChartModel();
		this.durationIsValid = this.outputForm.controls['duration'].hasError('required');
		if (this.outputForm.invalid) {
			console.log("validation error");
			return;
		};
		// assign form data to model
		this.formData = Object.assign({}, this.outputForm.value);
		this.insertData(this.formData);
	}
	// >> func for submit form data

	// << func for inserting form data to sqlite db
	insertData(data: OutputChartModel) {
		this.outputSelectionCtl.SelectedItems.forEach(element => {
            this.chartConfModel.name = element;
		});
		// this.chartConfModel.name = this.outputType[data.outputType];
		this.chartConfModel.duration = data.duration;
		this.chartConfModel.remark = data.remark;
		//this.chartConfModel.frequency = 2;
		const desc = `Check ${this.chartConfModel.name} for ${data.duration} days`;
		this.chartConfModel.desc = desc;
		// this.chartConfModel.startDate = data.startDate;
		// this.chartConfModel.foodInst = data.foodInst;
		let confString = JSON.stringify(this.chartConfModel);
		// set db model
		this.chartDbModel.uuid = PlatformHelper.API.getRandomUUID();
		this.chartDbModel.admission_uuid = this.passDataService.getAdmissionID();
		this.chartDbModel.conf = confString;
		this.chartDbModel.start_date = data.startDate.toISOString();
		this.chartDbModel.conf_type_code = ConfigCodeType.OUTPUT;
		console.log(this.chartDbModel);
		this.createActions(this.chartDbModel, confString);
		// get chart data from sqlite db
		// this.chartService.getChartList();
		// this.goBackPage();
	}


	// << func for creating form controls
	createFormControls(): void {
		this.outputForm = new FormGroup({
			// foodInst: new FormControl(),
			outputType: new FormControl(),
			duration: new FormControl('', [Validators.required, Validators.pattern(this.pattern), Validators.max(MAXIMUM_SCHEDULE_DURATION)]),
			startDate: new FormControl(),
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

	// fucntion for getting  output type form database
	public getOutputType() {
		this.chartService.getAllData('outputType').then(
			(success) => {
				let outputtype;
				if (success.length > 0) {
					outputtype = JSON.parse(success[0].conf);
					this.outputType = [];
					for (let item of outputtype) {
						this.outputType.push(item);
					}
				}
				this.selectionOutputData();
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}
}// end of class