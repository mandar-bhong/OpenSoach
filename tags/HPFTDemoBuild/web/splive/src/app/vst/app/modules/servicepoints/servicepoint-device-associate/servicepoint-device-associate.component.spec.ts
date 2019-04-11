import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointDeviceAssociateComponent } from './servicepoint-device-associate.component';

describe('ServicepointDeviceAssociateComponent', () => {
  let component: ServicepointDeviceAssociateComponent;
  let fixture: ComponentFixture<ServicepointDeviceAssociateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointDeviceAssociateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointDeviceAssociateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
