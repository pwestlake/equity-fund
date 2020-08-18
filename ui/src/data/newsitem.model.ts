export interface NewsItemModel {
    id: string;
    datetime: Date;
    catalogref: string;
    companycode: string;
    companyname: string;
    content: string;
    sentiment: number;
    title: string;
}

export interface NewsItemViewModel extends NewsItemModel {
    selected: Boolean;
}