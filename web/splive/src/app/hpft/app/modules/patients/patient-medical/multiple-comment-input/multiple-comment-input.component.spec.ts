import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MultipleCommentInputComponent } from './multiple-comment-input.component';


describe('MedicalDetailsComponent', () => {
  let component: MultipleCommentInputComponent;
  let fixture: ComponentFixture<MultipleCommentInputComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MultipleCommentInputComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MultipleCommentInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
