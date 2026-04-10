import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { User, InvoiceForm, Item } from '../types';

interface StoreState {
  user: User | null;
  form: InvoiceForm;
  currentStep: number;
  invoiceData: any | null;
  setUser: (user: User | null) => void;
  setForm: (data: Partial<InvoiceForm>) => void;
  addItem: (item: Item) => void;
  updateItemQuantity: (id: number, quantity: number) => void;
  removeItem: (id: number) => void;
  setStep: (step: number) => void;
  setInvoiceData: (data: any | null) => void;
  resetForm: () => void;
}

export const useStore = create<StoreState>()(
  persist(
    (set) => ({
      user: null,
      currentStep: 1,
      invoiceData: null,
      form: {
        sender_name: '',
        sender_address: '',
        receiver_name: '',
        receiver_address: '',
        items: [],
      },
      setUser: (user) => set({ user }),
      setForm: (data) => set((state) => ({ form: { ...state.form, ...data } })),
      addItem: (item) =>
        set((state) => {
          const exists = state.form.items.find((i) => i.id === item.id);
          if (exists) return state; // Mencegah duplikasi
          return {
            form: {
              ...state.form,
              items: [...state.form.items, { ...item, quantity: 1, subtotal: item.price }],
            },
          };
        }),
      updateItemQuantity: (id, quantity) =>
        set((state) => ({
          form: {
            ...state.form,
            items: state.form.items.map((i) =>
              i.id === id ? { ...i, quantity, subtotal: quantity * i.price } : i
            ),
          },
        })),
      removeItem: (id) =>
        set((state) => ({
          form: {
            ...state.form,
            items: state.form.items.filter((i) => i.id !== id),
          },
        })),
      setStep: (step) => set({ currentStep: step }),
      setInvoiceData: (data) => set({ invoiceData: data }),
      resetForm: () =>
        set({
          currentStep: 1,
          invoiceData: null,
          form: {
            sender_name: '',
            sender_address: '',
            receiver_name: '',
            receiver_address: '',
            items: [],
          },
        }),
    }),
    { name: 'fleetify-resi-storage' }
  )
);