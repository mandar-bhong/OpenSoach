import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CorporateListViewComponent } from './corporate-list-view.component';

describe('CorporateListViewComponent', () => {
  let component: CorporateListViewComponent;
  let fixture: ComponentFixture<CorporateListViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CorporateListViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CorporateListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
