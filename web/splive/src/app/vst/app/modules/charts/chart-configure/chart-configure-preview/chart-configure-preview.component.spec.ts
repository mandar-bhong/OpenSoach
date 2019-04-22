import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChartConfigurePreviewComponent } from './chart-configure-preview.component';

describe('ChartConfigurePreviewComponent', () => {
  let component: ChartConfigurePreviewComponent;
  let fixture: ComponentFixture<ChartConfigurePreviewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChartConfigurePreviewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChartConfigurePreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
