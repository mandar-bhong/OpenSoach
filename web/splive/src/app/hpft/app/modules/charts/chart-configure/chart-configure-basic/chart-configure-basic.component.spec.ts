import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChartConfigureBasicComponent } from './chart-configure-basic.component';

describe('ChartConfigureBasicComponent', () => {
  let component: ChartConfigureBasicComponent;
  let fixture: ComponentFixture<ChartConfigureBasicComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChartConfigureBasicComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChartConfigureBasicComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
