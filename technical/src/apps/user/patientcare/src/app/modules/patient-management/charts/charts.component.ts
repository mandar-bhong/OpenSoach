import { Component, OnInit, ViewChild } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { View, isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { ListViewLinearLayout, ListViewEventData, RadListView, ListViewItemSnapMode } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Observable } from 'tns-core-modules/data/observable';

export class DataItem {
	public pstatus: string;
	public title: string;
	public due_at: string;
	public has_details: boolean;
	public desc: string;
	public status: number;
}

@Component({
	moduleId: module.id,
	selector: 'charts',
	templateUrl: './charts.component.html',
	styleUrls: ['./charts.component.css']
})

export class ChartsComponent implements OnInit {

	private _dataItems: ObservableArray<any>;

	// >> seleced bottom button change color
	monitorbuttonClicked: boolean = false;
	intakebuttonClicked: boolean = true;
	medicinebuttonClicked: boolean = false;
	outputbuttonClicked: boolean = false;

	// >> finding grouping index then after click show in top
	intakeIndex;
	medicineIndex;
	monitorIndex;
	outputIndex;

	// >> grouping 
	private _funcGrouping: (item: DataItem) => DataItem;

	tempdata = new Array<DataItem>();

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

		this._dataItems = new ObservableArray<DataItem>();
	}

	public listLoaded() {

		console.log('list loaded');

		setTimeout(() => {
			this.initDataItems();
		}, 200);
	}

	public initDataItems() {
		const tempdata = new Array<DataItem>();
		this.tempdata.push({ title: "Saline",  desc:"200ml", due_at: "17:00:00", has_details: true, pstatus: "Intake", status: 1 });
		this.tempdata.push({ title: "Output", desc:"200ml", due_at: "15:00:00", has_details: true, pstatus: "Output", status: 1 });
		this.tempdata.push({ title: "Sinarest", desc:"3 times a day after meal" ,due_at: "16:00:00", has_details: false, pstatus: "Medicine", status: 1 });
		this.tempdata.push({ title: "Acetaminophen", desc:"Morning and evening before meal", due_at: "17:30:00", has_details: false, pstatus: "Medicine", status: 2 });
		this.tempdata.push({ title: "Aspirin", desc:"Incase of high body temperature", due_at: "17:50:00", has_details: false, pstatus: "Medicine", status: 1 });
		this.tempdata.push({ title: "Zofran", desc:"Incase of continuos vomitting and nausea", due_at: "18:00:00", has_details: false, pstatus: "Medicine", status: 1 });
		this.tempdata.push({ title: "Temperature", desc:"Monitor every 2 hours", due_at: "12:00:00", has_details: true, pstatus: "Monitor", status: 3 });
		this.tempdata.push({ title: "Blood pressure", desc:"Monitor every 3 hours", due_at: "12:30:00", has_details: true, pstatus: "Monitor", status: 1 });
		this.tempdata.push({ title: "Blood pressure", desc:"Monitor every 3 hours", due_at: "13:00:00", has_details: true, pstatus: "Monitor", status: 3 });
		this.tempdata.push({ title: "Pulse Rate", desc:"Monitor every 15 mins", due_at: "14:15:00", has_details: true, pstatus: "Monitor", status: 3 });
		this.tempdata.push({ title: "Respiration Rate", desc:"Monitor every 30 mins", due_at: "14:45:00", has_details: true, pstatus: "Monitor", status: 1 });

		this._dataItems = new ObservableArray(this.tempdata);
		this.getGroupIndex();
	}

	// >> Grouping position change 
	// >> Grouping intake scroll to top position change 
	public selectIntake() {
		const listView = this.listViewComponent.listView;
		listView.scrollToIndex(this.intakeIndex, false, ListViewItemSnapMode.Start);
		this.monitorbuttonClicked = false;
		this.intakebuttonClicked = true;
		this.medicinebuttonClicked = false;
		this.outputbuttonClicked = false;
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
	}
	// <<  Grouping medicine scroll to top position change

	// >> Calculate Grouping index value
	public getGroupIndex() {

		const medicine = this.tempdata.filter(a => a.pstatus === "Medicine");
		const medicineCount = medicine.length;

		const monitor = this.tempdata.filter(a => a.pstatus === "Monitor");
		const monitorCount = monitor.length;

		const intake = this.tempdata.filter(a => a.pstatus === "Intake");
		const intakeCount = intake.length;

		const output = this.tempdata.filter(a => a.pstatus === "Output");
		const outputCount = output.length;

		this.intakeIndex = 0;
		this.medicineIndex = intakeCount + 1;
		this.monitorIndex = intakeCount + medicineCount + 2;
		this.outputIndex = intakeCount + medicineCount + monitorCount + 3;

	}

}