import { Component, OnInit } from '@angular/core';

@Component({
	moduleId: module.id,
	selector: 'monitor',
	templateUrl: './monitor.component.html',
	styleUrls: ['./monitor.component.css']
})

export class MonitorComponent implements OnInit {
	// categoricalSource: { Country: string, Amount: any }[] = [
	// 	{ Country: "08:00 AM", Amount: 98 },
	// 	{ Country: "12:00 AM", Amount: 98.3 },
	// 	{ Country: "04:00 PM", Amount: 99 },
	// 	{ Country: "08:00 PM", Amount: 101 },
	// 	{ Country: "12:00 PM", Amount: 97 },
	// 	{ Country: "04:00 AM", Amount: 100 }
	// ];
	categoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "08:00 AM", Amount1: 20 },
        { Country: "12:00 AM", Amount1: 20 },
        { Country: "04:00 PM", Amount1: 20 },
        { Country: "12:00 PM", Amount1: 20 },
        { Country: "04:00 AM", Amount1: 20 },

    ];
    categoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "08:00 AM", Amount1: 5 },
        { Country: "12:00 AM", Amount1: 5 },
        { Country: "04:00 PM", Amount1: 5 },
        { Country: "12:00 PM", Amount1:5 },
        { Country: "04:00 AM", Amount1: 5 },

    ];
    categoricalSource: { Country: string, Amount: number, SecondVal: number, ThirdVal: number }[] = [
        { Country: "08:00 AM", Amount: 15, SecondVal: 14, ThirdVal: 24 },
        { Country: "12:00 AM", Amount: 13, SecondVal: 23, ThirdVal: 25 },
        { Country: "04:00 PM", Amount: 24, SecondVal: 17, ThirdVal: 23 },
        { Country: "12:00 PM", Amount: 11, SecondVal: 19, ThirdVal: 24 },
        { Country: "04:00 AM", Amount: 18, SecondVal: 8, ThirdVal: 21 }
	];
	
	bloodSourcecategoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "08:00 AM", Amount1: 80 },
        { Country: "12:00 AM", Amount1: 80 },
        { Country: "04:00 PM", Amount1: 80 },
        { Country: "12:00 PM", Amount1: 80 },
        { Country: "04:00 AM", Amount1: 80 },

    ];
    bloodSourcecategoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "08:00 AM", Amount1: 25 },
        { Country: "12:00 AM", Amount1: 25 },
        { Country: "04:00 PM", Amount1: 25 },
        { Country: "12:00 PM", Amount1: 25 },
        { Country: "04:00 AM", Amount1: 25 },

    ];
    bloodSourcecategoricalSource: { Country: string, Amount: number, SecondVal: number, ThirdVal: number }[] = [
        { Country: "08:00 AM", Amount: 40, SecondVal: 70, ThirdVal: 24 },
        { Country: "12:00 AM", Amount: 50, SecondVal: 23, ThirdVal: 25 },
        { Country: "04:00 PM", Amount: 40, SecondVal: 17, ThirdVal: 23 },
        { Country: "12:00 PM", Amount: 80, SecondVal: 19, ThirdVal: 24 },
        { Country: "04:00 AM", Amount: 70, SecondVal: 8, ThirdVal: 21 }
	];
	

	// bloodSource: { Country: string, Amount: any }[] = [
	// 	{ Country: "08:00 AM", Amount: 70 },
	// 	{ Country: "12:00 AM", Amount: 90 },
	// 	{ Country: "04:00 PM", Amount: 75 },
	// 	{ Country: "08:00 PM", Amount: 110 },
	// 	{ Country: "12:00 PM", Amount: 80 },
	// 	{ Country: "04:00 AM", Amount: 130 }
	// ];
	constructor() { }

	ngOnInit() {
		// alert('monitor');
	}

}