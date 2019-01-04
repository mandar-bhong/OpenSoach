import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { ListViewLinearLayout, ListViewEventData, RadListView, ListViewItemSnapMode } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Observable } from 'tns-core-modules/data/observable';
import { RouterExtensions } from 'nativescript-angular/router';

export class DataItem {
	public name: string;
	public description: string;

}
@Component({
	moduleId: module.id,
	selector: 'section-one',
	templateUrl: './section-one.component.html',
	styleUrls: ['./section-one.component.css']
})

export class SectionOneComponent implements OnInit {
	private _dataItems: ObservableArray<any>;
	tempdata = new Array<DataItem>();
	constructor(private routerExtensions: RouterExtensions) { }

	ngOnInit() { }
	public listLoaded() {

		console.log('list loaded');

		setTimeout(() => {
			this.initDataItems();
		}, 200);
	}

	public initDataItems() {
		const tempdata = new Array<DataItem>();
		this.tempdata.push({ name: "Amol Patil",  description:"200ml" });
		this.tempdata.push({ name: "Shubham Lunia",  description:"3 times a day after meal" });
		this.tempdata.push({ name: "Mayuri jain",  description:"Morning and evening before meal" });
		this.tempdata.push({ name: "Sanjay Mohan",  description:"Incase of high body temperature" });
		this.tempdata.push({ name: "Pooja Lokare",  description:"Incase of continuos vomitting and nausea" });
		this.tempdata.push({ name: "Jagdish Wagh",  description:"Monitor every 2 hours" });
		this.tempdata.push({ name: "Mandar Bhong",  description:"Monitor every 3 hours" });
		this.tempdata.push({ name: "Praveen Pandey",  description:"Monitor every 15 mins" });
		this.tempdata.push({ name: "Shashank Atre",  description:"Incase of high body temperature" });
		this.tempdata.push({ name: "Abhijeet Kalbhor",  description:"Morning and evening before meal" });
		this.tempdata.push({ name: "Sarjerao",  description:"Monitor every 15 mins" });
		this.tempdata.push({ name: "Rahul",  description:"Incase of high body temperature" });
		this.tempdata.push({ name: "Praveen",  description:"Morning and evening before meal" });

		this._dataItems = new ObservableArray(this.tempdata);
	}
	goCameras() {
		this.routerExtensions.navigate(['patientmgnt', 'cameras'], { clearHistory: true });

	}
}