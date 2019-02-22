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

export class DataItem {
	public name: string;
	public description: string;
	result: string;
	time: string;

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
		this.tempdata.push({ name: "Laboratory Report", description: "Morning and evening before meal", time:"12:00am", result:"test result" });
		this.tempdata.push({ name: "Radiology Report", description: "Incase of high body temperature", time:"01:00am",  result:"test result" });
		this.tempdata.push({ name: "Blood Test", description: "Incase of continuos vomitting and nausea", time:"02:00am", result:"test result" });
		this.tempdata.push({ name: "Blood Glucose Test", description: "Blood glucose tests are also sometimes called blood sugar tests.", time:"03:00am",  result:"test result" });
		this.tempdata.push({ name: "Calcium Test", description: "Calcium is important because it gives strength to your bones.", time:"04:00am",  result:"test result" });
		this.tempdata.push({ name: "D-dimer Test", description: " D-dimer test is a blood test usually used to help check for or monitor blood clotting problems. ", time:"05:00am",  result:"test result" });
		this.tempdata.push({ name: "ESR Test", description: "The erythrocyte sedimentation rate (ESR) test checks for inflammation in the body. ", time:"06:00am", result:"test result" });
		this.tempdata.push({ name: "Floate Test", description: "Folate is an important nutrient for making normal red blood cells", time:"07:00am", result:"test result" });
		this.tempdata.push({ name: "Full Blood Count", description: "tiredness or weakness", time:"08:00am", result:"test result" });
		this.tempdata.push({ name: "HbA1c", description: "HbA1c is a blood test that is used to help diagnose and monitor people with diabetes.", time:"09:00am", result:"test result" });
		this.tempdata.push({ name: "Vitamin B12 test", description: "You need vitamin B12 in your blood so you can make blood cells", time:"10:00am", result:"test result"});
		this.tempdata.push({ name: "Calcium Test", description: "Calcium is important because it gives strength to your bones.", time:"11:00am", result:"test result" });
		this.tempdata.push({ name: "D-dimer Test", description: " D-dimer test is a blood test usually used to help check for or monitor blood clotting problems. ", time:"12:00pm",result:"test result" });
		this.tempdata.push({ name: "ESR Test", description: "The erythrocyte sedimentation rate (ESR) test checks for inflammation in the body. ", time:"12:00am", result:"test result" });
		this.tempdata.push({ name: "Floate Test", description: "Folate is an important nutrient for making normal red blood cells", time:"12:00am", result:"test result" });
		this.tempdata.push({ name: "Full Blood Count", description: "tiredness or weakness", time:"12:00am", result:"test result" });
		this.tempdata.push({ name: "HbA1c", description: "HbA1c is a blood test that is used to help diagnose and monitor people with diabetes.", time:"12:00am", result:"test result" });
		this.tempdata.push({ name: "Vitamin B12 test", description: "You need vitamin B12 in your blood so you can make blood cells" , time:"12:00am", result:"test result"});


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
}