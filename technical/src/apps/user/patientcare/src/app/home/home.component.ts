import { Component, OnInit, ViewChild } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Router } from "@angular/router";
import { DatabaseService } from "../services/offline-store/database.service";
import { Page } from "ui/page";
import { RouterExtensions } from "nativescript-angular/router";
import { SearchBar } from "tns-core-modules/ui/search-bar";
import { isAndroid } from "platform";
import * as app from "application";

import { alert } from "tns-core-modules/ui/dialogs";
import { LocalNotifications } from "nativescript-local-notifications";
import { ObservableArray } from "tns-core-modules/data/observable-array";
import { RadListViewComponent } from "nativescript-ui-listview/angular";
import { ListViewEventData, RadListView, LoadOnDemandListViewEventData, ListViewLinearLayout } from "nativescript-ui-listview";
import { View } from "tns-core-modules/ui/core/view";
import { layout } from "tns-core-modules/utils/utils";
import { EventData } from "tns-core-modules/data/observable";

export class DataItem {
	public name: string;
	public ward: string;
	public mobile: string;
	public attended: number;

}

@Component({
	selector: "Home",
	moduleId: module.id,
	templateUrl: "./home.component.html",
	styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
	private _dataItems: ObservableArray<DataItem>;
	data = new Array<DataItem>();
	private leftItem: View;
	private rightItem: View;
	private mainView: View;
	private layout: ListViewLinearLayout;

	// var for search 
	public myItems: ObservableArray<DataItem> = new ObservableArray<DataItem>();
	tempdata = new Array<DataItem>();
	constructor(private routerExtensions: RouterExtensions,
		private page: Page) { }

	get dataItems(): ObservableArray<DataItem> {
		return this._dataItems;
	}

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this._dataItems = new ObservableArray<DataItem>();

		// console.log('this.data', this.data);
		// for (let i = 1; i < 100; i++) {
		// 	let newName = { ward: "3A/312", name: "Sumeet karande", mobile: "9878978980"};
		// 	this.data.push(newName);
		// }
		//this.initDataItems();
		console.log('init completed');
	}
	public sBLoaded(args) {
		var searchbar: SearchBar = <SearchBar>args.object;
		if (isAndroid) {
			searchbar.android.clearFocus();
		}
	}

	goBackPage() {
		this.routerExtensions.navigate(["/home"], { clearHistory: true });
	}
	details() {
		this.routerExtensions.navigate(["/patientmgnt/details"], { clearHistory: true });
	}
	camerasdetails() {
		this.routerExtensions.navigate(["/patientmgnt/cameras"], { clearHistory: true });
	}

	// onSubmit() {

	// }
	searchBarLoaded() {

	}
	onTextChange() {

	}

	// one time call push notification  code start
	showWithSound(): void {
		LocalNotifications.schedule([{
			id: 1,
			title: 'Sound & Badge',
			body: 'Who needs a push service anyway?',
			badge: 1,
			at: new Date(new Date().getTime() + (5 * 1000)) // 5 seconds from now
		}]);

		// adding a handler, so we can do something with the received notification.. in this case an alert
		LocalNotifications.addOnMessageReceivedCallback(data => {
			alert({
				title: "Local Notification received",
				message: `id: '${data.id}', title: '${data.title}'.`,
				okButtonText: "Roger that"
			});
		});
	}
	// one time call push notiffication  code end
	// continus call push notiffication after create 1min code start
	conti(): void {
		LocalNotifications.schedule([{
			id: 2,
			title: 'Cancel me, quickly!',
			body: 'Who thought this would be a good idea!?',
			interval: 'minute',
			sound: null,
			at: new Date(new Date().getTime() + (5 * 1000)), // 5 seconds from now
		}]);
	}
	// continus call push notiffication after create 1min code start

	// cancel all push notiffication code start
	cancelAll(): void {
		LocalNotifications.cancelAll();
	}
	// cancel all push notiffication code start


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
		this.tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980", attended: 1 });
		this.tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432", attended: 2 });
		this.tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352", attended: 2 });
		this.tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980", attended: 1 });
		this.tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665", attended: 1 });
		this.tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980", attended: 1 });
		this.tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432", attended: 2 });
		this.tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352", attended: 1 });
		this.tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980", attended: 1 });
		this.tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665", attended: 2 });
		this.tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980", attended: 1 });
		this.tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432", attended: 1 });
		this.tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352", attended: 2 });
		this.tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980", attended: 1 });
		this.tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665", attended: 1 });




		for (let i = 0; i < 20; i++) {
			this.tempdata.forEach(item => {
				this.data.push(item);
			})
		}

		// this.dataItems.push(this.tempdata);
		this.myItems = new ObservableArray<DataItem>(this.tempdata);
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

	// search record by list code start
	public onSubmit(args) {
		let searchBar = <SearchBar>args.object;
		let searchValue = searchBar.text.toLowerCase();

		this.myItems = new ObservableArray<DataItem>();
		if (searchValue !== "") {
			for (let i = 0; i < this.tempdata.length; i++) {
				if (this.tempdata[i].name.toLowerCase().indexOf(searchValue) !== -1) {
					this.myItems.push(this.tempdata[i]);
				}
			}
		}
	}
	// search record by list code end

	//clear search record list then show all list record code start
	public onClear(args) {
		let searchBar = <SearchBar>args.object;
		searchBar.text = "";
		// searchBar.hint = "Search for a country and press enter";
		this.myItems = new ObservableArray<DataItem>();
		this.tempdata.forEach(item => {
			this.myItems.push(item);
		});
	}
	//clear search record list then show all list record code end

}
