import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddTreatmentReportComponent } from './add-treatment-report.component';

describe('AddTreatmentReportComponent', () => {
  let component: AddTreatmentReportComponent;
  let fixture: ComponentFixture<AddTreatmentReportComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddTreatmentReportComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddTreatmentReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
