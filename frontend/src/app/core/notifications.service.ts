import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { AppHttpParams, contentHeaders } from '../helpers/http-config';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { AppNotification } from '../model/notification';
import { APIEndpoints } from '../helpers/app-settings'

@Injectable()
export class NotificationsService {

  constructor(private http: HttpClient) { }

  public getNotifications(): Observable<AppNotification[]> {
    var options = { headers: contentHeaders };
    return this.http.get<AppNotification[]>(`${APIEndpoints.NOTIFICATION_API}`, options);
  }

  public updateNotification(id: string): Observable<AppNotification[]> {
    var options = { headers: contentHeaders };
    return this.http.put<AppNotification[]>(`${APIEndpoints.NOTIFICATION_API + '/' + id}`, options);
  }
}
