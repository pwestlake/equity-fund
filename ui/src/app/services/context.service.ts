import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { StringValue } from 'src/data/stringvalue.model';


@Injectable({
  providedIn: 'root'
})
export class ContextService {

  constructor(private http: HttpClient) { }

  public getTitle() {
    return this.http.get<StringValue>(environment.host + '/equity-fund/uicontroller/title');
  }
}
