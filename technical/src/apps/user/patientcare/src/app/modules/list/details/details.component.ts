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
	public tabSelectedIndex: number;
	b1;
	selectedfirst = true;
	selectedsecond = false;
	public SelectedIndex: number;

	constructor(private routerExtensions: RouterExtensions) {
		this.tabSelectedIndex = 0;
	}

	ngOnInit() { }
	changeTab() {
		if (this.tabSelectedIndex === 0) {
			this.tabSelectedIndex = 1;
		} else if (this.tabSelectedIndex === 1) {
			this.tabSelectedIndex = 2;
		} else if (this.tabSelectedIndex === 2) {
			this.tabSelectedIndex = 0;
		}
	}
	goBackPage() {
		this.routerExtensions.navigate(["/list"], { clearHistory: true });
	}
	firstTab() {
		this.selectedfirst = true;
		this.selectedsecond = false;
	}
	secondTab() {
		this.selectedfirst = false;
		this.selectedsecond = true;
	}
	test() {
		var coll = document.getElementsByClassName("collapsible");
		var i;

		for (i = 0; i < coll.length; i++) {
			coll[i].addEventListener("click", function () {
				this.classList.toggle("active");
				var content = this.nextElementSibling;
				if (content.style.display === "block") {
					content.style.display = "none";
				} else {
					content.style.display = "block";
				}
			});
		}

	}





	
}