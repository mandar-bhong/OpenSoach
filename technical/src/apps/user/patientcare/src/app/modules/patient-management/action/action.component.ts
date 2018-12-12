import { Component, OnInit, ViewChild } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { View } from 'tns-core-modules/ui/page/page';
import { ListViewLinearLayout, ListViewEventData, RadListView } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { ObservableProperty } from '~/app/observable-property-decorator';

// import { SearchBar } from "tns-core-modules/ui/search-bar";
import { Observable } from 'tns-core-modules/data/observable';


export class DataItem {
	public name: string;
	public ward: string;
	public mobile: string;
	public attended: number;
	public pstatus: string;

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
	// data = [];

	// search var declaration
	// public myItems: ObservableArray<DataItem> = new ObservableArray<DataItem>();
	tempdata = new Array<DataItem>();

	// grouping 
	private _funcGrouping: (item: DataItem) => DataItem;
	// private _funcSorting: (item1: any, item2: any) => number;

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

	// public onCellSwiping(args: ListViewEventData) {
	// 	const swipeLimits = args.data.swipeLimits;
	// 	const swipeView = args['swipeView'];
	// 	this.mainView = args['mainView'];
	// 	this.leftItem = swipeView.getViewById('left-stack');
	// 	this.rightItem = swipeView.getViewById('right-stack');

	// 	if (args.data.x > 0) {
	// 		const leftDimensions = View.measureChild(
	// 			<View>this.leftItem.parent,
	// 			this.leftItem,
	// 			layout.makeMeasureSpec(Math.abs(args.data.x), layout.EXACTLY),
	// 			layout.makeMeasureSpec(this.mainView.getMeasuredHeight(), layout.EXACTLY));
	// 		View.layoutChild(<View>this.leftItem.parent, this.leftItem, 0, 0, leftDimensions.measuredWidth, leftDimensions.measuredHeight);
	// 		this.hideOtherSwipeTemplateView("left");
	// 	} else {
	// 		const rightDimensions = View.measureChild(
	// 			<View>this.rightItem.parent,
	// 			this.rightItem,
	// 			layout.makeMeasureSpec(Math.abs(args.data.x), layout.EXACTLY),
	// 			layout.makeMeasureSpec(this.mainView.getMeasuredHeight(), layout.EXACTLY));

	// 		View.layoutChild(<View>this.rightItem.parent, this.rightItem, this.mainView.getMeasuredWidth() - rightDimensions.measuredWidth, 0, this.mainView.getMeasuredWidth(), rightDimensions.measuredHeight);
	// 		this.hideOtherSwipeTemplateView("right");
	// 	}
	// }

	// private hideOtherSwipeTemplateView(currentSwipeView: string) {
	// 	switch (currentSwipeView) {
	// 		case "left":
	// 			if (this.rightItem.getActualSize().width !== 0) {
	// 				View.layoutChild(<View>this.rightItem.parent, this.rightItem, this.mainView.getMeasuredWidth(), 0, this.mainView.getMeasuredWidth(), 0);
	// 			}
	// 			break;
	// 		case "right":
	// 			if (this.leftItem.getActualSize().width !== 0) {
	// 				View.layoutChild(<View>this.leftItem.parent, this.leftItem, 0, 0, 0, 0);
	// 			}
	// 			break;
	// 		default:
	// 			break;
	// 	}
	// }
	// // << angular-listview-swipe-action-multiple

	// // >> angular-listview-swipe-action-multiple-limits
	// public onSwipeCellStarted(args: ListViewEventData) {
	// 	const swipeLimits = args.data.swipeLimits;
	// 	swipeLimits.threshold = args['mainView'].getMeasuredWidth() * 0.2; // 20% of whole width
	// 	swipeLimits.left = swipeLimits.right = args['mainView'].getMeasuredWidth() * 0.65; // 65% of whole width
	// }
	// // << angular-listview-swipe-action-multiple-limits

	// public onSwipeCellFinished(args: ListViewEventData) {
	// 	if (args.data.x > 200) {
	// 		console.log("Perform left action");
	// 	} else if (args.data.x < -200) {
	// 		console.log("Perform right action");
	// 	}
	// }

	// public onLeftSwipeClick(args: EventData) {
	// 	let itemView = args.object as View;
	// 	console.log("Button clicked: " + itemView.id + " for item with index: " + this.listViewComponent.listView.items.indexOf(itemView.bindingContext));
	// 	this.listViewComponent.listView.notifySwipeToExecuteFinished();
	// }

	// public onRightSwipeClick(args: EventData) {
	// 	let itemView = args.object as View;
	// 	console.log("Button clicked: " + itemView.id + " for item with index: " + this.listViewComponent.listView.items.indexOf(itemView.bindingContext));
	// 	this.listViewComponent.listView.notifySwipeToExecuteFinished();
	// }

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
		// const tempdata = new Array<DataItem>();
		this.tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980", attended: 1, pstatus: "Monitor" });
		this.tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432", attended: 2, pstatus: "Medicine" });
		this.tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352", attended: 2, pstatus: "Medicine" });
		this.tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980", attended: 1, pstatus: "Intake" });
		this.tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665", attended: 1, pstatus: "Intake" });
		this.tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980", attended: 1, pstatus: "Output" });
		this.tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432", attended: 2, pstatus: "Output" });
		this.tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352", attended: 1, pstatus: "Output" });
		this.tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980", attended: 1, pstatus: "Medicine" });
		this.tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665", attended: 2, pstatus: "Medicine" });
		this.tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980", attended: 1, pstatus: "Monitor" });
		this.tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432", attended: 1, pstatus: "Monitor" });
		this.tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352", attended: 2, pstatus: "Intake" });
		this.tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980", attended: 1, pstatus: "Monitor" });
		this.tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665", attended: 1, pstatus: "Monitor" });




		for (let i = 0; i < 20; i++) {
			this.tempdata.forEach(item => {
				this.data.push(item);
			})
		}
		this._dataItems = new ObservableArray(this.tempdata);
		// this.dataItems.push(this.tempdata);
		// this.myItems = new ObservableArray<DataItem>(this.tempdata);
	}

	// public addMoreItemsFromSource(chunkSize: number) {
	// 	// console.log('items loaded pre dataitems', this.dataItems.length);
	// 	// console.log('items loaded pre data', this.data.length);
	// 	let newItems = this.data.slice(this.dataItems.length, this.dataItems.length + chunkSize);
	// 	this.dataItems.push(newItems);
	// 	// console.log('items loaded post new items', newItems.length);
	// 	// console.log('items loaded post', this.dataItems.length);
	// }

	// public onLoadMoreItemsRequested(args: LoadOnDemandListViewEventData) {
	// 	// console.log('onLoadMoreItemsRequested');

	// 	// const that = new WeakRef(this);
	// 	const listView: RadListView = args.object;
	// 	if (this.dataItems.length < this.data.length) {
	// 		setTimeout(()=> {
	// 			this.addMoreItemsFromSource(20);
	// 			listView.notifyLoadOnDemandFinished();
	// 			//console.log('onLoadMoreItemsRequested', this.dataItems.length);
	// 		}, 200);
	// 	} else {
	// 		args.returnValue = false;
	// 		listView.notifyLoadOnDemandFinished(true);
	// 		// console.log('onLoadMoreItemsRequested', 'load on demand finished');
	// 	}
	// }

     // select list item 
	itemSelected(args: ListViewEventData) {
		// const item = this._dataItems.getItem(args.index);
		// alert(item.name);
		// item.selected = true;
	}
	// deselect list item 
	itemDeselected(args: ListViewEventData) {
		// const item = this._dataItems.getItem(args.index);
		// item.selected = false;
	}

	public onSubmit(args) {
		let searchBar = this.tempdata.filter(a => a.pstatus === "monitor");
		
		// let searchValue = searchBar.text;
		console.log('searchBar', searchBar);
		let searchValue = searchBar;
		this._dataItems = new ObservableArray(searchBar);
		// this._funcGrouping = (item: any) => {
		// 	return item.pstatus;
		// };	

		// this._dataItems = new ObservableArray<DataItem>();
		// if (searchValue !== null) {
		// 	for (let i = 0; i < this.tempdata.length; i++) {
		// 		if (this.tempdata[i].pstatus ) {
		// 			this._dataItems.push(this.tempdata[i]);
		// 			console.log('_dataItems', this._dataItems);
		// 		}
		// 	}
		// }
	}
	// medecine(){
	// 	let searchBar = this.tempdata.filter(a => a.pstatus === "medicine");
		
	// 	console.log('searchBar', searchBar);
	// 	let searchValue = searchBar;
	// 	this._dataItems = new ObservableArray(searchBar);
	// }
	

}