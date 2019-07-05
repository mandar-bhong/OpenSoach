import { Component, OnInit, ViewContainerRef, ChangeDetectorRef } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { requestPermissions, takePicture } from 'nativescript-camera';
import * as imagepicker from 'nativescript-imagepicker';
import { ListViewEventData } from 'nativescript-ui-listview';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import * as utils from 'tns-core-modules/utils/utils';
import { PassDataService } from '~/app/services/pass-data-service';
import { ReportsService } from '~/app/services/reports/reports-service';
import { TeatmentReportModel, TreatmentReportDocModel } from '~/app/models/ui/reports-models';
import { AppGlobalContext } from '~/app/app-global-context';
import { AppRepoService } from '~/app/services/app-repo.service';
const fileSystemModule = require("tns-core-modules/file-system");
import { Folder } from "tns-core-modules/file-system";
import { ModalDialogOptions, ModalDialogService } from 'nativescript-angular/modal-dialog'
import { getFile } from "tns-core-modules/http"
import { ImageModalComponent } from '../../image-modal/image-modal.component';
import { DownloadProgress, RequestOptions } from 'nativescript-download-progress';

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

	isLoading = true;
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
		private viewContainerRef: ViewContainerRef,
		private modalService: ModalDialogService,
		private changeDetectorRef: ChangeDetectorRef,
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
				this.isLoading = false;
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
				let treatmentReportDocModel = new TreatmentReportDocModel();
				treatmentReportDocModel = val;
				treatmentReportDocModel.progress = 0;
				teatmentReportModel.doclist.push(treatmentReportDocModel);
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
	download(document_name, document_uuid, treatment_uuid) {

		if (document_uuid) {
			console.log('tap document_uuid', document_uuid);
			const token1 = AppGlobalContext.Token;
			console.log('token', token1);
			const requestObj = new ApiParse();
			requestObj.uuid = document_uuid;
			requestObj.token = token1;

			const apiUrl = '/v1/document/download/ep';
			const apiURL = AppRepoService.Instance.API_APP_BASE_URL + apiUrl + "/" + document_name + '?params=' + JSON.stringify(requestObj);
			console.log('apiURL', apiURL);
			// utils.openUrl(apiURL);

			var externalStoragePath = android.os.Environment.getExternalStorageDirectory().toString();
			var downloadFolderPath = fileSystemModule.path.join(externalStoragePath, "PatientCare");
			if (!Folder.exists(downloadFolderPath)) {
				Folder.fromPath(downloadFolderPath);
			}

			let fileName = document_uuid + "." + document_name.split('.').pop()

			const filePath: string = fileSystemModule.path.join(downloadFolderPath, fileName);
			console.log("path", filePath);

			const exists = fileSystemModule.File.exists(filePath);
			console.log(`Does file exists: ${exists}`);
			if (exists) {
				const existingItem = this.teatmentReportModel.filter(e => e.uuid == treatment_uuid)[0];
				if (existingItem) {
					const existingDocItem = existingItem.doclist.filter(e => e.document_uuid == document_uuid)[0];
					if (existingDocItem) {
						existingDocItem.document_path = filePath;
					}
				}
			} else {

				//download plugin
				const download = new DownloadProgress();
				const requestOptions: RequestOptions = {
					method: "GET",
					headers: {
					}
				};

				download.downloadFile(apiURL, requestOptions, filePath).then(f => {
					console.log("download Success");
					const existingItem = this.teatmentReportModel.filter(e => e.uuid == treatment_uuid)[0];
					if (existingItem) {
						const existingDocItem = existingItem.doclist.filter(e => e.document_uuid == document_uuid)[0];
						if (existingDocItem) {
							existingDocItem.document_path = filePath;
						}
					}
				}).catch(e => {
					console.log("download Error:", e);
				})

				download.addProgressCallback(progress => {
					const existingItem = this.teatmentReportModel.filter(e => e.uuid == treatment_uuid)[0];
					if (existingItem) {
						const existingDocItem = existingItem.doclist.filter(e => e.document_uuid == document_uuid)[0];
						if (existingDocItem) {
							existingDocItem.progress = Number((progress * 100).toFixed());
							this.changeDetectorRef.detectChanges();
						}
					}
				})
				//end download plugin
			}


		}
	}

	showImageModal(document_uuid, pathology_record_uuid) {

		const pathologyRecordItem = this.teatmentReportModel.filter(e => e.uuid == pathology_record_uuid)[0];
		const existingDocItem = pathologyRecordItem.doclist.filter(e => e.document_uuid == document_uuid)[0];

		const options: ModalDialogOptions = {
			viewContainerRef: this.viewContainerRef,
			fullscreen: true,
			context: {
				docPath: existingDocItem.document_path,
				docType: existingDocItem.doctype,
				modalName:"Treatment Report",
			}
		};
		this.modalService.showModal(ImageModalComponent, options);
	}
	// >> document download
}