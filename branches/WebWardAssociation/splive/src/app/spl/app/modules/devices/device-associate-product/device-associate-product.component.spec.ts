import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceAssociateProductComponent } from './device-associate-product.component';

describe('DeviceAssociateProductComponent', () => {
  let component: DeviceAssociateProductComponent;
  let fixture: ComponentFixture<DeviceAssociateProductComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeviceAssociateProductComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeviceAssociateProductComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
