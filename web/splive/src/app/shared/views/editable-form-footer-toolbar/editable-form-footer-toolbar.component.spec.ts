import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditableFormFooterToolbarComponent } from './editable-form-footer-toolbar.component';

describe('EditableFormFooterToolbarComponent', () => {
  let component: EditableFormFooterToolbarComponent;
  let fixture: ComponentFixture<EditableFormFooterToolbarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditableFormFooterToolbarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditableFormFooterToolbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
