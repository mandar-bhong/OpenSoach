import { Component, OnInit, ViewChild } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { View, isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { ListViewLinearLayout, ListViewEventData, RadListView, LoadOnDemandListViewEventData, ListViewItemSnapMode } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { layout } from "tns-core-modules/utils/utils";
import { EventData } from "tns-core-modules/data/observable";
import { Observable } from 'tns-core-modules/data/observable';

// expand row 
import * as utils from "utils/utils";
declare var UIView, NSMutableArray, NSIndexPath;
// import { TextField } from "ui/text-field";

// SnackBar import 
import { alert } from "tns-core-modules/ui/dialogs";
import { SnackBar, SnackBarOptions } from "nativescript-snackbar";
import { Page } from "ui/page";
import { Subscription } from 'rxjs/internal/Subscription';
import { PassDataService } from '~/app/services/pass-data-service';
import { ChartService } from '~/app/services/chart/chart.service';
import { ChartListViewModel, ConfigData, Schedulardata, MedChartModel, AftrnFreqInfo, MornFreqInfo, NightFreqInfo } from '~/app/models/ui/chart-models';
import { error } from 'tns-core-modules/trace/trace';
import { medicine, freuencyzero } from '~/app/common-constants';
import { MedicineHelper } from '~/app/helpers/actions/medicine-helper';
import { ActionService } from '~/app/services/action/action.service';
import { ActionListViewModel, ActionTxnDBModel } from '~/app/models/ui/action-models';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ActionDataDBRequest } from '~/app/models/api/actions-models';
import { PlatformHelper } from '~/app/helpers/platform-helper';

export class DataItem {
	public pstatus: string;
	public title: string;
	public due_at: string;
	public has_details: boolean;
	public desc: string;
	public status: number;
}
export class DataActionItem {
	uuid: string;
	admission_uuid: string;
	conf_type_code: string;
	schedule_uuid: string;
	exec_time: Date;
	name: string;
	desc: string;
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
	data = new Array<DataItem>();
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
	tempdata = new Array<DataItem>();


	// >> grouping 
	public _funcGrouping: (item: DataItem) => DataItem;

	// >> exapnd row
	expanded: false;

	// >> finding grouping index then after click show in top
	intakeIndex;
	medicineIndex;
	monitorIndex;
	outputIndex;

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


	constructor(public page: Page,
		private actionService: ActionService,
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

	// get dataItems(): ObservableArray<DataItem> {
	// 	return this._dataItems;
	// }
	get dataItems(): ObservableArray<DataActionItem> {
		return this._dataItems;
	}


	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {

		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this._dataItems = new ObservableArray<DataItem>();
		this.getActionData();
		// for (let i = 1; i < 50; i++) {
		// 	let newName = { ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" };
		// 	this.data.push(newName);
		// }
		// alert('action');

		// subscription for create actions
		this.actionSubscription = this.passdataservice.createActionsSubject.subscribe((value) => {

		}); // end of subscriptions.
		// this.createActions();
		this.actionDbData = new ActionDataDBRequest();
		this.actiondata = new ActionDataDBRequest();
	}

	public listLoaded() {
		//return; 
		console.log('list loaded');

		// this._dataItems = new ObservableArray(this.data);
		// this.addMoreItemsFromSource(20);
		setTimeout(() => {
			this.initDataItems();
			//this.addMoreItemsFromSource(20);
		}, 200);
	}

	public initDataItems() {
		const tempdata = new Array<DataItem>();
		this.tempdata.push({ title: "Saline", desc: "200ml", due_at: "17:00:00", has_details: true, pstatus: "Intake", status: 1 });
		this.tempdata.push({ title: "Output", desc: "200ml", due_at: "15:00:00", has_details: true, pstatus: "Output", status: 1 });
		this.tempdata.push({ title: "Sinarest", desc: "3 times a day after meal", due_at: "16:00:00", has_details: false, pstatus: "Medicine", status: 1 });
		this.tempdata.push({ title: "Acetaminophen", desc: "Morning and evening before meal", due_at: "17:30:00", has_details: false, pstatus: "Medicine", status: 2 });
		this.tempdata.push({ title: "Aspirin", desc: "Incase of high body temperature", due_at: "17:50:00", has_details: false, pstatus: "Medicine", status: 1 });
		this.tempdata.push({ title: "Zofran", desc: "Incase of continuos vomitting and nausea", due_at: "18:00:00", has_details: false, pstatus: "Medicine", status: 1 });
		this.tempdata.push({ title: "Temperature", desc: "Monitor every 2 hours", due_at: "12:00:00", has_details: true, pstatus: "Monitor", status: 3 });
		this.tempdata.push({ title: "Blood pressure", desc: "Monitor every 3 hours", due_at: "12:30:00", has_details: true, pstatus: "Monitor", status: 1 });
		this.tempdata.push({ title: "Pulse Rate", desc: "Monitor every 15 mins", due_at: "14:15:00", has_details: true, pstatus: "Monitor", status: 3 });
		this.tempdata.push({ title: "Respiration Rate", desc: "Monitor every 30 mins", due_at: "14:45:00", has_details: true, pstatus: "Monitor", status: 1 });

		for (let i = 0; i < 20; i++) {
			this.tempdata.forEach(item => {
				this.data.push(item);
			})
		}
		this._dataItems = new ObservableArray(this.tempdata);
		this.getCount();
		// this.dataItems.push(this.tempdata);
		// this.myItems = new ObservableArray<DataItem>(this.tempdata);
	}

	public addMoreItemsFromSource(chunkSize: number) {
		// console.log('items loaded pre dataitems', this.dataItems.length);
		// console.log('items loaded pre data', this.data.length);
		let newItems = this.data.slice(this.dataItems.length, this.dataItems.length + chunkSize);
		// this.dataItems.push(newItems);
		// console.log('items loaded post new items', newItems.length);
		// console.log('items loaded post', this.dataItems.length);
	}

	public onLoadMoreItemsRequested(args: LoadOnDemandListViewEventData) {
		// console.log('onLoadMoreItemsRequested');

		// const that = new WeakRef(this);
		const listView: RadListView = args.object;
		if (this.dataItems.length < this.data.length) {
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

		const medicine = this.tempdata.filter(a => a.pstatus === "Medicine");
		const medicineCount = medicine.length;
		console.log("medicineCount", medicineCount);

		const monitor = this.tempdata.filter(a => a.pstatus === "Monitor");
		const monitorCount = monitor.length;
		console.log("monitorCount", monitorCount);

		const intake = this.tempdata.filter(a => a.pstatus === "Intake");
		const intakeCount = intake.length;
		console.log("intakeCount", intakeCount);

		const output = this.tempdata.filter(a => a.pstatus === "Output");
		const outputCount = output.length;
		console.log("outputCount", outputCount);

		this.intakeIndex = 0;
		this.medicineIndex = intakeCount + 1;
		this.monitorIndex = intakeCount + medicineCount + 2;
		this.outputIndex = intakeCount + medicineCount + monitorCount + 3;


		console.log("medicine index", this.medicineIndex);
		console.log("monitor index", this.monitorIndex);
		console.log("output index", this.outputIndex);

	}
	// << Calculate Grouping index value


	// >> discard patient in list 
	public onDeleteRow(event: ListViewEventData) {

		console.log("Right swipe click");
		this._dataItems.splice(this._dataItems.indexOf(event.object.bindingContext), 1);
		console.log(this._dataItems);
		console.log(this._dataItems.length);
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
					console.log("action list", this.actionListItem);
				});
				// this.getCount();
				this.getListDataById();
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}

	// >> get action list by id
	public getListDataById() {
		console.log('this.actionnListItem', this.actionListItem);
		// this.tempList = new Array<DataActionItem>();
		this.actionListItem.forEach(item => {
			const actionListDataItem = new DataActionItem();
			actionListDataItem.admission_uuid = item.dbmodel.admission_uuid;
			actionListDataItem.schedule_uuid = item.dbmodel.schedule_uuid;
			actionListDataItem.conf_type_code = item.dbmodel.conf_type_code;
			actionListDataItem.exec_time = item.dbmodel.exec_time;

			this.chartService.getChartByUUID(actionListDataItem.schedule_uuid).then(
				(val) => {
					console.log('val', val);
					let val1: any;
					val1 = val;
					const conf = JSON.parse(val1.conf);
					actionListDataItem.name = conf.name;
					actionListDataItem.desc = conf.desc;

				})

			this.tempList.push(actionListDataItem);
			// console.log('testItem array', this.tempList);
		});
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
			console.log('confString', this.confString1);
		} else {
			this.actionDbData.comment = this.actiondata.comment;
			this.actionDbData.value = this.actiondata.value;
			this.confString = JSON.stringify(this.actionDbData);
			console.log('confString', this.confString);
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
		this.formData.txn_date = null;
		this.formData.txn_state = null;

		// console.log('this.actionformData', this.formData);

		// after done data push one by one ietm in array hold data
		this.actionDbArray.push(this.formData);
		console.log('this.actionDbArray', this.actionDbArray);

	}
	// all action done and discard save in action-trn-table
	save() {
		// array hold entries one by one save
		this.actionformData = new ActionTxnDBModel();
		// insert Action db model to sqlite db
		this.actionDbArray.forEach(item => {
			this.actionformData.uuid = item.uuid;
			this.actionformData.schedule_uuid = item.schedule_uuid;
			this.actionformData.conf_type_code = item.conf_type_code;
			this.actionformData.txn_data = item.txn_data;
			this.actionformData.txn_date = item.txn_date;
			this.actionformData.txn_state = item.txn_state;
			this.actionformData.conf_type_code = item.conf_type_code;
			this.actionformData.runtime_config_data = item.runtime_config_data;

			console.log('item.conf_type_code', item.conf_type_code);
			console.log('this.actionformData.schedule_uuid', this.actionformData.conf_type_code);
			console.log('actionformData', this.actionformData);
			
			// save action done and discard in DB
			this.actionService.insertActionTxnItem(this.actionformData);
		})
		
		// check data save entries added in action trn table 
		this.gettrnlistdata();
	}

	// selected done and discard row change background color
	itemSelected(item){
		item.selected = true;
	}
	gettrnlistdata(){
		this.actionService.getActionTxnList();
	}
}
