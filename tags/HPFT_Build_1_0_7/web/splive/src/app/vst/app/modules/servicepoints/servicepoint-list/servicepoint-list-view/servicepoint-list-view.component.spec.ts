import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointListViewComponent } from './servicepoint-list-view.component';

describe('ServicepointListViewComponent', () => {
  let component: ServicepointListViewComponent;
  let fixture: ComponentFixture<ServicepointListViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointListViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
