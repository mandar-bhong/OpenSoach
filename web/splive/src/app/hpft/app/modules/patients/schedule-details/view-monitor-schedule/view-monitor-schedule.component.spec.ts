import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewMonitorScheduleComponent } from './view-monitor-schedule.component';

describe('ViewMonitorScheduleComponent', () => {
  let component: ViewMonitorScheduleComponent;
  let fixture: ComponentFixture<ViewMonitorScheduleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewMonitorScheduleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewMonitorScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
