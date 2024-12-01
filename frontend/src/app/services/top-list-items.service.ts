import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { API_URL } from '../constant/constant';

@Injectable({
  providedIn: 'root'
})
export class TopListItemsService {

  private apiUrl = API_URL + '/topListItems'; // Backend API URL

  constructor(private http: HttpClient) { }

  getTopListItems(pageId: number, pageSize: number): Observable<any> {
    const url = `${this.apiUrl}?page_id=${pageId}&page_size=${pageSize}`;
    return this.http.get<any>(url);
  }
}
