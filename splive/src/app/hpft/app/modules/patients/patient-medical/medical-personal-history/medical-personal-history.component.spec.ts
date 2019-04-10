import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MedicalPersonalHistoryComponent } from './medical-personal-history.component';

describe('MedicalPersonalHistoryComponent', () => {
  let component: MedicalPersonalHistoryComponent;
  let fixture: ComponentFixture<MedicalPersonalHistoryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MedicalPersonalHistoryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MedicalPersonalHistoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
