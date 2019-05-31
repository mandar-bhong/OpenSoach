import { Component, OnInit } from '@angular/core';
import { ImageAsset } from 'tns-core-modules/image-asset/image-asset';
import { PatientListViewModel } from '~/app/models/ui/patient-view-models';
import { PassDataService } from '~/app/services/pass-data-service';
import { ObservableArray } from 'tns-core-modules/data/observable-array/observable-array';
import { TokenModel } from 'nativescript-ui-autocomplete';

@Component({
	moduleId: module.id,
	selector: 'show-uploaded-image',
	templateUrl: './show-uploaded-image.component.html',
	styleUrls: ['./show-uploaded-image.component.css']
})

export class ShowUploadedImageComponent implements OnInit {

	imageAsset: ImageAsset[] = [];
	public uploadedImage: ImageAsset;
	labelText: string;
	selectedPatient: PatientListViewModel;
	patientName: string;


	autocompleteReports: ObservableArray<TokenModel> = new ObservableArray<TokenModel>([
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
	constructor(private passdataservice: PassDataService) { }

	ngOnInit() {
       // for getting header name label text
		this.patientName = this.passdataservice.getHeaderName();
		// this.page.actionBarHidden = true;
		console.log('show camera image iniated', this.passdataservice.pickedImage);
		this.imageAsset = this.passdataservice.uploadedImage;
		if (this.imageAsset.length >= 1) {
			this.uploadedImage = this.imageAsset[0];
		}

		this.showImage();
	}

	showImage() {

		// that.imageSrc = that.isSingleMode && selection.length > 0 ? selection[0] : null;

		// // set the images to be loaded from the assets with optimal sizes (optimize memory usage)
		// selection.forEach(function (element) {
		// 	element.options.width = that.isSingleMode ? that.previewSize : that.thumbSize;
		// 	element.options.height = that.isSingleMode ? that.previewSize : that.thumbSize;
		// });
		// this.imageAsset = selection;


		// this.cameraImage=this.imageAsset;
		// this.imageAsset.getImageAsync(function (nativeImage) {
		// 	if (this.imageAsset.android) {
		// 		// get the current density of the screen (dpi) and divide it by the default one to get the scale
		// 		// that.scale = nativeImage.getDensity() / android.util.DisplayMetrics.DENSITY_DEFAULT;
		// 		this.imageAsset.actualWidth = nativeImage.getWidth();
		// 		this.imageAsset.actualHeight = nativeImage.getHeight();
		// 	} else {
		// 		this.imageAsset.scale = nativeImage.scale;
		// 		this.imageAsset.actualWidth = nativeImage.size.width * this.imageAsset.scale;
		// 		this.imageAsset.actualHeight = nativeImage.size.height * this.imageAsset.scale;
		// 	}
		// 	this.labelText = `Displayed Size: ${this.imageAsset.actualWidth}x${this.imageAsset.actualHeight} with scale ${this.imageAsset.scale}\n` +`Image Size: ${Math.round(this.imageAsset.actualWidth / this.imageAsset.scale)}x${Math.round(this.imageAsset.actualHeight / this.imageAsset.scale)}`;
		// });
	}

}
