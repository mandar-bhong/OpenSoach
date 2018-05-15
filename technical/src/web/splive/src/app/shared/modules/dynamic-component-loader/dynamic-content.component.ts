import {
    Component,
    ComponentFactoryResolver,
    ComponentRef,
    Input,
    OnDestroy,
    OnInit,
    ViewChild,
    ViewContainerRef,
} from '@angular/core';
import { Subscription } from 'rxjs/Subscription';

import { DynamicContextService } from './dynamic-context.service';

@Component({
    selector: 'app-dynamic-content',
    template: `<div #dynamicContainer></div>`
})
export class DynamicContentComponent implements OnInit, OnDestroy {
    private componentRef: ComponentRef<{}>;

    // record the subscription with dynamic context service to change the component at runtime.
    subscriptionChangeDynamicComponent: Subscription;

    /**
     * specifies the type of component to render.
     * the type specified should match with one of the mappings declared
     */
    private _type: string;

    @Input()
    componentdata: any;

    @Input()
    mode: number;

    @Input()
    set type(type: string) {
        console.log('prev value: ', this._type);
        this.destroyComponent();
        console.log('new value name: ', type);
        this._type = type;
        this.initComponent();
    }

    constructor(private componentFactoryResolver: ComponentFactoryResolver,
        private dynamicContextService: DynamicContextService) {
    }

    /**
     * container is where the component is rendered.
     */
    @ViewChild('dynamicContainer', {
        read: ViewContainerRef
    })
    container: ViewContainerRef;

    /**
     * define a string key for each component which needs to be rendered
     * at runtime based on some condition.
     * for simplicity use, the selector of each component as the key.
     */
    private mappings = {
    };
    getComponentType(typeName: string) {
        let type = this.mappings[typeName];
        if (type) {
            return type;
        } else {
            type = this.dynamicContextService.getDynamicComponentMaping(typeName);
            return type;
        }
    }

    /**
     * Initialiazes the component based on the input parameter 'type' and
     * subcsribes with DynamicContextService for change in component.
     */
    ngOnInit() {
        this.subscriptionChangeDynamicComponent = this.dynamicContextService.triggerChangeDynamicComponent.subscribe((componentType) => {
            this.type = componentType;
        });
    }

    ngOnDestroy() {
        if (this.subscriptionChangeDynamicComponent) { this.subscriptionChangeDynamicComponent.unsubscribe(); }
        this.destroyComponent();
    }

    initComponent() {
        if (this._type) {
            const componentType = this.getComponentType(this._type);
            const factory = this.componentFactoryResolver.resolveComponentFactory(componentType);
            this.componentRef = this.container.createComponent(factory);
            this.componentRef.instance['mode'] = Number(this.mode);
            this.componentRef.instance['componentdata'] = this.componentdata;
        }
    }

    destroyComponent() {
        if (this.componentRef) {
            this.componentRef.destroy();
            this.componentRef = null;
        }
    }
}
