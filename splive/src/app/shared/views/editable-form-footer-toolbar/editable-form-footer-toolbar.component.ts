import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

import { EditRecordBase, FORM_MODE } from '../edit-record-base';

@Component({
  selector: 'app-editable-form-footer-toolbar',
  templateUrl: './editable-form-footer-toolbar.component.html',
  styleUrls: ['./editable-form-footer-toolbar.component.css']
})
export class EditableFormFooterToolbarComponent implements OnInit {

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
  }
}
