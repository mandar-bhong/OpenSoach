import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChartConfigureTimeComponent } from './chart-configure-time.component';

describe('ChartConfigureTimeComponent', () => {
  let component: ChartConfigureTimeComponent;
  let fixture: ComponentFixture<ChartConfigureTimeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChartConfigureTimeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChartConfigureTimeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
