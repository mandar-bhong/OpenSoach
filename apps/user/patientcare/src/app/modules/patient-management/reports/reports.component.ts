import { Component, OnInit } from '@angular/core';
const permissions = require("nativescript-permissions");

@Component({
	moduleId: module.id,
	selector: 'reports',
	templateUrl: './reports.component.html',
	styleUrls: ['./reports.component.css']
})

export class ReportsComponent implements OnInit {

	constructor() { }

	ngOnInit() { 
		
		// check storage permission
		this.getStoragePermission();
	}

	getStoragePermission() {
		const hasPermission = permissions.hasPermission(android.Manifest.permission.WRITE_EXTERNAL_STORAGE);
		console.log("hasPermission:", hasPermission);
		if (hasPermission == false) {
			permissions.requestPermission(android.Manifest.permission.WRITE_EXTERNAL_STORAGE, "Need for storing file").then(() => {
				console.log("Permission granted!");
			}).catch(() => {
				console.log("Permission is not granted");
			});
		}
	}
}