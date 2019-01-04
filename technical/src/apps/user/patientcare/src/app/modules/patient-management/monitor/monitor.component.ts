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
        { Country: "12:00 PM", Amount1:99.5 },
        { Country: "04:00 PM", Amount1: 99.5 },
        { Country: "08:00 PM", Amount1: 99.5 },
        { Country: "12:00 AM", Amount1: 99.5 },
        { Country: "04:00 AM", Amount1: 99.5 },

    ];
    categoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "12:00 PM", Amount1:  97.7 },
        { Country: "04:00 PM", Amount1:  97.7 },
        { Country: "08:00 PM", Amount1:  97.7 },
        { Country: "12:00 AM", Amount1:  97.7 },
        { Country: "04:00 AM", Amount1:  97.7 },

    ];
    categoricalSource: { Country: string, Amount: number }[] = [
        { Country: "12:00 PM", Amount: 98.6 },
        { Country: "04:00 PM", Amount: 97.3 },
        { Country: "08:00 PM", Amount: 99.0},
        { Country: "12:00 AM", Amount: 98.0 },
        { Country: "04:00 AM", Amount: 99.8 }
    ];

    bloodSourcecategoricalSourcehigh: { Country: string, Amount1: number }[] = [
        { Country: "12:00 PM", Amount1: 140 },
        { Country: "04:00 PM", Amount1: 140 },
        { Country: "08:00 PM", Amount1: 140 },
        { Country: "12:00 AM", Amount1: 140 },
        { Country: "04:00 AM", Amount1: 140 },

    ];
    bloodSourcecategoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "12:00 PM", Amount1: 90 },
        { Country: "04:00 PM", Amount1: 90 },
        { Country: "08:00 PM", Amount1: 90 },
        { Country: "12:00 AM", Amount1: 90 },
        { Country: "04:00 AM", Amount1: 89 },

    ];
    bloodSourcecategoricalSource: { Country: string, Amount: number }[] = [
        { Country: "12:00 PM", Amount: 100 },
        { Country: "04:00 PM", Amount: 150 },
        { Country: "08:00 PM", Amount: 130},
        { Country: "12:00 AM", Amount: 87 },
        { Country: "04:00 AM", Amount: 110 }
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