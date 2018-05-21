import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';

/**
 * DynamicContextService, holds the transferring context to the components rendered at runtime.
 */
@Injectable()
export class DynamicContextService {
  triggerSaveDynamicComponent: Subject<any> = new Subject<any>();
  triggerChangeDynamicComponent: Subject<string> = new Subject<string>();
  triggerToggleSaveButton: Subject<boolean> = new Subject<boolean>();
  triggerAction: Subject<any> = new Subject<any>();

  private dynamicComponentMappings = {};

  constructor() {
  }

  /**
   * Save event trigger
   */
  save(): void {
    this.triggerSaveDynamicComponent.next(null);
  }

  /**
   * Executes the triggerChangeDynamicComponent, to navigate the changes to dynamic-component
   */
  changeComponent(type: string): void {
    this.triggerChangeDynamicComponent.next(type);
  }

  toggleSaveButton(state: boolean) {
    this.triggerToggleSaveButton.next(state);
  }

  onAction(result: any): void {
    this.triggerAction.next(result);
  }

  addDynamicComponentMaping(typeName: string, component: any) {
    this.dynamicComponentMappings[typeName] = component;
  }

  getDynamicComponentMaping(typeName: string) {
    return this.dynamicComponentMappings[typeName];
  }
}
