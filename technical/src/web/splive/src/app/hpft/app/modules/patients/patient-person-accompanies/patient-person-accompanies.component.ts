import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { JSONBaseDataModel, PersonAccompanyingInfo, PersonalDetailsRequest } from 'app/models/api/patient-data-models';
import { Subscription } from 'rxjs';
import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { AppNotificationService } from '../../../../../shared/services/notification/app-notification.service';
import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../../../../../shared/views/edit-record-base';
import { PatientPersonalDetails } from '../../../models/ui/patient-models';
import { PatientService } from '../../../services/patient.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-patient-person-accompanies',
  templateUrl: './patient-person-accompanies.component.html',
  styleUrls: ['./patient-person-accompanies.component.css']
})
export class PatientPersonAccompaniesComponent extends EditRecordBase implements OnInit, OnDestroy {

  dataModel = new PatientPersonalDetails();
  dataModelOrg = new PatientPersonalDetails();
  routeSubscription: Subscription;
  patientStates: EnumDataSourceItem<number>[];
  personGender: EnumDataSourceItem<number>[];
  contact: string;
  personAccompanyingInfo = new PersonAccompanyingInfo();
  personAccompanyingInfoArray = new JSONBaseDataModel<PersonAccompanyingInfo[]>();
  personAccompanyDataItem = new JSONBaseDataModel<PersonAccompanyingInfo>();
  
  personAccompanyDataItemOrg = new JSONBaseDataModel<PersonAccompanyingInfo>();
  personaldetailsid: number;
  data: string[];
  constructor(
    private patientService: PatientService,
    private route: ActivatedRoute,
    private appNotificationService: AppNotificationService,
  ) {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Person Accompanying';
  }

  ngOnInit() {
    this.createControls();

    this.showBackButton = false;
    this.personGender = this.patientService.getPersonGender();
    this.personAccompanyDataItem.data = new PersonAccompanyingInfo();    
    this.personAccompanyDataItemOrg = this.deepClone(this.personAccompanyDataItem);
    this.personAccompanyingInfoArray.data = [] ;
      this.routeSubscription = this.route.queryParams.subscribe(params => {
        if (params['admissionid']) {
          this.dataModel.admissionid = Number(params['admissionid']);
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          if (this.patientService.admissionid) {
            this.getPatientPersonalDetailId();
          }
        } else if (this.patientService.admissionid) {
          if(this.getPatientPersonalDetailId() == null)
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


  deepClone(source):any{
    return JSON.parse( JSON.stringify(source));      
  }

  //Save function for Family History.
  save() {
    if (this.editableForm.invalid) { return; }
    this.inProgress = true;
    if (this.recordState === EDITABLE_RECORD_STATE.UPDATE) {
      const patientPersonAccompanyingDetail = new PersonalDetailsRequest();
      patientPersonAccompanyingDetail.admissionid = this.patientService.admissionid;
      patientPersonAccompanyingDetail.patientid = this.patientService.patientid;
      patientPersonAccompanyingDetail.personaldetailsid = this.personaldetailsid;


      if(JSON.stringify(this.personAccompanyDataItem) === JSON.stringify( this.personAccompanyDataItemOrg )){
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.inProgress = false;
        return;
      }

      if(this.personAccompanyingInfoArray.data.length > 0){
       this.personAccompanyingInfoArray.data[0] = this.personAccompanyDataItem.data;
      }else{
        this.personAccompanyingInfoArray.data.push(this.personAccompanyDataItem.data);
      }
      patientPersonAccompanyingDetail.personaccompanying = JSON.stringify(this.personAccompanyingInfoArray);
      this.patientService.personalAddAccompanying(patientPersonAccompanyingDetail).subscribe(payloadResponse => {
        if (payloadResponse && payloadResponse.issuccess) {
          this.patientService.personaldetailsid = this.dataModel.personaldetailsid;
          this.appNotificationService.success();
          this.recordState = EDITABLE_RECORD_STATE.UPDATE;
          this.setFormMode(FORM_MODE.VIEW);
          this.personAccompanyDataItemOrg = this.deepClone(this.personAccompanyDataItem);
          this.inProgress = false;
        }else{
          this.inProgress = false;
        }
      });

      }
    }


  onCancelHandler() {
    Object.assign( this.personAccompanyDataItem.data,this.personAccompanyDataItemOrg.data);
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
          
          if (payloadResponse.data.personaccompanying != null) {
            this.personAccompanyingInfoArray = JSON.parse(payloadResponse.data.personaccompanying);
            
            if (this.personAccompanyingInfoArray.data.length > 0) {
              this.personAccompanyDataItem.data =  this.personAccompanyingInfoArray.data[0]; 
              this.personAccompanyDataItemOrg = this.deepClone(this.personAccompanyDataItem);               
            }
          }else{
            this.recordState = EDITABLE_RECORD_STATE.UPDATE;
            this.setFormMode(FORM_MODE.EDITABLE);
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
