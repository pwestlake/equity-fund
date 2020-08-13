import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { NewsItemModel } from 'src/data/newsitem.model';

@Injectable({
  providedIn: 'root'
})
export class NewsService {

  constructor(private http: HttpClient) { }

  getNewsItems(count: number, key: string, sortkey: Date, catalogref: string): Observable<NewsItemModel[]> {
    return this.http.get<NewsItemModel[]>(`${environment.host}/equity-fund/uicontroller/newsitems`, {
      params: new HttpParams()
        .set('catalogref', catalogref)
        .set('key', key)
        .set('sortkey', sortkey == null ? "" : sortkey.toString())
        .set('count', count.toString())
    });
  }
}
