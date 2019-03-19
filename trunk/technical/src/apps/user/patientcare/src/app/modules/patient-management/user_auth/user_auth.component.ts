import { Component, OnInit, ViewChild, Input } from '@angular/core';
import { ActionListViewModel, UserAuthDBRequest, UserCreateFormRequest } from '~/app/models/ui/action-models';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { ActionService } from '~/app/services/action/action.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { RouterExtensions } from 'nativescript-angular/router';
import { PassDataService } from '~/app/services/pass-data-service';
import { ServerApiInterfaceService } from '~/app/services/server-api-interface.service';
import { API_SPL_BASE_URL } from '~/app/app-constants';
import { AppGlobalContext } from '~/app/app-global-context';
import { UserDetails, UserDetailDBModel } from '~/app/models/ui/user-auth-models';
import { UserAuthService } from '~/app/services/user-auth/user-auth-service';

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
	constructor(private UserAuthService: UserAuthService,
		private routerExtensions: RouterExtensions,
		private passdataservice: PassDataService,
		private serverApiInterfaceService: ServerApiInterfaceService) {
		console.log('user auth');
		this.datamodel = new UserAuthDBRequest();
	}
	get dataItems(): ObservableArray<ActionListViewModel> {
		return this._dataItemsaccount;
	}
	@Input() patientName: string;
	@ViewChild("myListView") listViewComponent: RadListViewComponent;

	ngOnInit() {
		this.passdataservice.backalert = true;
		this.createFormUserAuthControls();
		this.createFormUserPinControls();
		this.userPinCheckControls();
		this.getUserAccountData();
		this.UserAuthService.getUserAccountList1();

		// if (this._dataItemsaccount.length > 0) {
		// 	this.pinview = true;
		// 	this.newpinview = false;
		// } else {
		// 	this.newpinview = true;
		// 	this.pinview = false;
		// }
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
		this.UserAuthService.getUserAccountList().then(
			(val) => {
				val.forEach(item => {
					const actionListItems = new ActionListViewModel();
					actionListItems.dbmodel = item;
					this._dataItemsaccount.push(actionListItems);
					// console.log('User Account List', item);

				});
			},
			(error) => {
				console.log("getUserAccountData error:", error);
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
			// console.log('this.pin Right', formmodel.pin);
			this.passdataservice.getAuthUserId = this.selecedItemsData.userid;
			// console.log('call all ready add list user id', this.passdataservice.getAuthUserId);
			this.passdataservice.authResultReuested.onDeviceAuthSuccess(this.selecedItemsData.userid);
			this.routerExtensions.back();

		} else {
			console.log('this.pin wrong', formmodel.pin);
		}
		console.log('on submit back');

	}
	// new user create add email and password then compare to server user is valid or not. user is vaild then call API get fname and lname
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

		console.log('token', AppGlobalContext.Token);

		if (formmodel.email && formmodel.password) {
			this.serverApiInterfaceService.post<UserDetails>(API_SPL_BASE_URL + "/v1/endpoint/userauthorization",
				{
					'username': formmodel.email,
					'password': formmodel.password,
					'devicetoken': AppGlobalContext.Token,
				})
				.then(
					(res) => {
						console.log("POST Request is successful ", res);
						this.getuserdetails = res;

						this.newpinview = true;
						this.hidecheckpin = false;
					}, (error) => {
						console.log(error);
					}
				);
		}


	}
	// new user create account set new password 
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

		if (formmodel.newpin === formmodel.reenterpin) {
			const saveuserdetails = new UserDetailDBModel();
			saveuserdetails.userid = this.getuserdetails.userid;
			saveuserdetails.first_name = this.getuserdetails.firstname;
			saveuserdetails.last_name = this.getuserdetails.lastname;
			saveuserdetails.email = this.getuserdetails.username;
			saveuserdetails.pin = formmodel.reenterpin;
			// save action done and discard in DB
			// const userid = Number(saveuserdetails.userid)
			const userid = saveuserdetails.userid;
			this.passdataservice.getAuthUserId = userid;
			// console.log('call new user id', this.passdataservice.getAuthUserId);
			this.UserAuthService.insertDeviceAccessItem(saveuserdetails);
			this.passdataservice.authResultReuested.onDeviceAuthSuccess(userid);
			this.routerExtensions.back();
		}

	}
}