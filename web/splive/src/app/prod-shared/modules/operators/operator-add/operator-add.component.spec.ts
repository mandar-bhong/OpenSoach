import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OperatorAddComponent } from './operator-add.component';

describe('OperatorAddComponent', () => {
  let component: OperatorAddComponent;
  let fixture: ComponentFixture<OperatorAddComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OperatorAddComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OperatorAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
