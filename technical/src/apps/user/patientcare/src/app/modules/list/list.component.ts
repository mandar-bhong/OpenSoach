import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { SearchBar } from "tns-core-modules/ui/search-bar";
import { isAndroid } from "platform";
import * as app from "application";
import { Page } from "ui/page";

import { alert } from "tns-core-modules/ui/dialogs";
import { LocalNotifications } from "nativescript-local-notifications";
import { ObservableArray } from "tns-core-modules/data/observable-array";

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
	constructor(private routerExtensions: RouterExtensions,
		private page: Page) { }

		get dataItems(): ObservableArray<DataItem> {
			return this._dataItems;
	}
	ngOnInit() {

		const tempdata = new Array<DataItem>();
		tempdata.push({ ward: "3A/312", name: "Sumeet karande", mobile: "9878978980" });
		tempdata.push({ ward: "3B/323", name: "Amol Patil", mobile: "32423432432" });
		tempdata.push({ ward: "2A/643", name: "Sanjay Mohan", mobile: "453545352"});
		tempdata.push({ ward: "4A/515", name: "Sanjay Sawant", mobile: "9878978980"});
		tempdata.push({ ward: "5A/616", name: "Parveen", mobile: "5645645665"});
		

		

		for(let i=0;i<20;i++)
		{
			tempdata.forEach(item=>{
				this.data.push(item);
			})
		}

		this._dataItems = new ObservableArray(this.data);
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
	}
	public sBLoaded(args){
        var searchbar:SearchBar = <SearchBar>args.object;
        if(isAndroid){    
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
		this.searchiocn =  false;
		// this.page.actionBarHidden = true;
	}
	searchTabClose() {
		this.searchshow = false;
		this.searchiocn =  true;
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
			
	
}