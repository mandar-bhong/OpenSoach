import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MedicalContainerComponent } from './medical-container.component';

describe('MedicalContainerComponent', () => {
  let component: MedicalContainerComponent;
  let fixture: ComponentFixture<MedicalContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MedicalContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MedicalContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
