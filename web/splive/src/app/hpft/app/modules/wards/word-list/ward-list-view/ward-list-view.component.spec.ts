import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordListViewComponent } from './word-list-view.component';

describe('WordListViewComponent', () => {
  let component: WordListViewComponent;
  let fixture: ComponentFixture<WordListViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordListViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
