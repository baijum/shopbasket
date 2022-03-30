import { Injectable } from '@angular/core';
import { Inventory } from './inventory';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class InventoryService {
  private inv:Observable<Inventory>= new Observable<Inventory>();
  private inventoryUrl = 'api/inventory'; 
  constructor(private http: HttpClient) { }
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  getInventory(): Observable<Inventory[]> {
    return this.http.get<Inventory[]>(this.inventoryUrl);
  }

  createInventory(name: String, description: String, price: String): Observable<Inventory> {
    var inventory= {name:name, description:description, price:price, status:true}
    this.http.post<Inventory>(this.inventoryUrl,inventory, this.httpOptions).subscribe({
      next: data => {
        console.log("Success")
      },
      error: error => {
            console.error('There was an error!', error);
        }
    });
    return this.inv
  }

}
