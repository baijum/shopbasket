import { Component, OnInit } from '@angular/core';
import { InventoryService } from '../inventory.service';
import { Inventory } from '../inventory';

@Component({
  selector: 'app-inventory',
  templateUrl: './inventory.component.html',
  styleUrls: ['./inventory.component.scss']
})

export class InventoryComponent implements OnInit {
  inventory: Inventory[] = [];
  constructor(private inventoryService: InventoryService) { }

  ngOnInit(): void {
    this.getInventory();
  }
  getInventory(): void {
    this.inventoryService.getInventory()
        .subscribe(inventoryList => this.inventory = inventoryList);
  }
}
