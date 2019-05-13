import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointUpdateComponent } from './servicepoint-update.component';

describe('ServicepointUpdateComponent', () => {
  let component: ServicepointUpdateComponent;
  let fixture: ComponentFixture<ServicepointUpdateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointUpdateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
