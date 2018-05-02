export class FloatingMenuItem {
    key: string;
    title: string;
    icon: string;
    action: any;
    data: any;
    navigate: boolean;
    url: string;
}

export class FloatingMenu {
    items: FloatingMenuItem[];
    menuInstanceKey: string;
}

export class FloatingMenuAction {
    key: string;
    data: any;
}
