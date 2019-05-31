import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CorporateSearchComponent } from './corporate-search.component';

describe('CorporateSearchComponent', () => {
  let component: CorporateSearchComponent;
  let fixture: ComponentFixture<CorporateSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CorporateSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CorporateSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
