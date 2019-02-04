import { Component, OnInit } from '@angular/core';
import { Schedulardata, SchedularConfigData, MornFreqInfo, AftrnFreqInfo, NightFreqInfo } from '~/app/models/ui/chart-models';
import { MedicineHelper } from '~/app/helpers/actions/medicine-helper';
import { MonitorHelper } from '~/app/helpers/actions/monitor-helper';
import { IntakeHelper } from '~/app/helpers/actions/intake-helper';
import { ScheduleDatastoreModel } from '~/app/models/db/schedule-model';
import { ConfigCodeType } from '~/app/app-constants';

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
        { Country: "12:00 PM", Amount1: 99.5 },
        { Country: "04:00 PM", Amount1: 99.5 },
        { Country: "08:00 PM", Amount1: 99.5 },
        { Country: "12:00 AM", Amount1: 99.5 },
        { Country: "04:00 AM", Amount1: 99.5 },

    ];
    categoricalSourcelow: { Country: string, Amount1: number }[] = [
        { Country: "12:00 PM", Amount1: 97.7 },
        { Country: "04:00 PM", Amount1: 97.7 },
        { Country: "08:00 PM", Amount1: 97.7 },
        { Country: "12:00 AM", Amount1: 97.7 },
        { Country: "04:00 AM", Amount1: 97.7 },

    ];
    categoricalSource: { Country: string, Amount: number }[] = [
        { Country: "12:00 PM", Amount: 98.6 },
        { Country: "04:00 PM", Amount: 97.3 },
        { Country: "08:00 PM", Amount: 99.0 },
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
        { Country: "08:00 PM", Amount: 130 },
        { Country: "12:00 AM", Amount: 87 },
        { Country: "04:00 AM", Amount: 110 }
    ];
    schedulardata: Schedulardata;


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
      //  this.createActions();
    }
    createActions() {
        this.schedulardata = new Schedulardata();
        const dt = new Date();

        const scheduleDatastoreModel = new ScheduleDatastoreModel();
        scheduleDatastoreModel.uuid = '12';
        scheduleDatastoreModel.admission_uuid = '2';
        scheduleDatastoreModel.conf_type_code = ConfigCodeType.MEDICINE;
        scheduleDatastoreModel.conf = '';
        this.schedulardata.conf = new SchedularConfigData();
        this.schedulardata.conf.mornFreqInfo = new MornFreqInfo();
        this.schedulardata.conf.mornFreqInfo.freqMorn = true

        this.schedulardata.conf.aftrnFreqInfo = new AftrnFreqInfo();
        this.schedulardata.conf.aftrnFreqInfo.freqAftrn = true;
        this.schedulardata.conf.nightFreqInfo = new NightFreqInfo();
        this.schedulardata.conf.nightFreqInfo.freqNight = true;
        this.schedulardata.conf.desc = " Morning & Afternoon & Night before meal. \nTest.";
        this.schedulardata.conf.name = "Cipla";
        this.schedulardata.conf.quantity = 11;
        this.schedulardata.conf.startDate = "2019-01-29 12:04" // 23th 12.10 pm
        this.schedulardata.conf.duration = 3;
        this.schedulardata.conf.frequency = 1;
        this.schedulardata.conf.startTime = '21.30' // 16.30
        this.schedulardata.conf.intervalHrs = 180
        this.schedulardata.conf.foodInst = 1;
        this.schedulardata.conf.endTime = '12.30';
        this.schedulardata.conf.numberofTimes = 3
        this.schedulardata.conf.specificTimes = [];
        this.schedulardata.conf.specificTimes.push(11.30);
        this.schedulardata.conf.specificTimes.push(12.30);
        this.schedulardata.data = scheduleDatastoreModel;


        let medicineSchedular = new MedicineHelper();

        //	let mmonitorHelper = new MonitorHelper();
        //	mmonitorHelper.createMonitorActions(this.schedulardata);
        //	let intakeHelper = new IntakeHelper();
        //	intakeHelper.createIntakeActions(this.schedulardata);
        const test = medicineSchedular.createMedicineActions(this.schedulardata);
        console.log('received actions data',test);
        //	console.log('in create actions fucntions');
        // this.chartService.getChartList().then(
        // 	(val) => {
        // 		val.forEach(item => {
        // 			item.conf = JSON.parse(item.conf);
        // 		});					
        // 		if (this.schedulardata.conf_type_code = medicine) {
        // 			if (this.schedulardata.conf.frequency === freuencyzero) {

        // 			}
        // 		}				
        // 	},
        // 	(error) => {
        // 		console.log("getChartData error:", error);
        // 	}
        // );
    }
}