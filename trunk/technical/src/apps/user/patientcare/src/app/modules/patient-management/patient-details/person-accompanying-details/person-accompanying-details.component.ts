import { Component, OnInit } from '@angular/core';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { PassDataService } from '~/app/services/pass-data-service';

export class PersonAccompanyDetailsModel {
	personaccompany: PersonDetailsModel;
	age: string;
}
export class PersonDetailsModel {
	name: string;
	gender: number;
	personaddress: string;
	relationshipwithpatient: string;
	contact: string;
	alternatecontact: string;
	mobile_no: string;
}


export class JSONBaseDataModel<T> {
	version: number;
	data: T;
}
@Component({
	moduleId: module.id,
	selector: 'person-accompanying-details',
	templateUrl: './person-accompanying-details.component.html',
	styleUrls: ['./person-accompanying-details.component.css']
})

export class PersonAccompanyingDetailsComponent implements OnInit {
	personAccompanyDetailsModel = new PersonDetailsModel();
	jsonField;
	constructor(private patientListService: PatientListService,
		private passdataservice: PassDataService) { }

	ngOnInit() {
		this.getPersonAccompanyingByUUID();
		this.jsonField = new JSONBaseDataModel<PersonDetailsModel[]>();
	}
	public getPersonAccompanyingByUUID() {
		this.patientListService.getPersonAccompanyingByUUID(this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					// console.log("person_accompanying", item.person_accompanying);
					this.jsonField = new JSONBaseDataModel<PersonDetailsModel[]>();
					this.jsonField.data = [];
					Object.assign(this.jsonField, JSON.parse(item.person_accompanying));
					// console.log('JSON get data', this.jsonField.data);
					if (this.jsonField.data.length > 0) {
						// console.log('JSON', this.jsonField.data);
						this.personAccompanyDetailsModel = new PersonDetailsModel();
						this.personAccompanyDetailsModel.name = this.jsonField.data[0].name;
						this.personAccompanyDetailsModel.contact = this.jsonField.data[0].contact;
						this.personAccompanyDetailsModel.alternatecontact = this.jsonField.data[0].alternatecontact;
						this.personAccompanyDetailsModel.personaddress = this.jsonField.data[0].personaddress;
						this.personAccompanyDetailsModel.relationshipwithpatient = this.jsonField.data[0].relationshipwithpatient;
						this.personAccompanyDetailsModel.gender = this.jsonField.data[0].gender;
						// const testdata = this.jsonField.data[0].contact;
						// console.log('testdata', testdata);
					}
				});
			},
			(error) => {
				console.log("person_accompanying error:", error);
			}
		);
	}
}