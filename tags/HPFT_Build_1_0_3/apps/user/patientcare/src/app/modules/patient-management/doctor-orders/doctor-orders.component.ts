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
import { ChartService } from '~/app/services/chart/chart.service';


import { animate, state, style, transition, trigger } from "@angular/animations";
import { SegmentedBarItem } from 'tns-core-modules/ui/segmented-bar/segmented-bar';
import { VALIDATION_REQUIRED_FIELD } from '~/app/common-constants';
const imageSourceModule = require("tns-core-modules/image-source");
const fileSystemModule = require("tns-core-modules/file-system");


var mime = require('mime-types')
@Component({
	moduleId: module.id,
	selector: 'doctor-orders',
	templateUrl: './doctor-orders.component.html',
	styleUrls: ['./doctor-orders.component.css'],
	animations: [
		trigger("from-bottom", [
			state("in", style({
				"opacity": 1,
				transform: "translateY(0)"
			})),
			state("void", style({
				"opacity": 0,
				transform: "translateY(20%)"
			})),
			transition("void => *", [animate("1600ms 700ms ease-out")]),
			transition("* => void", [animate("600ms ease-in")])
		]),
		trigger("fade-in", [
			state("in", style({
				"opacity": 1
			})),
			state("void", style({
				"opacity": 0
			})),
			transition("void => *", [animate("800ms 2000ms ease-out")])
		]),
		trigger("scale-in", [
			state("in", style({
				"opacity": 1,
				transform: "scale(1)"
			})),
			state("void", style({
				"opacity": 0,
				transform: "scale(0.9)"
			})),
			transition("void => *", [animate("1100ms ease-out")])
		])
	]
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
	public orderType: Array<string> = [];
	doctorName = new FormControl('', [Validators.required]);
	private doctors: DoctorsList[] = [];
	isDescRequired = false;
	isDrNameRequired = false;
	isCommentRequired = false;
	isSingleMode: boolean;



	// drowing
	plugins: Array<SegmentedBarItem> = [];
	selectedPlugin: string = "Drawing";
	drawings: Array<any> = [];
	getDrawingPad: boolean;
	getFormView: boolean;
	pencolor;
	penWidth;
	pencilbuttonClicked: boolean = true;
	eraserbuttonClicked: boolean = false;
	alleraserbuttonClicked: boolean = false;
	drawImg: boolean = false;
	drawDocName: string;
	drawImgtype: string;

	// show last item in array
	fileArray: Array<any> = [];
	getFileItem: Array<any> = [];

	VALIDATION_REQUIRED_FIELD = VALIDATION_REQUIRED_FIELD;
	constructor(private params: ModalDialogParams,
		private chartService: ChartService,
		private passDataService: PassDataService) {
		this.getFormView = true;
	}
	// @ViewChild('aut') Item: RadAutoCompleteTextViewComponent;
	ngOnInit() {
		this.createFormControls();
		this.getOrderType();
		this.getDoctorNames();

		this.pencolor = '#e66465';
		this.penWidth = '1';
		// drowing
		this.addPluginToSegmentedBar("Drawing");
	}
	// << func for creating form controls
	createFormControls(): void {
		this.doctorOrdersForm = new FormGroup({
			desc: new FormControl('', [Validators.required]),
			comment: new FormControl('', [Validators.required]),
			order_type: new FormControl()
		});
		this.doctorOrdersForm.addControl('doctorName', this.doctorName);
	}
	// >> func for creating form controls
	// will execute on submit
	onSubmit() {
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

		const serverDataStoreModel = new ServerDataStoreDataModel<DoctorsOrdersDatastoreModel>();
		serverDataStoreModel.datastore = SYNC_STORE.DOCTORS_ORDERS;
		serverDataStoreModel.data = new DoctorsOrdersDatastoreModel();
		serverDataStoreModel.data.admission_uuid = this.passDataService.getAdmissionID();
		serverDataStoreModel.data.client_updated_at = new Date().toISOString();
		serverDataStoreModel.data.doctor_id = doctorName;
		serverDataStoreModel.data.doctors_orders = desc;
		serverDataStoreModel.data.status = 0;
		serverDataStoreModel.data.order_type = this.orderType[this.doctorOrdersForm.get('order_type').value];
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

			if (this.drawImg == true) {
				// drawimg
				serverDocumentDataStoreModel.data.doc_path = this.docPath;
				serverDocumentDataStoreModel.data.doc_name = this.drawDocName;
				serverDocumentDataStoreModel.data.doc_type = this.drawImgtype;
			}
			else {
				// upload and click img
				serverDocumentDataStoreModel.data.doc_path = this.docPath;
				serverDocumentDataStoreModel.data.doc_name = 'test';
				serverDocumentDataStoreModel.data.doc_type = mime.lookup('xlsx');
			}

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
				this.getFileItem = [];
				this.getFileItem.push(this.imageTaken);
				this.fileArray = this.getFileItem;
				// console.log("Size: " + imageAsset.options.width + "x" + imageAsset.options.height);
				// console.log('picked image', this.imageTaken);
				if (application.android) {
					this.docPath = this.imageTaken.android;
					console.log('this.docPath', this.docPath);
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

	// private initDataItems() {
	// 	this._items = new ObservableArray<DoctorInfo>();

	// 	for (let i = 0; i < this.doctors.length; i++) {
	// 		this._items.push(new DoctorInfo(this.doctors[i].name, undefined, this.doctors[i].id));
	// 	}
	// }
	public onDidAutoComplete(args) {
		const doctorName = args.text
		//  to do 
		// if same doctors have same names then append email id before or after doctor name so it will easy to find id from doctor list.
		//  fetching id from name bcz auto complete returns onley text.
		if (doctorName != '') {
			const item = this.doctors.filter(doctor => doctor.formatedName === doctorName)[0] || null;
			if (item) {
				this.doctorOrdersForm.controls['doctorName'].setValue(item.usr_id);
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
				this.getFileItem = [];
				this.getFileItem.push(that.imageTaken);
				this.fileArray = this.getFileItem;
				if (application.android) {
					that.docPath = this.imageTaken.android;
				} else if (application.ios) {
					that.docPath = this.imageTaken.ios;
				}
			}).catch(function (e) {
				console.log(e);
			});
	} // end of select image
	// fucntion for getting  doctor order type form database
	public getOrderType() {
		this.chartService.getAllData('orderType').then(
			(success) => {
				if (success.length > 0) {
					const medicineType = JSON.parse(success[0].conf);
					this.orderType = [];
					for (let item of medicineType) {
						this.orderType.push(item);
					}
				}
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}// end of block
	// code block for getting doctors names 
	public getDoctorNames() {
		this.chartService.getAllData('doctorName').then(
			(success) => {
				this.doctors = success;
				if (this.doctors.length > 0) {
					// console.log('doctors names', success);
					this._items = new ObservableArray<DoctorInfo>();
					success.forEach((item) => {
						item.formatedName = item.fname + " " + item.lname;
						this._items.push(new DoctorInfo(item.formatedName, undefined, item.usr_id));
					});
				}
			},
			(error) => {
				console.log("getChartData error:", error);
			}
		);
	}


	private addPluginToSegmentedBar(name: string) {
		let drawingPad = new SegmentedBarItem();
		drawingPad.title = name;
		this.plugins.push(drawingPad);
	}
	// get Drawing Img 
	getMyDrawing(pad: any) {
		// then get the drawing (Bitmap on Android) of the drawingpad
		pad.getDrawing().then(data => {
			this.drawings.push(data);

			this.getFileItem = this.drawings;
			this.clearMyDrawing(pad);
			this.getformview();

		}, error => {
			console.log(error);
		});



		// fun for get drawing to convert PNG formate
		this.drawingImgToPng(pad);
	}
	drawingImgToPng(pad) {
		this.drawImg = true;
		let name;
		let path;
		// grabs the  path for Downloads (string value)
		pad.getDrawing().then(function (result) {
			var img2 = imageSourceModule.fromNativeSource(result);
			var androidDownloadsPath = android.os.Environment.getExternalStoragePublicDirectory(
				android.os.Environment.DIRECTORY_DOWNLOADS).toString();

			// creates PATH for folder called MyFolder in /Downloads (string value)
			var myFolderPath = fileSystemModule.path.join(androidDownloadsPath, "MyFolder");

			const ticks = new Date;
			name = 'drawing-pad-' + ticks.getTime() + '.png';
			// creates a path of kind ../Downloads/MyFolder/my-file-name.jpg
			path = fileSystemModule.path.join(myFolderPath, name);
			var saved = img2.saveToFile(path, "png");
		});
		// settime out for convrting img take time for avoiding this  
		setTimeout(() => {
			this.drawDocName = name;
			this.drawImgtype = '.png';
			this.docPath = path;
		}, 1);

	}
	// show drawingpad view
	onDrawingPad() {
		this.getDrawingPad = true;
		this.getFormView = false;
		// this.getDrawingPad = !this.getDrawingPad;
	}
	// show form view
	getformview() {
		this.getDrawingPad = false;
		this.getFormView = true;
	}
	// erase text
	getEraser() {
		this.pencolor = '#ffffff';
		this.penWidth = '20';
		this.pencilbuttonClicked = false;
		this.eraserbuttonClicked = true;
		this.alleraserbuttonClicked = false
	}
	// all clear drawing 
	clearMyDrawing(pad: any) {
		pad.clearDrawing();
		this.pencilbuttonClicked = false;
		this.eraserbuttonClicked = false;
		this.alleraserbuttonClicked = true;
	}
	// after clear drawing then select pen 
	getDrow() {
		this.pencolor = '#e66465';
		this.penWidth = '1';
		this.pencilbuttonClicked = true;
		this.eraserbuttonClicked = false;
		this.alleraserbuttonClicked = false;
	}

	// TODO find use this fun if unused then remove
	protected getScreenName(): string {
		return "Input";
	}


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
	usr_id: number;
	usr_name: string;
	urole_name: string;
	fname: string;
	lname: string;
	formatedName: string;
}
