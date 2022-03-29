import { Component, OnInit } from '@angular/core';
import { AddInventory } from '../inventory';
@Component({
  selector: 'app-inventory-form',
  templateUrl: './inventory-form.component.html',
  styleUrls: ['./inventory-form.component.scss']
})
export class InventoryFormComponent implements OnInit {
  model= new AddInventory(0,"","",false,0)
  constructor() { }
  ngOnInit(): void {
  }

}
