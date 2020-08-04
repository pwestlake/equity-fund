import { Injectable } from '@angular/core';
import { EquityCatalogItem } from 'src/data/equitycatalogitem'
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class EquitycatalogService {

  constructor(private http: HttpClient) { }

  addEquityCatalogItem(item: EquityCatalogItem) {
    const jsonBody: string = JSON.stringify(item);

    const httpOptions = {
    headers: new HttpHeaders({
        'Content-Type':  'application/json'
      })
    };
    const obs: Observable<any> = this.http.post(`${environment.host}/equity-fund/uicontroller/equitycatalogitem`, 
      jsonBody, httpOptions);
    obs.subscribe(
      x => console.log('Next:' + x),
      err => console.log('Error:' + err)
    );
  }

  getAllEquityCatalogItems(): Observable<EquityCatalogItem[]> {
    return this.http.get<EquityCatalogItem[]>(`${environment.host}/equity-fund/uicontroller/equitycatalogitem`);
  }
}
