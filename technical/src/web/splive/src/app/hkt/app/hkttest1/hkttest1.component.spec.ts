import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { Hkttest1Component } from './hkttest1.component';

describe('Hkttest1Component', () => {
  let component: Hkttest1Component;
  let fixture: ComponentFixture<Hkttest1Component>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ Hkttest1Component ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(Hkttest1Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
