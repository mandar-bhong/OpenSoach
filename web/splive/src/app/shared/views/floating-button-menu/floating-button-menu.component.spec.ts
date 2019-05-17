import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FloatingButtonMenuComponent } from './floating-button-menu.component';

describe('FloatingButtonMenuComponent', () => {
  let component: FloatingButtonMenuComponent;
  let fixture: ComponentFixture<FloatingButtonMenuComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FloatingButtonMenuComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FloatingButtonMenuComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
