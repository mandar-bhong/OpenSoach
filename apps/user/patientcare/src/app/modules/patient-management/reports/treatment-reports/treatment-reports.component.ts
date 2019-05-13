import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { requestPermissions, takePicture } from 'nativescript-camera';
import * as imagepicker from 'nativescript-imagepicker';
import { ListViewEventData } from 'nativescript-ui-listview';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import * as utils from 'tns-core-modules/utils/utils';
import { PassDataService } from '~/app/services/pass-data-service';
import { ReportsService } from '~/app/services/reports/reports-service';
import { TeatmentReportModel } from '~/app/models/ui/reports-models';
import { AppGlobalContext } from '~/app/app-global-context';
import { API_APP_BASE_URL } from '~/app/app-constants';

export class ApiParse {
	uuid: any;
	token: any;
}

export class DataItem {
	treatmentdone: string;
	treatmentdetails: string;
	doctorsinvolved: string;
	time: string;
	post_observations: string;
}

@Component({
	moduleId: module.id,
	selector: 'treatment-reports',
	templateUrl: './treatment-reports.component.html',
	styleUrls: ['./treatment-reports.component.css']
})

export class TreatmentReportsComponent implements OnInit {

	teatmentReportModel: TeatmentReportModel[] = [];
	teatmentReportList = new ObservableArray<TeatmentReportModel>();
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



	constructor(private routerExtensions: RouterExtensions,
		private passdataservice: PassDataService,
		private reportsService: ReportsService) { }

	ngOnInit() {

	}
	public listLoaded() {

		console.log('list loaded');

		setTimeout(() => {
			this.getTreatmentReportByUUID();
		}, 200);
	}

	public getTreatmentReportByUUID() {
		this.teatmentReportList = new ObservableArray<TeatmentReportModel>();
		this.reportsService.getTreatmentReportByUUID(this.passdataservice.getAdmissionID()).then(
			(val) => {
				// val.forEach(item => {
				// console.log("get getPathlogyReportByUUID item ", val);
				this.teatmentReportModel = val;
				// console.log(this.newMethod(), this.pathlogyReportList);
				this.getDoc();
			},
			(error) => {
				console.log("admistion details error:", error);
			}
		);
	}// end 

	getDoc() {
		this.teatmentReportModel.forEach(async (teatmentReportModel: TeatmentReportModel) => {
			teatmentReportModel.doclist = [];
			const x = await this.reportsService.getTreatmentReportDoc(teatmentReportModel.uuid);
			x.forEach((val) => {
				teatmentReportModel.doclist.push(val);
			});

			this.teatmentReportList.push(teatmentReportModel);
			// console.log('this.teatmentReportList', this.teatmentReportList);
		});


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

	// << document download
	download(document_name,document_uuid) {

		console.log('tap document_uuid', document_uuid);
		const token1 = AppGlobalContext.Token;
		console.log('token', token1);
		const requestObj = new ApiParse();
		requestObj.uuid = document_uuid;
		requestObj.token = token1;

		requestObj.uuid = document_uuid;
		requestObj.token = token1;
		const apiUrl = '/v1/document/download/ep';
		const apiURL = API_APP_BASE_URL + apiUrl +"/" + document_name +  '?params=' + JSON.stringify(requestObj);
		console.log('apiURL', apiURL);
		utils.openUrl(apiURL);

	}
	// >> document download
}