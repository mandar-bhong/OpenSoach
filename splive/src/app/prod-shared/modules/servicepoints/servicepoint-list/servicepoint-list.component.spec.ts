import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ServicepointListComponent } from './servicepoint-list.component';

describe('ServicepointListComponent', () => {
  let component: ServicepointListComponent;
  let fixture: ComponentFixture<ServicepointListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ServicepointListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ServicepointListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
