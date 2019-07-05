import { Component, OnInit } from '@angular/core';
import { PatientListService } from '~/app/services/patient-list/patient-list.service';
import { PassDataService } from '~/app/services/pass-data-service';
import { JSONBaseDataModel } from '~/app/models/ui/json-base-data-model';
import { PersonAccompanyModel } from '~/app/models/ui/person-accompany-model';
import { PERSON_ACCMPANYING_GENDER } from '~/app/app-constants';

@Component({
	moduleId: module.id,
	selector: 'person-accompanying-details',
	templateUrl: './person-accompanying-details.component.html',
	styleUrls: ['./person-accompanying-details.component.css']
})

export class PersonAccompanyingDetailsComponent implements OnInit {
	personAccompanyDetailsModel = new PersonAccompanyModel();
	jsonField: JSONBaseDataModel<PersonAccompanyModel[]>;
	constructor(private patientListService: PatientListService,
		private passdataservice: PassDataService) { }

	ngOnInit() {
		this.getPersonAccompanyingByUUID();
	}
	public getPersonAccompanyingByUUID() {
		this.patientListService.getPersonAccompanyingByUUID(this.passdataservice.getAdmissionID()).then(
			(val) => {
				val.forEach(item => {
					this.jsonField = new JSONBaseDataModel<PersonAccompanyModel[]>();
					this.jsonField.data = [];
					Object.assign(this.jsonField, JSON.parse(item.person_accompanying));
					if (this.jsonField.data.length > 0) {
						switch (this.jsonField.data[0].gender) {
							case 0:
								this.jsonField.data[0].genderString = PERSON_ACCMPANYING_GENDER.GENDER_NOT_SELECTED;
								break;
							case 1:
								this.jsonField.data[0].genderString = PERSON_ACCMPANYING_GENDER.GENDER_MALE;
								break;
							case 2:
								this.jsonField.data[0].genderString = PERSON_ACCMPANYING_GENDER.GENDER_FEMALE;
								break;
						}
						this.personAccompanyDetailsModel = this.jsonField.data[0];

					}
				});
			},
			(error) => {
				console.log("person_accompanying error:", error);
			}
		);
	}
	prepareContact(contact, alternateContact) {
		let contactString;
		if (contact) {
			contactString = contact;
		}
		if (alternateContact) {
			contactString += ',' + alternateContact
		}
		return contactString;
	}
}