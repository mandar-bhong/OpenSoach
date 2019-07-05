import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceDetailsViewComponent } from './device-details-view.component';

describe('SharedDeviceDetailsViewComponent', () => {
  let component: DeviceDetailsViewComponent;
  let fixture: ComponentFixture<DeviceDetailsViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeviceDetailsViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeviceDetailsViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
