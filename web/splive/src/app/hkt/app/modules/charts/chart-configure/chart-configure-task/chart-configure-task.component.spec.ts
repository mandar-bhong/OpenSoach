import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChartConfigureTaskComponent } from './chart-configure-task.component';

describe('ChartConfigureTaskComponent', () => {
  let component: ChartConfigureTaskComponent;
  let fixture: ComponentFixture<ChartConfigureTaskComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChartConfigureTaskComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChartConfigureTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
