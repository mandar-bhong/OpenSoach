import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PatientsPersonalDetailComponent } from './patients-personal-detail.component';

describe('PatientsPersonalDetailComponent', () => {
  let component: PatientsPersonalDetailComponent;
  let fixture: ComponentFixture<PatientsPersonalDetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PatientsPersonalDetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PatientsPersonalDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
