import { Component, OnInit, ViewChild } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { View, isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { ListViewLinearLayout, ListViewEventData, RadListView, ListViewItemSnapMode } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Observable } from 'tns-core-modules/data/observable';
import { ChartService } from "~/app/services/chart/chart.service";
import { ChartListViewModel } from "~/app/models/ui/chart-models";
import { RouterExtensions } from 'nativescript-angular/router';

@Component({
	moduleId: module.id,
	selector: 'charts',
	templateUrl: './charts.component.html',
	styleUrls: ['./charts.component.css']
})

export class ChartsComponent implements OnInit {

	private chartListItems = new ObservableArray<ChartListViewModel>();

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
	public _funcGrouping: (item: ChartListViewModel) => ChartListViewModel;

	dialogOpen = false;
	constructor(private chartService: ChartService,
		private routerExtensions: RouterExtensions) {
		//  list grouping
		this._funcGrouping = (item: any) => {
			return item.dbmodel.conf_type_code;
		};
	}

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {

		this.getChartData();

	}
	showDialog() {
		this.dialogOpen = true;
	}

	closeDialog() {
		this.dialogOpen = false;
	}

	public listLoaded() {

		console.log('list loaded');
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

		const medicine = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Medicine");
		const medicineCount = medicine.length;

		const monitor = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Monitor");
		const monitorCount = monitor.length;

		const intake = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Intake");
		const intakeCount = intake.length;

		const output = this.chartListItems.filter(a => a.dbmodel.conf_type_code === "Output");
		const outputCount = output.length;

		this.intakeIndex = 0;
		this.medicineIndex = intakeCount + 1;
		this.monitorIndex = intakeCount + medicineCount + 2;
		this.outputIndex = intakeCount + medicineCount + monitorCount + 3;

	}

	get _listItems(): ObservableArray<ChartListViewModel> {
		return this.chartListItems;
	}

	public getChartData() {
		this.chartService.getChartList().then(
			(val) => {
				val.forEach(item => {
					let chartListItem = new ChartListViewModel();
					chartListItem.dbmodel = item;
					chartListItem.dbmodel.conf = JSON.parse(item.conf);
					this.chartListItems.push(chartListItem);
				});
				this.getGroupIndex();
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}
	monitorForm() {
		this.routerExtensions.navigate(['patientmgnt', 'monitor-chart'], { clearHistory: false });	
	}
	medicineForm() {
		this.routerExtensions.navigate(['patientmgnt', 'medicine-chart'], { clearHistory: false });			
	}
	intakeForm() {
		this.routerExtensions.navigate(['patientmgnt', 'intake-chart'], { clearHistory: false });		
	}

}