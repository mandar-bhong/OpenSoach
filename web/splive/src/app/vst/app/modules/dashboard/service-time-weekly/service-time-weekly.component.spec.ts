import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServiceTimeWeeklyComponent } from './service-time-weekly.component';

describe('ServiceTimeWeeklyComponent', () => {
  let component: ServiceTimeWeeklyComponent;
  let fixture: ComponentFixture<ServiceTimeWeeklyComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServiceTimeWeeklyComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServiceTimeWeeklyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
