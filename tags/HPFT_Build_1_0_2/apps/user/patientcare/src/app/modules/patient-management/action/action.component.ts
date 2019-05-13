import { Component, OnDestroy, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { ModalDialogOptions, ModalDialogService } from 'nativescript-angular/modal-dialog';
import { RouterExtensions } from 'nativescript-angular/router';
import { ListViewEventData, ListViewItemSnapMode, ListViewLinearLayout } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Subscription } from 'rxjs/internal/Subscription';
import * as appSettings from "tns-core-modules/application-settings";
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { Page } from 'ui/page';
import { ActionStatus, ACTION_STATUS, APP_MODE, ConfigCodeType, MonitorType, SERVER_WORKER_MSG_TYPE, SYNC_STORE } from '~/app/app-constants';
import { ActionStatusHelper } from '~/app/helpers/action-status-helper';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { TimeConversion } from '~/app/helpers/time-conversion-helper';
import { ServerDataProcessorMessageModel } from '~/app/models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { ActionDataStoreModel } from '~/app/models/db/action-datastore';
import { ActionTxnDatastoreModel } from '~/app/models/db/action-txn-model';
import { IDatastoreModel } from '~/app/models/db/idatastore-model';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { ActionItemVMModel } from '~/app/models/ui/action-item-vm-model';
import { BloodPressureValueModel, DataActionItem, GetJsonModel } from '~/app/models/ui/action-model';
import { ActionListViewModel, ActionProcess, ActionSubmitDiscardModel, ActionTxnDBModel } from '~/app/models/ui/action-models';
import { Schedulardata } from '~/app/models/ui/chart-models';
import { ActionService } from '~/app/services/action/action.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { WorkerService } from '~/app/services/worker.service';
import { IDeviceAuthResult } from '../../idevice-auth-result';
import { ActionFabComponent } from '../action-fab/action-fab.component';
import { DoctorOrdersComponent } from '../doctor-orders/doctor-orders.component';
import { MedicineActionsComponent } from './medicine-actions/medicine-actions.component';
import { DoctorsOrdersDatastoreModel } from '~/app/models/db/doctors-orders-model';
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

	//objetcs
	scheduleDatastoreModel: ScheduleDatastoreModel;
	// blood pressure high and low value model	
	private layout: ListViewLinearLayout;

	actionItems = new ObservableArray<ActionItemVMModel>();
	displayModeAll = false;
	monitorType = MonitorType;


	//arrays
	tempdata = new Array<ActionTxnDBModel>();
	public _dataItems: ObservableArray<any>;
	actionDbArray: ActionProcess[] = [];
	tempActionTxnDataArray: ActionSubmitDiscardModel[] = [];
	monitorschedulardata: Schedulardata[] = []
	outputschedulardata: Schedulardata[] = []
	_dataItemsaccount = new ObservableArray<ActionListViewModel>();
	//for view 
	uiList = new ObservableArray<DataActionItem>();

	activeAction = new ObservableArray<DataActionItem>();
	allAction = new ObservableArray<DataActionItem>();
	ServerDataStoreDataModelArray: ServerDataStoreDataModel<any>[] = [];
	public actionListItem = new ObservableArray<ActionListViewModel>();

	// fucntions
	// for make group of items 
	public _funcGrouping: (item: ActionItemVMModel) => string;
	public activeAllFilter: (item: ActionItemVMModel) => boolean;

	// saveViewOpen = false;
	// subscriptions
	actioncreationSubscription: Subscription;
	schedulecreationSubscription: Subscription;
	actionTxnDataReceivedSubject: Subscription;
	doctorOrderSubscription: Subscription;

	// >> seleced bottom button change color
	monitorbuttonClicked = false;
	intakebuttonClicked = true;
	medicinebuttonClicked = false;
	doctorOrderButtonClicked = false;
	outputbuttonClicked = false;
	chartbuttonClicked = false;


	//constants
	conf_type_code_const = ConfigCodeType;
	action_Status = ACTION_STATUS;

	// >> exapnd row
	expanded = false;
	viewexpand = false;
	// >> finding grouping index then after click show in top
	intakeIndex: any;
	medicineIndex: any;
	monitorIndex: any;
	outputIndex: any;
	doctorOrderIndex: any;
	// >>  bottom snackbar msg
	confString;
	confString1;
	exectime;
	bloddname: any;
	// switch active and complited
	completeorpending: string;
	iscompleted: boolean;
	// filter buttton 
	buttonClicked = true;
	buttonCompleted = false;
	getAllFlag = false;
	// filter data complited UI readonly Mode
	editMode = false;
	actionStatus: any;
	conf_type_code: any;


	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	constructor(
		public page: Page,
		private actionService: ActionService,
		public workerService: WorkerService,
		private modalService: ModalDialogService,
		private passdataservice: PassDataService,
		private workerservice: WorkerService,
		private viewContainerRef: ViewContainerRef,
		private routerExtensions: RouterExtensions) {
		//  list item grouping based on config_type_code.
		this._funcGrouping = (item: ActionItemVMModel) => {
			return item.conf_type_code;
		};
		// object creation
		// this.actionDbData = new ActionDataDBRequest();
		// this.formData = new ActionTxnDBModel();


		this.activeAllFilter = (item: ActionItemVMModel) => {
			return item.isActionActive;
		}
	}

	ngOnInit() {
		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		//	this.getActionData();

		// this.actionDbData = new ActionDataDBRequest();
		// this.actiondata = new ActionDataDBRequest();
		this.completeorpending = "Active Actions";

		const listView = this.listViewComponent.listView;
		listView.filteringFunction = function (item: ActionItemVMModel) {
			return true;
		};

		// subject for handel action notifications.
		this.actioncreationSubscription = this.workerservice.actionDataReceivedSubject.subscribe((value) => {
			this.handelActionNotification(value);
		});
		// subject for action transaction notification.
		this.actionTxnDataReceivedSubject = this.workerservice.actionTxnDataReceivedSubject.subscribe((value) => {
			this.handelActionTransaction(value);
		});
		// subscription for adding newly  created doctors orders in action list.
		this.doctorOrderSubscription = this.workerService.doctorOrderDataReceivedSubject.subscribe((value) => {
			this.handelDoctorOrderNotification(value);
		});
		this.prepareData();
		this.activeList();

	}// end of ng init.

	async prepareData() {
		this.actionService.getallActionActiveList(this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					const actionItem = this.prepareActionItem(null, item, false)
					console.log('actionStatus', actionItem.actionStatus);
					console.log('hasTxnData', actionItem.hasTxnData);
					console.log('scheduled_time', actionItem.dbModel.scheduled_time);
					console.log('isItemActive', actionItem.isActionActive);
					console.log('	item.conf_type_code', actionItem.conf_type_code);

					this.actionItems.push(actionItem);
				});
				console.log('this.actionItems', this.actionItems.length);
			},
			(error) => {
				console.log("getActinData error:", error);
			});


		this.actionService.getDoctorsList('getdoctororders', this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					let actionItemVM = this.prepareActionItem(null, item, true);
					this.actionItems.push(actionItemVM);
				});
			},
			(error) => {
				console.log("getActinData error:", error);
			});


	};
	get dataItems(): ObservableArray<DataActionItem> {
		return this._dataItems;
	}
	prepareActionItem(actionItemVM: ActionItemVMModel, data: any, isDoctorOrderData: boolean) {
		let actionItemVMModel = (actionItemVM == null) ? new ActionItemVMModel() : actionItemVM;

		if (isDoctorOrderData) {
			actionItemVMModel.conf_type_code = this.conf_type_code_const.DOCTOR_ORDERS;
			actionItemVMModel.doctorOrderModel = data;
			actionItemVMModel.hasTxnData = false;
			actionItemVMModel.isActionActive = (actionItemVMModel.doctorOrderModel.status == 0) ? true : false;

		} else {
			actionItemVMModel.dbModel = data;
			actionItemVMModel.configData = JSON.parse(data.conf)
			actionItemVMModel.actionStatus = this.getActionStatus(actionItemVMModel.dbModel);
			actionItemVMModel.isActionActive = this.isActiveAction(actionItemVMModel);
			actionItemVMModel.txnData = this.prepareTransactionModel(actionItemVMModel);
			actionItemVMModel.conf_type_code = actionItemVMModel.dbModel.conf_type_code;

			if (actionItemVMModel.dbModel.action_txn_uuid != null) {
				actionItemVMModel.hasTxnData = true;
			} else {
				actionItemVMModel.hasTxnData = false;
			}
		}


		//	console.log('actionItemVMModel', actionItemVMModel);

		return actionItemVMModel;
	}

	prepareTransactionModel(actionItemVMModel: ActionItemVMModel) {

		let txn_data;
		if (actionItemVMModel.dbModel.txn_data == null) {
			switch (actionItemVMModel.dbModel.conf_type_code) {
				case ConfigCodeType.MONITOR:
					txn_data = new GetJsonModel();
					if (actionItemVMModel.configData.name == MonitorType.BLOOD_PRESSURE) {
						txn_data.value = new BloodPressureValueModel();
					}
					return txn_data;
				default:
					txn_data = new GetJsonModel();
					return txn_data;
			}
		} else {
			txn_data = JSON.parse(actionItemVMModel.dbModel.txn_data);
		}

		//	console.log('transaction data in return ', txn_data);
		return txn_data;
	}
	// >> expand row code start
	templateSelector(item: any, index: number, items: any): string {
		return item.expanded ? "expanded" : "default";
	}
	//expand list list view item
	onItemTap(event: ListViewEventData) {
		const listView = event.object,
			rowIndex = event.index,
			dataItem = event.view.bindingContext;

		dataItem.expanded = !dataItem.expanded;
		if (isIOS) {
			var indexPaths = NSMutableArray.new();
			indexPaths.addObject(NSIndexPath.indexPathForRowInSection(rowIndex, event.groupIndex));
			listView.ios.reloadItemsAtIndexPaths(indexPaths);
		}
		if (isAndroid) {
			listView.androidListView.getAdapter().notifyItemChanged(rowIndex);

		}
	}

	// TODO: restructure needed.
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
	// action txn nottification code
	handelActionTransaction(actionTxnDatastoreModel: ActionTxnDatastoreModel) {
		console.log('action transaction notification received', actionTxnDatastoreModel);
		// checking is schedule is created for particular patient

		if (actionTxnDatastoreModel.admission_uuid != this.passdataservice.getAdmissionID()) {
			return;
		}
		let filteredItems = this.actionItems.filter(data =>
			data.conf_type_code != this.conf_type_code_const.DOCTOR_ORDERS &&
			data.dbModel.scheduled_time === actionTxnDatastoreModel.scheduled_time &&
			data.dbModel.schedule_uuid === actionTxnDatastoreModel.schedule_uuid);

		let stickyActionitem = this.actionItems.filter(data =>
			data.conf_type_code != this.conf_type_code_const.DOCTOR_ORDERS &&
			data.dbModel.schedule_uuid === actionTxnDatastoreModel.schedule_uuid &&
			data.dbModel.scheduled_time === null)[0] || null;
		if (stickyActionitem != null) {
			//stickyActionitem.
			this.resetActionItem(stickyActionitem);
		}

		if (filteredItems.length > 0) {
			let itemndex = this.actionItems.indexOf(filteredItems[0]);
			let actionItemVM = filteredItems[0];
			this.actionService.getallActionById(this.passdataservice.getAdmissionID(), actionItemVM.dbModel.action_uuid).then(
				(val) => {
					val.forEach(item => {
						this.prepareActionItem(actionItemVM, item, false);
						this.refreshListView();
					});
				});

		}
	}


	// >> on submit one bye one item data
	onSubmit(item: ActionItemVMModel) {
		this.itemSelected(item);
		const txnItem = this.prepareActionTxnDataStore(item, 1, 1);
		this.tempActionTxnDataArray.push(txnItem)
	}

	// >> on discard one bye one item data
	onDiscard(item) {
		this.itemSelected(item);
		const txnItem = this.prepareActionTxnDataStore(item, 0, 2);
		this.tempActionTxnDataArray.push(txnItem);
	}

	// all action done and discard save in action-trn-table
	savetoUserAuth() {
		this.passdataservice.authResultReuested = this;
		this.routerExtensions.navigate(['patientmgnt', 'user-auth'], { clearHistory: false });
	}

	save() {

		this.tempActionTxnDataArray.forEach((actionItemVM) => {
			const txnDataStoreModel = this.generateDataStoreModel(actionItemVM);
			switch (actionItemVM.actionItem.dbModel.conf_type_code) {
				case ConfigCodeType.OUTPUT:
					const actionDataStoreModel = this.getActionDatastoreModel(actionItemVM);
					this.ServerDataStoreDataModelArray.push(actionDataStoreModel);
					txnDataStoreModel.data.scheduled_time = actionDataStoreModel.data.scheduled_time;
					this.ServerDataStoreDataModelArray.push(txnDataStoreModel);
					break;
				default:
					this.ServerDataStoreDataModelArray.push(txnDataStoreModel);
					break;
			}
			this.unSelectItem(actionItemVM.actionItem);
		});
		this.tempActionTxnDataArray = [];
		console.log('this.ServerDataStoreDataModelArray', this.ServerDataStoreDataModelArray);
		const appMode = appSettings.getNumber("APP_MODE", APP_MODE.NONE);
		if (appMode == APP_MODE.USER_DEVICE) {
			console.log('appSettings.getNumber("USER_ID")', appSettings.getNumber("USER_ID"));
			this.onDeviceAuthSuccess(appSettings.getNumber("USER_ID"));
		} else {
			this.savetoUserAuth();
		}


	}//end of code block

	generateDataStoreModel(actionSubmitDiscardModel: ActionSubmitDiscardModel) {
		const serverDataStoreModel = new ServerDataStoreDataModel<ActionTxnDatastoreModel>();
		serverDataStoreModel.datastore = SYNC_STORE.ACTION_TXN;
		serverDataStoreModel.data = new ActionTxnDatastoreModel();
		serverDataStoreModel.data.uuid = actionSubmitDiscardModel.actionTxnData.uuid;
		serverDataStoreModel.data.sync_pending = 1
		serverDataStoreModel.data.admission_uuid = actionSubmitDiscardModel.actionTxnData.admission_uuid;
		serverDataStoreModel.data.schedule_uuid = actionSubmitDiscardModel.actionTxnData.schedule_uuid;
		serverDataStoreModel.data.conf_type_code = actionSubmitDiscardModel.actionTxnData.conf_type_code;
		serverDataStoreModel.data.txn_data = actionSubmitDiscardModel.actionTxnData.txn_data;
		serverDataStoreModel.data.scheduled_time = actionSubmitDiscardModel.actionTxnData.scheduled_time;
		serverDataStoreModel.data.txn_state = actionSubmitDiscardModel.actionTxnData.txn_state;
		serverDataStoreModel.data.runtime_config_data = actionSubmitDiscardModel.actionTxnData.runtime_config_data;
		serverDataStoreModel.data.client_updated_at = new Date().toISOString();
		return serverDataStoreModel;

	}


	// gettrnlistdata() {
	// 	setTimeout(() => {
	// 		console.log(this.actionService.getActionTxnList());
	// 	}, 300);

	// }
	showDialog() {
		this.createDoctorModalView(ActionFabComponent, false).then((dialogResult: string) => {
			if (dialogResult) {
				switch (dialogResult) {
					case 'DoctorOrdersComponent':
						setTimeout(() => {
							this.openModal(DoctorOrdersComponent);
						});
						break;
					case 'MedicineActionsComponent':
						setTimeout(() => {
							this.openModal(MedicineActionsComponent);
						});
						break;
					default:
						break;

				}
			}
		});
	}

	// code block for opening component in modal.
	openModal(value: any) {
		this.createDoctorModalView(value, true).then((dialogResult: ServerDataStoreDataModel<ScheduleDatastoreModel>[]) => {
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
		this.buttonCompleted = false;
		this.buttonClicked = true;
		this.displayModeAll = false;
		const listView = this.listViewComponent.listView;
		listView.filteringFunction = this.activeAllFilter;
		this.completeorpending = "Active Actions";
	}
	public compilitedList() {
		this.buttonCompleted = true;
		this.buttonClicked = false;
		this.displayModeAll = true;
		const listView = this.listViewComponent.listView;
		listView.filteringFunction = function (item) { return true };
		this.completeorpending = "All Actions";
	}
	handelDoctorOrderNotification(doctorsOrders: DoctorsOrdersDatastoreModel) {

		let dataActionItem = new DataActionItem();
		Object.assign(dataActionItem, doctorsOrders);

		if (dataActionItem.admission_uuid != this.passdataservice.getAdmissionID()) {
			return;
		}

		this.actionService.getDoctorOrderByID('getdoctororderbyid', dataActionItem.uuid).then(
			(val) => {
				val.forEach(item => {
					let actionItemVM = this.prepareActionItem(null, item, true);
					this.actionItems.push(actionItemVM);
				});
			},
			(error) => {
				console.log("getActinData error:", error);
			});

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

	// will get executed once sticky action is performed
	resetActionItem(item) {
		item.txnData = new GetJsonModel();
		item.expanded = false;
		item.selected = false;

	}
	// auth user functinal
	// if authorization is successfull
	onDeviceAuthSuccess(userid: number): void {
		this.tempActionTxnDataArray = [];
		this.ServerDataStoreDataModelArray.forEach(element => {
			element.data.updated_by = userid;
		});
		// posting to worker for  submit into database.
		const initModel = new ServerDataProcessorMessageModel();
		initModel.data = this.ServerDataStoreDataModelArray
		initModel.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;
		this.workerservice.ServerDataProcessorWorker.postMessage(initModel);
		this.ServerDataStoreDataModelArray = [];
	}
	// if autorizations is fail
	onDeviceAuthError(error: any): void {
		throw new Error("Method not implemented.");
	}
	// on discarded user auth
	onSubmitDiscarded(): void {
		throw new Error("Method not implemented.");
	}

	getActionStatus(item: DataActionItem) {
		const recivedDateFormDB = new Date(item.scheduled_time);
		const recivedDateDb = recivedDateFormDB.getMinutes();
		recivedDateFormDB.setMinutes(recivedDateDb);
		const reciveTimeDb = recivedDateFormDB.toLocaleString();
		const Dbdate = new Date(reciveTimeDb);
		return ActionStatusHelper.getActionStatus(Dbdate);
	}


	//TODO: handle transaction item
	isActiveAction(item: ActionItemVMModel) {
		// for completed actions.
		if (item.dbModel.action_txn_uuid !== null) {
			return false;
		}
		if (item.actionStatus == ACTION_STATUS.ACTIVE_NORMAL ||
			item.actionStatus == ACTION_STATUS.ACTIVE_DELAYED ||
			item.actionStatus == ACTION_STATUS.ACTIVE_NEEDS_ATTENTION) {
			// caculating action execution state
			const scheduleTime = this.calculateActiveActionTime(item.dbModel.scheduled_time);
			if (scheduleTime) {
				return true;
			}
		}

		return false;
	}
	// displayAction(item: ActionItemVMModel): boolean {
	// 	//displayMode === 'All' item.isActionAction

	// 	return false;
	// }

	// getActionItemValue(item: ActionItemVMModel, filter: string): any {
	// 	if (item.dbModel.name !== 'Blood Pressure') {
	// 		return item.txnData.value;
	// 	} else {
	// 		switch (filter) {
	// 			case "systolic":
	// 				return JSON.parse(item.txnData.value).systolic;
	// 			case "diastolic":
	// 				return JSON.parse(item.txnData.value).diastolic;
	// 			default:
	// 				return "";
	// 		}

	// 	}
	// }

	getItemCSSClass(item: ActionItemVMModel): string {
		let cssClass: string;
		switch (item.actionStatus) {
			case ACTION_STATUS.ACTIVE_DELAYED:
				cssClass = 'active_delayed_missed';
				break;
			case ACTION_STATUS.ACTIVE_NEEDS_ATTENTION:
				cssClass = 'active_needs_attention';
				break;
			case ACTION_STATUS.ACTIVE_NORMAL:
				cssClass = 'active_normal';
				break;
			case ACTION_STATUS.MISSED:
				cssClass = 'active_missed';
				break
			case ACTION_STATUS.ACTIVE_FUTURE:
				cssClass = 'active_future';
				break

		}
		return cssClass;
		//	{active_delayed_missed: item.actionStatus === action_Status.ACTIVE_DELAYED,active_needs_attention: item.actionStatus === action_Status.ACTIVE_NEEDS_ATTENTION,active_normal: item.actionStatus === action_Status.ACTIVE_NORMAL, active_missed: item.actionStatus === action_Status.MISSED}
	}

	async handelActionNotification(actionDataStoreModel: ActionDataStoreModel) {
		console.log('action notification executed', actionDataStoreModel);
		if (actionDataStoreModel.admission_uuid != this.passdataservice.getAdmissionID()) {
			return;
		}

		switch (actionDataStoreModel.is_deleted) {
			case 1://actiond deleted due to schedule canceled  
				let filteredItems = this.actionItems.filter(data =>
					data.conf_type_code != this.conf_type_code_const.DOCTOR_ORDERS &&
					data.dbModel.action_uuid === actionDataStoreModel.uuid);

				if (filteredItems.length > 0) {
					let itemndex = this.actionItems.indexOf(filteredItems[0]);
					if (itemndex) {
						this.actionItems.splice(itemndex, 1);
					}
				}
				break;
			default:
				await this.actionService.getallActionById(this.passdataservice.getAdmissionID(), actionDataStoreModel.uuid).then(
					(val) => {
						val.forEach(item => {
							const actionItem = this.prepareActionItem(null, item, false);
							console.log('action item', actionItem);
							this.actionItems.push(actionItem);
						});
						console.log('action  list', this.actionItems);
					});
				break;
		}


	}
	// selected done and discard row change background color
	itemSelected(item) {
		item.selected = true;
	}
	unSelectItem(item) {
		item.selected = false;
	}
	//for new actions againesh sticky actions.
	getActionDatastoreModel(actionSubmitDiscardModel: ActionSubmitDiscardModel): ServerDataStoreDataModel<ActionDataStoreModel> {
		const scheduleExecutionTime = TimeConversion.getServerShortTimeFormat(new Date());
		const actionStoreModel = new ServerDataStoreDataModel<ActionDataStoreModel>();
		actionStoreModel.data = new ActionDataStoreModel();
		actionStoreModel.data.scheduled_time = scheduleExecutionTime;
		actionStoreModel.data.admission_uuid = actionSubmitDiscardModel.actionItem.dbModel.admission_uuid;
		actionStoreModel.data.schedule_uuid = actionSubmitDiscardModel.actionItem.dbModel.schedule_uuid;
		actionStoreModel.data.conf_type_code = actionSubmitDiscardModel.actionItem.dbModel.conf_type_code;
		actionStoreModel.data.is_deleted = ActionStatus.ACTION_ACTIVE;
		actionStoreModel.data.sync_pending = 1;
		actionStoreModel.data.client_updated_at = new Date().toISOString();
		actionStoreModel.data.uuid = PlatformHelper.API.getRandomUUID();
		actionStoreModel.datastore = SYNC_STORE.ACTION;
		return actionStoreModel;
	}
	prepareActionTxnDataStore(item: ActionItemVMModel, status: number, txn_state: number): ActionSubmitDiscardModel {
		const actionTxnDBModel = new ActionTxnDBModel();
		actionTxnDBModel.txn_data = JSON.stringify(item.txnData);
		actionTxnDBModel.status = status;
		actionTxnDBModel.scheduled_time = item.dbModel.scheduled_time;
		actionTxnDBModel.txn_state = txn_state;
		actionTxnDBModel.uuid = PlatformHelper.API.getRandomUUID();
		actionTxnDBModel.schedule_uuid = item.dbModel.schedule_uuid;
		actionTxnDBModel.runtime_config_data = null;
		actionTxnDBModel.conf_type_code = item.dbModel.conf_type_code;
		actionTxnDBModel.admission_uuid = item.dbModel.admission_uuid;
		const actionSubmitDiscardModel = new ActionSubmitDiscardModel();
		actionSubmitDiscardModel.actionTxnData = actionTxnDBModel;
		actionSubmitDiscardModel.actionItem = item;
		return actionSubmitDiscardModel;
	}

	refreshListView() {
		let filterfunc = this.listViewComponent.listView.filteringFunction;
		this.listViewComponent.listView.filteringFunction = undefined;
		this.listViewComponent.listView.filteringFunction = filterfunc;
	}

	// clean up
	ngOnDestroy(): void {
		//Called once, before the instance is destroyed.

		if (this.doctorOrderSubscription) { this.doctorOrderSubscription.unsubscribe(); }
		if (this.schedulecreationSubscription) { this.schedulecreationSubscription.unsubscribe(); }
		if (this.actioncreationSubscription) { this.actioncreationSubscription.unsubscribe(); }
		if (this.actionTxnDataReceivedSubject) { this.actionTxnDataReceivedSubject.unsubscribe(); }

	}
}
