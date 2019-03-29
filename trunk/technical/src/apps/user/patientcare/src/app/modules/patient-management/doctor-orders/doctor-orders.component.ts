import { Component, OnInit, ViewChild } from '@angular/core';
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
import { RadAutoCompleteTextViewComponent } from 'nativescript-ui-autocomplete/angular/autocomplete-directives';
import * as imagepicker from "nativescript-imagepicker";
import { DocumentUploadDatastore } from '~/app/models/db/document-upload-datastore';
var mime = require('mime-types')
@Component({
	moduleId: module.id,
	selector: 'doctor-orders',
	templateUrl: './doctor-orders.component.html',
	styleUrls: ['./doctor-orders.component.css']
})
export class DoctorOrdersComponent implements OnInit {
	public imageTaken: ImageAsset;
	imageSrc: any;
	public saveToGallery: boolean = true;
	public keepAspectRatio: boolean = true;
	public width: number = 300;
	public height: number = 300;
	docPath: string;
	thumbSize: number = 80;
	previewSize: number = 300;
	doctorOrdersForm: FormGroup;
	private _items: ObservableArray<DoctorInfo>;
	doctorName = new FormControl('', [Validators.required]);
	private doctors: DoctorsList[] = [
		{ name: "Amol Patil", id: '11' },
		{ name: "Ganesh Patil", id: '12' },
		{ name: "mahesh Patil", id: '13' },
		{ name: "Sarjerao", id: '14' }]
	isDescRequired = false;
	isDrNameRequired = false;
	isCommentRequired = false;
	isSingleMode: boolean;
	constructor(private params: ModalDialogParams,
		private passDataService: PassDataService) {
		this.initDataItems();
	}
	// @ViewChild('aut') Item: RadAutoCompleteTextViewComponent;
	ngOnInit() {
		this.createFormControls();

	}
	// << func for creating form controls
	createFormControls(): void {
		this.doctorOrdersForm = new FormGroup({
			desc: new FormControl('', [Validators.required]),
			comment: new FormControl('', [Validators.required])
		});
		this.doctorOrdersForm.addControl('doctorName', this.doctorName);
	}
	// >> func for creating form controls
	// will execute on submit
	onSubmit() {
		console.log('this.doctorOrdersForm', this.doctorOrdersForm.status);
		this.isDescRequired = this.doctorOrdersForm.controls['desc'].hasError('required');
		this.isDrNameRequired = this.doctorOrdersForm.controls['doctorName'].hasError('required');
		this.isCommentRequired = this.doctorOrdersForm.controls['comment'].hasError('required');
		if (this.doctorOrdersForm.invalid) {
			console.log("validation error");
			return;
		}
		const desc = this.doctorOrdersForm.controls['desc'].value
		const doctorName = this.doctorOrdersForm.controls['doctorName'].value;
		const doctorComment = this.doctorOrdersForm.controls['comment'].value;
		console.log(this.autocomplete);
		const serverDataStoreModel = new ServerDataStoreDataModel<DoctorsOrdersDatastoreModel>();
		serverDataStoreModel.datastore = SYNC_STORE.DOCTORS_ORDERS;
		serverDataStoreModel.data = new DoctorsOrdersDatastoreModel();
		serverDataStoreModel.data.admission_uuid = this.passDataService.getAdmissionID();
		serverDataStoreModel.data.client_updated_at = new Date().toISOString();
		serverDataStoreModel.data.doctor_id = doctorName;
		serverDataStoreModel.data.doctors_orders = desc;
		serverDataStoreModel.data.status = 0;
		serverDataStoreModel.data.comment = doctorComment;
		serverDataStoreModel.data.order_created_time = new Date().toISOString();
		// to do call api of upload document and set received rec id here.
		serverDataStoreModel.data.sync_pending = SYNC_PENDING.TRUE;
		serverDataStoreModel.data.uuid = PlatformHelper.API.getRandomUUID();
		const TempDataStore = new Array();
		// create document datastore 
		if (this.docPath && this.docPath != null) {
			const serverDocumentDataStoreModel = new ServerDataStoreDataModel<DocumentUploadDatastore>();
			serverDocumentDataStoreModel.datastore = SYNC_STORE.DOCUMENT;
			serverDocumentDataStoreModel.data = new DocumentUploadDatastore();
			serverDocumentDataStoreModel.data.client_updated_at = new Date().toISOString();
			serverDocumentDataStoreModel.data.doc_path = this.docPath;
			serverDocumentDataStoreModel.data.doc_name = 'test';
			serverDocumentDataStoreModel.data.doc_type = mime.lookup('xlsx');
			serverDocumentDataStoreModel.data.datastore = SYNC_STORE.DOCTORS_ORDERS;
			serverDocumentDataStoreModel.data.sync_pending = SYNC_PENDING.TRUE;
			serverDocumentDataStoreModel.data.uuid = PlatformHelper.API.getRandomUUID();
			serverDataStoreModel.data.document_uuid = serverDocumentDataStoreModel.data.uuid;
			TempDataStore.push(serverDocumentDataStoreModel);
		}
		TempDataStore.push(serverDataStoreModel);
		this.params.closeCallback(TempDataStore);
	}

	//end of code block
	onTakePhoto() {
		requestPermissions().then(
			() => this.capturePicture(),
			() => console.log('Permission Rejected')
		);
	}
	onRequestPermissions() {
		requestPermissions();
	}
	capturePicture() {
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
					this.docPath = this.imageTaken.android;
				} else if (application.ios) {
					this.docPath = this.imageTaken.ios;
				}
			}).catch(err => {
				console.log(err.message);
			});
	}
	// will executed on back forms back buttons
	goBackPage() {
		this.params.closeCallback([]);
	}
	//end of code block

	@ViewChild("autocomplete") autocomplete: RadAutoCompleteTextViewComponent;

	get dataItems(): ObservableArray<DoctorInfo> {
		return this._items;
	}

	private initDataItems() {
		this._items = new ObservableArray<DoctorInfo>();

		for (let i = 0; i < this.doctors.length; i++) {
			this._items.push(new DoctorInfo(this.doctors[i].name, undefined, this.doctors[i].id));
		}
	}
	public onDidAutoComplete(args) {
		console.log("DidAutoComplete with: " + args.text);
		const doctorName = args.text
		//  to do 
		// if same doctors have same names then append email id before or after doctor name so it will easy to find id from doctor list.
		//  fetching id from name bcz auto complete returns onley text.
		if (doctorName != '') {
			const item = this.doctors.filter(doctor => doctor.name === doctorName)[0]
			if (item) {
				this.doctorOrdersForm.controls['doctorName'].setValue(item.id);
				this.doctorOrdersForm.controls['doctorName'].updateValueAndValidity();
			}
		}
	}
	//  fucntion for selectiong image from gallery.
	public onSelectSingleTap() {
		this.isSingleMode = true;
		let context = imagepicker.create({
			mode: "single"
		});
		// code block for start selection.
		this.startSelection(context);
	}
	private startSelection(context) {
		let that = this;
		context
			.authorize()
			.then(() => {
				that.imageTaken = null
				that.imageSrc = null;
				return context.present();
			})
			.then((selection) => {
				that.imageSrc = that.isSingleMode && selection.length > 0 ? selection[0] : null;
				// set the images to be loaded from the assets with optimal sizes (optimize memory usage)
				selection.forEach(function (element) {
					element.options.width = that.isSingleMode ? that.previewSize : that.thumbSize;
					element.options.height = that.isSingleMode ? that.previewSize : that.thumbSize;
				});
				that.imageTaken = selection[0];
				console.log('select image  from galary');
				console.log(that.imageTaken);
				if (application.android) {
					that.docPath = this.imageTaken.android;
				} else if (application.ios) {
					that.docPath = this.imageTaken.ios;
				}
			}).catch(function (e) {
				console.log(e);
			});
	} // end of select image 
} // end of class 

export class DoctorInfo extends TokenModel {
	//name: string;
	id: string;
	constructor(text: string, image: string, id: string) {
		super(text, image);
		this.id = id;
	}
}
export class DoctorsList {
	name: string;
	id: string;
}