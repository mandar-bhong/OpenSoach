import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';

import { FloatingMenu } from '../models/ui/floating-menu';

@Injectable()
export class FloatingButtonMenuService {

    floatingMenuSubject = new Subject<FloatingMenu>();

    setFloatingMenu(menu: FloatingMenu) {
        this.floatingMenuSubject.next(menu);
    }
}
