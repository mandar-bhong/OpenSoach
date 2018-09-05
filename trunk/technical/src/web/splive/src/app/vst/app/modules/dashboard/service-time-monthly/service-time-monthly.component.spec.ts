import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServiceTimeMonthlyComponent } from './service-time-monthly.component';

describe('ServiceTimeMonthlyComponent', () => {
  let component: ServiceTimeMonthlyComponent;
  let fixture: ComponentFixture<ServiceTimeMonthlyComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServiceTimeMonthlyComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServiceTimeMonthlyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
