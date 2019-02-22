import { Component, OnInit } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { requestPermissions, takePicture } from 'nativescript-camera';
import * as imagepicker from 'nativescript-imagepicker';
import { ListViewEventData } from 'nativescript-ui-listview';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { isAndroid, isIOS } from 'tns-core-modules/ui/page/page';
import * as utils from 'tns-core-modules/utils/utils';
import { PassDataService } from '~/app/services/pass-data-service';


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
			this.tempdata.push({ treatmentdone: "Laboratory", treatmentdetails: "Morning and evening before meal", time:"12:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Radiology", treatmentdetails: "Incase of high body temperature", time:"01:00am",  doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Blood ", treatmentdetails: "Incase of continuos vomitting and nausea", time:"02:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Blood Glucose", treatmentdetails: "Blood glucose tests are also sometimes called blood sugar tests.", time:"03:00am",  doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Calcium ", treatmentdetails: "Calcium is important because it gives strength to your bones.", time:"04:00am",  doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "D-dimer ", treatmentdetails: " D-dimer test is a blood test usually used to help check for or monitor blood clotting problems. ", time:"05:00am",  doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "ESR ", treatmentdetails: "The erythrocyte sedimentation rate (ESR) test checks for inflammation in the body. ", time:"06:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Floate ", treatmentdetails: "Folate is an important nutrient for making normal red blood cells", time:"07:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Full Blood ", treatmentdetails: "tiredness or weakness", time:"08:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "HbA1c", treatmentdetails: "HbA1c is a blood test that is used to help diagnose and monitor people with diabetes.", time:"09:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "D-dimer ", treatmentdetails: " D-dimer test is a blood test usually used to help check for or monitor blood clotting problems. ", time:"12:00pm",doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "ESR Test", treatmentdetails: "The erythrocyte sedimentation rate (ESR) test checks for inflammation in the body. ", time:"12:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Full Blood Count", treatmentdetails: "tiredness or weakness", time:"12:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "HbA1c", treatmentdetails: "HbA1c is a blood test that is used to help diagnose and monitor people with diabetes.", time:"12:00am", doctorsinvolved:"test result",post_observations:"xyz" });
			this.tempdata.push({ treatmentdone: "Vitamin B12 ", treatmentdetails: "You need vitamin B12 in your blood so you can make blood cells" , time:"12:00am", doctorsinvolved:"test result",post_observations:"xyz"});
	
	
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