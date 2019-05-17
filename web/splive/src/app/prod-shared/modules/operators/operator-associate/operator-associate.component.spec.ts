import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OperatorAssociateComponent } from './operator-associate.component';

describe('OperatorAssociateComponent', () => {
  let component: OperatorAssociateComponent;
  let fixture: ComponentFixture<OperatorAssociateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OperatorAssociateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OperatorAssociateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
