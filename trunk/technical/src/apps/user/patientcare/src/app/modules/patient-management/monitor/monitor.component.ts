import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'monitor',
	templateUrl: './monitor.component.html',
	styleUrls: ['./monitor.component.css']
})

export class MonitorComponent implements OnInit {
	categoricalSource: { Country: string, Amount: any }[] = [
		{ Country: "08:00 AM", Amount: 98 },
		{ Country: "12:00 AM", Amount: 98.3 },
		{ Country: "04:00 PM", Amount: 99 },
		{ Country: "08:00 PM", Amount: 101 },
		{ Country: "12:00 PM", Amount: 97 },
		{ Country: "04:00 AM", Amount: 100 }
	];
	bloodSource: { Country: string, Amount: any }[] = [
		{ Country: "08:00 AM", Amount: 70 },
		{ Country: "12:00 AM", Amount: 90 },
		{ Country: "04:00 PM", Amount: 75 },
		{ Country: "08:00 PM", Amount: 110 },
		{ Country: "12:00 PM", Amount: 80 },
		{ Country: "04:00 AM", Amount: 130 }
	];
	constructor() { }

	ngOnInit() {
		// alert('monitor');
	}

}