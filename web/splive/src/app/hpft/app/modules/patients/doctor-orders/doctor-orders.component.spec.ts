import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DoctorOrdersComponent } from './doctor-orders.component';

describe('DoctorOrdersComponent', () => {
  let component: DoctorOrdersComponent;
  let fixture: ComponentFixture<DoctorOrdersComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DoctorOrdersComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DoctorOrdersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
