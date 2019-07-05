import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

import { EDITABLE_RECORD_STATE, EditRecordBase, FORM_MODE } from '../edit-record-base';


@Component({
  selector: 'app-editable-form-header',
  templateUrl: './editable-form-header.component.html',
  styleUrls: ['./editable-form-header.component.css']
})
export class EditableFormHeaderComponent implements OnInit {

  @Input()
  editRecordBase: EditRecordBase;
  @Output()
  editClick = new EventEmitter<null>();

  constructor() { }

  ngOnInit() {
  }

  edit() {
    this.editClick.emit();
  }

  close() {
    this.editRecordBase.closeForm();
  }

  cancel() {
    this.editRecordBase.setFormMode(FORM_MODE.VIEW);
    this.editRecordBase.onCancelHandler();
  }
}
