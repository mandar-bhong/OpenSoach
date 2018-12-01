import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { SearchBar } from "tns-core-modules/ui/search-bar";
import { isAndroid } from "platform";
import * as app from "application";
import { Page } from "ui/page";

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

}

@Component({
	moduleId: module.id,
	selector: 'list',
	templateUrl: './list.component.html',
	styleUrls: ['./list.component.css']
})


export class ListComponent implements OnInit {
	private _dataItems: ObservableArray<DataItem>;
	data = new Array<DataItem>();
	searchshow = false;
	searchiocn = true;
	private leftItem: View;
	private rightItem: View;
	private mainView: View;
	private layout: ListViewLinearLayout;
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

		// this.data.push({ ward: "7A/244", name: "Om", mobile: "323434355445"});
		// this.data.push({ ward: "3B/324", name: "Ahubham", mobile: "9809878679"});
		// this.data.push({ ward: "2A/454", name: "Suraj", mobile: "76568768778"});
		// this.data.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665"});
		// this.data.push({ ward: "6A/897", name: "Mandar bhong", mobile: "98789909090"});
		// this.data.push({ ward: "7A/244", name: "Om", mobile: "323434355445"});
		// this.data.push({ ward: "3B/324", name: "Ahubham", mobile: "9809878679"});

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
		this.routerExtensions.navigate(["/list/details"], { clearHistory: true });
	}

	searchTab() {
		this.searchshow = true;
		this.searchiocn = false;
		// this.page.actionBarHidden = true;
	}
	searchTabClose() {
		this.searchshow = false;
		this.searchiocn = true;
	}
	onSubmit() {

	}
	searchBarLoaded() {

	}
	onTextChange() {

	}


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

	cancelAll(): void {
		LocalNotifications.cancelAll();
	}

	public onCellSwiping(args: ListViewEventData) {
		const swipeLimits = args.data.swipeLimits;
		const swipeView = args['swipeView'];
		this.mainView = args['mainView'];
		this.leftItem = swipeView.getViewById('left-stack');
		this.rightItem = swipeView.getViewById('right-stack');

		if (args.data.x > 0) {
			const leftDimensions = View.measureChild(
				<View>this.leftItem.parent,
				this.leftItem,
				layout.makeMeasureSpec(Math.abs(args.data.x), layout.EXACTLY),
				layout.makeMeasureSpec(this.mainView.getMeasuredHeight(), layout.EXACTLY));
			View.layoutChild(<View>this.leftItem.parent, this.leftItem, 0, 0, leftDimensions.measuredWidth, leftDimensions.measuredHeight);
			this.hideOtherSwipeTemplateView("left");
		} else {
			const rightDimensions = View.measureChild(
				<View>this.rightItem.parent,
				this.rightItem,
				layout.makeMeasureSpec(Math.abs(args.data.x), layout.EXACTLY),
				layout.makeMeasureSpec(this.mainView.getMeasuredHeight(), layout.EXACTLY));

			View.layoutChild(<View>this.rightItem.parent, this.rightItem, this.mainView.getMeasuredWidth() - rightDimensions.measuredWidth, 0, this.mainView.getMeasuredWidth(), rightDimensions.measuredHeight);
			this.hideOtherSwipeTemplateView("right");
		}
	}

	private hideOtherSwipeTemplateView(currentSwipeView: string) {
		switch (currentSwipeView) {
			case "left":
				if (this.rightItem.getActualSize().width !== 0) {
					View.layoutChild(<View>this.rightItem.parent, this.rightItem, this.mainView.getMeasuredWidth(), 0, this.mainView.getMeasuredWidth(), 0);
				}
				break;
			case "right":
				if (this.leftItem.getActualSize().width !== 0) {
					View.layoutChild(<View>this.leftItem.parent, this.leftItem, 0, 0, 0, 0);
				}
				break;
			default:
				break;
		}
	}
	// << angular-listview-swipe-action-multiple

	// >> angular-listview-swipe-action-multiple-limits
	public onSwipeCellStarted(args: ListViewEventData) {
		const swipeLimits = args.data.swipeLimits;
		swipeLimits.threshold = args['mainView'].getMeasuredWidth() * 0.2; // 20% of whole width
		swipeLimits.left = swipeLimits.right = args['mainView'].getMeasuredWidth() * 0.65; // 65% of whole width
	}
	// << angular-listview-swipe-action-multiple-limits

	public onSwipeCellFinished(args: ListViewEventData) {
		if (args.data.x > 200) {
			console.log("Perform left action");
		} else if (args.data.x < -200) {
			console.log("Perform right action");
		}
	}

	public onLeftSwipeClick(args: EventData) {
		let itemView = args.object as View;
		console.log("Button clicked: " + itemView.id + " for item with index: " + this.listViewComponent.listView.items.indexOf(itemView.bindingContext));
		this.listViewComponent.listView.notifySwipeToExecuteFinished();
	}

	public onRightSwipeClick(args: EventData) {
		let itemView = args.object as View;
		console.log("Button clicked: " + itemView.id + " for item with index: " + this.listViewComponent.listView.items.indexOf(itemView.bindingContext));
		this.listViewComponent.listView.notifySwipeToExecuteFinished();
	}

	public listLoaded() {
		//return; 
		console.log('list loaded');
		
		//this._dataItems = new ObservableArray(this.data);
		//this.addMoreItemsFromSource(20);
		setTimeout(() => {
			this.initDataItems();
			this.addMoreItemsFromSource(20);
		}, 200);
	}

	public initDataItems() {
		const tempdata = new Array<DataItem>();
		tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" });
		tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432" });
		tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352" });
		tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980" });
		tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665" });




		for (let i = 0; i < 20; i++) {
			tempdata.forEach(item => {
				this.data.push(item);
			})
		}
	}

	public addMoreItemsFromSource(chunkSize: number) {
		console.log('items loaded pre dataitems', this.dataItems.length);
		console.log('items loaded pre data', this.data.length);
		let newItems = this.data.slice(this.dataItems.length, this.dataItems.length + chunkSize);
		this.dataItems.push(newItems);
		console.log('items loaded post new items', newItems.length);
		console.log('items loaded post', this.dataItems.length);
	}

	public onLoadMoreItemsRequested(args: LoadOnDemandListViewEventData) {
		console.log('onLoadMoreItemsRequested');

		// const that = new WeakRef(this);
		const listView: RadListView = args.object;
		if (this.dataItems.length < this.data.length) {
			setTimeout(()=> {
				this.addMoreItemsFromSource(20);
				listView.notifyLoadOnDemandFinished();
				//console.log('onLoadMoreItemsRequested', this.dataItems.length);
			}, 200);
		} else {
			args.returnValue = false;
			listView.notifyLoadOnDemandFinished(true);
			console.log('onLoadMoreItemsRequested', 'load on demand finished');
		}
	}
}