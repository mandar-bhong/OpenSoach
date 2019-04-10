import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ComplaintTrendComponent } from './complaint-trend.component';

describe('ComplaintTrendComponent', () => {
  let component: ComplaintTrendComponent;
  let fixture: ComponentFixture<ComplaintTrendComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ComplaintTrendComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ComplaintTrendComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
