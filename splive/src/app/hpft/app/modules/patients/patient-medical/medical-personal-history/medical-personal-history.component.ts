import { Component, OnInit, Input, Output, EventEmitter, OnDestroy } from '@angular/core';
// import { MedicalDetailsModel } from 'app/models/ui/patient-models';
import { Subscription } from 'rxjs';
import { PersonalHistoryInfo, WeightData, AlcoholData, SmokData, JSONBaseDataModel } from '../../../../../app/models/api/patient-data-models';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../../shared/views/edit-record-base';
import { FormGroup, FormControl } from '@angular/forms';
import { MedicalDetailsModel } from '../../../../models/ui/patient-models';

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
    this.iconCss = 'fa fa-address-card-o';
    this.pageTitle = 'Personal History';
  }

  ngOnInit() {
    this.createControls();
    this.tendency = 'Increasing';
    this.showBackButton = false;
    this.alcoholcheck = false;
    this.smokCheck = false;
    setTimeout(() => {
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
    if (Object.keys(this.itemPersonList).length > 0) {
      this.dataList();
      this.recordState = EDITABLE_RECORD_STATE.UPDATE;
      this.setFormMode(FORM_MODE.VIEW);
    }
  }

  dataList() {
    this.weight = this.itemPersonList.data.weight.weight;
    this.tendency = this.itemPersonList.data.weight.weight_tendency;
    this.alcoholcheck = this.itemPersonList.data.alcohol.applicable;
    this.alcoholquantity = this.itemPersonList.data.alcohol.quantity;
    this.alcoholcomment = this.itemPersonList.data.alcohol.remarks;
    this.smokCheck = this.itemPersonList.data.smoking.applicable;
    this.smokingquantity = this.itemPersonList.data.smoking.quantity;
    this.smokingcomment = this.itemPersonList.data.smoking.remarks;
    this.other = this.itemPersonList.data.others;
  }

  onCancelHandler() {
    this.dataList();
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
      personalHistoryData.data.weight.weight_tendency = this.tendency;

      personalHistoryData.data.alcohol = new AlcoholData();
      personalHistoryData.data.alcohol.applicable = this.alcoholcheck;
      personalHistoryData.data.alcohol.quantity = this.alcoholquantity;
      personalHistoryData.data.alcohol.remarks = this.alcoholcomment;

      personalHistoryData.data.smoking = new SmokData();
      personalHistoryData.data.smoking.applicable = this.smokCheck;
      personalHistoryData.data.smoking.quantity = this.smokingquantity;
      personalHistoryData.data.smoking.remarks = this.smokingcomment;

      personalHistoryData.data.others = this.other;

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