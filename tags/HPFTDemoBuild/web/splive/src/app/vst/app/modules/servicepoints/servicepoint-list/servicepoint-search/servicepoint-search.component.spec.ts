import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointSearchComponent } from './servicepoint-search.component';

describe('ServicepointSearchComponent', () => {
  let component: ServicepointSearchComponent;
  let fixture: ComponentFixture<ServicepointSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
