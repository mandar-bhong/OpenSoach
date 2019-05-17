import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditableFormHeaderComponent } from './editable-form-header.component';

describe('EditableFormHeaderComponent', () => {
  let component: EditableFormHeaderComponent;
  let fixture: ComponentFixture<EditableFormHeaderComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditableFormHeaderComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditableFormHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
