import { Component, OnInit } from '@angular/core';
import * as camera from "nativescript-camera";
import { RouterExtensions } from "nativescript-angular/router";

import { takePicture, requestPermissions } from 'nativescript-camera';
import { ImageAsset } from 'tns-core-modules/image-asset';


import * as imagepicker from "nativescript-imagepicker";

@Component({
	moduleId: module.id,
	selector: 'cameras',
	templateUrl: './cameras.component.html',
	styleUrls: ['./cameras.component.css']
})

export class CamerasComponent implements OnInit {
	public saveToGallery: boolean = true;
    public keepAspectRatio: boolean = true;
    public width: number = 320;
    public height: number = 240;
    public cameraImage: ImageAsset;
    public actualWidth: number;
    public actualHeight: number;
    public scale: number = 1;
    public labelText: string;


    imageAssets = [];
    imageSrc: any;
    isSingleMode: boolean = true;
    thumbSize: number = 80;
    previewSize: number = 300;
    imgpicker: boolean = true;
    public removedImageUrl: string;

	constructor(private routerExtensions: RouterExtensions,) { }

	ngOnInit() {
        // alert('cameras component load ');
     }
	goBackPage() {
		// this.routerExtensions.back();
		this.routerExtensions.navigate(["/patientmgnt"], { clearHistory: true });
	}
    // public goBackPage() {
    //     this.routerExtensions.back();
    //     console.log("click back button");
    // }

    // take picture 
	onTakePictureTap(args) {
        requestPermissions().then(
            () => {
                takePicture({ width: this.width, height: this.height, keepAspectRatio: this.keepAspectRatio, saveToGallery: this.saveToGallery })
                    .then((imageAsset: any) => {
                        this.cameraImage = imageAsset;
                        let that = this;
                        imageAsset.getImageAsync(function (nativeImage) {
                            if (imageAsset.android) {
                                // get the current density of the screen (dpi) and divide it by the default one to get the scale
                                // that.scale = nativeImage.getDensity() / android.util.DisplayMetrics.DENSITY_DEFAULT;
                                that.actualWidth = nativeImage.getWidth();
                                that.actualHeight = nativeImage.getHeight();
                            } else {
                                that.scale = nativeImage.scale;
                                that.actualWidth = nativeImage.size.width * that.scale;
                                that.actualHeight = nativeImage.size.height * that.scale;
                            }
                            that.labelText = `Displayed Size: ${that.actualWidth}x${that.actualHeight} with scale ${that.scale}\n` +
                                `Image Size: ${Math.round(that.actualWidth / that.scale)}x${Math.round(that.actualHeight / that.scale)}`;

                            console.log(`${that.labelText}`);
                        });
                    }, (error) => {
                        console.log("Error: " + error);
                    });
            },
            () => alert('permissions rejected')
        );
	}





    // image file picker code start 


    public onSelectMultipleTap() {
        this.isSingleMode = false;

        let context = imagepicker.create({
            mode: "multiple"
        });
        this.startSelection(context);
    }

    public onSelectSingleTap() {
        this.isSingleMode = true;

        let context = imagepicker.create({
            mode: "single"
        });
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
        }).catch(function (e) {
            console.log(e);
        });
    }
    onRemoveImageButtonTap(): void {
        this.imageSrc = null;
        this.imageAssets = []

		// if (this.currentImageSource) {
		// 	this.currentImageSource = null;
		// }
	}
}