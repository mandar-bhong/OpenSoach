import { Component, OnInit, Input, Output, EventEmitter, OnDestroy } from '@angular/core';
import { MedicalDetailsModel } from 'app/models/ui/patient-models';
import { Subscription } from 'rxjs';
import { PersonalHistoryInfo, WeightData, AlcoholData, SmokData, JSONBaseDataModel } from 'app/models/api/patient-models';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../../shared/views/edit-record-base';
import { FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-medical-personal-history',
  templateUrl: './medical-personal-history.component.html',
  styleUrls: ['./medical-personal-history.component.css']
})
export class MedicalPersonalHistoryComponent extends EditRecordBase implements OnInit, OnDestroy {

  @Input() itemPersonList: JSONBaseDataModel<PersonalHistoryInfo>;
  @Input() trueValue = true;
  @Input() falseValue = false;
  @Input() placeHolderTextPerson: string;
  @Input() headerTextPerson: string;
  @Output() onItemAddPerson = new EventEmitter();
  dataModel = new MedicalDetailsModel();
  routeSubscription: Subscription;
  contextValue: string;
  medicaldetialsid: number;
  weight: string;
  alcoholaplicable: string;
  alcoholquantity: string;
  alcoholcomment: string;
  smokingaplicable: string;
  smokingquantity: string;
  smokingcomment: string;
  other: string;
  tendency: string;
  alcoholcheck: boolean;
  smokCheck: boolean;
  constructor() {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Personal History';
  }

  ngOnInit() {
    this.createControls();
    this.tendency = 'Increasing';
    this.showBackButton = false;
    this.alcoholcheck = false;
    this.smokCheck = false;
    setTimeout(() => {
      console.log('this.itemPersonList',this.itemPersonList);
      if (Object.keys(this.itemPersonList).length > 0) {
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
        this.getData();
      } else {
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.EDITABLE);
        this.getData();
      }
    });
  }
  getData() {
    
    // setTimeout(() => {
    if (Object.keys(this.itemPersonList).length > 0) {

      this.weight = this.itemPersonList.data.weight.weight;
      this.tendency = this.itemPersonList.data.weight.weighttendency;

      this.alcoholcheck = this.itemPersonList.data.alcohol.aplicable;
      this.alcoholquantity = this.itemPersonList.data.alcohol.alcoholquantity;
      this.alcoholcomment = this.itemPersonList.data.alcohol.alcoholcomment;

      this.smokCheck = this.itemPersonList.data.smoking.aplicable;
      this.smokingquantity = this.itemPersonList.data.smoking.smokingquantity;
      this.smokingcomment = this.itemPersonList.data.smoking.smokingcomment;

      this.other = this.itemPersonList.data.other;
      this.recordState = EDITABLE_RECORD_STATE.UPDATE;
      this.setFormMode(FORM_MODE.VIEW);
    }
  // });
  }

  toggleVisibility() {
    if (this.alcoholcheck == false) {
      this.alcoholquantity = null;
    }
  }
  toggleVisibilitySmok(f) {
    if (this.smokCheck == false) {
      this.smokingquantity = null;
    }
  }

  itemAdd() {
    if (this.weight || this.tendency || this.alcoholquantity || this.alcoholcomment || this.smokingquantity || this.smokingcomment || this.other) {
      const personalHistoryData = new JSONBaseDataModel<PersonalHistoryInfo>();

      personalHistoryData.version = 1;

      personalHistoryData.data = new PersonalHistoryInfo();
      personalHistoryData.data.weight = new WeightData();
      personalHistoryData.data.weight.weight = this.weight;
      personalHistoryData.data.weight.weighttendency = this.tendency;

      personalHistoryData.data.alcohol = new AlcoholData();
      personalHistoryData.data.alcohol.aplicable = this.alcoholcheck;
      personalHistoryData.data.alcohol.alcoholquantity = this.alcoholquantity;
      personalHistoryData.data.alcohol.alcoholcomment = this.alcoholcomment;

      personalHistoryData.data.smoking = new SmokData();
      personalHistoryData.data.smoking.aplicable = this.smokCheck;
      personalHistoryData.data.smoking.smokingquantity = this.smokingquantity;
      personalHistoryData.data.smoking.smokingcomment = this.smokingcomment;

      personalHistoryData.data.other = this.other;

      this.onItemAddPerson.emit(personalHistoryData);
      this.recordState = EDITABLE_RECORD_STATE.UPDATE;
      this.setFormMode(FORM_MODE.VIEW);
    }
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      weightControls: new FormControl(''),
      tendencyControls: new FormControl(''),
      alcoholcheckControls: new FormControl(''),
      alcohalquantityControls: new FormControl(''),
      alcohalcommentControls: new FormControl(''),
      smokCheckControls: new FormControl(''),
      smokingquantityControls: new FormControl(''),
      smokingcommentControls: new FormControl(''),
      otherControls: new FormControl(''),
    });
  }
  closeForm() { }
  ngOnDestroy() {
    if (this.routeSubscription) {
      this.routeSubscription.unsubscribe();
    }
  }

}