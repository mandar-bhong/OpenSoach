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

export class DataItem {
	public name: string;
	public description: string;

}
@Component({
	moduleId: module.id,
	selector: 'section-one',
	templateUrl: './section-one.component.html',
	styleUrls: ['./section-one.component.css']
})

export class SectionOneComponent implements OnInit {
	public _dataItems: ObservableArray<any>;
	tempdata = new Array<DataItem>();

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
		private passdataservice: PassDataService, ) { }

	ngOnInit() { }
	public listLoaded() {

		console.log('list loaded');

		setTimeout(() => {
			this.initDataItems();
		}, 200);
	}

	public initDataItems() {
		const tempdata = new Array<DataItem>();
		this.tempdata.push({ name: "Laboratory Report", description: "Morning and evening before meal" });
		this.tempdata.push({ name: "Radiology Report", description: "Incase of high body temperature" });
		this.tempdata.push({ name: "Blood Test", description: "Incase of continuos vomitting and nausea" });
		this.tempdata.push({ name: "Blood Glucose Test", description: "Blood glucose tests are also sometimes called blood sugar tests." });
		this.tempdata.push({ name: "Calcium Test", description: "Calcium is important because it gives strength to your bones." });
		this.tempdata.push({ name: "D-dimer Test", description: " D-dimer test is a blood test usually used to help check for or monitor blood clotting problems. " });
		this.tempdata.push({ name: "ESR Test", description: "The erythrocyte sedimentation rate (ESR) test checks for inflammation in the body. " });
		this.tempdata.push({ name: "Floate Test", description: "Folate is an important nutrient for making normal red blood cells" });
		this.tempdata.push({ name: "Full Blood Count", description: "tiredness or weakness" });
		this.tempdata.push({ name: "HbA1c", description: "HbA1c is a blood test that is used to help diagnose and monitor people with diabetes." });
		this.tempdata.push({ name: "Vitamin B12 test", description: "You need vitamin B12 in your blood so you can make blood cells" });
		this.tempdata.push({ name: "Calcium Test", description: "Calcium is important because it gives strength to your bones." });
		this.tempdata.push({ name: "D-dimer Test", description: " D-dimer test is a blood test usually used to help check for or monitor blood clotting problems. " });
		this.tempdata.push({ name: "ESR Test", description: "The erythrocyte sedimentation rate (ESR) test checks for inflammation in the body. " });
		this.tempdata.push({ name: "Floate Test", description: "Folate is an important nutrient for making normal red blood cells" });
		this.tempdata.push({ name: "Full Blood Count", description: "tiredness or weakness" });
		this.tempdata.push({ name: "HbA1c", description: "HbA1c is a blood test that is used to help diagnose and monitor people with diabetes." });
		this.tempdata.push({ name: "Vitamin B12 test", description: "You need vitamin B12 in your blood so you can make blood cells" });


		this._dataItems = new ObservableArray(this.tempdata);
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
}