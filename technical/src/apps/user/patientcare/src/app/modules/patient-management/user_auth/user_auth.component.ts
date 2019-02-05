import { Component, OnInit, ViewChild } from '@angular/core';
import { ActionListViewModel, UserAuthDBRequest } from '~/app/models/ui/action-models';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { ActionService } from '~/app/services/action/action.service';

@Component({
	moduleId: module.id,
	selector: 'user_auth',
	templateUrl: './user_auth.component.html',
	styleUrls: ['./user_auth.component.css']
})

export class UserAuthComponent implements OnInit {

	login = true;
	listaccount = true;
	removeAccount = false;
	pinview = true;
	pin: any;
	pincheck: any;
	newpinview = false;
	hidecheckpin = true;
	datamodel = new  UserAuthDBRequest();

	_dataItemsaccount = new ObservableArray<ActionListViewModel>();
	constructor(private actionService: ActionService) {
		console.log('user auth');
	}
	get dataItems(): ObservableArray<ActionListViewModel> {
		return this._dataItemsaccount;
	}
	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		// this.layout = new ListViewLinearLayout();
		// this.layout.scrollDirection = "Vertical";
		this.getUserAccountData();
	}
	allreadyaccount() {
		this.listaccount = !this.listaccount;
	}
	removeaccount() {
		this.removeAccount = !this.removeAccount;
	}

	public getUserAccountData() {
		this.actionService.getUserAccountList().then(
			(val) => {
				val.forEach(item => {
					const actionListItems = new ActionListViewModel();
					actionListItems.dbmodel = item;
					this._dataItemsaccount.push(actionListItems);
					console.log('User Account List', item);

				});
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}
	selectedUserData(item) {
		console.log('item', item);

		this.pincheck = item.dbmodel.pin;
		this.pinview = !this.pinview;
	}
	onsubmitPin() {
		if (this.pin == this.pincheck) {
			console.log('this.pin Right', this.pin);
		} else {
			console.log('this.pin wrong', this.pin);
		}
	}
	createAccount(){
		// this.datamodel = new  UserAuthDBRequest();
		console.log('this.datamodel.user_fname', this.datamodel.user_fname);
		console.log('this.datamodel.user_lname', this.datamodel.user_lname);
		console.log('this.datamodel.email', this.datamodel.email);
		this.newpinview = true;
		this.hidecheckpin = false;
	}
}