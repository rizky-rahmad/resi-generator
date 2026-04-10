import React, { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { Search, Trash2, ArrowRight, ArrowLeft } from 'lucide-react';
import { useStore } from '@/store/useStore';
import { useDebounce } from '@/hooks/useDebounce';
import { apiFetch } from '@/lib/api';
import { Item } from '@/types';

export const Step2 = () => {
  const { form, addItem, removeItem, updateItemQuantity, setStep } = useStore();
  const [searchTerm, setSearchTerm] = useState('');
  const [isFocused, setIsFocused] = useState(false);
  const debouncedTerm = useDebounce(searchTerm, 500);

  const { data: searchResults, isFetching, isError } = useQuery({
    queryKey: ['items', debouncedTerm],
    queryFn: async ({ signal }) => {
      const data = await apiFetch(`/items`, { signal });
      const items = data.data || data;
      const allItems = (Array.isArray(items) ? items : []) as Item[];

      if (!debouncedTerm) return allItems;
      const lowerTerm = debouncedTerm.toLowerCase();
      return allItems.filter(item => 
        item.code.toLowerCase().includes(lowerTerm) || 
        item.name.toLowerCase().includes(lowerTerm)
      );
    },
  });

  return (
    <div className="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
      <div className="bg-gray-50 p-4 rounded-lg border border-gray-200">
        <label className="block text-sm font-medium text-gray-700 mb-2">Cari & Tambah Barang</label>
        <div className="relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Search className="h-5 w-5 text-gray-400" />
          </div>
          <input
            type="text"
            className="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
            placeholder="Klik di sini atau ketik kode/nama barang..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            onFocus={() => setIsFocused(true)}
            onBlur={() => setIsFocused(false)} 
          />
        </div>

        {isFocused && (
          <ul className="mt-1 max-h-60 overflow-auto rounded-md bg-white border border-gray-200 shadow-lg absolute w-full z-10 sm:max-w-xl">
            {isFetching ? (
              <li className="py-3 px-4 text-sm text-gray-500 flex items-center justify-center">Mencari barang...</li>
            ) : isError ? (
              <li className="py-3 px-4 text-sm text-red-500 text-center bg-red-50">Gagal memuat data server.</li>
            ) : searchResults && searchResults.length > 0 ? (
              searchResults.map((item) => (
                <li key={item.id} className="cursor-pointer select-none relative py-2 pl-3 pr-9 hover:bg-blue-50 border-b last:border-b-0"
                  onMouseDown={(e) => {
                    e.preventDefault();
                    addItem(item);
                    setSearchTerm('');
                    setIsFocused(false);
                  }}
                >
                  <div className="flex justify-between items-center">
                    <span className="font-medium block truncate text-gray-900">{item.code} - {item.name}</span>
                    <span className="text-sm text-gray-500">Rp {item.price.toLocaleString('id-ID')}</span>
                  </div>
                </li>
              ))
            ) : (
              <li className="py-4 px-4 text-sm text-gray-500 text-center">Barang tidak ditemukan.</li>
            )}
          </ul>
        )}
      </div>

      <div className="overflow-x-auto shadow ring-1 ring-black ring-opacity-5 rounded-lg">
        <table className="min-w-full divide-y divide-gray-300">
          <thead className="bg-gray-50">
            <tr>
              <th className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900">Kode</th>
              <th className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Nama Barang</th>
              <th className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Harga</th>
              <th className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900 w-24">Qty</th>
              <th className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Subtotal</th>
              <th className="relative py-3.5 pl-3 pr-4 sm:pr-6 w-10"></th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200 bg-white">
            {form.items.length === 0 ? (
              <tr><td colSpan={6} className="py-8 text-center text-gray-500 text-sm">Keranjang kosong.</td></tr>
            ) : (
              form.items.map((item) => (
                <tr key={item.id}>
                  <td className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900">{item.code}</td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{item.name}</td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">Rp {item.price.toLocaleString('id-ID')}</td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                    <input type="number" min="1" className="w-16 p-1 border rounded bg-white text-gray-900" value={item.quantity} onChange={(e) => updateItemQuantity(item.id, parseInt(e.target.value) || 1)} />
                  </td>
                  <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-900 font-medium">Rp {item.subtotal.toLocaleString('id-ID')}</td>
                  <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm">
                    <button onClick={() => removeItem(item.id)} className="text-red-600 hover:text-red-900"><Trash2 className="h-4 w-4" /></button>
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>

      <div className="flex justify-between pt-4 border-t">
        <button onClick={() => setStep(1)} type="button" className="flex items-center px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 transition"><ArrowLeft className="mr-2 h-4 w-4" /> Kembali</button>
        <button onClick={() => setStep(3)} disabled={form.items.length === 0} className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50">Review Invoice <ArrowRight className="ml-2 h-4 w-4" /></button>
      </div>
    </div>
  );
};