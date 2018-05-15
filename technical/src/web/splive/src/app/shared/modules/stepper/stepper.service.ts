import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Subject } from 'rxjs/Subject';
@Injectable()
export class StepperService {
  stepperCountSubject: Subject<number> = new Subject<number>();

  constructor() { }
  stepperCount(stepcount: number) {
    this.stepperCountSubject.next(stepcount);
  }
}
