import { Component, OnInit, ViewChild} from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import * as Toast from 'nativescript-toast';
import { ListViewEventData } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { isAndroid } from 'tns-core-modules/platform';
import * as appSettings from "tns-core-modules/application-settings";
import { ObservableArray } from 'tns-core-modules/data/observable-array';
import { SearchBar } from 'tns-core-modules/ui/search-bar';
import { Switch } from 'tns-core-modules/ui/switch/switch';
import { ApiRequestModel, DataDBModel, FilterRequest, MonitoredRequest, UiViewModel } from '~/app/models/ui/patient-monitore-unmonitore-model';
import { ServerApiInterfaceService } from '~/app/services/server-api-interface.service';
import { AppRepoService } from '~/app/services/app-repo.service';
var dialogs = require("tns-core-modules/ui/dialogs");

@Component({
	moduleId: module.id,
	selector: 'patient-monitore-unmonitore-list',
	templateUrl: './patient-monitore-unmonitore-list.component.html',
	styleUrls: ['./patient-monitore-unmonitore-list.component.css']
})

export class PatientMonitoreUnmonitoreListComponent implements OnInit {
	// >> server list
	PATIENT_STATUS_MONITORING = 1;
	_dataItems = new ObservableArray<UiViewModel>();
	searchValue = "";
	isBusy = true;
	isDisabled = false;
	@ViewChild("patientListview", {static: false}) listViewComponent: RadListViewComponent;

	public funcGroupingFilter: (item: DataDBModel) => DataDBModel;

	constructor(private serverApiInterfaceService: ServerApiInterfaceService,
		private routerExtensions: RouterExtensions
		) {
		this.funcGroupingFilter = (item: any) => {
			if (item) {
				return item.dbmodel.spname;
			}
		};
	}


	ngOnInit() {
		this.getDataFormServerApi();
	}


	// single item toggle button 
	public onToggleSingleItem(args, item: any) {
		var selectedItem = item;
		let firstSwitch = <Switch>args.object;


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

	// location toggle button bind value 
	groupingSwitch(value) {
		const filterLocationItem = this._dataItems.filter(a => a.dbmodel.spname === value);
		let bindFilterAllGrouping = this._dataItems.filter(data => data.dbmodel.upmmidspid == null && data.dbmodel.upmmidpatientid == null);
		if (bindFilterAllGrouping.length == this._dataItems.length) {
			bindFilterAllGrouping.forEach((item) => {
				item.isDisabled = true;
			});
			this.isDisabled = true;
			return true;
		} else {
			if (filterLocationItem.length > 0) {
				let bindFilterItemLoactionId = filterLocationItem.filter(data => data.dbmodel.upmmidspid !== null && data.dbmodel.upmmidpatientid == null);
				if (bindFilterItemLoactionId.length == filterLocationItem.length) {
					bindFilterItemLoactionId.forEach((item) => {
						item.isDisabled = true;
					});

					this.isDisabled = false;
					return true;
				}
				else {
					this.isDisabled = false;
					return false;
				}
			}
			return false;
		}

	}

	// error msg funcation
	getErrorMsg(error) {
		let msg: string;
		if (error && error.handled == false) {
			msg = "Operation failed. Please check your internet connection.";
		} else {
			msg = "Server error occured please try after some time.";
		}
		return msg;
	}

	// pull to list then call getDataFormServerApi funcation
	public onPullToRefreshInitiatedGrid(args: ListViewEventData) {
		const that = new WeakRef(this);
		setTimeout(function () {
			that.get().getDataFormServerApi();
			const listView = args.object;
			listView.notifyPullToRefreshFinished();
		}, 1000);
	}


	// >> server list
	getDataFormServerApi() {
		this.isBusy = true;
		// after click button then change list mode
		const request = new ApiRequestModel();
		request.filter = new FilterRequest();
		this._dataItems = new ObservableArray<UiViewModel>();
		if (this.searchValue != "") {
			request.filter.fname = this.searchValue;
			request.filter.lname = this.searchValue;
			request.filter.patientregno = this.searchValue;
			request.filter.bedno = this.searchValue;
			request.filter.spname = this.searchValue;
		}
		request.page = 1;
		request.limit = 300;
		request.orderby = "fname";
		request.orderdirection = "asc";

		// console.log('test', request);
		this.serverApiInterfaceService.get(AppRepoService.Instance.API_APP_BASE_URL + "/v1/endpoint/list/patient", request)
			.then((result) => {
				let data: any = result;
				data.records.forEach(element => {
					let viewModel = new UiViewModel();
					viewModel.dbmodel = new DataDBModel();
					viewModel.dbmodel = element;
					if (element.monitored === this.PATIENT_STATUS_MONITORING) {
						viewModel.checked = true;
					} else {
						viewModel.checked = false;
					}
					this.isBusy = false;
					this._dataItems.push(viewModel);
				});

			}, (error) => {
				this.isBusy = false;
				this.getErrorMsg(error)
				var toast = Toast.makeText(this.getErrorMsg(error));
				toast.show();
			});
	}

	apiUpdateMonitorData(isMonitor: boolean, isGroup: boolean, selectedItem: UiViewModel, onResult: any) {
		this.isBusy = true;
		var requestURL = (isMonitor == true) ? '/v1/endpoint/user/associatepatient' : '/v1/endpoint/user/deassociatepatient';

		const monitoredRequest = new MonitoredRequest();
		monitoredRequest.usrid = appSettings.getNumber("USER_ID");

		monitoredRequest.spid = selectedItem.dbmodel.spid;

		if (isGroup == false) {
			monitoredRequest.patientid = selectedItem.dbmodel.patientid;
		}
		console.log('monitoredRequest', monitoredRequest);
		this.serverApiInterfaceService.post<any>(AppRepoService.Instance.API_APP_BASE_URL + requestURL, monitoredRequest).then(
			(success) => {
				onResult(true, 0, "");
				this.isBusy = false;
			}, (error) => {
				console.log('POST Request is Failed', error);
				this.isBusy = false;
				if (error && error.handled == false) {
					onResult(false, 1, "Operation failed. Please check your internet connection.");
					return;
				}

				console.log(`Server returned failure. Error Code : ${error.code}`);
				onResult(false, error.code, "Server error occured please try after some time.");
			});


	}

}