import { Component, OnInit } from '@angular/core';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { ListViewLinearLayout, ListViewEventData, RadListView, ListViewItemSnapMode } from 'nativescript-ui-listview';
import { RadListViewComponent } from 'nativescript-ui-listview/angular/listview-directives';
import { Observable } from 'tns-core-modules/data/observable';
import { RouterExtensions } from 'nativescript-angular/router';
import * as camera from "nativescript-camera";


import { takePicture, requestPermissions } from 'nativescript-camera';
import { ImageAsset } from 'tns-core-modules/image-asset';
import { PassDataService } from '~/app/services/pass-data-service';
import * as imagepicker from "nativescript-imagepicker";
import { Page } from 'tns-core-modules/ui/page/page';
import * as utils from "tns-core-modules/utils/utils";
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import { ReportsService } from '~/app/services/reports/reports-service';
import { PathlogyReportModel } from '~/app/models/ui/reports-models';
import { HttpClient, HttpParams } from '@angular/common/http';
import { AppGlobalContext } from '~/app/app-global-context';
import { ServerApiInterfaceService } from '~/app/services/server-api-interface.service';
import { API_SPL_BASE_URL, API_APP_BASE_URL } from '~/app/app-constants';



export class ApiParse {
	uuid: any;
	token: any;
}
@Component({
	moduleId: module.id,
	selector: 'section-one',
	templateUrl: './section-one.component.html',
	styleUrls: ['./section-one.component.css']
})

export class SectionOneComponent implements OnInit {
	// public _dataItems: ObservableArray<any>;
	// tempdata = new Array<DataItem>();

	pathlogyReportModel: PathlogyReportModel[] = [];
	pathlogyReportList = new ObservableArray<PathlogyReportModel>();
	public saveToGallery: boolean = true;
	public keepAspectRatio: boolean = true;
	public width: number = 320;
	public height: number = 240;

	imageAssets = [];
	imageSrc: any;
	isSingleMode: boolean = true;
	thumbSize: number = 80;
	previewSize: number = 300;
	imgpicker: boolean = true;
	public removedImageUrl: string;
	http: any;
	httpModule: any;


	constructor(private routerExtensions: RouterExtensions,
		private passdataservice: PassDataService,
		private reportsService: ReportsService,
		public httpClient: HttpClient,
		private serverApiInterfaceService: ServerApiInterfaceService) { }

	ngOnInit() {

	}
	public listLoaded() {

		console.log('list loaded');

		setTimeout(() => {
			this.getPathlogyReportByUUID()
		}, 200);
	}
	public getPathlogyReportByUUID() {
		this.pathlogyReportList = new ObservableArray<PathlogyReportModel>();
		this.reportsService.getPathlogyReportByUUID(this.passdataservice.getAdmissionID()).then(
			(val) => {
				// val.forEach(item => {
				// console.log("get getPathlogyReportByUUID item ", val);
				this.pathlogyReportModel = val;
				// console.log(this.newMethod(), this.pathlogyReportList);
				this.getDoc();
			},
			(error) => {
				console.log("admistion details error:", error);
			}
		);
	}// end 

	getDoc() {

		this.pathlogyReportModel.forEach(async (pathlogyReportModel: PathlogyReportModel) => {
			pathlogyReportModel.doclist = [];
			const document = await this.reportsService.getPathlogyReportDoc(pathlogyReportModel.uuid);
			document.forEach((val) => {
				pathlogyReportModel.doclist.push(val);
			});

			this.pathlogyReportList.push(pathlogyReportModel);
			console.log('this.pathlogyReportList', this.pathlogyReportList);
		});


	}



	private newMethod(): any {
		return 'pathlogyReportList';
	}

	onTakePictureTap() {
		requestPermissions().then(
			() => {
				takePicture({ width: this.width, height: this.height, keepAspectRatio: this.keepAspectRatio, saveToGallery: this.saveToGallery })
					.then((imageAsset: any) => {
						console.log('imageAsset', imageAsset);
						this.passdataservice.pickedImage = imageAsset;
						// navigate to show camera image component once image has been picked.
						this.routerExtensions.navigate(['patientmgnt', 'showcameraimage']);
					}, (error) => {
						console.log("Error: " + error);
					});
			},
			() => alert('permissions rejected')
		);
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
				that.imageAssets = [];
				that.imageSrc = null;
				return context.present();
			})
			.then((selection) => {
				console.log("Selection done: " + JSON.stringify(selection));
				that.imageSrc = that.isSingleMode && selection.length > 0 ? selection[0] : null;

				// set the images to be loaded from the assets with optimal sizes (optimize memory usage)
				selection.forEach(function (element) {
					element.options.width = that.isSingleMode ? that.previewSize : that.thumbSize;
					element.options.height = that.isSingleMode ? that.previewSize : that.thumbSize;
				});

				that.imageAssets = selection;
				console.log('select image  from galary');
				console.log(that.imageAssets);
				this.passdataservice.uploadedImage = selection;
				this.routerExtensions.navigate(['patientmgnt', 'showuoploadedimage']);
			}).catch(function (e) {
				console.log(e);
			});
	} // end of select image 



	onRemoveImageButtonTap(): void {
		this.imageSrc = null;
		this.imageAssets = []

		// if (this.currentImageSource) {
		// 	this.currentImageSource = null;
		// }
	}
	public openPDFFile() {
		utils.openUrl("https://www.princexml.com/samples/flyer/flyer.pdf");
	}
	// >> expand row code start
	templateSelector(item: any, index: number, items: any): string {
		return item.expanded ? "expanded" : "default";

	}

	onItemTap(event: ListViewEventData) {
		const listView = event.object,
			rowIndex = event.index,
			dataItem = event.view.bindingContext;

		dataItem.expanded = !dataItem.expanded;
		if (isIOS) {
			// Uncomment the lines below to avoid default animation
			// UIView.animateWithDurationAnimations(0, () => {
			var indexPaths = NSMutableArray.new();
			indexPaths.addObject(NSIndexPath.indexPathForRowInSection(rowIndex, event.groupIndex));
			listView.ios.reloadItemsAtIndexPaths(indexPaths);
			// });
		}
		if (isAndroid) {
			listView.androidListView.getAdapter().notifyItemChanged(rowIndex);

		}
	}
	// << expand row code end
	download(document_name,document_uuid) {

		console.log('tap document_uuid', document_uuid);
		const token1 = AppGlobalContext.Token;
		console.log('token', token1);
		const requestObj = new ApiParse();
		requestObj.uuid = document_uuid;
		requestObj.token = token1;

		let result;


		// this.serverApiInterfaceService.get(API_APP_BASE_URL + "/v1/document/download/ep",
		// 	{
		// 		uuid: document_uuid,
		// 		token: AppGlobalContext.Token

		// 	})
		// 	.then((result) => {
		// 		console.log('result', result);
		// 		result = result;

		// 	})
		// utils.openUrl(result);


		requestObj.uuid = document_uuid;
		requestObj.token = token1;
		const apiUrl = '/v1/document/download/ep';
		const apiURL = API_APP_BASE_URL + apiUrl +"/" + document_name +  '?params=' + JSON.stringify(requestObj);
		console.log('apiURL', apiURL);
		utils.openUrl(apiURL);

	}

}