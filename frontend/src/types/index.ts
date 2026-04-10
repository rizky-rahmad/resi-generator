export interface User {
  token: string;
  role: 'admin' | 'kerani';
  name: string;
}

export interface Item {
  id: number;
  code: string;
  name: string;
  price: number;
}

export interface SelectedItem extends Item {
  quantity: number;
  subtotal: number;
}

export interface InvoiceForm {
  sender_name: string;
  sender_address: string;
  receiver_name: string;
  receiver_address: string;
  items: SelectedItem[];
}