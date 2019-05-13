import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChartConfigureComponent } from './chart-configure.component';

describe('ChartConfigureComponent', () => {
  let component: ChartConfigureComponent;
  let fixture: ComponentFixture<ChartConfigureComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChartConfigureComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChartConfigureComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
