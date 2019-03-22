import { Component, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { SnackBar, SnackBarOptions } from 'nativescript-snackbar';
import {
	ListViewEventData,
	ListViewItemSnapMode,
	ListViewLinearLayout,
	LoadOnDemandListViewEventData,
	RadListView,
} from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Subscription } from 'rxjs/internal/Subscription';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { error } from 'tns-core-modules/trace/trace';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { layout } from 'tns-core-modules/utils/utils';
import { Page } from 'ui/page';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { ActionListViewModel, ActionTxnDBModel, ActionDataDBRequest } from '~/app/models/ui/action-models';
import { Schedulardata, SchedularConfigData, MornFreqInfo, AftrnFreqInfo, NightFreqInfo } from '~/app/models/ui/chart-models';
import { ActionService } from '~/app/services/action/action.service';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { IntakeHelper } from '~/app/helpers/actions/intake-helper';
import { MedicineHelper } from '~/app/helpers/actions/medicine-helper';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { SYNC_STORE, SERVER_WORKER_MSG_TYPE, ConfigCodeType, ACTION_STATUS } from '~/app/app-constants';
import { ActionTxnDatastoreModel } from '~/app/models/db/action-txn-model';
import { WorkerService } from '~/app/services/worker.service';
import { RouterExtensions } from 'nativescript-angular/router';
import { IDeviceAuthResult } from '../../idevice-auth-result';
import { Switch } from 'tns-core-modules/ui/switch/switch';
import { ModalDialogService, ModalDialogOptions } from 'nativescript-angular/modal-dialog';
import { DoctorOrdersComponent } from '../doctor-orders/doctor-orders.component';
import { IDatastoreModel } from '~/app/models/db/idatastore-model';
import { ActionFabComponent } from '../action-fab/action-fab.component';
import { ActionStatusHelper } from '~/app/helpers/action-status-helper';
import { DataActionItem, BloodPressureValueModel, GetJsonmodel } from '~/app/models/ui/action-model';


// expand row 
declare var UIView, NSMutableArray, NSIndexPath;
// import { TextField } from "ui/text-field";

// SnackBar import 

@Component({
	moduleId: module.id,
	selector: 'action',
	templateUrl: './action.component.html',
	styleUrls: ['./action.component.css']
})

export class ActionComponent implements OnInit, IDeviceAuthResult {
	dialogOpen = false;
	conf_type_code_const = ConfigCodeType;
	onDeviceAuthSuccess(userid: number): void {
		console.log('user auth id', userid);
		console.log('chart componenent onDeviceAuthSuccess executed');
		this.ServerDataStoreDataModelArray.forEach(element => {
			element.data.updated_by = userid;
			console.log('get data userid', userid);
		});
		const initModel = new ServerDataProcessorMessageModel();
		initModel.data = this.ServerDataStoreDataModelArray
		initModel.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;
		this.workerservice.ServerDataProcessorWorker.postMessage(initModel);

		//	this.save();
		//	throw new Error("Method not implemented.");
	}
	onDeviceAuthError(error: any): void {
		throw new Error("Method not implemented.");
	}
	onSubmitDiscarded(): void {
		throw new Error("Method not implemented.");
	}
	public actionListItem = new ObservableArray<ActionListViewModel>();
	ServerDataStoreDataModelArray: ServerDataStoreDataModel<any>[] = [];
	chartbuttonClicked: boolean = false;
	public _dataItems: ObservableArray<any>;
	private layout: ListViewLinearLayout;
	schedulardata: Schedulardata;
	monitorschedulardata: Schedulardata[] = []
	outputschedulardata: Schedulardata[] = []

	// >> seleced bottom button change color
	monitorbuttonClicked: boolean = false;
	intakebuttonClicked: boolean = true;
	medicinebuttonClicked: boolean = false;
	doctorOrderButtonClicked = false;
	outputbuttonClicked: boolean = false;
	actionSubscription: Subscription;
	doctorOrderSubscription: Subscription;
	// >> search var declaration
	tempdata = new Array<ActionTxnDBModel>();


	// >> grouping 
	public _funcGrouping: (item: DataActionItem) => DataActionItem;
	// >> exapnd row
	expanded: false;
	viewexpand = false;
	// >> finding grouping index then after click show in top
	intakeIndex: any;
	medicineIndex: any;
	monitorIndex: any;
	outputIndex: any;
	doctorOrderIndex: any;

	// >>  bottom snackbar msg
	private snackbar: SnackBar;

	actionListItems = new ActionListViewModel();
	tempList = new ObservableArray<DataActionItem>();

	// >> details form field
	// actionForm: FormGroup;
	formData: ActionTxnDBModel;
	actionformData: ActionTxnDBModel;
	actiondata: ActionDataDBRequest;
	actionDbData: ActionDataDBRequest;
	actionDbArray: ActionTxnDBModel[] = [];
	confString;
	confString1;
	saveViewOpen = false;
	exectime;
	// listaccount = true;
	// removeAccount = false;
	_dataItemsaccount = new ObservableArray<ActionListViewModel>();
	// blood pressure high and low value model
	bloodPressureValueModel = new BloodPressureValueModel();
	bloddname: any;
	// switch active and complited
	completeorpending: string;
	iscompleted: boolean;
	// filter buttton 
	buttonClicked: boolean = true;
	buttonCompleted: boolean = false;
	getAllFlag = false;

	// filter data complited UI readonly Mode
	editMode = false;

	activeAction = new ObservableArray<DataActionItem>();
	allAction = new ObservableArray<DataActionItem>();
	action_Status = ACTION_STATUS;
	get_Value: BloodPressureValueModel;
	actionStatus: any;
	conf_type_code: any;
	constructor(public page: Page,
		private actionService: ActionService,
		public workerService: WorkerService,
		private modalService: ModalDialogService,
		private passdataservice: PassDataService,
		private workerservice: WorkerService,
		private viewContainerRef: ViewContainerRef,
		private chartService: ChartService,
		private routerExtensions: RouterExtensions) {
		//  list grouping
		this._funcGrouping = (item: any) => {
			return item.conf_type_code;
		};
		this.formData = new ActionTxnDBModel();
		// >>  bottom snackbar msg
		this.snackbar = new SnackBar();
		this.actionDbData = new ActionDataDBRequest();
	}
	get dataItems(): ObservableArray<DataActionItem> {
		return this._dataItems;
	}
	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {

		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this.getActionData();

		// subscription for create actions
		this.actionSubscription = this.passdataservice.createActionsSubject.subscribe((value) => {

		}); // end of subscriptions.
		this.actionDbData = new ActionDataDBRequest();
		this.actiondata = new ActionDataDBRequest();

		this.completeorpending = "Active Actions";


		// subscription for adding newly  created doctors orders in action list.
		this.doctorOrderSubscription = this.workerService.doctorOrderSubject.subscribe((value) => {
			this.pushDoctorOredrs(value);
		});
		this.passdataservice.backalert = false;
	}// end of ng init.


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

	// >> Grouping position change 

	// >> Grouping intake scroll to top position change 
	public selectIntake() {
		if (this.conf_type_code === "Intake") {
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(0, false, ListViewItemSnapMode.Start);

			console.log("Clicked select intake", this.intakeIndex);
		}
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = true;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
		this.doctorOrderButtonClicked = false;
	}
	// <<  Grouping intake scroll to top position change 

	// >>  Grouping monitor scroll to top position change 
	public selectMonitor() {
		if (this.conf_type_code === "Monitor") {
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(this.monitorIndex, false, ListViewItemSnapMode.Start);

			console.log("Clicked select monitor", this.monitorIndex);
		}
		this.monitorbuttonClicked = true;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
		this.doctorOrderButtonClicked = false;
	}
	// <<  Grouping monitor scroll to top position change 

	// >>  Grouping medicine scroll to top position change 
	public selectMedicine() {
		console.log('this.conf_type_code button', this.conf_type_code);
		if (this.conf_type_code === "Medicine") {
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(this.medicineIndex, false, ListViewItemSnapMode.Start);

			console.log("Clicked select medicine", this.medicineIndex);
		}
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = true;
		this.outputbuttonClicked = false;
		this.doctorOrderButtonClicked = false;

	}
	// <<  Grouping medicine scroll to top position change 

	// >>  Grouping medicine scroll to top position change
	public selectOutput() {
		if (this.conf_type_code === "Output") {
			console.log('this.conf_type_code button output', this.conf_type_code);
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(this.outputIndex, false, ListViewItemSnapMode.Start);
			console.log("Clicked select output", this.outputIndex);
		}
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = true;
		this.doctorOrderButtonClicked = false;
	}
	selectDoctorOrder() {
		console.log();
		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(this.doctorOrderIndex, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = false;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
		this.doctorOrderButtonClicked = true;
	}
	// <<  Grouping medicine scroll to top position change

	// >> Calculate Grouping index value
	public getCount() {

		const medicine = this.tempList.filter(a => a.conf_type_code === "Medicine");
		const medicineCount = medicine.length;
		// console.log("medicineCount", medicineCount);

		const monitor = this.tempList.filter(a => a.conf_type_code === "Monitor");
		const monitorCount = monitor.length;
		// console.log("monitorCount", monitorCount);

		const intake = this.tempList.filter(a => a.conf_type_code === "Intake");
		const intakeCount = intake.length;
		// console.log("intakeCount", intakeCount);

		const output = this.tempList.filter(a => a.conf_type_code === "Output");
		const outputCount = output.length;
		const DOrder = this.tempList.filter(a => a.conf_type_code === ConfigCodeType.DOCTOR_ORDERS);
		const DOrderCount = DOrder.length;
		console.log("DOrderCount", DOrderCount);


		this.doctorOrderIndex = 0;
		this.intakeIndex = DOrderCount + 1;
		this.medicineIndex = DOrderCount + intakeCount + 1;
		this.monitorIndex = DOrderCount + intakeCount + medicineCount + 1;
		this.outputIndex = DOrderCount + intakeCount + medicineCount + monitorCount + 1;
		console.log("medicine index", this.medicineIndex);
		console.log("monitor index", this.monitorIndex);
		console.log("output index", this.outputIndex);
		console.log("DR OR index", this.doctorOrderIndex);

	}
	// << Calculate Grouping index value

	// clean up
	ngOnDestroy(): void {
		//Called once, before the instance is destroyed.
		if (this.actionSubscription) { this.actionSubscription.unsubscribe(); }
		if (this.doctorOrderSubscription) { this.doctorOrderSubscription.unsubscribe(); }

	}

	public getActionData() {
		console.log('getActinData')
		this.actionListItem = new ObservableArray<ActionListViewModel>();
		this.actionService.getallActionActiveList(this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					console.log("get action item ", item);
					let actionListItem = new ActionListViewModel();
					actionListItem.dbmodel = item;
					this.actionListItem.push(actionListItem);
				});
				// console.log('this.actionListItem', this.actionListItem);
				this.getListDataById();
			},
			(error) => {
				console.log("getActinData error:", error);
			}
		);
	}

	// >> get action list by id
	public getListDataById() {
		this.tempList = new ObservableArray<DataActionItem>();
		this.activeAction = new ObservableArray<DataActionItem>();
		this.allAction = new ObservableArray<DataActionItem>();

		this.actionListItem.forEach(item => {
			console.log('txn item', item);
			const actionListDataItem = new DataActionItem();
			const gettxn_data = new GetJsonmodel();
			gettxn_data.jsonvalue = new BloodPressureValueModel();
			actionListDataItem.admission_uuid = item.dbmodel.admission_uuid;
			actionListDataItem.schedule_uuid = item.dbmodel.schedule_uuid;
			actionListDataItem.conf_type_code = item.dbmodel.conf_type_code;
			this.conf_type_code = actionListDataItem.conf_type_code;
			console.log('this.conf_type_code', this.conf_type_code);
			actionListDataItem.scheduled_time = item.dbmodel.scheduled_time;
			actionListDataItem.fname = item.dbmodel.fname;
			actionListDataItem.lname = item.dbmodel.lname;
			// const a = item.dbmodel.scheduled_time;


			// >> will display list of actions to be performed in another 12 hours and after 1 hours
			const recivedDateFormDB = new Date(item.dbmodel.scheduled_time);
			const recivedDateDb = recivedDateFormDB.getMinutes();
			recivedDateFormDB.setMinutes(recivedDateDb);
			const reciveTimeDb = recivedDateFormDB.toLocaleString();
			const Dbdate = new Date(reciveTimeDb);
			// console.log('Dbdate',Dbdate);

			const tempEndTime = new Date();
			const next12Hours = tempEndTime.getMinutes() + 720;
			tempEndTime.setMinutes(next12Hours);
			const tempEnd = tempEndTime.toLocaleString();
			const endTime = new Date(tempEnd);
			// console.log('endTime',endTime);

			const tempStartTime = new Date();
			const after1Hours = tempStartTime.getMinutes() - 60;
			tempStartTime.setMinutes(after1Hours);
			const tempStart = tempStartTime.toLocaleString();
			const startTime = new Date(tempStart);
			// console.log('startTime',startTime);
			actionListDataItem.txn_state = item.dbmodel.txn_state;
			Object.assign(gettxn_data, JSON.parse(item.dbmodel.txn_data));
			actionListDataItem.txn_data = gettxn_data;
			actionListDataItem.txn_data.comment = gettxn_data.comment;
			console.log('actionListDataItem.value.systolic', gettxn_data.jsonvalue.systolic);




			const conf = JSON.parse(item.dbmodel.conf);
			// console.log('conf ', conf);
			actionListDataItem.name = conf.name;
			// console.log('conf name', conf.name);
			actionListDataItem.desc = conf.desc;



			if (actionListDataItem.name === 'Blood Pressure') {
				// console.log('if value');
				// Object.assign(gettxnData.value, JSON.parse(gettxn_data.value));
				// console.log('gettxnData.value.systolic', gettxnData.value.systolic);
				// actionListDataItem.value.systolic = gettxnData.value.systolic;
				// console.log('************systolic', gettxnData.value.systolic);
				// actionListDataItem.value.diastolic = gettxnData.value.diastolic;
				// console.log('************diastolic', gettxnData.value.diastolic);
			} else {
				actionListDataItem.txn_data.value = gettxn_data.value;
				console.log('else value', actionListDataItem.txn_data.value);
			}

			//  condition for type set action list active and all
			if (item.dbmodel.action_txn_uuid != null) {
				// set type  Completed action 2
				const recivedDateFormDB = new Date(item.dbmodel.client_updated_at);
				const recivedDateDb = recivedDateFormDB.getMinutes();
				recivedDateFormDB.setMinutes(recivedDateDb);
				const reciveTimeDb = recivedDateFormDB.toLocaleString();
				actionListDataItem.client_updated_at = new Date(reciveTimeDb);
				actionListDataItem.type = 2;
				console.log('type 2 =======');
				this.allAction.push(actionListDataItem);
			} else {
				console.log('type 1 ******');
				// set type Active action1 
				actionListDataItem.type = 1;
				actionListDataItem.actionStatus = ActionStatusHelper.getActionStatus(Dbdate);
				this.actionStatus = actionListDataItem.actionStatus;
				if (actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_NORMAL || actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_DELAYED || actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_NEEDS_ATTENTION) {
					if (Dbdate >= startTime && Dbdate <= endTime) {
						this.activeAction.push(actionListDataItem);
					}

				}
				this.allAction.push(actionListDataItem);
			}
			// status color set
			if (actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_NORMAL) {
				actionListDataItem.status = 3;
			} else if (actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_DELAYED) {
				actionListDataItem.status = 1;
			} else if (actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_NEEDS_ATTENTION) {
				actionListDataItem.status = 2;
			} else if (actionListDataItem.actionStatus == ACTION_STATUS.ACTIVE_FUTURE) {
				actionListDataItem.status = 4;
			} else if (actionListDataItem.actionStatus == ACTION_STATUS.MISSED) {
				actionListDataItem.status = 1;
			}

			if (this.getAllFlag == true) {
				// all
				this.tempList = this.allAction;
			} else {
				//  active
				this.tempList = this.activeAction;
			}

		});
		// get doctor orders.
		this.getDoctorsOrders();

	}


	// >> on submit one bye one item data
	onSubmit(item) {
		console.log('on sumbit', item);
		//set action conf model
		this.passdataservice.backalert = true;
		this.itemSelected(item);
		this.saveViewOpen = true;
		this.formData = new ActionTxnDBModel();

		// >> check condition medicine data not add comment and value entries
		if (item.conf_type_code === 'Medicine') {
			this.actionDbData.comment = this.actiondata.comment;
			this.actionDbData.value = null;
			this.confString1 = JSON.stringify(this.actionDbData);
		} else {
			if (item.name === 'Blood Pressure') {
				console.log('if  condition');
				const bloodPressureValueModel = new BloodPressureValueModel()
				bloodPressureValueModel.systolic = this.bloodPressureValueModel.systolic;
				bloodPressureValueModel.diastolic = this.bloodPressureValueModel.diastolic;
				this.actionDbData.value = JSON.stringify(bloodPressureValueModel);
				this.actionDbData.comment = this.actiondata.comment;
			} else {
				console.log('else condition');
				this.actionDbData.value = this.actiondata.value;
				this.actionDbData.comment = this.actiondata.comment;
			}
			this.confString = JSON.stringify(this.actionDbData);
		}

		// set db model 
		this.formData.uuid = PlatformHelper.API.getRandomUUID();
		this.formData.schedule_uuid = item.schedule_uuid;

		// >> check condition medicine data in josn format push
		if (item.conf_type_code === 'Medicine') {
			this.formData.txn_data = this.confString1;
		} else {
			this.formData.txn_data = this.confString;
		}
		this.formData.conf_type_code = item.conf_type_code;
		this.formData.runtime_config_data = null;
		this.formData.scheduled_time = item.scheduled_time;
		this.formData.txn_state = 1;
		this.formData.status = 1;
		this.formData.admission_uuid = item.admission_uuid;

		// after done data push one by one ietm in array hold data
		this.actionDbArray.push(this.formData);
		console.log('this.actionDbArray', this.actionDbArray);



	}
	// >> on discard one bye one item data
	onDiscard(item) {
		console.log(item);
		//set action conf model

		this.passdataservice.backalert = true;
		this.itemSelected(item);
		this.saveViewOpen = true;
		this.formData = new ActionTxnDBModel();

		// >> check condition medicine data not add comment and value entries
		if (item.conf_type_code === 'Medicine') {
			this.actionDbData.comment = null;
			this.actionDbData.value = null;
			this.confString1 = JSON.stringify(this.actionDbData);
		} else {
			this.actionDbData.comment = this.actiondata.comment;
			this.actionDbData.value = this.actiondata.value;
			this.confString = JSON.stringify(this.actionDbData);
		}

		// set db model 
		this.formData.uuid = PlatformHelper.API.getRandomUUID();
		this.formData.schedule_uuid = item.schedule_uuid;

		// >> check condition medicine data in josn format push
		if (item.conf_type_code === 'Medicine') {
			this.formData.txn_data = this.confString1;
		} else {
			this.formData.txn_data = this.confString;
		}
		this.formData.conf_type_code = item.conf_type_code;
		this.formData.runtime_config_data = null;
		this.formData.scheduled_time = item.scheduled_time;
		this.formData.txn_state = 2;
		this.formData.status = 0;
		this.formData.admission_uuid = item.admission_uuid;

		// after done data push one by one ietm in array hold data
		this.actionDbArray.push(this.formData);
		// console.log('this.actionDbArray', this.actionDbArray);


	}
	// all action done and discard save in action-trn-table
	savetoUserAuth() {
		this.passdataservice.authResultReuested = this;
		this.routerExtensions.navigate(['patientmgnt', 'user-auth'], { clearHistory: false });
	}
	save() {
		// array hold entries one by one save
		this.actionformData = new ActionTxnDBModel();
		// insert Action db model to sqlite db
		this.actionDbArray.forEach(item => {
			console.log('item', item);
			const actionModel = new ServerDataProcessorMessageModel();
			const serverDataStoreModel = new ServerDataStoreDataModel<ActionTxnDatastoreModel>();
			serverDataStoreModel.datastore = SYNC_STORE.ACTION_TXN;
			serverDataStoreModel.data = new ActionTxnDatastoreModel();

			serverDataStoreModel.data.uuid = item.uuid;
			serverDataStoreModel.data.sync_pending = 1
			serverDataStoreModel.data.admission_uuid = item.admission_uuid;
			serverDataStoreModel.data.schedule_uuid = item.schedule_uuid;
			console.log('item.schedule_uuid;', item.schedule_uuid);
			serverDataStoreModel.data.conf_type_code = item.conf_type_code;
			serverDataStoreModel.data.txn_data = item.txn_data;
			console.log('item.txn_data;', item.txn_data);
			serverDataStoreModel.data.scheduled_time = item.scheduled_time;
			serverDataStoreModel.data.txn_state = item.txn_state;
			serverDataStoreModel.data.runtime_config_data = item.runtime_config_data;
			serverDataStoreModel.data.client_updated_at = new Date().toISOString();
			console.log('created data', serverDataStoreModel.data);
			this.ServerDataStoreDataModelArray.push(serverDataStoreModel);
		});
		console.log('his.ServerDataStoreDataModelArray', this.ServerDataStoreDataModelArray);
		this.savetoUserAuth();
		// check data save entries added in action trn table 
		this.gettrnlistdata();

	}

	// selected done and discard row change background color
	itemSelected(item) {
		item.selected = true;

	}
	gettrnlistdata() {
		setTimeout(() => {
			console.log(this.actionService.getActionTxnList());
		}, 300);

	}
	showDialog() {
		this.createDoctorModalView(ActionFabComponent, false).then((dialogResult: string) => {
			if (dialogResult) {
				switch (dialogResult) {
					case 'DoctorOrdersComponent':
						setTimeout(() => {
							this.openModal();
						});
						break;
					default:
						break;

				}
			}
		});
		//	this.dialogOpen = true;
	}

	// code block for opening component in modal.
	openModal() {
		console.log('doctors order  tapped');
		// this.dialogOpen = false;
		this.createDoctorModalView(DoctorOrdersComponent, true).then((dialogResult: ServerDataStoreDataModel<ScheduleDatastoreModel>[]) => {
			console.log('dialogResult', dialogResult);
			if (dialogResult) {
				this.ServerDataStoreDataModelArray = dialogResult;
				if (this.ServerDataStoreDataModelArray.length > 0) {
					setTimeout(() => {
						this.savetoUserAuth();
					});
				}
			}
		});

	}//end of fucntion

	private createDoctorModalView(Component, isFullScreen): Promise<any> {
		let options: ModalDialogOptions = {
			context: { promptMsg: "This is the prompt message!" },
			fullscreen: isFullScreen,
			viewContainerRef: this.viewContainerRef
		};
		return this.modalService.showModal(Component, options);
	}

	public activeList() {
		this.completeorpending = "Active Actions";
		this.iscompleted = false;
		this.viewexpand = false;
		this.saveViewOpen = false;
		this.buttonCompleted = false;
		this.getAllFlag = false;
	}
	public compilitedList() {
		this.completeorpending = "All Actions";
		this.iscompleted = true;
		this.viewexpand = true;
		this.saveViewOpen = false;
		this.buttonClicked = false;
		this.getAllFlag = true;		
	}

	// code block for closing opened dialog
	closeDialog() {
		this.dialogOpen = false;
	}// end 

	public getDoctorsOrders() {
		// console.log('getDoctors Orders');
		this.actionService.getDoctorsList('getdoctororders', this.passdataservice.getAdmissionID()).then(
			(val) => {
				// console.log('doctor order received', this.tempList);
				val.forEach(item => {
					console.log('getdoctororders item', item);
					let actionListItem = new DataActionItem();
					actionListItem = item;
					actionListItem.conf_type_code = ConfigCodeType.DOCTOR_ORDERS;
					try {
						this.tempList.push(actionListItem);
					} catch (e) {
						console.log(e.error);
					}
				});
				this.getCount();
			},
			(error) => {
				console.log("getActinData error:", error);
			}
		);

	} // end of fucntion
	pushDoctorOredrs(doctorsOrders: ServerDataStoreDataModel<IDatastoreModel>) {
		let actionListItem = new DataActionItem();
		Object.assign(actionListItem, doctorsOrders.data);
		actionListItem.conf_type_code = ConfigCodeType.DOCTOR_ORDERS;
		console.log('pushDoctorOredrs executed actionListItem', actionListItem);
		const item = this.tempList.filter(data => data.uuid == actionListItem.uuid)[0];
		//  if record found in list  
		if (item) {
			const index = this.tempList.indexOf(item);
			this.tempList[index] = item;
		} else {
			this.tempList.push(actionListItem);
		}
		this.getCount();
	} // end of code block.

}
