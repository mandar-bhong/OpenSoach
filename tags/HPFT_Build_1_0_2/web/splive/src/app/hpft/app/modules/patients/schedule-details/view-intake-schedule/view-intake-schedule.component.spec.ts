import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewIntakeScheduleComponent } from './view-intake-schedule.component';

describe('ViewIntakeScheduleComponent', () => {
  let component: ViewIntakeScheduleComponent;
  let fixture: ComponentFixture<ViewIntakeScheduleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewIntakeScheduleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewIntakeScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
