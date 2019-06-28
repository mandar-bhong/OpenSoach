import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FeedbackSummaryComponent } from './feedback-summary.component';

describe('FeedbackSummaryComponent', () => {
  let component: FeedbackSummaryComponent;
  let fixture: ComponentFixture<FeedbackSummaryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FeedbackSummaryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FeedbackSummaryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
