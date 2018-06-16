import { EnumDataSourceItem } from '../models/ui/enum-datasource-item';

export class EnumNumberDatasource {
    static getDataSource(prefix: string, source: any): EnumDataSourceItem<number>[] {
        const dataSource: EnumDataSourceItem<number>[] = [];
        const options = Object.keys(source);
        const optionValues = options.slice(0, options.length / 2);
        for (let i = 0; i < optionValues.length; i++) {
            dataSource.push({ text: prefix + optionValues[i], value: Number(optionValues[i]) });
        }
        return dataSource;
    }
}
