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
import { PatientListService } from "~/app/services/patient-list/patient-list.service";
import { PatientListViewModel } from "~/app/models/ui/patient-view-models";
import { PatientDetails } from "~/app/models/ui/patient-details";
import { PassDataService } from "~/app/services/pass-data-service";
import { DataItem } from "~/app/modules/patient-management/reports/section-one/section-one.component";

@Component({
	selector: "Home",
	moduleId: module.id,
	templateUrl: "./home.component.html",
	styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
	patientname: string;
	public patientListItems = new ObservableArray<PatientListViewModel>();
	public patientListSource = new ObservableArray<PatientListViewModel>();
	public isBusy = true;
	private layout: ListViewLinearLayout;
	// >> grouping 
	public _funcGrouping: (item: PatientListViewModel) => PatientListViewModel;

	constructor(private routerExtensions: RouterExtensions,
		private patientListService: PatientListService,
		private page: Page,
		private passdataservice: PassDataService) {
			console.log("home");
			this._funcGrouping = (item: any) => {
				return item.dbmodel.sp_name;
			};
	}

	get _patientListItems(): ObservableArray<PatientListViewModel> {
		return this.patientListItems;
	}

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";

		this.patientListSource = new ObservableArray<PatientListViewModel>();
		
			this.getPatientListData();
			//this.addMoreItemsFromSource(20);
		

		// this.getPatientListData();

		console.log('init completed');
	}

	public sBLoaded(args) {
		var searchbar: SearchBar = <SearchBar>args.object;
		if (isAndroid) {
			searchbar.android.clearFocus();
		}
	}

	//  fucntion for view patient details.
	details(listItem) {
		console.log(listItem);
		// assigning data to service object.
		this.passdataservice.setPatientData(listItem);
		this.routerExtensions.navigate(["/patientmgnt"], { clearHistory: true });
	}// end of code block.
	camerasdetails() {
		this.routerExtensions.navigate(["/patientmgnt/cameras"], { clearHistory: true });
	}

	onTextChange(event) {

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

	public getPatientListData() {
		this.patientListService.getData().then(
		  (val) => {
				this.isBusy = false;
				val.forEach(item => {
					const patientListItem = new PatientListViewModel();
					patientListItem.dbmodel = item;
					this.patientListSource.push(patientListItem);
					this.patientListItems.push(patientListItem);
					// console.log("patient list",this.patientListSource);
					// this.patientListSource = new ObservableArray(patientListItem);
				});
			},
			(error) => {
				console.log("patientListService error:", error);
			}

		);
	}



	// search record by list code start
	public onSubmit(args) {

		let searchBar = <SearchBar>args.object;
		let searchValue = searchBar.text.toLowerCase();
		this.patientListSource = new ObservableArray<PatientListViewModel>()

		if (searchValue !== "") {

			this.patientListItems.forEach(item => {
				if (item.dbmodel.fname.toLowerCase().indexOf(searchValue) !== -1 || item.dbmodel.lname.toLowerCase().indexOf(searchValue) !== -1 || item.dbmodel.bed_no.toLowerCase().indexOf(searchValue) !== -1 || item.dbmodel.sp_name.toLowerCase().indexOf(searchValue) !== -1) {
					this.patientListSource.push(item);
				}
			});
		}

	}
	// search record by list code end

	//clear search record list then show all list record code start
	public onClear(args) {
		let searchBar = <SearchBar>args.object;
		searchBar.text = "";
		this.patientListSource = new ObservableArray<PatientListViewModel>();
		this.patientListItems.forEach(item => {
			this.patientListSource.push(item);
		});
	}
	//clear search record list then show all list record code end

}
