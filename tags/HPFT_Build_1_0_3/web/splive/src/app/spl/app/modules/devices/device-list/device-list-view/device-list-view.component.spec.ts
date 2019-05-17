import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeviceListViewComponent } from './device-list-view.component';

describe('DeviceListViewComponent', () => {
  let component: DeviceListViewComponent;
  let fixture: ComponentFixture<DeviceListViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeviceListViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeviceListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
