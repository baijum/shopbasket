export interface Inventory {
    id: number;
    name: string;
    description: string;
    status:boolean;
    price:number;
  }

  export class AddInventory {

    constructor(
      public id: number,
      public name: string,
      public description: string,
      public status: boolean,
      public price:number
    ) {  }
  
  }