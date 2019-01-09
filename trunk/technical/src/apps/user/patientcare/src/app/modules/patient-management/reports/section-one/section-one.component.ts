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
		this.tempdata.push({ name: "Amol Patil", description: "200ml" });
		this.tempdata.push({ name: "Shubham Lunia", description: "3 times a day after meal" });
		this.tempdata.push({ name: "Mayuri jain", description: "Morning and evening before meal" });
		this.tempdata.push({ name: "Sanjay Mohan", description: "Incase of high body temperature" });
		this.tempdata.push({ name: "Pooja Lokare", description: "Incase of continuos vomitting and nausea" });
		this.tempdata.push({ name: "Jagdish Wagh", description: "Monitor every 2 hours" });
		this.tempdata.push({ name: "Mandar Bhong", description: "Monitor every 3 hours" });
		this.tempdata.push({ name: "Praveen Pandey", description: "Monitor every 15 mins" });
		this.tempdata.push({ name: "Shashank Atre", description: "Incase of high body temperature" });
		this.tempdata.push({ name: "Abhijeet Kalbhor", description: "Morning and evening before meal" });
		this.tempdata.push({ name: "Sarjerao", description: "Monitor every 15 mins" });
		this.tempdata.push({ name: "Rahul", description: "Incase of high body temperature" });
		this.tempdata.push({ name: "Praveen", description: "Morning and evening before meal" });

		this._dataItems = new ObservableArray(this.tempdata);
	}
	goCameras() {
		this.routerExtensions.navigate(['patientmgnt', 'cameras'], { clearHistory: true });

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
}