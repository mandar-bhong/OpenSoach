import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable()
export class SidebarToggleService {
  menuToggleSubject: Subject<boolean> = new Subject<boolean>();
  constructor() { }

  toggleMenu(state: boolean) {
    this.menuToggleSubject.next(state);
  }
}
