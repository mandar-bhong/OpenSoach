import { Component, OnInit, ViewChild } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { LocalNotifications } from 'nativescript-local-notifications';
import { ListViewLinearLayout } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular';
import { isAndroid } from 'platform';
import { ObservableArray } from 'tns-core-modules/data/observable-array';
import { alert } from 'tns-core-modules/ui/dialogs';
import { SearchBar } from 'tns-core-modules/ui/search-bar';
import { layout } from 'tns-core-modules/utils/utils';
import { Page } from 'ui/page';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { PassDataService } from '~/app/services/pass-data-service';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { DataListingInterface } from '../data-listing-interface';
import { Subscription } from 'rxjs';
import { ServerDataProcessorMessageModel } from '../models/api/server-data-processor-message-model';
import { SERVER_WORKER_MSG_TYPE, SYNC_STORE } from '../app-constants';
import { PatientMasterDatastoreModel } from '../models/db/patient-master-model';
import { ServerDataStoreDataModel } from '../models/api/server-data-store-data-model';
import { WorkerService } from '../services/worker.service';
import { PatientAdmissionDatastoreModel } from '../models/db/patient-admission-model';

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
		private workerService: WorkerService) {
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


	// get _patientListItems(): ObservableArray<PatientListViewModel> {
	// 	return this.listItems;
	// }

	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this.getData();
		console.log('init completed');

		// setTimeout(() => {
		// 	// const patientMasterAdd = new ServerDataProcessorMessageModel();
		// 	// patientMasterAdd.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;

		// 	// const masterModel = new PatientMasterDatastoreModel();
		// 	// masterModel.uuid = "PM003";
		// 	// masterModel.patient_reg_no = "P12B12213";
		// 	// masterModel.fname = "Shubham",
		// 	// 	masterModel.lname = "Lunia",
		// 	// 	masterModel.mob_no = "9832345333",
		// 	// 	masterModel.age = "28";
		// 	// masterModel.blood_grp = "B+ve";
		// 	// masterModel.gender = "Male";
		// 	// masterModel.updated_on = new Date;
		// 	// masterModel.sync_pending = 0;
		// 	// const serverDataStoreDataModelForMaster = new ServerDataStoreDataModel();
		// 	// serverDataStoreDataModelForMaster.datastore = SYNC_STORE.PATIENT_MASTER;
		// 	// serverDataStoreDataModelForMaster.data = masterModel;

		// 	// patientMasterAdd.data = [serverDataStoreDataModelForMaster];
		// 	// this.workerService.postMessageToServerDataProcessorWorker(patientMasterAdd);

		// 	const patientAdmissionAdd = new ServerDataProcessorMessageModel();
		// 	patientAdmissionAdd.msgtype = SERVER_WORKER_MSG_TYPE.SEND_MESSAGE;

		// 	const admissionModel = new PatientAdmissionDatastoreModel();
		// 	admissionModel.uuid = "PA003";
		// 	admissionModel.patient_uuid = "PM003";
		// 	admissionModel.patient_reg_no = "P12B12213";
		// 	admissionModel.bed_no = "2A/666";
		// 	admissionModel.status = "1";
		// 	admissionModel.sp_uuid = "SP001";
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
		this.listItems = new ObservableArray<PatientListViewModel>();;
		if (this.searchValue !== "") {
			this.listSource.forEach(item => {
				if (item.dbmodel.fname.toLowerCase().indexOf(this.searchValue) !== -1 || item.dbmodel.lname.toLowerCase().indexOf(this.searchValue) !== -1 || item.dbmodel.bed_no.toLowerCase().indexOf(this.searchValue) !== -1 || item.dbmodel.sp_name.toLowerCase().indexOf(this.searchValue) !== -1) {
					this.listItems.push(item);
				}
			});
		} else {
			this.listSource.forEach(item => {
				this.listItems.push(item);
			});

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
		// check if this item exists in listSource by admission_uuid

		// if(exits)
		//{
		// delete from listSource
		//}


		items.forEach(item => {
			const existingItems = this.listSource.filter(a => a.dbmodel.admission_uuid === item.dbmodel.admission_uuid)[0];
					const lenght = this.listSource.length;
			console.log('listsoure lenght', lenght);
			const index = this.listSource.indexOf(existingItems);
			// console.log(' index', index);

			if (existingItems) {
				// console.log('this.listSource[index].dbmodel', this.listSource[index].dbmodel);
				this.listSource[index].dbmodel = item.dbmodel;
				console.log('received item data', item.dbmodel);
				// console.log('admissionuuid', existingItems);
				// remove existingItems[0]
			}
			// this.bindList();
			// this.listSource.push(item);



		});

		// add item to listSource

		// if searchvalue= ""
		// add to listItems
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
