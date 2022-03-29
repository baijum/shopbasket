import { Component, OnInit } from '@angular/core';
import { AddInventory } from '../inventory';
import { InventoryService } from '../inventory.service';
@Component({
  selector: 'app-inventory-form',
  templateUrl: './inventory-form.component.html',
  styleUrls: ['./inventory-form.component.scss']
})
export class InventoryFormComponent implements OnInit {
  model= new AddInventory(0,"","",false,0)
  constructor(private inventoryService: InventoryService) { }
  ngOnInit(): void {
  }
  newInventory(){
    this.inventoryService.createInventory(this.model.name,this.model.description, this.model.price)
  }
  onSubmit(){

  }
}
