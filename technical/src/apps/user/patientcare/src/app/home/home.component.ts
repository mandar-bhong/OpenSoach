import { Component, NgZone, OnDestroy, OnInit, ViewChild, ChangeDetectorRef } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { ListViewLinearLayout } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular';
import { isAndroid } from 'platform';
import { Subscription } from 'rxjs';
import { ObservableArray } from 'tns-core-modules/data/observable-array';
import { SearchBar } from 'tns-core-modules/ui/search-bar';
import { layout } from 'tns-core-modules/utils/utils';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { PassDataService } from '~/app/services/pass-data-service';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';

import { DataListingInterface } from '../data-listing-interface';
import { NextActionService } from '../services/action/next-action-service';
import { JSONBaseDataModel } from '../models/ui/json-base-data-model';
import { PersonAccompanyModel } from '../models/ui/person-accompany-model';
import { ACTION_STATUS, APP_MODE } from '../app-constants';
import { RadSideDrawerComponent } from "nativescript-ui-sidedrawer/angular";
import { RadSideDrawer } from 'nativescript-ui-sidedrawer';
import * as appSettings from "tns-core-modules/application-settings";
import { AppGlobalContext } from '../app-global-context';
import { DatabaseHelper } from '../helpers/database-helper';
import { screen } from "tns-core-modules/platform/platform"


@Component({
	selector: "Home",
	moduleId: module.id,
	templateUrl: "./home.component.html",
	styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit, OnDestroy, DataListingInterface<PatientListViewModel> {
	listSource = new Array<PatientListViewModel>();
	listItems = new ObservableArray<PatientListViewModel>();

	public isBusy = true;
	private layout: ListViewLinearLayout;
	appMode: number;
	appModeEnum = APP_MODE;
	@ViewChild("myListView") listViewComponent: RadListViewComponent;
	// view child 
	@ViewChild(RadSideDrawerComponent) public drawerComponent: RadSideDrawerComponent;
	private drawer: RadSideDrawer;
	private _mainContentText: string;
	// end 
	// >> grouping 
	public _funcGrouping: (item: PatientListViewModel) => PatientListViewModel;


	// serach 
	searchValue = "";
	patientListChanged: Subscription;
	patientListItemMaster = new PatientListViewModel();

	jsonField;
	ACTION_STATUS = ACTION_STATUS;
	constructor(private routerExtensions: RouterExtensions,
		private patientListService: PatientListService,
		private passdataservice: PassDataService,
		private _changeDetectionRef: ChangeDetectorRef,
		private ngZone: NgZone,
		private nextActionService: NextActionService) {
		this._funcGrouping = (item: any) => {
			if (item) {
				return item.dbmodel.sp_name;
			}
		};
		this.patientListChanged = this.patientListService.patientListChangedSubject.subscribe((listItem) => {
			this.onDataReceived(listItem);
		});


	}


	ngOnInit() {
		this.appMode = appSettings.getNumber("APP_MODE", APP_MODE.NONE);
		console.log('home component init');
		this.layout = new ListViewLinearLayout();
		this.layout.scrollDirection = "Vertical";
		this.getData();
		this.jsonField = new JSONBaseDataModel<PersonAccompanyModel[]>();

	}

	bindList() {
		console.log('bindList');
		this.listItems = new ObservableArray<PatientListViewModel>();
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
			// console.log('this.listItems', this.listItems);
		}
	}

	getData() {
		this.patientListService.getData().then(
			(val) => {
				this.isBusy = false;
				val.forEach(item => {
					this.listSource.push(item);
				});

				this.getNextActionTimeForAll();
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
			this._funcGrouping = (item: any) => {
				if (item) {
					return item.dbmodel.sp_name;
				}
			};
			// console.log('on data received in home');
			items.forEach(item => {

				if (item.deleteuuid) {
					this.listSource = this.listSource.filter(e => e.dbmodel.admission_uuid != item.deleteuuid);
					this.bindList();
				} else {
					const existingItem = this.listSource.find(a => a.dbmodel.admission_uuid === item.dbmodel.admission_uuid);
					if (existingItem) {
						Object.assign(existingItem, item);
					}
					else {
						this.listSource.push(item);
						if (this.searchValue == "") {
							this.listItems.push(item);
						}
					}

					// associate nextActionItems from NextActionServiceMap
					const nextActionItems = this.nextActionService.nextActionTimesMap.get(item.dbmodel.admission_uuid);
					if (nextActionItems) {
						console.log('setting next action items', nextActionItems);
						item.nextActionTimes = nextActionItems;
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
		if (this.searchValue != "") {
			this.searchValue = "";
			this.bindList();
		}
	}

	details(listItem) {
		// console.log(listItem);
		this.passdataservice.setPatientData(listItem);
		// console.log('listItem',listItem);
		this.routerExtensions.navigate(["patientmgnt"], { clearHistory: false });
	}

	// clean up
	ngOnDestroy(): void {
		if (this.patientListChanged) {
			this.patientListChanged.unsubscribe();
		}
	}

	getNextActionTimeForAll() {
		this.nextActionService.getNextActionsForAllPatients().then(result => {
			this.listSource.forEach(item => {
				const nextActionTimes = result.get(item.dbmodel.admission_uuid);
				if (nextActionTimes) {
					item.nextActionTimes = nextActionTimes;
					//console.log('action times for admission', item);
				}
			});
		}, error => {

		});
		this.nextActionService.nextActionMapChanged.subscribe(entry => {
			const viewModel = this.listSource.find(a => a.dbmodel.admission_uuid == entry.admission_uuid);
			if (viewModel) {
				//console.log('setting next action items receieved from action service notification', entry);
				viewModel.nextActionTimes = entry.nextActionTimes;
			}
		});
	}

	//  drawing side bar fucntional
	get mainContentText() {
		return this._mainContentText;
	}

	set mainContentText(value: string) {
		this._mainContentText = value;
	}

	public openDrawer() {
		this.drawer.toggleDrawerState();
	}

	ngAfterViewInit() {
		this.drawer = this.drawerComponent.sideDrawer;
		this._changeDetectionRef.detectChanges();
	}
	//log out 
	async logout() {
		AppGlobalContext.Token = null;
		AppGlobalContext.WebsocketUrl = null;
		appSettings.clear();
		const deleted = await this.deleteUsers();
		if (deleted) {
			this.routerExtensions.navigate(['login'], { clearHistory: true });
		}

		console.log('logout executed');
	}
	// fucntion for delete user from database.
	deleteUsers(): Promise<any> {
		return new Promise((resolve, reject) => {
			DatabaseHelper.selectAll('deleteuser').then((success) => {
				console.log('success', success);
				resolve(success);
			}, (error) => {
				console.log('getuser response Failed', error);
				reject(error);
			});
		});
	}


	getServerList() {
		this.routerExtensions.navigate(['home', 'monitore'], { clearHistory: false });
	}// end of fucntions.
}
