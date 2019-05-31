import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewMedicineScheduleComponent } from './view-medicine-schedule.component';

describe('ViewMedicineScheduleComponent', () => {
  let component: ViewMedicineScheduleComponent;
  let fixture: ComponentFixture<ViewMedicineScheduleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewMedicineScheduleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewMedicineScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
