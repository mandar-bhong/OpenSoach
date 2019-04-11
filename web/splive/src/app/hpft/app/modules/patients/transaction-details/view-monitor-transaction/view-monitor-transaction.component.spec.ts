import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewMonitorTransactionComponent } from './view-monitor-transaction.component';

describe('ViewMonitorTransactionComponent', () => {
  let component: ViewMonitorTransactionComponent;
  let fixture: ComponentFixture<ViewMonitorTransactionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewMonitorTransactionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewMonitorTransactionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
