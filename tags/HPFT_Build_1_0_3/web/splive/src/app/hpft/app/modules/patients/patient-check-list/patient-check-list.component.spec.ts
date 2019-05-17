import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PatientCheckListComponent } from './patient-check-list.component';

describe('PatientCheckListComponent', () => {
  let component: PatientCheckListComponent;
  let fixture: ComponentFixture<PatientCheckListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PatientCheckListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PatientCheckListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
