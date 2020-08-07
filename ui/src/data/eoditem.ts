export interface EndOfDayItem {
    id: string;
    date: Date;
    adj_close: number;
    adj_high: number;
    adj_low: number;
    adj_open: number;
    adj_volume: number;
    exchange: string;
    close: number;
    high: number;
    low: number;
    open: number;
    symbol: string;
    volume: number
    close_chg: number;
}