import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointDetailsComponent } from './servicepoint-details.component';

describe('ServicepointDetailsComponent', () => {
  let component: ServicepointDetailsComponent;
  let fixture: ComponentFixture<ServicepointDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
