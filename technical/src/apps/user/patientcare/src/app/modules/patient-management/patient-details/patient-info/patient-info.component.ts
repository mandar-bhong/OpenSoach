import { Component, OnInit } from '@angular/core';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { PatientDetailsModel } from '~/app/models/ui/patient-details';
import { PERSON_ACCMPANYING_GENDER } from '~/app/app-constants';


@Component({
	moduleId: module.id,
	selector: 'patient-info',
	templateUrl: './patient-info.component.html',
	styleUrls: ['./patient-info.component.css']
})

export class PatientInfoComponent implements OnInit {
	patientDetailsModel = new PatientDetailsModel();
	patientAddress: string;
	constructor(
		private patientListService: PatientListService,
		private passdataservice: PassDataService
	) { }

	ngOnInit() {
		this.getAdmissionDetailsByUUID()
		this.getPersonAccompanyingByUUID();
	}
	public getAdmissionDetailsByUUID() {
		this.patientListService.getPatientDetailsByUUID(this.passdataservice.getPatientID()).then(
			(val) => {
				val.forEach(item => {

					this.patientDetailsModel.fname = item.fname;
					this.patientDetailsModel.lname = item.lname;
					this.patientDetailsModel.blood_grp = item.blood_grp;

					this.patientDetailsModel.age = item.age;
					this.patientDetailsModel.mob_no = item.mob_no;
					let gender: any;
					if (item.gender != null) {
						switch (item.gender) {
							case 0:
								gender = PERSON_ACCMPANYING_GENDER.GENDER_NOT_SELECTED;
								break;
							case 1:
								gender = PERSON_ACCMPANYING_GENDER.GENDER_MALE;
								break;
							case 2:
								gender = PERSON_ACCMPANYING_GENDER.GENDER_FEMALE;
								break;
						}
						this.patientDetailsModel.gender = gender;
					}
				});
			},
			(error) => {
				console.log("Patient details error:", error);
			}
		);
	}
	public getPersonAccompanyingByUUID() {
		this.patientListService.getPersonAccompanyingByUUID(this.passdataservice.getAdmissionID()).then(
			(val) => {				
				if (val.length > 0 && val[0].other_details != null) {
					const address = JSON.parse(val[0].other_details);
					if (address && address.data) {
						this.patientAddress = address.data.address;
						console.log('this.patientAddress',	this.patientAddress);
					}
				}				
			},
			(error) => {
				console.log("person_accompanying error:", error);
			}
		);
	}
}