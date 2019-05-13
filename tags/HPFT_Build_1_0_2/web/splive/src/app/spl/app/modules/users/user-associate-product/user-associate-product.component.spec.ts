import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UserAssociateProductComponent } from './user-associate-product.component';

describe('UserAssociateProductComponent', () => {
  let component: UserAssociateProductComponent;
  let fixture: ComponentFixture<UserAssociateProductComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UserAssociateProductComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UserAssociateProductComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
