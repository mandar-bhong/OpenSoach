import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerServicepointAssociateComponent } from './customer-servicepoint-associate.component';

describe('CustomerServicepointAssociateComponent', () => {
  let component: CustomerServicepointAssociateComponent;
  let fixture: ComponentFixture<CustomerServicepointAssociateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CustomerServicepointAssociateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CustomerServicepointAssociateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
