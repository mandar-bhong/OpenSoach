import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from "nativescript-angular/router";
import { SelectedIndexChangedEventData } from "tns-core-modules/ui/tab-view";
import { SegmentedBar, SegmentedBarItem } from "tns-core-modules/ui/segmented-bar";

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

	// >> seleced bottom button change color
	buttonClicked: boolean = true;
	actionbuttonClicked: boolean = false;
	chartbuttonClicked: boolean = false;
	reportbuttonClicked: boolean = false;

	constructor(private routerExtensions: RouterExtensions) {

	}

	ngOnInit() { }

	goBackPage() {
		this.routerExtensions.navigate(["/home"], { clearHistory: true });
	}
	
	patientdetail(){
		this.routerExtensions.navigate(['patientmgnt', 'patient'], { clearHistory: true });
	}
	monitorData() {
		this.monitor = true;
		this.action = false;
		this.chart = false;
		this.actionbuttonClicked= false;
		this.chartbuttonClicked = false;
		this.reportbuttonClicked = false;
	}
	actionData() {
		this.monitor = false;
		this.action = true;
		this.chart =false;
		this.buttonClicked= false;
		this.chartbuttonClicked = false;
		this.reportbuttonClicked = false;
	}
	chartData() {
		this.monitor = false;
		this.action = false;
		this.chart =true;
		this.buttonClicked= false;
		this.actionbuttonClicked = false;
		this.reportbuttonClicked = false;
	}
	reportData(){
		this.monitor = false;
		this.action = false;
		this.chart =false;
		this.buttonClicked= false;
		this.actionbuttonClicked = false;
		this.chartbuttonClicked = false;
	}
	
}