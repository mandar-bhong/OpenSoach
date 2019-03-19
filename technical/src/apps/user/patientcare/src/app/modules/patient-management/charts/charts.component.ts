import { Component, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { ModalDialogOptions, ModalDialogService } from 'nativescript-angular/modal-dialog';
import { RouterExtensions } from 'nativescript-angular/router';
import { ListViewEventData, ListViewItemSnapMode } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Subscription } from 'rxjs';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { ConfigCodeType, freuencyone, freuencyzero, SERVER_WORKER_MSG_TYPE, SYNC_STORE, ScheuldeStatus } from '~/app/app-constants';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { IDatastoreModel } from '~/app/models/db/idatastore-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { ChartListViewModel, ConfigData } from '~/app/models/ui/chart-models';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';

import { IDeviceAuthResult } from '../../idevice-auth-result';
import { SchedularFabComponent } from '../schedular-fab/schedular-fab.component';
import { IntakeChartComponent } from './intake-chart/intake-chart.component';
import { MedicineChartComponent } from './medicine-chart/medicine-chart.component';
import { MonitorChartComponent } from './monitor-chart/monitor-chart.component';
import { TraceCustomCategory } from '~/app/helpers/trace-helper';
import * as trace from 'trace';
@Component({
	moduleId: module.id,
	selector: 'charts',
	templateUrl: './charts.component.html',
	styleUrls: ['./charts.component.css']
})

export class ChartsComponent implements OnInit, IDeviceAuthResult {

	chartListItems: ObservableArray<ChartListViewModel>;
	ServerDataStoreDataModelArray: ServerDataStoreDataModel<ScheduleDatastoreModel>[] = [];
	//chartListItemsSource = new ObservableArray<ChartListViewModel>();
	// >> seleced bottom button change color
	monitorbuttonClicked: boolean = false;
	intakebuttonClicked: boolean = true;
	medicinebuttonClicked: boolean = false;
	outputbuttonClicked: boolean = false;
	isSchedularDataReceived = false;
	schedulecreationSubscription: Subscription;
	scheduleDataContext: Subscription;
	// >> finding grouping index then after click show in top
	intakeIndex;
	medicineIndex;
	monitorIndex;
	outputIndex;

	// >> grouping 
	public _funcGrouping: (item: ChartListViewModel) => ChartListViewModel;

	dialogOpen = false;
	expanded = false;
	configCodeType = ConfigCodeType;
	completeorpending: string;
	iscompleted: boolean;
	freuencyZero = freuencyzero;
	freuencyOne = freuencyone;

	buttonClicked: boolean = true;
	buttonCompleted: boolean = false;
	constructor(private chartService: ChartService,
		public workerservice: WorkerService,
		public passdataservice: PassDataService,
		private routerExtensions: RouterExtensions,
		private modalService: ModalDialogService,
		private viewContainerRef: ViewContainerRef) {
		//  list grouping
		this._funcGrouping = (item: any) => {
			return item.dbmodel.conf_type_code;
		};
	}

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.getChartData('getScheduleListActive');
		this.schedulecreationSubscription = this.workerservice.scheduleDataReceivedSubject.subscribe((value) => {
			trace.write('notified to schedule list page', TraceCustomCategory.SCHEDULE, trace.messageType.info);
			console.log('<======================notified to schedule list page===> ', value);
			this.pushAddedSchedule(value);
		});
		this.completeorpending = "Active Schedules";
		this.scheduleDataContext = this.chartService.scheduleDataContext.subscribe((value) => {
			// checking if schedulearrray  have any records.
			if (value.length > 0) {
				this.ServerDataStoreDataModelArray = value;
				this.isSchedularDataReceived = true;
				//	this.savetoUserAuth();
			}
		});
	}
	// code for showing fab button dialog.
	showDialog() {
		this.createModalView(SchedularFabComponent, false).then((dialogResult) => {
			if (dialogResult) {
				setTimeout(() => {
					this.openModel(dialogResult);
				});
			}
		});
		//	this.dialogOpen = true;
	}

	closeDialog() {
		this.dialogOpen = false;
	}


	// >> Grouping position change 
	// >> Grouping intake scroll to top position change 
	public selectIntake() {
		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(this.intakeIndex, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = true;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
	}
	// <<  Grouping intake scroll to top position change 

	// >>  Grouping monitor scroll to top position change 
	public selectMonitor() {
		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(this.monitorIndex, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = true;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
	}
	// <<  Grouping monitor scroll to top position change 

	// >>  Grouping medicine scroll to top position change 
	public selectMedicine() {
		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(this.medicineIndex, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = true;
		this.outputbuttonClicked = false;
	}
	// <<  Grouping medicine scroll to top position change 

	// >>  Grouping medicine scroll to top position change
	public selectOutput() {
		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(this.outputIndex, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = true;
	}
	// <<  Grouping medicine scroll to top position change

	// >> Calculate Grouping index value
	public getGroupIndex() {

		const medicine = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Medicine");
		const medicineCount = medicine.length;

		const monitor = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Monitor");
		const monitorCount = monitor.length;

		const intake = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Intake");
		const intakeCount = intake.length;

		const output = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Output");
		const outputCount = output.length;

		this.intakeIndex = 0;
		this.medicineIndex = intakeCount + 1;
		this.monitorIndex = intakeCount + medicineCount + 2;
		this.outputIndex = intakeCount + medicineCount + monitorCount + 3;

	}

	get _listItems(): ObservableArray<ChartListViewModel> {
		return this.chartListItems;
	}

	public getChartData(key: string) {
		this.chartListItems = new ObservableArray<ChartListViewModel>();
		this.chartService.getScheduleList(key, this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					let chartListItem = new ChartListViewModel();
					chartListItem.dbmodel = item;
					chartListItem.conf = new ConfigData();
					Object.assign(chartListItem.conf, JSON.parse(chartListItem.dbmodel.conf));
					this.chartListItems.push(chartListItem);
				});
				this.getGroupIndex();
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}
	monitorForm() {
		this.openModel(MonitorChartComponent);
		this.dialogOpen = false;
		//this.routerExtensions.navigate(['patientmgnt', 'monitor-chart'], { clearHistory: false });
	}
	medicineForm() {
		this.openModel(MedicineChartComponent);
		this.dialogOpen = false;
		//	this.routerExtensions.navigate(['patientmgnt', 'medicine-chart'], { clearHistory: false, });
	}
	intakeForm() {
		this.openModel(IntakeChartComponent);
		this.dialogOpen = false;
		//this.routerExtensions.navigate(['patientmgnt', 'intake-chart'], { clearHistory: false });
	}
	ngOnDestroy(): void {
		//Called once, before the instance is destroyed.
		//Add 'implements OnDestroy' to the class.
		if (this.schedulecreationSubscription) {
			this.schedulecreationSubscription.unsubscribe();
		}
	}// end 

	pushAddedSchedule(Schedule: ScheduleDatastoreModel) {
		let scheduleDatastoreModel = new ScheduleDatastoreModel();
		scheduleDatastoreModel = Schedule;
		const chartListViewModel = new ChartListViewModel();
		chartListViewModel.dbmodel = scheduleDatastoreModel;
		chartListViewModel.conf=JSON.parse(scheduleDatastoreModel.conf);
		if (scheduleDatastoreModel.status == ScheuldeStatus.SCHEDULE_ACTIVE) {
			const end_date = new Date(chartListViewModel.dbmodel.end_date);
			const todaysdate = new Date();
			// checking current mode of list view based on that decide wheater we have to add newly created scchedule in list or not.
			if (this.iscompleted) {
				if (todaysdate < end_date) {
					return;
				}
			} else if (!this.iscompleted) {
				if (todaysdate > end_date) {
					return;
				}
			} // end
			try {
				const scheduleitem = this.chartListItems.filter(data => data.dbmodel.uuid === chartListViewModel.dbmodel.uuid)[0];
				// ittem found in array 
				if (scheduleitem && scheduleitem != null) {
					const itemIndex = this.chartListItems.indexOf(scheduleitem);
					this.chartListItems[itemIndex].dbmodel = chartListViewModel;
					this.chartListItems[itemIndex].conf = JSON.stringify(chartListViewModel.dbmodel.conf)
				} else {
					this.chartListItems.push(chartListViewModel);
				}
			} catch (e) {
				console.log(e.error);
			}
		} else if (scheduleDatastoreModel.status == ScheuldeStatus.SCHEDULE_CANCELLED) {  // schedule is cancelled 
			const scheduleitem = this.chartListItems.filter(data => data.dbmodel.uuid === scheduleDatastoreModel.uuid)[0];
			if (scheduleitem) {
				const itemIndex = this.chartListItems.indexOf(scheduleitem);
				this.chartListItems.splice(itemIndex, 1);
			}
		}

	}

	// to do
	// remove this after action  heleper code testing
	test() {
		const initModel = new ServerDataProcessorMessageModel();
		const serverDataStoreModel = new ServerDataStoreDataModel<ScheduleDatastoreModel>();
		serverDataStoreModel.datastore = SYNC_STORE.SCHEDULE;
		serverDataStoreModel.data = new ScheduleDatastoreModel();
		serverDataStoreModel.data.uuid = 'cb86aeee-c21b-475e-ab68-1335f97c9b9b'
		serverDataStoreModel.data.sync_pending = 1
		serverDataStoreModel.data.admission_uuid = "11";
		serverDataStoreModel.data.conf_type_code = ConfigCodeType.MEDICINE
		serverDataStoreModel.data.conf = '{"mornFreqInfo":{"freqMorn":true},"aftrnFreqInfo":{"freqAftrn":true},"nightFreqInfo":{"freqNight":true},"desc":" Morning & Afternoon & Night before meal Test.","name":"Cipla ks","quantity":11,"startDate":"2019-01-23T08:30:00.438Z","duration":3,"frequency":1,"startTime":"20.30","intervalHrs":180,"foodInst":1,"endTime":"12.30","numberofTimes":3,"specificTimes":[11.3,12.3]}';
		initModel.data = [serverDataStoreModel];
		initModel.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;
		this.workerservice.ServerDataProcessorWorker.postMessage(initModel);

	}
	public listLoaded() {
		// console.log('this.isSchedularDataReceived', this.isSchedularDataReceived);
		// if (this.isSchedularDataReceived) {
		// 	this.savetoUserAuth();
		// }
	}
	public activeList() {
		this.completeorpending = "Active Schedules";
		this.iscompleted = false;
		this.buttonCompleted = false;
		this.getChartData('getScheduleListActive');
	}
	public compilitedList() {
		this.iscompleted = true;
		this.buttonClicked = false;
		this.completeorpending = "Completed Schedules";
		this.getChartData('getScheduleListComplated');

	}

	// public onListSorting(args) {
	// 	let firstSwitch = <Switch>args.object;
	// 	if (firstSwitch.checked) {
	// 		this.completeorpending = "Completed Schedules";
	// 		this.iscompleted = true;
	// 		this.getChartData('getScheduleListComplated');
	// 	} else {
	// 		this.completeorpending = "Active Schedules";
	// 		this.iscompleted = false;
	// 		this.getChartData('getScheduleListActive');

	// 	}
	// }
	// on click of button set current service name in passdata service for ref invocatios
	// like authResultReuested=this or name of your component.
	savetoUserAuth() {
		setTimeout(() => {
			this.passdataservice.authResultReuested = this;
			this.routerExtensions.navigate(['patientmgnt', 'user-auth'], { clearHistory: false });
		}, 2000);

	}
	onDeviceAuthSuccess(userid: number): void {
		console.log('chart componenent onDeviceAuthSuccess executed');
		this.ServerDataStoreDataModelArray.forEach(element => {
			element.data.updated_by = userid;
		});
		const initModel = new ServerDataProcessorMessageModel();
		initModel.data = this.ServerDataStoreDataModelArray
		initModel.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;
		this.workerservice.ServerDataProcessorWorker.postMessage(initModel);
	}
	onDeviceAuthError(error: any): void {
		throw new Error("Method not implemented.");
	}
	onSubmitDiscarded(): void {
		throw new Error("Method not implemented.");
	}
	// >> expand row code start
	templateSelector(item: any, index: number, items: any): string {
		return item.expanded ? "expanded" : "default";

	}

	onItemTap(event: ListViewEventData) {
		const listView = event.object,
			rowIndex = event.index,
			dataItem = event.view.bindingContext;

		dataItem.expanded = !dataItem.expanded;
		if (isIOS) {
			// Uncomment the lines below to avoid default animation
			// UIView.animateWithDurationAnimations(0, () => {
			var indexPaths = NSMutableArray.new();
			indexPaths.addObject(NSIndexPath.indexPathForRowInSection(rowIndex, event.groupIndex));
			listView.ios.reloadItemsAtIndexPaths(indexPaths);
			// });
		}
		if (isAndroid) {
			listView.androidListView.getAdapter().notifyItemChanged(rowIndex);

		}
	}
	// << expand row code end
	private createModalView(Component, isFullScreen): Promise<any> {
		let options: ModalDialogOptions = {
			context: { promptMsg: "This is the prompt message!" },
			fullscreen: isFullScreen,
			viewContainerRef: this.viewContainerRef
		};
		return this.modalService.showModal(Component, options);
	}

	openModel(componentName) {
		this.createModalView(componentName, true).then((dialogResult: ServerDataStoreDataModel<ScheduleDatastoreModel>[]) => {
			console.log('dialogResult', dialogResult);
			this.ServerDataStoreDataModelArray = dialogResult;
			if (this.ServerDataStoreDataModelArray.length > 0) {
				setTimeout(() => {
					this.savetoUserAuth();
				});
			}
			this.dialogOpen = false;
		});

	}// end 
	//code block for cancel schedule
	cancelScheudle(scheduleItem: ChartListViewModel) {
		// let chartListViewModel = new ChartListViewModel();
		// chartListViewModel.serverdbmodal = scheduleItem.serverdbmodal;
		// console.log('chartListViewModel', chartListViewModel);
		// chartListViewModel.serverdbmodal.status = 1;
		const serverDataStoreModel = new ServerDataStoreDataModel<ScheduleDatastoreModel>();
		serverDataStoreModel.datastore = SYNC_STORE.SCHEDULE;
		serverDataStoreModel.data = new ScheduleDatastoreModel();
		Object.assign(serverDataStoreModel.data, scheduleItem.dbmodel);
		console.log('schedule data',scheduleItem.dbmodel);
		serverDataStoreModel.data.status = 1;
		serverDataStoreModel.data.sync_pending = 1
		serverDataStoreModel.data.client_updated_at = new Date().toISOString();
		this.ServerDataStoreDataModelArray = [];
		this.ServerDataStoreDataModelArray.push(serverDataStoreModel);
		this.savetoUserAuth();
	}

} 