import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ModalDialogParams } from 'nativescript-angular/modal-dialog';

import { PassDataService } from '~/app/services/pass-data-service';
import { ImageAsset } from 'tns-core-modules/image-asset/image-asset';
import { takePicture, requestPermissions } from 'nativescript-camera';
import * as application from "tns-core-modules/application";
import { ServerDataStoreDataModel } from '~/app/models/api/server-data-store-data-model';
import { SYNC_STORE, SYNC_PENDING } from '~/app/app-constants';
import { DoctorsOrdersDatastoreModel } from '~/app/models/db/doctors-orders-model';
import { PlatformHelper } from '~/app/helpers/platform-helper';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { TokenModel } from 'nativescript-ui-autocomplete';
@Component({
	moduleId: module.id,
	selector: 'doctor-orders',
	templateUrl: './doctor-orders.component.html',
	styleUrls: ['./doctor-orders.component.css']
})

export class DoctorOrdersComponent implements OnInit {
	public imageTaken: ImageAsset;
	public saveToGallery: boolean = true;
	public keepAspectRatio: boolean = true;
	public width: number = 300;
	public height: number = 300;
	documentId: string;
	doctorOrdersForm: FormGroup;
	autocompleDoctors: ObservableArray<TokenModel> = new ObservableArray<TokenModel>([
		new TokenModel("Laboratory Report", undefined),
		new TokenModel("Radiology Report", undefined),
		new TokenModel("Blood Test", undefined),
		new TokenModel("Blood Glucose Test", undefined),
		new TokenModel("Calcium Test", undefined),
		new TokenModel("D-dimer Test", undefined),
		new TokenModel("ESR Test", undefined),
		new TokenModel("Floate Test", undefined),
		new TokenModel("Full Blood Count", undefined),
		new TokenModel("HbA1c", undefined),
		new TokenModel("Vitamin B12 test", undefined),
		new TokenModel("Calcium Test", undefined),
	]);
	constructor(private params: ModalDialogParams,
		private passDataService: PassDataService) { }

	ngOnInit() {
		this.createFormControls();

	}
	// << func for creating form controls
	createFormControls(): void {
		this.doctorOrdersForm = new FormGroup({
			desc: new FormControl(),
			doctorName: new FormControl()
		});
	}
	// >> func for creating form controls
	// will execute on submit
	onSubmit() {
		if (this.doctorOrdersForm.valid) {
			const desc = this.doctorOrdersForm.controls['desc'].value
			const doctorName = this.doctorOrdersForm.controls['doctorName'].value;
			console.log('doctorName',doctorName,'desc      ',desc);	
			
			const serverDataStoreModel = new ServerDataStoreDataModel<DoctorsOrdersDatastoreModel>();
			serverDataStoreModel.datastore = SYNC_STORE.DOCTORS_ORDERS;
			serverDataStoreModel.data = new DoctorsOrdersDatastoreModel();
			serverDataStoreModel.data.admission_uuid = this.passDataService.getAdmissionID();
			serverDataStoreModel.data.client_updated_at = new Date();
			serverDataStoreModel.data.doctor_id = doctorName;
			serverDataStoreModel.data.doctors_orders = desc;
			// to do call api of upload document and set received rec id here.
			serverDataStoreModel.data.document_uuid = this.documentId;
			serverDataStoreModel.data.sync_pending = SYNC_PENDING.TRUE;
			serverDataStoreModel.data.uuid = PlatformHelper.API.getRandomUUID();
			this.params.closeCallback([serverDataStoreModel]);
		}

	}
	//end of code block
	onTakePhoto() {
		this.onRequestPermissions();
		let options = {
			width: this.width,
			height: this.height,
			keepAspectRatio: this.keepAspectRatio,
			saveToGallery: this.saveToGallery
		};
		takePicture(options)
			.then(imageAsset => {
				this.imageTaken = imageAsset;
				console.log("Size: " + imageAsset.options.width + "x" + imageAsset.options.height);
				console.log('picked image', this.imageTaken);
				if (application.android) {
					this.documentId = this.imageTaken.android;
				} else if (application.ios) {
					this.documentId = this.imageTaken.ios;
				}
			}).catch(err => {
				console.log(err.message);
			});
	}
	onRequestPermissions() {
		requestPermissions();
	}
	// will executed on back forms back buttons
	goBackPage() {
		this.params.closeCallback([]);
	}
	//end of code block
}