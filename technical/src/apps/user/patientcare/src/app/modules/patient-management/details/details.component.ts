import { Component, OnInit } from '@angular/core';
import { RouterExtensions, PageRoute } from "nativescript-angular/router";
import { SelectedIndexChangedEventData } from "tns-core-modules/ui/tab-view";
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";
import { switchMap } from "rxjs/operators";
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { PassDataService } from '~/app/services/pass-data-service';
@Component({
	moduleId: module.id,
	selector: 'details',
	templateUrl: './details.component.html',
	styleUrls: ['./details.component.css']
})

export class DetailsComponent implements OnInit {
	// public tabSelectedIndex: number;
	// data = [];
	monitor = true;
	action = false;
	chart = false;
	report = false;

	// >> seleced bottom button change color
	buttonClicked: boolean = true;
	actionbuttonClicked: boolean = false;
	chartbuttonClicked: boolean = false;
	reportbuttonClicked: boolean = false;
	selectedPatient: PatientListViewModel;
	patientName: string;
	constructor(
		private routerExtensions: RouterExtensions,
		private passdataservice: PassDataService,
		private pageRoute: PageRoute) {
	}

	ngOnInit() {
		// getting patient data form service 
	      // for getting header name label text
		this.patientName = this.passdataservice.getHeaderName();
		console.log('patient name',this.patientName)
		// goBackPage() {
		// 	this.routerExtensions.navigate(["/home"], { clearHistory: true });
		// }

		// patientdetail(){
		// 	this.routerExtensions.navigate(['patientmgnt', 'patient'], { clearHistory: true });
		// }
	}
	monitorData() {
		this.monitor = true;
		this.action = false;
		this.chart = false;
		this.report = false;
		this.actionbuttonClicked = false;
		this.chartbuttonClicked = false;
		this.reportbuttonClicked = false;
	}
	actionData() {
		this.monitor = false;
		this.action = true;
		this.chart = false;
		this.report = false;
		this.buttonClicked = false;
		this.chartbuttonClicked = false;
		this.reportbuttonClicked = false;
	}
	chartData() {
		this.monitor = false;
		this.action = false;
		this.chart = true;
		this.report = false;
		this.buttonClicked = false;
		this.actionbuttonClicked = false;
		this.reportbuttonClicked = false;
	}
	reportData() {
		this.monitor = false;
		this.action = false;
		this.chart = false;
		this.report = true;
		this.buttonClicked = false;
		this.actionbuttonClicked = false;
		this.chartbuttonClicked = false;
	}

}