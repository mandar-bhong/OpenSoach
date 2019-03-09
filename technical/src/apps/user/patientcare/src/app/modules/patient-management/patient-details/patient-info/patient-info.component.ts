import { Component, OnInit } from '@angular/core';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { PassDataService } from '~/app/services/pass-data-service';

export class PatientDetailsModel{
	fname: string;
	lname:string;
	mob_no: string;
	age: string;
	blood_grp: string;
	gender: number;
}
@Component({
	moduleId: module.id,
	selector: 'patient-info',
	templateUrl: './patient-info.component.html',
	styleUrls: ['./patient-info.component.css']
})

export class PatientInfoComponent implements OnInit {
	patientDetailsModel = new PatientDetailsModel();
	constructor(
		private patientListService: PatientListService,
		private passdataservice: PassDataService
	) { }

	ngOnInit() { 
		this.getAdmissionDetailsByUUID()
		
	}
	public getAdmissionDetailsByUUID() {
		this.patientListService.getPatientDetailsByUUID(this.passdataservice.getPatientID()).then(
			(val) => {
				val.forEach(item => {
					// console.log("Patient details", val);
					
					this.patientDetailsModel.fname = item.fname;
					this.patientDetailsModel.lname = item.lname;
					this.patientDetailsModel.blood_grp = item.blood_grp;
					this.patientDetailsModel.gender = item.gender;
					this.patientDetailsModel.age = item.age;
					this.patientDetailsModel.mob_no = item.mob_no;

					// console.log('admissionDetailsModel', this.patientDetailsModel);
				});
			},
			(error) => {
				console.log("Patient details error:", error);
			}
		);
	}
}