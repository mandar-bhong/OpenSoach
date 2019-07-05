import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewOutputTransactionComponent } from './view-output-transaction.component';

describe('ViewOutputTransactionComponent', () => {
  let component: ViewOutputTransactionComponent;
  let fixture: ComponentFixture<ViewOutputTransactionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewOutputTransactionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewOutputTransactionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
