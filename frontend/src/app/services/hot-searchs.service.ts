import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_URL } from '../constant/constant';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class HotSearchsService {


  private apiUrl = API_URL+'/hotSearches'; // Backend API URL
  constructor(private http: HttpClient) {}
  getHotSearchs(pageId: number, pageSize: number): Observable<any> {
    const url = `${this.apiUrl}?page_id=${pageId}&page_size=${pageSize}`;
    return this.http.get<any>(url);
  }
}

