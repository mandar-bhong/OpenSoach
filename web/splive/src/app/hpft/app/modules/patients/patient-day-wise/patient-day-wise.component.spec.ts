import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PatientDayWiseComponent } from './patient-day-wise.component';

describe('PatientDayWiseComponent', () => {
  let component: PatientDayWiseComponent;
  let fixture: ComponentFixture<PatientDayWiseComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PatientDayWiseComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PatientDayWiseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
