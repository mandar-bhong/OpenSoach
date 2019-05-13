import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SpecificScheduleDetailsExpandViewComponent } from './specific-schedule-details-expand-view.component';

describe('SpecificScheduleDetailsExpandViewComponent', () => {
  let component: SpecificScheduleDetailsExpandViewComponent;
  let fixture: ComponentFixture<SpecificScheduleDetailsExpandViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SpecificScheduleDetailsExpandViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SpecificScheduleDetailsExpandViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
