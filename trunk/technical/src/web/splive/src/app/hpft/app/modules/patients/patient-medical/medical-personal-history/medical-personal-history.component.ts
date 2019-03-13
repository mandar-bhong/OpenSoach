import { Component, OnInit, Input, Output, EventEmitter, OnDestroy } from '@angular/core';
import { MedicalDetailsModel } from 'app/models/ui/patient-models';
import { Subscription } from 'rxjs';
import { PersonalHistoryInfo } from 'app/models/api/patient-models';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../../shared/views/edit-record-base';
import { FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-medical-personal-history',
  templateUrl: './medical-personal-history.component.html',
  styleUrls: ['./medical-personal-history.component.css']
})
export class MedicalPersonalHistoryComponent extends EditRecordBase implements OnInit, OnDestroy {

  @Input() itemPersonList: PersonalHistoryInfo;

  @Input() placeHolderTextPerson: string;
  @Input() headerTextPerson: string;
  @Output() onItemAddPerson = new EventEmitter();
  dataModel = new MedicalDetailsModel();
  routeSubscription: Subscription;
  contextValue: string;
  medicaldetialsid: number;
  weight: string;
  alcohalquantity: string;
  alcohalcomment: string;
  smokingquantity: string;
  smokingcomment: string;
  other: string;
  tendency: string;
  alcoholcheck = false;
  smokCheck = false;
  constructor() {
    super();
    this.iconCss = 'fa fa-user';
    this.pageTitle = 'Personal History';
  }

  ngOnInit() {
    this.createControls();
    this.tendency = 'Increasing';
    this.showBackButton = false;
    setTimeout(() => {
      if (Object.keys(this.itemPersonList).length > 0) {
        // update mode
        this.getData();
        this.recordState = EDITABLE_RECORD_STATE.UPDATE;
        this.setFormMode(FORM_MODE.VIEW);
      } else {
        // add mode.
        this.recordState = EDITABLE_RECORD_STATE.ADD;
        this.setFormMode(FORM_MODE.EDITABLE);
      }
    });
  }
  getData() {
    this.weight = this.itemPersonList.weight;
    this.tendency = this.itemPersonList.weighttendency;
    this.alcohalquantity = this.itemPersonList.alcohalquantity;
    this.alcohalcomment = this.itemPersonList.alcohalcomment;
    this.smokingquantity = this.itemPersonList.smokingquantity;
    this.smokingcomment = this.itemPersonList.smokingcomment;
    this.other = this.itemPersonList.other;
    if (this.alcohalquantity != null) {
      this.alcoholcheck = true;
    }
    if (this.smokingquantity != null) {
      this.smokCheck = true;
    }
  }
  itemAdd() {
    if (this.weight || this.tendency || this.alcohalquantity || this.alcohalcomment || this.smokingquantity || this.smokingcomment || this.other) {
      const personalHistoryInfo = new PersonalHistoryInfo();
      personalHistoryInfo.weight = this.weight;
      personalHistoryInfo.weighttendency = this.tendency;
      personalHistoryInfo.alcohalquantity = this.alcohalquantity;
      personalHistoryInfo.alcohalcomment = this.alcohalcomment;
      personalHistoryInfo.smokingquantity = this.smokingquantity;
      personalHistoryInfo.smokingcomment = this.smokingcomment;
      personalHistoryInfo.other = this.other;
      this.onItemAddPerson.emit(personalHistoryInfo);
    }
    this.recordState = EDITABLE_RECORD_STATE.UPDATE;
    this.setFormMode(FORM_MODE.VIEW);
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      weightControls: new FormControl(''),
      tendencyControls: new FormControl(''),
      alcohalquantityControls: new FormControl(''),
      alcohalcommentControls: new FormControl(''),
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