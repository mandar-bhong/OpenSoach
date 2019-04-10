import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PatientCheckSearchComponent } from './patient-check-search.component';

describe('PatientCheckSearchComponent', () => {
  let component: PatientCheckSearchComponent;
  let fixture: ComponentFixture<PatientCheckSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PatientCheckSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PatientCheckSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
