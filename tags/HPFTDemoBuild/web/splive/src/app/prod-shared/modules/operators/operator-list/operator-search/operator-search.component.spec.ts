import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OperatorSearchComponent } from './operator-search.component';

describe('OperatorSearchComponent', () => {
  let component: OperatorSearchComponent;
  let fixture: ComponentFixture<OperatorSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OperatorSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OperatorSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
