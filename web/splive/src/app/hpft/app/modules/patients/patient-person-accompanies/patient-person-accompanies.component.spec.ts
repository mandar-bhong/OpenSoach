import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PatientPersonAccompaniesComponent } from './patient-person-accompanies.component';

describe('PatientPersonAccompaniesComponent', () => {
  let component: PatientPersonAccompaniesComponent;
  let fixture: ComponentFixture<PatientPersonAccompaniesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PatientPersonAccompaniesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PatientPersonAccompaniesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
