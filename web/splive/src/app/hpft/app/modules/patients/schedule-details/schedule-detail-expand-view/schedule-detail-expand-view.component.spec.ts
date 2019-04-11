import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScheduleDetailExpandViewComponent } from './schedule-detail-expand-view.component';

describe('ScheduleDetailExpandViewComponent', () => {
  let component: ScheduleDetailExpandViewComponent;
  let fixture: ComponentFixture<ScheduleDetailExpandViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScheduleDetailExpandViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScheduleDetailExpandViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
