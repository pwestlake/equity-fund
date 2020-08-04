import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { DateValueModel } from 'src/data/date-value.model'
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class EodService {

  constructor(private http: HttpClient) { }

  getClosePriceTimeSeries(id: string): Observable<DateValueModel[]> {
    return this.http.get<DateValueModel[]>(`${environment.host}/equity-fund/uicontroller/timeseries/close/${id}`);
  }
}
