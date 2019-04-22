import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UserMasterDetailsComponent } from './user-master-details.component';

describe('UserMasterDetailsComponent', () => {
  let component: UserMasterDetailsComponent;
  let fixture: ComponentFixture<UserMasterDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UserMasterDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UserMasterDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
