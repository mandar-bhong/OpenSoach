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

export class DataItem {
	// public name: string;
	// public ward: string;
	// public mobile: string;
	// public attended: number;
	public pstatus: string;
	public title: string;
	public due_at: string;
	public has_details: boolean;

}
@Component({
	moduleId: module.id,
	selector: 'action',
	templateUrl: './action.component.html',
	styleUrls: ['./action.component.css']
})

export class ActionComponent implements OnInit {

	private _dataItems: ObservableArray<any>;
	data = new Array<DataItem>();
	private leftItem: View;
	private rightItem: View;
	private mainView: View;
	private layout: ListViewLinearLayout;

	// >> seleced bottom button change color
	monitorbuttonClicked: boolean = false;
	intakebuttonClicked: boolean = true;
	medicinebuttonClicked: boolean = false;
	outputbuttonClicked: boolean = false;

	// >> search var declaration
	// public myItems: ObservableArray<DataItem> = new ObservableArray<DataItem>();
	tempdata = new Array<DataItem>();


	// >> grouping 
	private _funcGrouping: (item: DataItem) => DataItem;

	// >> exapnd row
	expanded: false;

	// >> swap delete
	private leftThresholdPassed = false;
	private rightThresholdPassed = false;

	// >> finding grouping index then after click show in top
	intakeIndex;
	medicineIndex;
	monitorIndex;
	outputIndex;

	constructor() {
		//  list grouping
		this._funcGrouping = (item: any) => {
			return item.pstatus;
		};
	}

	get dataItems(): ObservableArray<DataItem> {
		return this._dataItems;
	}

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {

		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this._dataItems = new ObservableArray<DataItem>();

		// for (let i = 1; i < 50; i++) {
		// 	let newName = { ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" };
		// 	this.data.push(newName);
		// }
		// alert('action');
	}
	// >> angular-listview-swipe-action-thresholds

	// >> html code selectionBehavior="None" (itemSwipeProgressEnded)="onSwipeCellFinished($event)" (itemSwipeProgressStarted)="onSwipeCellStarted($event)" (itemSwipeProgressChanged)="onCellSwiping($event)" swipeActions="true"
	public onCellSwiping(args: ListViewEventData) {
		const swipeLimits = args.data.swipeLimits;
		const swipeView = args['swipeView'];
		const mainView = args['mainView'];
		const leftItem = swipeView.getViewById('mark-view');
		const rightItem = swipeView.getViewById('delete-view');

		if (args.data.x > swipeView.getMeasuredWidth() / 4 && !this.leftThresholdPassed) {
			console.log("Notify perform left action");
			const markLabel = leftItem.getViewById('mark-text');
			this.leftThresholdPassed = true;
		} else if (args.data.x < -swipeView.getMeasuredWidth() / 4 && !this.rightThresholdPassed) {
			const deleteLabel = rightItem.getViewById('delete-text');
			console.log("Notify perform right action");
			this.rightThresholdPassed = true;
		}
		if (args.data.x > 0) {
			const leftDimensions = View.measureChild(
				leftItem.parent,
				leftItem,
				layout.makeMeasureSpec(Math.abs(args.data.x), layout.EXACTLY),
				layout.makeMeasureSpec(mainView.getMeasuredHeight(), layout.EXACTLY));
			View.layoutChild(leftItem.parent, leftItem, 0, 0, leftDimensions.measuredWidth, leftDimensions.measuredHeight);
		} else {
			const rightDimensions = View.measureChild(
				rightItem.parent,
				rightItem,
				layout.makeMeasureSpec(Math.abs(args.data.x), layout.EXACTLY),
				layout.makeMeasureSpec(mainView.getMeasuredHeight(), layout.EXACTLY));

			View.layoutChild(rightItem.parent, rightItem, mainView.getMeasuredWidth() - rightDimensions.measuredWidth, 0, mainView.getMeasuredWidth(), rightDimensions.measuredHeight);
		}
	}
	// << angular-listview-swipe-action-thresholds

	// >> angular-listview-swipe-action-thresholds-limits
	public onSwipeCellStarted(args: ListViewEventData) {
		const swipeLimits = args.data.swipeLimits;
		const swipeView = args['object'];
		const leftItem = swipeView.getViewById('mark-view');
		const rightItem = swipeView.getViewById('delete-view');
		swipeLimits.left = swipeLimits.right = args.data.x > 0 ? swipeView.getMeasuredWidth() / 2 : swipeView.getMeasuredWidth() / 2;
		swipeLimits.threshold = swipeView.getMeasuredWidth();
	}
	// << angular-listview-swipe-action-thresholds-limits

	// >> angular-listview-swipe-actions-execute
	public onSwipeCellFinished(args: ListViewEventData) {
		const swipeView = args['object'];
		const leftItem = swipeView.getViewById('mark-view');
		const rightItem = swipeView.getViewById('delete-view');
		if (this.leftThresholdPassed) {
			console.log("Perform left action");
		} else if (this.rightThresholdPassed) {
			console.log("Perform right action");
			// this.onRightSwipeClick(args);
		}
		this.leftThresholdPassed = false;
		this.rightThresholdPassed = false;
	}
	// << angular-listview-swipe-actions-execute

	public onLeftSwipeClick(args: ListViewEventData) {
		console.log("Left swipe click");
		this.listViewComponent.listView.notifySwipeToExecuteFinished();
	}

	public onRightSwipeClick(args) {
		console.log("Right swipe click");
		this._dataItems.splice(this._dataItems.indexOf(args.object.bindingContext), 1);
		console.log(this._dataItems);
		// const index = this._dataItems.indexOf(args);
		// this._dataItems.splice(index, 1);
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
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Crocin", due_at: "16:00:00", has_details: false, pstatus: "Medicine" });
		this.tempdata.push({ title: "Crocin", due_at: "16:00:00", has_details: false, pstatus: "Medicine" });
		this.tempdata.push({ title: "Saline", due_at: "17:00:00", has_details: true, pstatus: "Intake" });
		this.tempdata.push({ title: "Saline", due_at: "17:00:00", has_details: true, pstatus: "Intake" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Output", due_at: "15:00:00", has_details: true, pstatus: "Output" });
		this.tempdata.push({ title: "Crocin", due_at: "16:00:00", has_details: false, pstatus: "Medicine" });
		this.tempdata.push({ title: "Crocin", due_at: "16:00:00", has_details: false, pstatus: "Medicine" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Saline", due_at: "17:00:00", has_details: true, pstatus: "Intake" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });
		this.tempdata.push({ title: "Monitor Temperature", due_at: "15:00:00", has_details: true, pstatus: "Monitor" });


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
		this.dataItems.push(newItems);
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

	// >> select list item >> multipleSelection="true" selectionBehavior="Press" (itemSelected)="itemSelected($event)" (itemDeselected)="itemDeselected($event)"
	itemSelected(args: ListViewEventData) {
		// const item = this._dataItems.getItem(args.index);
		// alert(item.name);
		// item.selected = true;
	}
	// >> deselect list item 
	itemDeselected(args: ListViewEventData) {
		// const item = this._dataItems.getItem(args.index);
		// item.selected = false;
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


}