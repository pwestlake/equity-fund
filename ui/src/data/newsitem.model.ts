export interface NewsItemModel {
    id: string;
    datetime: Date;
    catelogref: string;
    companycode: string;
    companyname: string;
    content: string;
    sentiment: number;
    title: string;
}