import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TaskTrendComponent } from './task-trend.component';

describe('TaskTrendComponent', () => {
  let component: TaskTrendComponent;
  let fixture: ComponentFixture<TaskTrendComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TaskTrendComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TaskTrendComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
