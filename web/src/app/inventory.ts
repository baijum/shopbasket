export interface Inventory {
    id: number;
    name: string;
    description: string;
    status:boolean;
    price: string;
  }

  export class AddInventory {

    constructor(
      public id: number,
      public name: string,
      public description: string,
      public status: boolean,
      public price: string
    ) {  }
  
  }