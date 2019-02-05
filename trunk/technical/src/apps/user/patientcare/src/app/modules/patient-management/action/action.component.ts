import { Component, OnInit, ViewChild } from '@angular/core';
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
import { SYNC_STORE, SERVER_WORKER_MSG_TYPE } from '~/app/app-constants';
import { ActionTxnDatastoreModel } from '~/app/models/db/action-txn-model';
import { WorkerService } from '~/app/services/worker.service';

// expand row 
declare var UIView, NSMutableArray, NSIndexPath;
// import { TextField } from "ui/text-field";

// SnackBar import 
export class DataActionItem {
	uuid: string;
	admission_uuid: string;
	conf_type_code: string;
	schedule_uuid: string;
	exec_time: String;
	name: string;
	desc: string;
	status: number;
}

@Component({
	moduleId: module.id,
	selector: 'action',
	templateUrl: './action.component.html',
	styleUrls: ['./action.component.css']
})

export class ActionComponent implements OnInit {

	public actionListItem = new ObservableArray<ActionListViewModel>();
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
	outputbuttonClicked: boolean = false;
	actionSubscription: Subscription;
	// >> search var declaration
	// public myItems: ObservableArray<DataItem> = new ObservableArray<DataItem>();
	// tempdata = new Array<DataItem>();
	tempdata = new Array<ActionTxnDBModel>();


	// >> grouping 
	public _funcGrouping: (item: DataActionItem) => DataActionItem;

	// >> exapnd row
	expanded: false;

	// >> finding grouping index then after click show in top
	intakeIndex: any;
	medicineIndex: any;
	monitorIndex: any;
	outputIndex: any;

	// >>  bottom snackbar msg
	private snackbar: SnackBar;

	actionListItems = new ActionListViewModel();;
	tempList = new Array<DataActionItem>();

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

	login = true;
	listaccount = true;
	removeAccount = false;

	_dataItemsaccount = new ObservableArray<ActionListViewModel>();

	constructor(public page: Page,
		private actionService: ActionService,
		public workerService: WorkerService,
		private passdataservice: PassDataService,
		private chartService: ChartService) {
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
		// for (let i = 1; i < 50; i++) {
		// 	let newName = { ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" };
		// 	this.data.push(newName);
		// }
		// alert('action');

		// subscription for create actions
		this.actionSubscription = this.passdataservice.createActionsSubject.subscribe((value) => {

		}); // end of subscriptions.
		this.actionDbData = new ActionDataDBRequest();
		this.actiondata = new ActionDataDBRequest();

		// this.getUserAccountData();
	}

	public listLoaded() {
		//return; 
		// console.log('list loaded');

		setTimeout(() => {
		}, 200);
	}

	public addMoreItemsFromSource(chunkSize: number) {
		let newItems = this.tempList.slice(this.dataItems.length, this.dataItems.length + chunkSize);
	}

	public onLoadMoreItemsRequested(args: LoadOnDemandListViewEventData) {
		// console.log('onLoadMoreItemsRequested');

		// const that = new WeakRef(this);
		const listView: RadListView = args.object;
		if (this.dataItems.length < this.tempList.length) {
			setTimeout(() => {
				this.addMoreItemsFromSource(20);
				listView.notifyLoadOnDemandFinished();
				//console.log('onLoadMoreItemsRequested', this.dataItems.length);
			}, 200);
		} else {
			args.returnValue = false;
			listView.notifyLoadOnDemandFinished(true);
			// console.log('onLoadMoreItemsRequested', 'load on demand finished');
		}
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

	// >> Grouping position change 

	// >> Grouping intake scroll to top position change 
	public selectIntake() {

		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(0, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = true;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
		console.log("Clicked select intake", this.intakeIndex);

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

		console.log("Clicked select monitor", this.monitorIndex);

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
		console.log("Clicked select medicine", this.medicineIndex);
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
		console.log("Clicked select output", this.outputIndex);
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
		// console.log("outputCount", outputCount);

		this.intakeIndex = 0;
		this.medicineIndex = intakeCount + 1;
		this.monitorIndex = intakeCount + medicineCount + 2;
		this.outputIndex = intakeCount + medicineCount + monitorCount + 3;


		// console.log("medicine index", this.medicineIndex);
		// console.log("monitor index", this.monitorIndex);
		// console.log("output index", this.outputIndex);

	}
	// << Calculate Grouping index value


	// >> discard patient in list 
	public onDeleteRow(event: ListViewEventData) {

		// console.log("Right swipe click");
		this._dataItems.splice(this._dataItems.indexOf(event.object.bindingContext), 1);
		// console.log(this._dataItems);
		// console.log(this._dataItems.length);
		// this.showAction();
	}
	// << discard patient in list 

	// >> snackbar mes show bottom 
	// showMessage(): void {
	// 	this.snackbar.simple("Have a snack(bar)!");
	// }
	public showAction(event: ListViewEventData) {
		const abcd = this._dataItems.splice(this._dataItems.indexOf(event.object.bindingContext), 1);
		console.log('delete element', abcd);
		const item = this._dataItems.getItem(event.index);
		// this._dataItems.push(this._dataItems.indexOf(abcd), 1);
		console.log(item);

		// console.log(this._dataItems.length);
		// this._dataItems.splice(this._dataItems.indexOf(event.object.bindingContext), 1);
		// console.log(this._dataItems);
		// console.log(this._dataItems.length);


		let options: SnackBarOptions = {
			actionText: "Undo",
			actionTextColor: "#FF8910", // Android only
			snackText: "Patient action have been discard!",
			hideDelay: 3500
		}
		this.snackbar.action(options).then(args => {
			if (args.command === "Action") {
				// this._dataItems.push(abcd);
				// this._dataItems.push(this._dataItems.indexOf(abcd), 1);
				// console.log(this._dataItems.length);
				// alert({
				//   title: "Well hello there!",
				//   message: "That Snackbar seems useful, right?",
				//   okButtonText: "Uhm, I guess..",
				//   cancelable: true
				// });
			}
		});
	}
	// << snackbar mes show bottom 



	// end of code block

	// clean up
	ngOnDestroy(): void {
		//Called once, before the instance is destroyed.
		if (this.actionSubscription) { this.actionSubscription.unsubscribe(); }

	}

	// >>get action list 
	public getActionData() {
		this.actionService.getActionList().then(
			(val) => {
				val.forEach(item => {
					this.actionListItems = new ActionListViewModel();
					this.actionListItems.dbmodel = item;
					this.actionListItem.push(this.actionListItems);

				});
				//	console.log("received action list", this.actionListItem);
				this.getListDataById();
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}

	// >> get action list by id
	public getListDataById() {
		// console.log('this.actionnListItem', this.actionListItem);
		// this.tempList = new Array<DataActionItem>();
		this.actionListItem.forEach(item => {
			const actionListDataItem = new DataActionItem();
			actionListDataItem.admission_uuid = item.dbmodel.admission_uuid;
			actionListDataItem.schedule_uuid = item.dbmodel.schedule_uuid;
			actionListDataItem.conf_type_code = item.dbmodel.conf_type_code;

			// const exectime = new Date(item.dbmodel.exec_time * 1000);
			// console.log('exectime', exectime);

			actionListDataItem.exec_time = item.dbmodel.exec_time;
			// console.log('actionListDataItem.exec_time', actionListDataItem.exec_time);


			// date to timestramp convert
			// const todaydate = new Date().valueOf() / 1E3 | 0;
			// console.log('timestramp', todaydate);

			// >> today date live time
			const thetodayDate = new Date();
			const todayhr = thetodayDate.getHours();
			const liveh = todayhr * 60;
			const todaym = thetodayDate.getMinutes();
			const totaltime = liveh + todaym;
			// console.log('totaltime', totaltime);
			// << today date

			// >> today date live time decress time 15 min 
			const thetodayDate15dec = new Date();
			const todayhr15dec = thetodayDate15dec.getHours();
			const liveh15dec = todayhr15dec * 60;
			const todaym15dec = thetodayDate15dec.getMinutes() - 15;
			const totaltime15dec = liveh15dec + todaym15dec;
			// console.log('time _today_15_dec', totaltime15dec);
			// << decress time 15 min 

			// >> today date live time increass time 15 min 
			const thetodayDate15 = new Date();
			const todayhr15 = thetodayDate15.getHours();
			const liveh15 = todayhr15 * 60;
			const todaym15 = thetodayDate15.getMinutes() + 15;
			const totaltime15 = liveh15 + todaym15;
			// console.log('time _today_15_inc', totaltime15);
			// << increass time 15 min 

			// >> Db Date timestramp convert in date 
			// const theDate = new Date(item.dbmodel.exec_time * 1000);
			const theDate = new Date(item.dbmodel.exec_time);
			const hr = theDate.getHours();
			const h = hr * 60;
			const m = theDate.getMinutes();
			const DBtotaltime = h + m;
			// console.log('DBtotaltime ', DBtotaltime);
			// << Db Date timestramp convert 


			if (totaltime15dec > DBtotaltime) {
				// console.log('red');
				actionListDataItem.status = 1;
			} else if (totaltime15 > DBtotaltime && DBtotaltime > totaltime15dec) {
				// console.log('yellow');
				actionListDataItem.status = 2;
			} else if (DBtotaltime > totaltime15) {
				// console.log('green');
				actionListDataItem.status = 3;
			}

			this.chartService.getChartByUUID(actionListDataItem.schedule_uuid).then(
				(val) => {
					val.forEach(item => {
						// console.log('val', val);
						const conf = JSON.parse(item.conf);
						actionListDataItem.name = conf.name;
						actionListDataItem.desc = conf.desc;
					});
				})

			this.tempList.push(actionListDataItem);
			// console.log('testItem array', this.tempList);
		});
		this.getCount();
	}


	// >> on submit one bye one item data
	onSubmit(item) {
		console.log(item);
		//set action conf model

		this.itemSelected(item);
		this.saveViewOpen = true;
		this.formData = new ActionTxnDBModel();

		// >> check condition medicine data not add comment and value entries
		if (item.conf_type_code === 'Medicine') {
			this.actionDbData.comment = null;
			this.actionDbData.value = null;
			this.confString1 = JSON.stringify(this.actionDbData);
			// console.log('confString', this.confString1);
		} else {
			this.actionDbData.comment = this.actiondata.comment;
			this.actionDbData.value = this.actiondata.value;
			this.confString = JSON.stringify(this.actionDbData);
			// console.log('confString', this.confString);
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
		this.formData.txn_date = new Date;
		this.formData.txn_state = null;
		this.formData.status = 1;

		// console.log('this.actionformData', this.formData);

		// after done data push one by one ietm in array hold data
		this.actionDbArray.push(this.formData);
		// console.log('this.actionDbArray', this.actionDbArray);

	}
	// >> on discard one bye one item data
	onDiscard(item) {
		console.log(item);
		//set action conf model

		this.itemSelected(item);
		this.saveViewOpen = true;
		this.formData = new ActionTxnDBModel();

		// >> check condition medicine data not add comment and value entries
		if (item.conf_type_code === 'Medicine') {
			this.actionDbData.comment = null;
			this.actionDbData.value = null;
			this.confString1 = JSON.stringify(this.actionDbData);
			// console.log('confString', this.confString1);
		} else {
			this.actionDbData.comment = this.actiondata.comment;
			this.actionDbData.value = this.actiondata.value;
			this.confString = JSON.stringify(this.actionDbData);
			// console.log('confString', this.confString);
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
		this.formData.txn_date = new Date;
		this.formData.txn_state = null;
		this.formData.status = 0;

		// console.log('this.actionformData', this.formData);

		// after done data push one by one ietm in array hold data
		this.actionDbArray.push(this.formData);
		// console.log('this.actionDbArray', this.actionDbArray);

	}
	// all action done and discard save in action-trn-table
	savemove() {
		this.login = !this.login;

	}
	save() {
		// array hold entries one by one save
		this.actionformData = new ActionTxnDBModel();
		// insert Action db model to sqlite db
		this.actionDbArray.forEach(item => {
			// this.actionformData.uuid = item.uuid;
			// this.actionformData.schedule_uuid = item.schedule_uuid;
			// this.actionformData.conf_type_code = item.conf_type_code;
			// this.actionformData.txn_data = item.txn_data;
			// this.actionformData.txn_date = item.txn_date;
			// this.actionformData.txn_state = item.txn_state;
			// this.actionformData.conf_type_code = item.conf_type_code;
			// this.actionformData.runtime_config_data = item.runtime_config_data;
			// this.actionformData.status = item.status

			// console.log('item.conf_type_code', item.conf_type_code);
			// console.log('this.actionformData.schedule_uuid', this.actionformData.conf_type_code);
			// console.log('actionformData', this.actionformData);


			const actionModel = new ServerDataProcessorMessageModel();
			const serverDataStoreModel = new ServerDataStoreDataModel<ActionTxnDatastoreModel>();
			serverDataStoreModel.datastore = SYNC_STORE.ACTION_TXN;
			serverDataStoreModel.data = new ActionTxnDatastoreModel();

			serverDataStoreModel.data.uuid = item.uuid;
			serverDataStoreModel.data.sync_pending = 1
			serverDataStoreModel.data.schedule_uuid = item.schedule_uuid;
			serverDataStoreModel.data.conf_type_code = item.conf_type_code;
			serverDataStoreModel.data.txn_data = item.txn_data;
			serverDataStoreModel.data.txn_date = item.txn_date;
			serverDataStoreModel.data.txn_state = Number(item.txn_state);
			serverDataStoreModel.data.runtime_config_data = item.runtime_config_data;

			console.log('created data', serverDataStoreModel.data)
			actionModel.data = [serverDataStoreModel];
			actionModel.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;
			this.workerService.ServerDataProcessorWorker.postMessage(actionModel);
			// save action done and discard in DB
			//	this.actionService.insertActionTxnItem(this.actionformData);
		});

		// check data save entries added in action trn table 
		this.gettrnlistdata();
	}

	// selected done and discard row change background color
	itemSelected(item) {
		item.selected = true;
	}
	gettrnlistdata() {
		setTimeout(() => {
			this.actionService.getActionTxnList();
		}, 300);

	}
}
