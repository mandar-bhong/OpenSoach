import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointServiceAssociateComponent } from './servicepoint-service-associate.component';

describe('ServicepointServiceAssociateComponent', () => {
  let component: ServicepointServiceAssociateComponent;
  let fixture: ComponentFixture<ServicepointServiceAssociateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointServiceAssociateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointServiceAssociateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
