import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PathologyReportComponent } from './pathology-report.component';

describe('PathologyReportComponent', () => {
  let component: PathologyReportComponent;
  let fixture: ComponentFixture<PathologyReportComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PathologyReportComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PathologyReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
