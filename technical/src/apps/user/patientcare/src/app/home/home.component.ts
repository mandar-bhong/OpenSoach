import { Component, NgZone, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { ListViewLinearLayout } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular';
import { isAndroid } from 'platform';
import { Subscription } from 'rxjs';
import { ObservableArray } from 'tns-core-modules/data/observable-array';
import { SearchBar } from 'tns-core-modules/ui/search-bar';
import { layout } from 'tns-core-modules/utils/utils';
import { Page } from 'ui/page';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { PassDataService } from '~/app/services/pass-data-service';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';

import { SERVER_WORKER_MSG_TYPE, SYNC_STORE } from '../app-constants';
import { DataListingInterface } from '../data-listing-interface';
import { ServerDataProcessorMessageModel } from '../models/api/server-data-processor-message-model';
import { ServerDataStoreDataModel } from '../models/api/server-data-store-data-model';
import { PatientAdmissionDatastoreModel } from '../models/db/patient-admission-model';
import { PatientMasterDatastoreModel } from '../models/db/patient-master-model';
import { WorkerService } from '../services/worker.service';

@Component({
	selector: "Home",
	moduleId: module.id,
	templateUrl: "./home.component.html",
	styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit, DataListingInterface<PatientListViewModel> {
	listSource = new Array<PatientListViewModel>();
	listItems = new ObservableArray<PatientListViewModel>();

	public isBusy = true;
	private layout: ListViewLinearLayout;

	// >> grouping 
	public _funcGrouping: (item: PatientListViewModel) => PatientListViewModel;


	// serach 
	searchValue = "";
	patientListChanged: Subscription;
	patientListItemMaster = new PatientListViewModel();

	constructor(private routerExtensions: RouterExtensions,
		private patientListService: PatientListService,
		private page: Page,
		private passdataservice: PassDataService,
		private workerService: WorkerService,
		private ngZone: NgZone) {
		console.log("home");
		this._funcGrouping = (item: any) => {
			if (item) {
				return item.dbmodel.sp_name;
			}
		};

		this.patientListChanged = this.patientListService.patientListChangedSubject.subscribe((listItem) => {
			this.onDataReceived(listItem);
		});
	}


	get _patientListItems(): ObservableArray<PatientListViewModel> {
		return this.listItems;
	}

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this.getData();
		console.log('init completed');

		// setTimeout(() => {
		// 	const patientMasterAdd = new ServerDataProcessorMessageModel();
		// 	patientMasterAdd.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;

		// 	const masterModel = new PatientMasterDatastoreModel();
		// 	masterModel.uuid = "PM003";
		// 	masterModel.patient_reg_no = "P12B12223";
		// 	masterModel.fname = "xyz",
		// 		masterModel.lname = "Lunia",
		// 		masterModel.mob_no = "9832345333",
		// 		masterModel.age = "28";
		// 	masterModel.blood_grp = "B+ve";
		// 	masterModel.gender = "Male";
		// 	masterModel.updated_on = new Date;
		// 	masterModel.sync_pending = 0;
		// 	const serverDataStoreDataModelForMaster = new ServerDataStoreDataModel();
		// 	serverDataStoreDataModelForMaster.datastore = SYNC_STORE.PATIENT_MASTER;
		// 	serverDataStoreDataModelForMaster.data = masterModel;

		// 	patientMasterAdd.data = [serverDataStoreDataModelForMaster];
		// 	this.workerService.postMessageToServerDataProcessorWorker(patientMasterAdd);

		// 	const patientAdmissionAdd = new ServerDataProcessorMessageModel();
		// 	patientAdmissionAdd.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;

		// 	const admissionModel = new PatientAdmissionDatastoreModel();
		// 	admissionModel.uuid = "PA003";
		// 	admissionModel.patient_reg_no = "P12B12223";
		// 	admissionModel.patient_uuid = "PM003";
		// 	admissionModel.bed_no = "2A/666";
		// 	// admissionModel.uuid = "PA013";
		// 	// admissionModel.patient_uuid = "PM013";
		// 	// admissionModel.patient_reg_no = "P12B12223";
		// 	// admissionModel.bed_no = "2A/888";
		// 	admissionModel.status = "1";
		// 	admissionModel.sp_uuid = "SP002";
		// 	admissionModel.dr_incharge = 1;
		// 	admissionModel.admitted_on = new Date;
		// 	admissionModel.discharged_on = new Date;
		// 	admissionModel.updated_on = new Date;
		// 	admissionModel.sync_pending = 0;
		// 	const serverDataStoreDataModelForAdmission = new ServerDataStoreDataModel();
		// 	serverDataStoreDataModelForAdmission.datastore = SYNC_STORE.PATIENT_ADMISSION;
		// 	serverDataStoreDataModelForAdmission.data = admissionModel;

		// 	patientAdmissionAdd.data = [serverDataStoreDataModelForAdmission];
		// 	this.workerService.postMessageToServerDataProcessorWorker(patientAdmissionAdd);
		// }, 15000);
	}

	bindList() {
		console.log('bindList');
		// this.listItems = new ObservableArray<PatientListViewModel>();
		if (this.searchValue != "") {
			this.listSource.forEach(item => {
				if (item.dbmodel.fname.toLowerCase().indexOf(this.searchValue) !== -1 || item.dbmodel.lname.toLowerCase().indexOf(this.searchValue) !== -1 || item.dbmodel.bed_no.toLowerCase().indexOf(this.searchValue) !== -1 || item.dbmodel.sp_name.toLowerCase().indexOf(this.searchValue) !== -1) {
					this.listItems.push(item);
				}
			});
		} else {
			this.listSource.forEach(item => {
				this.listItems.push(item);

			});
			console.log('this.listItems', this.listItems);
		}
		// if condition check key
		// search
		// else bind whole list
	}

	getData() {
		this.patientListService.getData().then(
			(val) => {
				this.isBusy = false;
				val.forEach(item => {
					// console.log("val", val);
					const patientListItem = new PatientListViewModel();
					patientListItem.dbmodel = item;
					this.listSource.push(patientListItem);

				});
				this.bindList();
			},
			(error) => {
				console.log("patientListService error:", error);
			}

		);
	}

	onDataReceived(items: PatientListViewModel[]) {

		this.ngZone.run(() => {
			// check if this item exists in listSource by admission_uuid
			console.log('on data received in home');
			items.forEach(item => {
				const existingItems = this.listSource.filter(a => a.dbmodel.admission_uuid === item.dbmodel.admission_uuid);
				if (existingItems.length > 0) {
					const lenght = this.listSource.length;
					// console.log('listsoure lenght', lenght);
					const index = this.listSource.indexOf(existingItems[0]);
					console.log(' index', index);
					this.listSource[index].dbmodel = item.dbmodel;
					// console.log('received item data', item);
				}
				else {
					// console.log('else condition new item add', item);
					this.listSource.push(item);
					if (this.searchValue == "")
					{
					this.listItems.push(item);
					}
				}
			});
			if (this.searchValue != "") {
				this.bindList();
			}
		});

	}

	public sBLoaded(args) {
		var searchbar: SearchBar = <SearchBar>args.object;
		if (isAndroid) {
			searchbar.android.clearFocus();
		}
	}


	public onSubmit(args) {
		let searchBar = <SearchBar>args.object;
		console.log('searchBar', searchBar);
		this.searchValue = searchBar.text.toLowerCase();
		console.log('searchValue ', this.searchValue);
		this.bindList();
	}

	public onClear(args) {
		let searchBar = <SearchBar>args.object;
		searchBar.text = "";
		this.searchValue = "";
		this.bindList();

	}

	details(listItem) {
		console.log(listItem);
		this.passdataservice.setPatientData(listItem);
		this.routerExtensions.navigate(["patientmgnt"], { clearHistory: false });
	}



}
