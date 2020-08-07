import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { DateValueModel } from 'src/data/date-value.model'
import { EndOfDayItem } from 'src/data/eoditem'
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class EodService {

  constructor(private http: HttpClient) { }

  getClosePriceTimeSeries(id: string): Observable<DateValueModel[]> {
    return this.http.get<DateValueModel[]>(`${environment.host}/equity-fund/uicontroller/timeseries/close/${id}`);
  }

  getLatestEndOfDayItems(): Observable<EndOfDayItem[]> {
    return this.http.get<EndOfDayItem[]>(`${environment.host}/equity-fund/uicontroller/latest-eod`);
  }

  getLatestEndOfDayItem(): Observable<EndOfDayItem> {
    return this.http.get<EndOfDayItem>(`${environment.host}/equity-fund/uicontroller/latest-eod-item`);
  }
}
