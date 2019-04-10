import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerAssociateProductComponent } from './customer-associate-product.component';

describe('CustomerAssociateProductComponent', () => {
  let component: CustomerAssociateProductComponent;
  let fixture: ComponentFixture<CustomerAssociateProductComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CustomerAssociateProductComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CustomerAssociateProductComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
