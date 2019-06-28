import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewMedicineTransactionComponent } from './view-medicine-transaction.component';

describe('ViewMedicineTransactionComponent', () => {
  let component: ViewMedicineTransactionComponent;
  let fixture: ComponentFixture<ViewMedicineTransactionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewMedicineTransactionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewMedicineTransactionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
