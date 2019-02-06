import { Component, OnInit, ViewChild } from '@angular/core';
import { ActionListViewModel, UserAuthDBRequest, UserCreateFormRequest } from '~/app/models/ui/action-models';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { ActionService } from '~/app/services/action/action.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
	moduleId: module.id,
	selector: 'user_auth',
	templateUrl: './user_auth.component.html',
	styleUrls: ['./user_auth.component.css']
})

export class UserAuthComponent implements OnInit {

	userAuthForm: FormGroup;
	emailIsValid: boolean;
	passwordIsValid: boolean;

	userAuthPinForm: FormGroup;
	newpinIsValid: boolean;
	re_enterpinIsValid: boolean;

	userPinCheckForm: FormGroup;
	userpinIsValid: boolean;

	login = true;
	listaccount = true;
	removeAccount = false;
	pinview = true;
	pincheck: any;
	newpinview = false;
	hidecheckpin = true;
	datamodel = new UserAuthDBRequest();

	_dataItemsaccount = new ObservableArray<ActionListViewModel>();
	constructor(private actionService: ActionService) {
		console.log('user auth');
		this.datamodel = new UserAuthDBRequest();
	}
	get dataItems(): ObservableArray<ActionListViewModel> {
		return this._dataItemsaccount;
	}
	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		// this.layout = new ListViewLinearLayout();
		// this.layout.scrollDirection = "Vertical";
		this.createFormUserAuthControls();
		this.createFormUserPinControls();
		this.userPinCheckControls();
		this.getUserAccountData();


	}
	createFormUserAuthControls(): void {
		this.userAuthForm = new FormGroup({
			email: new FormControl('', [Validators.required]),
			password: new FormControl('', [Validators.required]),
		});
	}
	createFormUserPinControls(): void {
		this.userAuthPinForm = new FormGroup({
			newpin: new FormControl('', [Validators.required]),
			re_enterpin: new FormControl('', [Validators.required]),
		});
	}
	userPinCheckControls(): void {
		this.userPinCheckForm = new FormGroup({
			userpin: new FormControl('', [Validators.required])
		});
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
		this.userpinIsValid = this.userPinCheckForm.controls['userpin'].hasError('required');
		if (this.userPinCheckForm.invalid) {
			console.log("validation error");
			return;
		}
		const formmodel = new UserAuthDBRequest();
		formmodel.pin = this.userPinCheckForm.get('userpin').value;
		if (formmodel.pin == this.pincheck) {
			console.log('this.pin Right', formmodel.pin);
		} else {
			console.log('this.pin wrong', formmodel.pin);
		}
	}
	userAuthAccount() {
		this.emailIsValid = this.userAuthForm.controls['email'].hasError('required');
		this.passwordIsValid = this.userAuthForm.controls['password'].hasError('required');
		if (this.userAuthForm.invalid) {
			console.log("validation error");
			return;
		}
		const formmodel = new UserCreateFormRequest();
		formmodel.email = this.userAuthForm.get('email').value;
		formmodel.password = this.userAuthForm.get('password').value;
		this.newpinview = true;
		this.hidecheckpin = false;
		// this.datamodel = new  UserAuthDBRequest();
		console.log('this.datamodel.user_fname', formmodel.email = this.userAuthForm.get('email').value);
		console.log('this.datamodel.user_lname', formmodel.password = this.userAuthForm.get('password').value);

	}
	userSign() {
		this.newpinIsValid = this.userAuthPinForm.controls['newpin'].hasError('required');
		this.re_enterpinIsValid = this.userAuthPinForm.controls['re_enterpin'].hasError('required');
		if (this.userAuthPinForm.invalid) {
			console.log("validation error");
			return;
		}
		const formmodel = new UserCreateFormRequest();
		formmodel.newpin = this.userAuthForm.get('newpin').value;
		formmodel.reenterpin = this.userAuthForm.get('re_enterpin').value;
	}
}