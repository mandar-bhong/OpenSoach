import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TreatmentReportComponent } from './treatment-report.component';

describe('TreatmentReportComponent', () => {
  let component: TreatmentReportComponent;
  let fixture: ComponentFixture<TreatmentReportComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TreatmentReportComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TreatmentReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
