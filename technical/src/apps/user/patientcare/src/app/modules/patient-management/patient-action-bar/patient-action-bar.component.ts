import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';
import { RouterExtensions } from 'nativescript-angular/router';
import { PassDataService } from '~/app/services/pass-data-service';
import { alert } from "tns-core-modules/ui/dialogs";
import {
	CFAlertDialog,
	DialogOptions,
	CFAlertActionAlignment,
	CFAlertActionStyle,
	CFAlertStyle
} from "nativescript-cfalert-dialog";
@Component({
	moduleId: module.id,
	selector: 'patient-action-bar',
	templateUrl: './patient-action-bar.component.html',
	styleUrls: ['./patient-action-bar.component.css']
})

export class PatientActionBarComponent implements OnInit {
	//  process variable 
	patientname: string;
	private cfalertDialog: CFAlertDialog;
	constructor(private routerExtensions: RouterExtensions,
		private passDataService: PassDataService) {
		this.cfalertDialog = new CFAlertDialog();
	}
	@Input() patientName: string;
	ngOnInit() {
		this.patientname = this.patientName;
	}

	goBackPage() {
		if (this.passDataService.backalert.length) {
			this.showNotification();
		} else {
			this.routerExtensions.back();
		}

		console.log('this.passDataService.backalert.length', this.passDataService.backalert.length);
	}

	patientdetail() {
		this.routerExtensions.navigate(['patientmgnt', 'patient'], { clearHistory: false });
	}

	showNotification(): void {
		let onSelection = response => {
			this.routerExtensions.back();
		};
		let onSelection1 = response => {
		};
		const options: DialogOptions = {
			dialogStyle: CFAlertStyle.NOTIFICATION,
			title: "Unsaved Changes!",
			message: "Do you wish to Disactive changes?",
			backgroundBlur: true,
			onDismiss: () => console.log("showAlert dismissed"),
			buttons: [
				{
					text: "Yes",
					buttonStyle: CFAlertActionStyle.POSITIVE,
					buttonAlignment: CFAlertActionAlignment.JUSTIFIED,
					onClick: onSelection
				},
				{
					text: "No, Thanks.",
					buttonStyle: CFAlertActionStyle.NEGATIVE,
					buttonAlignment: CFAlertActionAlignment.JUSTIFIED,
					onClick: onSelection1
				}]
		};
		this.cfalertDialog.show(options);
	}
}