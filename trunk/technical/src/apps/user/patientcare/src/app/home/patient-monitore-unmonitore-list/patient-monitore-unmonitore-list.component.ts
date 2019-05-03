import { Component, NgZone, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { CFAlertActionAlignment, CFAlertActionStyle, CFAlertDialog, CFAlertStyle, DialogOptions } from "nativescript-cfalert-dialog";
import { isAndroid } from 'platform';
import * as appSettings from "tns-core-modules/application-settings";
import { ObservableArray } from 'tns-core-modules/data/observable-array';
import { SearchBar } from 'tns-core-modules/ui/search-bar';
import { Switch } from 'tns-core-modules/ui/switch/switch';
import { API_APP_BASE_URL } from '~/app/app-constants';
import { ServerApiInterfaceService } from '~/app/services/server-api-interface.service';
import { DataDBModel, UiViewModel, ApiRequestModel, FilterRequest, MonitoredRequest } from '~/app/models/ui/patient-monitore-unmonitore-model';
import * as Toast from 'nativescript-toast';
import { ListViewEventData } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';


@Component({
	moduleId: module.id,
	selector: 'patient-monitore-unmonitore-list',
	templateUrl: './patient-monitore-unmonitore-list.component.html',
	styleUrls: ['./patient-monitore-unmonitore-list.component.css']
})

export class PatientMonitoreUnmonitoreListComponent implements OnInit {
	// >> server list
	_dataItems = new ObservableArray<UiViewModel>();
	data = new ObservableArray<UiViewModel>();
	mainPatientList: boolean = true;
	monitraPatientList: boolean = false;

	private cfalertDialog: CFAlertDialog;
	passGroupingSpid: number;
	passSingleSpid: number;
	passSinglePatientid: number;
	alertMeg: string;
	alertFuncation: boolean;
	toggleOptionSelection: boolean;
	clickSingleItemData: boolean;

	getToggleGroupingSpanme: string;
	args: any;
	//selectedItem: any;
	// >> server list
	public _funcGroupingServer: (item: DataDBModel) => DataDBModel;
	switchValue: boolean;
	// <<

	searchValue = "";
	constructor(private serverApiInterfaceService: ServerApiInterfaceService,
		private routerExtensions: RouterExtensions,
		private ngZone: NgZone) {
		// >> server list
		this._funcGroupingServer = (item: any) => {
			if (item) {
				return item.dbmodel.spname;
			}
		};
		this.cfalertDialog = new CFAlertDialog();
		// << 
	}


	ngOnInit() {
		this.getDataFormServerApi();
	}
	// >> server list
	getDataFormServerApi() {
		// after click button then change list mode
		this.mainPatientList = false;
		this.monitraPatientList = true;
		const request = new ApiRequestModel();
		request.filter = new FilterRequest();
		this.data = new ObservableArray<UiViewModel>();
		if (this.searchValue != "") {
			request.filter.fname = this.searchValue;
			request.filter.lname = this.searchValue;
			request.filter.patientregno = this.searchValue;
			request.filter.bedno = this.searchValue;
			request.filter.spname = this.searchValue;
			// console.log('api filter value___________', request.filter);
		}
		request.page = 1;
		request.limit = 10;
		request.orderby = "fname";
		request.orderdirection = "asc";

		// console.log('test', request);
		this.serverApiInterfaceService.get(API_APP_BASE_URL + "/v1/endpoint/list/patient", request)
			.then((result) => {
				let data: any = result;
				data.records.forEach(element => {
					let viewModel = new UiViewModel();
					viewModel.dbmodel = new DataDBModel();
					viewModel.dbmodel = element;
					if (element.monitored === 1) { /// 1 - patient is in monitoring
						viewModel.checked = true;
						viewModel.isGrouping = false;
					} else {
						viewModel.checked = false;
						viewModel.isGrouping = false;
					}
					this.data.push(viewModel);
				});
				// console.log('data.records', data.records);

			}, (error) => {
				if (!error.handled) {
					//TODO: error condition
				}
			});
		this.toggleListChange();
	}
	// get data pass to binding array 
	public toggleListChange() {
		this._dataItems = this.data;
	}

	// single item toggle button 
	public onToggleSingleItem(args, item: any) {
		var selectedItem = item;
		let firstSwitch = <Switch>args.object;


		var dialogs = require("tns-core-modules/ui/dialogs");
		let confirmationMsg = selectedItem.checked ? "Do you want to Unmonitored?" : "Do you want to monitored?";

		let options = {
			title: "Confirmation",
			message: confirmationMsg,
			okButtonText: "Yes",
			cancelButtonText: "Cancel"

		};

		dialogs.confirm(options).then((result: boolean) => {

			if (result == false) {
				firstSwitch.checked = selectedItem.checked = !firstSwitch.checked;
				return;
			}

			this.apiUpdateMonitorData(firstSwitch.checked, false, selectedItem, function (isSuccess: boolean, errorCode: number, errorMsg: string) {

				if (isSuccess == false) {
					firstSwitch.checked = selectedItem.checked = !firstSwitch.checked;
					var toast = Toast.makeText(errorMsg);
					toast.show();
					return;
				}

			});
		});

	}



	// grouping toggle button get data
	public onToggleLocation(args, spname) {
		const filterItemList = this._dataItems.filter(a => a.dbmodel.spname === spname);

		if (filterItemList.length == 0) {
			console.log("Filter item list have 0 items");
			return;
		}

		var selectedItem = filterItemList[0];

		let firstSwitch = <Switch>args.object;
		selectedItem.checked = firstSwitch.checked;

		console.log(`Group switch : ${firstSwitch}`);

		var dialogs = require("tns-core-modules/ui/dialogs");
		let confirmationMsg = selectedItem.checked ? "Do you want to Unmonitor?" : "Do you want to monitored?";

		let options = {
			title: "Confirmation",
			message: confirmationMsg,
			okButtonText: "Yes",
			cancelButtonText: "Cancel"

		};

		dialogs.confirm(options).then((result: boolean) => {

			if (result == false) {
				firstSwitch.checked = selectedItem.checked = !firstSwitch.checked;
				return;
			}
			this.apiUpdateMonitorData(firstSwitch.checked, true, selectedItem, function (isSuccess: boolean, errorCode: number, errorMsg: string) {

				if (isSuccess == false) {
					firstSwitch.checked = selectedItem.checked = !firstSwitch.checked;
					var toast = Toast.makeText(errorMsg);
					toast.show();
					return;
				}
				if (firstSwitch.checked) {
					filterItemList.forEach(element => {
						element.isDisabled = true;
						element.checked = true;						
					});
				} else {
					filterItemList.forEach(element => {
						element.isDisabled = false;
						element.checked = false;
					});
				}

			});

		});
	}
	// back to show main patient list view
	goToMainList() {
		this.routerExtensions.back();
	}

	apiUpdateMonitorData(isMonitor: boolean, isGroup: boolean, selectedItem: UiViewModel, onResult: any) {

		var requestURL = (isMonitor == true) ? '/v1/endpoint/user/associatepatient' : '/v1/endpoint/user/deassociatepatient';

		const monitoredRequest = new MonitoredRequest();
		monitoredRequest.usrid = appSettings.getNumber("USER_ID");

		monitoredRequest.spid = selectedItem.dbmodel.spid;

		if (isGroup == false) {
			monitoredRequest.patientid = selectedItem.dbmodel.patientid;
		}
		console.log('monitoredRequest', monitoredRequest);
		this.serverApiInterfaceService.post<any>(API_APP_BASE_URL + requestURL, monitoredRequest).then(
			(success) => {
				onResult(true, 0, "");
			}, (error) => {
				console.log('POST Request is Failed', error);

				if (error && error.handled == false) {
					onResult(false, 1, "Operation failed. Please check your internet connection.");
					return;
				}

				console.log(`Server returned failure. Error Code : ${error.code}`);
				onResult(false, error.code, "Server error occured please try after some time.");
			});


	}


	public onSubmitServer(args) {
		let searchBar = <SearchBar>args.object;
		this.searchValue = searchBar.text.toLowerCase();
		console.log('searchValue ', this.searchValue);
		// this.bindList();
		if (this.searchValue != "") {
			this.getDataFormServerApi();
		}
	}

	public onClearServer(args) {
		let searchBar = <SearchBar>args.object;
		searchBar.text = "";
		if (this.searchValue != "") {
			this.searchValue = "";
			this.getDataFormServerApi();
		}
	}
	public sBLoaded(args) {
		var searchbar: SearchBar = <SearchBar>args.object;
		if (isAndroid) {
			searchbar.android.clearFocus();
		}
	}

}