import { Component, OnDestroy, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { ModalDialogOptions, ModalDialogService } from 'nativescript-angular/modal-dialog';
import { RouterExtensions } from 'nativescript-angular/router';
import { SnackBar } from 'nativescript-snackbar';
import { ListViewEventData, ListViewItemSnapMode, ListViewLinearLayout } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Subscription } from 'rxjs/internal/Subscription';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { Page } from 'ui/page';
import { ACTION_STATUS, ConfigCodeType, SERVER_WORKER_MSG_TYPE, SYNC_STORE, MonitorType, ActionStatus, APP_MODE } from '~/app/app-constants';
import { ActionStatusHelper } from '~/app/helpers/action-status-helper';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ActionDataStoreModel } from '~/app/models/db/action-datastore';
import { ActionTxnDatastoreModel } from '~/app/models/db/action-txn-model';
import { IDatastoreModel } from '~/app/models/db/idatastore-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { BloodPressureValueModel, DataActionItem, GetJsonModel } from '~/app/models/ui/action-model';
import { ActionDataDBRequest, ActionListViewModel, ActionTxnDBModel, ActionProcess } from '~/app/models/ui/action-models';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { ActionService } from '~/app/services/action/action.service';
import { ChartService } from '~/app/services/chart/chart.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';
import { IDeviceAuthResult } from '../../idevice-auth-result';
import { ActionFabComponent } from '../action-fab/action-fab.component';
import { DoctorOrdersComponent } from '../doctor-orders/doctor-orders.component';
import { TimeConversion } from '~/app/helpers/time-conversion-helper';
import * as trace from 'trace';
import { TraceCustomCategory } from '~/app/helpers/trace-helper';
import * as appSettings from "tns-core-modules/application-settings";
import { AppNotificationService } from '~/app/services/app-notification-service';
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

export class ActionComponent implements OnInit, OnDestroy, IDeviceAuthResult {
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
	uiList = new ObservableArray<DataActionItem>();

	// >> details form field
	// actionForm: FormGroup;
	formData: ActionTxnDBModel;
	actionformData: ActionTxnDBModel;
	actiondata: ActionDataDBRequest;
	actionDbData: ActionDataDBRequest;
	actionDbArray: ActionProcess[] = [];
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


	// action notification handler

	actioncreationSubscription: Subscription;
	schedulecreationSubscription: Subscription;
	actionTxnDataReceivedSubject: Subscription;
	scheduleDatastoreModel: ScheduleDatastoreModel;
	actionListNottificationDataItem = new DataActionItem();
	constructor(public page: Page,
		private actionService: ActionService,
		public workerService: WorkerService,
		private modalService: ModalDialogService,
		private passdataservice: PassDataService,
		private workerservice: WorkerService,
		private viewContainerRef: ViewContainerRef,
		private chartService: ChartService,
		private appNotificationService:AppNotificationService,
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
		this.actionDbData = new ActionDataDBRequest();
		this.actiondata = new ActionDataDBRequest();
		this.completeorpending = "Active Actions";

		// action notification handler
		// this.schedulecreationSubscription = this.workerservice.scheduleDataReceivedSubject.subscribe((value) => {
		// 	this.scheduleDatastoreModel = value;
		// });
		this.actioncreationSubscription = this.workerservice.actionDataReceivedSubject.subscribe((value) => {
			this.handelActionNotification(value);
		});
		this.actionTxnDataReceivedSubject = this.workerservice.actionTxnDataReceivedSubject.subscribe((value) => {
			this.handelActionTransaction(value);
		});
		// subscription for adding newly  created doctors orders in action list.
		this.doctorOrderSubscription = this.workerService.doctorOrderSubject.subscribe((value) => {
			this.handelDoctorOrderNotification(value);
		});
		this.passdataservice.backalert = false;
	}// end of ng init.


	// >> expand row code start
	templateSelector(item: any, index: number, items: any): string {
		return item.expanded ? "expanded" : "default";

	}
	tapDataitem(item) {
		console.log('tap item data ', item)
		if (item) {
			this.actiondata = new ActionDataDBRequest();
		}
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
		if (this.conf_type_code === ConfigCodeType.INTAKE) {
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(0, false, ListViewItemSnapMode.Start);

			// console.log("Clicked select intake", this.intakeIndex);
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
		if (this.conf_type_code === ConfigCodeType.MONITOR) {
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(this.monitorIndex, false, ListViewItemSnapMode.Start);

			// console.log("Clicked select monitor", this.monitorIndex);
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
		// console.log('this.conf_type_code button', this.conf_type_code);
		if (this.conf_type_code === ConfigCodeType.MEDICINE) {
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(this.medicineIndex, false, ListViewItemSnapMode.Start);

			// console.log("Clicked select medicine", this.medicineIndex);
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
		if (this.conf_type_code === ConfigCodeType.OUTPUT) {
			// console.log('this.conf_type_code button output', this.conf_type_code);
			const listView = this.listViewComponent.listView;
			listView.scrollToIndex(this.outputIndex, false, ListViewItemSnapMode.Start);
			// console.log("Clicked select output", this.outputIndex);
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

		const medicine = this.uiList.filter(a => a.conf_type_code === ConfigCodeType.MEDICINE);
		const medicineCount = medicine.length;

		const monitor = this.uiList.filter(a => a.conf_type_code === ConfigCodeType.MONITOR);
		const monitorCount = monitor.length;
		const intake = this.uiList.filter(a => a.conf_type_code === ConfigCodeType.INTAKE);
		const intakeCount = intake.length;
		const output = this.uiList.filter(a => a.conf_type_code === ConfigCodeType.OUTPUT);
		const outputCount = output.length;
		const DOrder = this.uiList.filter(a => a.conf_type_code === ConfigCodeType.DOCTOR_ORDERS);
		const DOrderCount = DOrder.length;
		this.doctorOrderIndex = 0;
		this.intakeIndex = DOrderCount + 1;
		this.medicineIndex = DOrderCount + intakeCount + 1;
		this.monitorIndex = DOrderCount + intakeCount + medicineCount + 1;
		this.outputIndex = DOrderCount + intakeCount + medicineCount + monitorCount + 1;

	}
	// << Calculate Grouping index value

	// clean up
	ngOnDestroy(): void {
		//Called once, before the instance is destroyed.

		if (this.doctorOrderSubscription) { this.doctorOrderSubscription.unsubscribe(); }
		if (this.schedulecreationSubscription) { this.schedulecreationSubscription.unsubscribe(); }
		if (this.actioncreationSubscription) { this.actioncreationSubscription.unsubscribe(); }
		if (this.actionTxnDataReceivedSubject) { this.actionTxnDataReceivedSubject.unsubscribe(); }

	}

	public getActionData() {
		this.actionListItem = new ObservableArray<ActionListViewModel>();
		this.actionService.getallActionActiveList(this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					// console.log("get action item ", item);
					let currentDateTime = new Date();
					trace.write(`Processing action item for display. Date: ${currentDateTime}, Db Data Item: ${item}`, TraceCustomCategory.SCHEDULE, trace.messageType.info);
					console.log('Processing action  Db Data Item: ', item);
					if ((item.schedule_time === null) && (item.start_date > TimeConversion.getServerShortTimeFormat(currentDateTime) ||
						item.end_date < TimeConversion.getServerShortTimeFormat(currentDateTime))) {
						console.log('skipping item ', item);
						return; // Skipping this because if schedule is expired then need to remove sticky action
					}
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
		this.uiList = new ObservableArray<DataActionItem>();
		this.activeAction = new ObservableArray<DataActionItem>();
		this.allAction = new ObservableArray<DataActionItem>();

		this.actionListItem.forEach(item => {
			// console.log('txn item', item);
			const actionListDataItem = new DataActionItem();
			const gettxn_data = new GetJsonModel();
			actionListDataItem.value = new BloodPressureValueModel();
			actionListDataItem.admission_uuid = item.dbmodel.admission_uuid;
			actionListDataItem.schedule_uuid = item.dbmodel.schedule_uuid;
			actionListDataItem.uuid = item.dbmodel.action_uuid;
			actionListDataItem.conf_type_code = item.dbmodel.conf_type_code;
			this.conf_type_code = actionListDataItem.conf_type_code;
			// console.log('this.conf_type_code', this.conf_type_code);
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
			const conf = JSON.parse(item.dbmodel.conf);
			// console.log('conf ', conf);
			actionListDataItem.name = conf.name;
			// console.log('conf name', conf.name);
			actionListDataItem.desc = conf.desc;

			if (actionListDataItem.name === 'Blood Pressure') {
				if (gettxn_data.value != null) {
					const jsonvalue = new BloodPressureValueModel();
					Object.assign(jsonvalue, JSON.parse(gettxn_data.value));
					actionListDataItem.value.systolic = jsonvalue.systolic;
					actionListDataItem.value.diastolic = jsonvalue.diastolic;
				}
			} else {
				actionListDataItem.txn_data.value = gettxn_data.value;
			}

			//  condition for type set action list active and all
			if (item.dbmodel.action_txn_uuid != null) {
				// set type  Completed action 2
				actionListDataItem.client_updated_at = item.dbmodel.client_updated_at;
				actionListDataItem.type = 2;
				this.allAction.push(actionListDataItem);
			} else {
				// set type Active action1 
				actionListDataItem.type = 1;
				actionListDataItem.actionStatus = ActionStatusHelper.getActionStatus(Dbdate);
				this.actionStatus = actionListDataItem.actionStatus;

				this.allAction.push(actionListDataItem);
			}

		});
		// for creating active items list
		this.allAction.forEach((item) => {
			if (item.actionStatus == ACTION_STATUS.ACTIVE_NORMAL || item.actionStatus == ACTION_STATUS.ACTIVE_DELAYED || item.actionStatus == ACTION_STATUS.ACTIVE_NEEDS_ATTENTION) {
				// caculating action execution state
				const scheduleTime = this.calculateActiveActionTime(item.scheduled_time);
				if (scheduleTime) {
					this.activeAction.push(item);
				}
			}
		});
		//  updating to ui list		
		// get doctor orders.
		this.getDoctorsOrders();
		// this.toggoleActionList();

	}
	toggoleActionList() {
		if (this.getAllFlag == true) {
			// all
			console.log('all action list', this.allAction.length);
			this.uiList = this.allAction;
		} else {
			//  active
			this.uiList = this.activeAction;
			console.log('active action list', this.activeAction.length);
		}
		this.getCount();
	}
	// action txn nottification code
	handelActionTransaction(actionTxnDatastoreModel: ActionTxnDatastoreModel) {
		console.log('action transaction notification received', actionTxnDatastoreModel);
		// checking is schedule is created for particular patient
		if (actionTxnDatastoreModel.admission_uuid == this.passdataservice.getAdmissionID()) {

			let actionitem = this.allAction.filter(data => data.schedule_uuid === actionTxnDatastoreModel.schedule_uuid && data.scheduled_time === actionTxnDatastoreModel.scheduled_time)[0] || null;
			let stickyActionitem = this.allAction.filter(data => data.schedule_uuid === actionTxnDatastoreModel.schedule_uuid && data.scheduled_time === null)[0] || null;
			if (stickyActionitem != null) {
				//stickyActionitem.
				this.resetActionItem(stickyActionitem);
			} else {
				trace.write(`unalbel to find sticky action item for schedule: ${actionTxnDatastoreModel.schedule_uuid}`, TraceCustomCategory.SCHEDULE, trace.messageType.error);
			}
			// check if item is already exist			
			if (actionitem && actionitem != null) {
				const gettxn_data = new GetJsonModel();
				actionitem.txn_state = actionTxnDatastoreModel.txn_state;
				actionitem.type = 2;
				actionitem.actionStatus = null;
				Object.assign(gettxn_data, JSON.parse(actionTxnDatastoreModel.txn_data));
				actionitem.txn_data = gettxn_data;
				actionitem.txn_data.comment = gettxn_data.comment;
				actionitem.client_updated_at = actionTxnDatastoreModel.client_updated_at;
				// updated by get fname and lname by updated_on id
				this.actionService.getUserByUserid(actionTxnDatastoreModel.updated_by).then(
					(val) => {
						val.forEach(item => {
							actionitem.fname = item.fname;
							actionitem.lname = item.lname;
						});
					});
				if (actionitem.name === MonitorType.BLOOD_PRESSURE) {
					if (gettxn_data.value != null) {
						const jsonvalue = new BloodPressureValueModel();
						Object.assign(jsonvalue, JSON.parse(gettxn_data.value));
						actionitem.value.systolic = jsonvalue.systolic;
						actionitem.value.diastolic = jsonvalue.diastolic;
					}
				} else {
					actionitem.txn_data.value = gettxn_data.value;
				}
				let x = this.activeAction.filter(data => data.schedule_uuid === actionTxnDatastoreModel.schedule_uuid && data.scheduled_time === actionTxnDatastoreModel.scheduled_time)[0] || null;
				if (x && x != null) {
					const itemindex = this.activeAction.indexOf(x);
					if (itemindex >= 0) {
						this.activeAction.splice(itemindex, 1);
					}
				}
				this.toggoleActionList();
			}

		}
	}
	async handelActionNotification(actionDataStoreModel: ActionDataStoreModel) {
		console.log('action notification received', actionDataStoreModel)
		// if schedule added for specific patient.
		console.log('this.passdataservice.getAdmissionID()', this.passdataservice.getAdmissionID());
		if (actionDataStoreModel.admission_uuid == this.passdataservice.getAdmissionID()) {
			// check for  item exist
			const scheduleConfData = await this.actionService.getScheduleDetails('getScheduleData', actionDataStoreModel.schedule_uuid);
			let actionitem = this.allAction.filter(data => data.uuid === actionDataStoreModel.uuid)[0] || null;
			console.log('scheduleConfData received', scheduleConfData)
			if (actionitem && actionitem != null) {
				// item found				
				//	actionitem = <DataActionItem><any>actionDataStoreModel;
				actionitem.admission_uuid = actionDataStoreModel.admission_uuid;
				actionitem.client_updated_at = actionDataStoreModel.client_updated_at;
				actionitem.conf_type_code = actionDataStoreModel.conf_type_code;
				actionitem.is_deleted = actionDataStoreModel.is_deleted;
				actionitem.schedule_uuid = actionDataStoreModel.schedule_uuid;
				actionitem.scheduled_time = actionDataStoreModel.scheduled_time;
				actionitem.uuid = actionDataStoreModel.uuid
				// // for handeling tranasction mode			
				// fetching schedule name and its description
				if (scheduleConfData.length > 0) {
					const conf = JSON.parse(scheduleConfData[0].conf);
					actionitem.name = conf.name;
					actionitem.desc = conf.desc;
				} else {
					// disscuss with mandar
				}
				// assign empty value for avoiding can not read of deficned.
				if (actionitem.name === 'Blood Pressure') {
					actionitem.txn_data = new GetJsonModel();
					actionitem.value = new BloodPressureValueModel();
					actionitem.value.systolic = null
					actionitem.value.diastolic = null;
					actionitem.txn_data.comment = null;
				} else {
					actionitem.txn_data = new GetJsonModel();
					actionitem.txn_data.value = null;
					actionitem.txn_data.comment = null;
				}
			} else {
				// item not found 			
				actionitem = new DataActionItem();
				//actionitem = <DataActionItem><unknown>actionDataStoreModel;
				actionitem.admission_uuid = actionDataStoreModel.admission_uuid;
				actionitem.client_updated_at = actionDataStoreModel.client_updated_at;
				actionitem.conf_type_code = actionDataStoreModel.conf_type_code;
				actionitem.is_deleted = actionDataStoreModel.is_deleted;
				actionitem.schedule_uuid = actionDataStoreModel.schedule_uuid;
				actionitem.scheduled_time = actionDataStoreModel.scheduled_time;
				actionitem.uuid = actionDataStoreModel.uuid
				// assign empty value for avoiding can not read of deficned.
				if (actionitem.name === MonitorType.BLOOD_PRESSURE) {
					actionitem.txn_data = new GetJsonModel();
					actionitem.value = new BloodPressureValueModel();
					actionitem.value.systolic = null
					actionitem.value.diastolic = null;
					actionitem.txn_data.comment = null;
				} else {
					actionitem.txn_data = new GetJsonModel();
					actionitem.txn_data.value = null;
					actionitem.txn_data.comment = null;
				}
				// fetching schedule name and its description
				if (scheduleConfData.length > 0) {
					const conf = JSON.parse(scheduleConfData[0].conf);
					actionitem.name = conf.name;
					actionitem.desc = conf.desc;
				} else {
					// disscuss with mandar
				}
				const recivedDateFormDB = new Date(actionitem.scheduled_time);
				const recivedDateDb = recivedDateFormDB.getMinutes();
				recivedDateFormDB.setMinutes(recivedDateDb);
				const reciveTimeDb = recivedDateFormDB.toLocaleString();
				const Dbdate = new Date(reciveTimeDb);
				actionitem.actionStatus = ActionStatusHelper.getActionStatus(Dbdate);
				this.allAction.push(actionitem);

			}
			// if action is deleted
			if (actionitem.is_deleted === 1) {
				console.log('in isdeleted fucntion', this.activeAction);
				const itemindex = this.activeAction.indexOf(actionitem);

				console.log('itemindex in actiive action list', itemindex);
				if (itemindex >= 0) {
					this.activeAction.splice(itemindex, 1);
				}
				// removing from all action list
				const index = this.allAction.indexOf(actionitem);
				if (index >= 0) {
					this.allAction.splice(index, 1);
				}
			} else {
				// based on action  status calculating its schedule time/state
				if (actionitem.actionStatus == ACTION_STATUS.ACTIVE_NORMAL || actionitem.actionStatus == ACTION_STATUS.ACTIVE_DELAYED || actionitem.actionStatus == ACTION_STATUS.ACTIVE_NEEDS_ATTENTION) {
					const scheduleTime = this.calculateActiveActionTime(actionitem.scheduled_time);
					if (scheduleTime) {
						this.activeAction.push(actionitem);
					}
				}
			}
			// updating to ui list.
			this.toggoleActionList();

		} else {
			console.log('action for different patinet');
		}


	}

	// >> on submit one bye one item data
	onSubmit(item) {
		//set action conf model
		this.formData = new ActionTxnDBModel();
		this.passdataservice.backalert = true;
		this.itemSelected(item);
		this.saveViewOpen = true;

		// >> check condition medicine data not add comment and value entries
		if (item.conf_type_code === ConfigCodeType.MEDICINE) {
			this.actionDbData.comment = item.txn_data.comment;
			this.actionDbData.value = null;

			this.confString1 = JSON.stringify(this.actionDbData);
		} else {
			if (item.name === MonitorType.BLOOD_PRESSURE) {
				// console.log('if  condition');
				const bloodPressureValueModel = new BloodPressureValueModel()
				bloodPressureValueModel.systolic = item.value.systolic;
				bloodPressureValueModel.diastolic = item.value.diastolic;
				this.actionDbData.value = JSON.stringify(bloodPressureValueModel);
				this.actionDbData.comment = item.txn_data.comment;

			} else {
				// console.log('else condition');
				this.actionDbData.value = item.txn_data.value;
				this.actionDbData.comment = item.txn_data.comment;
			}
			this.confString = JSON.stringify(this.actionDbData);
		}



		// set db model 
		this.formData.uuid = PlatformHelper.API.getRandomUUID();
		this.formData.schedule_uuid = item.schedule_uuid;

		// >> check condition medicine data in josn format push
		if (item.conf_type_code === ConfigCodeType.MEDICINE) {
			console.log(' this.confString1;', this.confString1);
			this.formData.txn_data = this.confString1;
		} else {
			this.formData.txn_data = this.confString;
			console.log(' this.confString', this.confString);
		}
		this.formData.conf_type_code = item.conf_type_code;
		this.formData.runtime_config_data = null;
		this.formData.scheduled_time = item.scheduled_time;
		this.formData.txn_state = 1;
		this.formData.status = 1;
		this.formData.admission_uuid = item.admission_uuid;

		// after done data push one by one ietm in array hold data

		const actionProcess = new ActionProcess();
		actionProcess.actionTxnData = this.formData;
		actionProcess.actionItem = item;
		console.log('add data', this.actionDbArray);
		this.actionDbArray.push(actionProcess);

	}
	// >> on discard one bye one item data
	async onDiscard(item) {
		const isConfrim=await this.appNotificationService.confirm('Do You Want To Discard ?');
		if(isConfrim){
		console.log('itrem discard clicked');
		//set action conf model

		this.passdataservice.backalert = true;
		this.itemSelected(item);
		this.saveViewOpen = true;
		this.formData = new ActionTxnDBModel();

		// >> check condition medicine data not add comment and value entries
		if (item.conf_type_code === ConfigCodeType.MEDICINE) {
			this.actionDbData.comment = null;
			this.actionDbData.value = null;
			this.confString1 = JSON.stringify(this.actionDbData);
		} else {
			this.actionDbData.comment = item.txn_data.comment;
			this.actionDbData.value = item.txn_data.value;
			this.confString = JSON.stringify(this.actionDbData);
		}

		// set db model 
		this.formData.uuid = PlatformHelper.API.getRandomUUID();
		this.formData.schedule_uuid = item.schedule_uuid;

		// >> check condition medicine data in josn format push
		if (item.conf_type_code === ConfigCodeType.MEDICINE) {
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
		const actionProcess = new ActionProcess();
		actionProcess.actionTxnData = this.formData;
		actionProcess.actionItem = item;
		console.log('add data', this.actionDbArray);
		this.actionDbArray.push(actionProcess);
	}

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
		console.log('this.actionDbArray', this.actionDbArray);
		this.actionDbArray.forEach(actionData => {
			console.log('save item', actionData);

			const txnDataStoreModel = this.generateDataStoreModel(actionData);

			switch (actionData.actionItem.conf_type_code) {
				case ConfigCodeType.OUTPUT:

					if (actionData.actionTxnData.scheduled_time == null) {
						//const txnDataStoreModel = this.generateDataStoreModel(actionData);
						// updating schedule execution time with current date

						const scheduleExecutionTime = TimeConversion.getServerShortTimeFormat(new Date());

						const actionStoreModel = new ServerDataStoreDataModel<ActionDataStoreModel>();
						actionStoreModel.data = new ActionDataStoreModel();
						actionStoreModel.data.scheduled_time = scheduleExecutionTime;
						actionStoreModel.data.admission_uuid = actionData.actionItem.admission_uuid;
						actionStoreModel.data.schedule_uuid = actionData.actionItem.schedule_uuid;
						actionStoreModel.data.conf_type_code = actionData.actionItem.conf_type_code;

						actionStoreModel.data.is_deleted = ActionStatus.ACTION_ACTIVE;
						actionStoreModel.data.sync_pending = 1;
						actionStoreModel.data.client_updated_at = new Date().toISOString();
						//	actionList.updated_by = actionData.actionTxnData.updated_by;
						actionStoreModel.data.uuid = PlatformHelper.API.getRandomUUID();
						actionStoreModel.datastore = SYNC_STORE.ACTION;
						this.ServerDataStoreDataModelArray.push(actionStoreModel);


						txnDataStoreModel.data.scheduled_time = scheduleExecutionTime;
						this.ServerDataStoreDataModelArray.push(txnDataStoreModel);
					}
					break;
				default:
					//const dataStoreModel = this.generateDataStoreModel(actionData)
					this.ServerDataStoreDataModelArray.push(txnDataStoreModel);
					break;
			}
		});

		const appMode = appSettings.getNumber("APP_MODE", APP_MODE.NONE);
		if (appMode == APP_MODE.USER_DEVICE) {
			console.log('appSettings.getNumber("USER_ID")', appSettings.getNumber("USER_ID"));
			this.onDeviceAuthSuccess(appSettings.getNumber("USER_ID"));
		} else {
			this.savetoUserAuth();
		}


		this.saveViewOpen = false;
		// check data save entries added in action trn table 
		this.gettrnlistdata();

	}//end of code block

	generateDataStoreModel(actionData) {
		const serverDataStoreModel = new ServerDataStoreDataModel<ActionTxnDatastoreModel>();
		serverDataStoreModel.datastore = SYNC_STORE.ACTION_TXN;
		serverDataStoreModel.data = new ActionTxnDatastoreModel();
		serverDataStoreModel.data.uuid = actionData.actionTxnData.uuid;
		serverDataStoreModel.data.sync_pending = 1
		serverDataStoreModel.data.admission_uuid = actionData.actionTxnData.admission_uuid;
		serverDataStoreModel.data.schedule_uuid = actionData.actionTxnData.schedule_uuid;
		serverDataStoreModel.data.conf_type_code = actionData.actionTxnData.conf_type_code;
		serverDataStoreModel.data.txn_data = actionData.actionTxnData.txn_data;
		serverDataStoreModel.data.scheduled_time = actionData.actionTxnData.scheduled_time;
		serverDataStoreModel.data.txn_state = actionData.actionTxnData.txn_state;
		serverDataStoreModel.data.runtime_config_data = actionData.actionTxnData.runtime_config_data;
		serverDataStoreModel.data.client_updated_at = new Date().toISOString();
		return serverDataStoreModel;

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
	}

	// code block for opening component in modal.
	openModal() {
		this.createDoctorModalView(DoctorOrdersComponent, true).then((dialogResult: ServerDataStoreDataModel<ScheduleDatastoreModel>[]) => {
			if (dialogResult) {
				this.ServerDataStoreDataModelArray = dialogResult;
				if (this.ServerDataStoreDataModelArray.length > 0) {
					setTimeout(() => {
						const appMode = appSettings.getNumber("APP_MODE", APP_MODE.NONE);
						if (appMode == APP_MODE.USER_DEVICE) {
							console.log('appSettings.getNumber("USER_ID")', appSettings.getNumber("USER_ID"));
							this.onDeviceAuthSuccess(appSettings.getNumber("USER_ID"));
						} else {
							this.savetoUserAuth();
						}

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
		this.toggoleActionList();
	}
	public compilitedList() {
		this.completeorpending = "All Actions";
		this.iscompleted = true;
		this.viewexpand = true;
		this.saveViewOpen = false;
		this.buttonClicked = false;
		this.getAllFlag = true;
		this.toggoleActionList();
	}



	public getDoctorsOrders() {
		this.actionService.getDoctorsList('getdoctororders', this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					let actionListItem = new DataActionItem();
					actionListItem = item;
					actionListItem.conf_type_code = ConfigCodeType.DOCTOR_ORDERS;
					try {
						this.allAction.push(actionListItem);
						this.activeAction.push(actionListItem);
					} catch (e) {
						console.log(e.error);
					}
				});
				this.toggoleActionList();
			},
			(error) => {
				console.log("getActinData error:", error);
				this.toggoleActionList();
			});


	} // end of fucntion
	handelDoctorOrderNotification(doctorsOrders: ServerDataStoreDataModel<IDatastoreModel>) {
		let actionListItem = new DataActionItem();
		Object.assign(actionListItem, doctorsOrders.data);
		actionListItem.conf_type_code = ConfigCodeType.DOCTOR_ORDERS;
		if (actionListItem.admission_uuid == this.passdataservice.getAdmissionID()) {
			let item = this.allAction.filter(data => data.uuid == actionListItem.uuid)[0] || null;
			//  if record found in list  
			if (item) {
				item = actionListItem;
			} else {
				this.allAction.push(actionListItem);
				this.activeAction.push(actionListItem);
			}
			this.getCount();
		}
	} // end of code block.

	calculateActiveActionTime(scheduled_time) {
		if (scheduled_time != null) {
			const recivedDateFormDB = new Date(scheduled_time);
			const recivedDateDb = recivedDateFormDB.getMinutes();
			recivedDateFormDB.setMinutes(recivedDateDb);
			const reciveTimeDb = recivedDateFormDB.toLocaleString();
			const Dbdate = new Date(reciveTimeDb);

			const tempStartTime = new Date();
			const after1Hours = tempStartTime.getMinutes() - 60;
			tempStartTime.setMinutes(after1Hours);
			const tempStart = tempStartTime.toLocaleString();
			const startTime = new Date(tempStart);

			const tempEndTime = new Date();
			const next12Hours = tempEndTime.getMinutes() + 720;
			tempEndTime.setMinutes(next12Hours);
			const tempEnd = tempEndTime.toLocaleString();
			const endTime = new Date(tempEnd);
			if (Dbdate >= startTime && Dbdate <= endTime) {
				return true;
			} else {
				false;
			}
		} else {
			// schedule time is null means this action is sticky so return true.
			return true
		}
	}


	resetActionItem(actionItem: DataActionItem) {
		actionItem.txn_data = new GetJsonModel();
		actionItem.expanded = false;
		actionItem.selected = false;

	}
}
