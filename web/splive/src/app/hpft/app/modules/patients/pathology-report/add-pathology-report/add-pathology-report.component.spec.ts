import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddPathologyReportComponent } from './add-pathology-report.component';

describe('AddPathologyReportComponent', () => {
  let component: AddPathologyReportComponent;
  let fixture: ComponentFixture<AddPathologyReportComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddPathologyReportComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddPathologyReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
