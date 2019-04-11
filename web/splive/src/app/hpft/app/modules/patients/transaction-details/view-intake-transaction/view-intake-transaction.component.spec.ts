import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewIntakeTransactionComponent } from './view-intake-transaction.component';

describe('ViewIntakeTransactionComponent', () => {
  let component: ViewIntakeTransactionComponent;
  let fixture: ComponentFixture<ViewIntakeTransactionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewIntakeTransactionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewIntakeTransactionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
