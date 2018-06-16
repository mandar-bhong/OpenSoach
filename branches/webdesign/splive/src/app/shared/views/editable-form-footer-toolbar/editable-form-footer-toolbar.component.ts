import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

import { FORM_MODE } from '../edit-record-base';

@Component({
  selector: 'app-editable-form-footer-toolbar',
  templateUrl: './editable-form-footer-toolbar.component.html',
  styleUrls: ['./editable-form-footer-toolbar.component.css']
})
export class EditableFormFooterToolbarComponent implements OnInit {

  @Input()
  formMode: FORM_MODE;
  @Output()
  editClick = new EventEmitter<null>();
  @Output()
  closeClick = new EventEmitter<null>();

  constructor() { }

  ngOnInit() {
  }

  edit() {
    this.editClick.emit();
  }

  close() {
    this.closeClick.emit();
  }
}
