import { Component, OnInit, ViewChild, Input } from '@angular/core';
import { ActionListViewModel, UserAuthDBRequest, UserCreateFormRequest } from '~/app/models/ui/action-models';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { ActionService } from '~/app/services/action/action.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { RouterExtensions } from 'nativescript-angular/router';
import { PassDataService } from '~/app/services/pass-data-service';
import { HttpClient } from "@angular/common/http";

export class UserData{
	data: UserDetails;
}
export class UserDetails{
	userid : string;
	firstname: string;
	lastname: string;
	username: string;
	pin: string;
}
export class UserDetailDBModel{
	userid : string;
	first_name: string;
	last_name: string;
	email: string;
	pin: string;
}
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

	// patientname: string;
	getuserdetails = new UserDetails();

	login = true;
	listaccount = true;
	removeAccount = false;
	pinview = true;
	pincheck: any;
	newpinview = false;
	hidecheckpin = true;
	datamodel = new UserAuthDBRequest();

	selecedItemsData: any;

	_dataItemsaccount = new ObservableArray<ActionListViewModel>();
	constructor(private actionService: ActionService,
		private routerExtensions: RouterExtensions,
		private passdataservice: PassDataService,
		private httpClient: HttpClient) {
		console.log('user auth');
		this.datamodel = new UserAuthDBRequest();
	}
	get dataItems(): ObservableArray<ActionListViewModel> {
		return this._dataItemsaccount;
	}
	@Input() patientName: string;
	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.createFormUserAuthControls();
		this.createFormUserPinControls();
		this.userPinCheckControls();
		this.getUserAccountData();


		this.patientName = this.passdataservice.getHeaderName();
		// console.log('patient name',this.patientName)
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
					// console.log('User Account List', item);

				});
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}
	selectedUserData(item) {
		console.log('item', item);

		this.selecedItemsData = item.dbmodel;
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
			this.passdataservice.authResultReuested.onDeviceAuthSuccess(this.selecedItemsData.userid);
			this.routerExtensions.back();

		} else {
			console.log('this.pin wrong', formmodel.pin);
		}
		console.log('on submit back');

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

		// this.datamodel = new  UserAuthDBRequest();
		console.log('this.datamodel.user_fname', formmodel.email = this.userAuthForm.get('email').value);
		console.log('this.datamodel.user_lname', formmodel.password = this.userAuthForm.get('password').value);

		console.log('token', this.passdataservice.token);

		if (formmodel.email && formmodel.password) {
			this.httpClient.post("http://172.105.232.148/api/v1/endpoint/userauthorization",
				{
					'username': formmodel.email,
					'password': formmodel.password,
					'devicetoken':  this.passdataservice.token,
				})
				.subscribe(
					(res: UserData) => {
						console.log("POST Request is successful ", res);
						this.getuserdetails.firstname = res.data.firstname;
						console.log('fname',this.getuserdetails.firstname);
						this.getuserdetails.lastname = res.data.lastname;
						this.getuserdetails.userid = res.data.userid;
						this.getuserdetails.username = res.data.username;

						this.newpinview = true;
						this.hidecheckpin = false;
					}, (error) => {
						console.log(error);
					}
				);
		}


	}
	userSign() {
		this.newpinIsValid = this.userAuthPinForm.controls['newpin'].hasError('required');
		this.re_enterpinIsValid = this.userAuthPinForm.controls['re_enterpin'].hasError('required');
		if (this.userAuthPinForm.invalid) {
			console.log("validation error");
			return;
		}
		const formmodel = new UserCreateFormRequest();
		formmodel.newpin = this.userAuthPinForm.get('newpin').value;
		formmodel.reenterpin = this.userAuthPinForm.get('re_enterpin').value;

		if(formmodel.newpin === formmodel.reenterpin){
			const saveuserdetails = new UserDetailDBModel();
			saveuserdetails.userid = this.getuserdetails.userid;
			saveuserdetails.first_name = this.getuserdetails.firstname;
			saveuserdetails.last_name = this.getuserdetails.lastname;
			saveuserdetails.email = this.getuserdetails.username;
			saveuserdetails.pin = formmodel.reenterpin;
			// save action done and discard in DB
			const userid = Number(saveuserdetails.userid)
			this.actionService.insertDeviceAccessItem(saveuserdetails);
			this.passdataservice.authResultReuested.onDeviceAuthSuccess(userid);
			this.routerExtensions.back();
		}

	}
}