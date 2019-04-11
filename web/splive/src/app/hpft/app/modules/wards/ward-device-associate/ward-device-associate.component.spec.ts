import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WardDeviceAssociateComponent } from './ward-device-associate.component';

describe('WardDeviceAssociateComponent', () => {
  let component: WardDeviceAssociateComponent;
  let fixture: ComponentFixture<WardDeviceAssociateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WardDeviceAssociateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WardDeviceAssociateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
