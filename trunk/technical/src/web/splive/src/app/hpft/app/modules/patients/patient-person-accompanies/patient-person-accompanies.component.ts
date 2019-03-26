import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { TranslatePipe } from '../../../../../shared/pipes/translate/translate.pipe';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { PatientPersonalDetails } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { PersonAccompanyingInfo, JSONBaseDataModel, PersonalDetailsRequest } from 'app/models/api/patient-data-models';

@Component({
  selector: 'app-patient-person-accompanies',
  templateUrl: './patient-person-accompanies.component.html',
  styleUrls: ['./patient-person-accompanies.component.css']
})
export class PatientPersonAccompaniesComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new PatientPersonalDetails();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  personGender: EnumDataSourceItem<number>[];
  contact: string;
  personAccompanyingInfo = new JSONBaseDataModel<PersonAccompanyingInfo>();
  name: string;
  age: number;
  address: string;
  gender: any;
  relationshipwithpatient: string;
  alternatecontact: string;
  personaldetailsid: number;
  data: string[];
  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private router: Router,
    private appNotificationService: AppNotificationService,
    private translatePipe: TranslatePipe,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Person Accompanying';
  }

  ngOnInit() {
    this.createControls();
    this.showBackButton = false;
    this.personGender = this.patientService.getPersonGender();
    this.personAccompanyingInfo = new JSONBaseDataModel<PersonAccompanyingInfo>();
    this.routeSubscription = this.route.queryParams.subscribe(params => {
      if (params['admissionid']) {
        this.dataModel.admissionid = Number(params['admissionid']);
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        if (this.patientService.admissionid) {
          this.getPatientPersonalDetailId();
        }
      } else if (this.patientService.admissionid) {
        this.getPatientPersonalDetailId();
        {
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.EDITABLE);
        }
      }

      this.callbackUrl = params['callbackurl'];
    });
    this.dataModel.personAccompanyingData = new JSONBaseDataModel<PersonAccompanyingInfo[]>();
    this.dataModel.personAccompanyingData.data = [];
  }

  //Save function for Family History.
  save() {
    const patientPersonAccompanyingDetail = new PersonalDetailsRequest();
    patientPersonAccompanyingDetail.admissionid = this.patientService.admissionid;
    patientPersonAccompanyingDetail.patientid = this.patientService.patientid;
    patientPersonAccompanyingDetail.personaldetailsid = this.personaldetailsid;
    if (this.name || this.age || this.gender || this.address || this.relationshipwithpatient || this.contact || this.alternatecontact) {
      const personAccompanyingInfo = new JSONBaseDataModel<PersonAccompanyingInfo>();
      personAccompanyingInfo.data = new PersonAccompanyingInfo();
      personAccompanyingInfo.data.name = this.name;
      personAccompanyingInfo.data.age = this.age;
      personAccompanyingInfo.data.gender = this.gender;
      personAccompanyingInfo.data.address = this.address;
      personAccompanyingInfo.data.relationshipwithpatient = this.relationshipwithpatient;
      personAccompanyingInfo.data.contact = this.contact;
      personAccompanyingInfo.data.alternatecontact = this.alternatecontact;
      const tempPerson = new JSONBaseDataModel<PersonAccompanyingInfo[]>();
      tempPerson.data = [];
      tempPerson.data.push(personAccompanyingInfo.data);
      patientPersonAccompanyingDetail.personaccompanying = JSON.stringify(tempPerson);
      this.patientService.personalAddAccompanying(patientPersonAccompanyingDetail).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
        }
      });
    }
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      personnameControl: new FormControl('', [Validators.required]),
      personageControl: new FormControl(''),
      genderControl: new FormControl('', [Validators.required]),
      personaladdressControl: new FormControl(''),
      reletionpatientsControl: new FormControl(''),
      mobilenoControl: new FormControl('', [Validators.required]),
      alternateContactControl: new FormControl('')
    });
  }

  //Getting data from database
  getPatientPersonalDetailId() {
    this.patientService.getPatientMedicalID({ recid: this.patientService.admissionid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.personaldetailsid = payloadResponse.data.personaldetails.personaldetailsid;
          if (this.personaldetailsid) {
            this.getPatientPersonDetails();

          }
        }
      }
    });
  }
  //Getting data from database
  getPatientPersonDetails() {
    this.patientService.getPatientPersonalDetail({ recid: this.personaldetailsid }).subscribe(payloadResponse => {
      if (payloadResponse && payloadResponse.issuccess) {
        if (payloadResponse.data) {
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.personaldetailsid = payloadResponse.data.personaldetailsid;
          this.personaldetailsid = payloadResponse.data.personaldetailsid;
          const personAccompanyingInfo = new JSONBaseDataModel<PersonAccompanyingInfo[]>();
          if (payloadResponse.data.personaccompanying != null) {
            const tempPersonAccompanying = JSON.parse(payloadResponse.data.personaccompanying);
            personAccompanyingInfo.data = [];
            personAccompanyingInfo.data = tempPersonAccompanying.data || null;
            personAccompanyingInfo.version = tempPersonAccompanying.version;
            this.name = personAccompanyingInfo.data[0].name;
            this.age = personAccompanyingInfo.data[0].age;
            this.gender = personAccompanyingInfo.data[0].gender;
            this.address = personAccompanyingInfo.data[0].address;
            this.relationshipwithpatient = personAccompanyingInfo.data[0].relationshipwithpatient;
            this.contact = personAccompanyingInfo.data[0].contact;
            this.alternatecontact = personAccompanyingInfo.data[0].alternatecontact;
          }
        }
      }
    });
  }

  closeForm() { }

  //gender value
  getgender(value: number) {
    if (this.personGender && value) {
      return this.personGender.find(a => a.value === value).text;
    }
  }

  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }
}
