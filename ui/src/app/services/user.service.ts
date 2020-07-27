import { Injectable } from '@angular/core';
import { HttpClient, HttpUrlEncodingCodec, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';
import { UserID } from 'src/data/userid.model';
import { StringValue } from 'src/data/stringvalue.model';
import { UserPreference } from 'src/data/user-preference.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  public getAuthenticatedUserID(): Observable<UserID> {
    return this.http.get<UserID>(environment.host + '/userservice/authenticated-user');
  }

  public getPreference(userid: string, preference: string): Observable<StringValue> {
    let codec = new HttpUrlEncodingCodec()
    let encodedUserID = codec.encodeValue(userid);
    return this.http.get<StringValue>(`${environment.host}/userservice/preference/${encodedUserID}/${preference}`)
  }

  public setPreference(userid: string, key: string, value: string) {
    const preference: UserPreference =  {
      userid: userid,
      key: key,
      value: value,
      datecreated: new Date()
    };
    const jsonBody: string = JSON.stringify(preference);

    const httpOptions = {
    headers: new HttpHeaders({
        'Content-Type':  'application/json'
      })
    };
    const obs: Observable<any> = this.http.post(`${environment.host}/userservice/preference/${userid}/${key}`, 
      jsonBody, httpOptions);
    obs.subscribe(
      x => console.log('Next:' + x),
      err => console.log('Error:' + err)
    );
  }
}
