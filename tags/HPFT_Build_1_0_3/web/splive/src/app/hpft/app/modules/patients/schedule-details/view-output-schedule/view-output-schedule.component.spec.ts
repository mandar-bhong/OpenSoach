import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewOutputScheduleComponent } from './view-output-schedule.component';

describe('ViewOutputScheduleComponent', () => {
  let component: ViewOutputScheduleComponent;
  let fixture: ComponentFixture<ViewOutputScheduleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewOutputScheduleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewOutputScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
