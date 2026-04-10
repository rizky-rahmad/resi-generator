import React, { FormEvent } from 'react';
import { ArrowRight } from 'lucide-react';
import { useStore } from '@/store/useStore';

export const Step1 = () => {
  const { form, setForm, setStep } = useStore();

  const handleNext = (e: FormEvent) => {
    e.preventDefault();
    setStep(2);
  };

  return (
    <form onSubmit={handleNext} className="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-4">
          <h3 className="text-lg font-semibold text-gray-800 border-b pb-2">Data Pengirim</h3>
          <div>
            <label className="block text-sm font-medium text-gray-700">Nama Pengirim</label>
            <input required type="text" value={form.sender_name} onChange={(e) => setForm({ sender_name: e.target.value })} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm border p-2 bg-white text-gray-900 focus:ring-blue-500 focus:border-blue-500" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Alamat Pengirim</label>
            <textarea required rows={3} value={form.sender_address} onChange={(e) => setForm({ sender_address: e.target.value })} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm border p-2 bg-white text-gray-900 focus:ring-blue-500 focus:border-blue-500" />
          </div>
        </div>
        <div className="space-y-4">
          <h3 className="text-lg font-semibold text-gray-800 border-b pb-2">Data Penerima</h3>
          <div>
            <label className="block text-sm font-medium text-gray-700">Nama Penerima</label>
            <input required type="text" value={form.receiver_name} onChange={(e) => setForm({ receiver_name: e.target.value })} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm border p-2 bg-white text-gray-900 focus:ring-blue-500 focus:border-blue-500" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Alamat Penerima</label>
            <textarea required rows={3} value={form.receiver_address} onChange={(e) => setForm({ receiver_address: e.target.value })} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm border p-2 bg-white text-gray-900 focus:ring-blue-500 focus:border-blue-500" />
          </div>
        </div>
      </div>
      <div className="flex justify-end pt-4 border-t">
        <button type="submit" className="flex items-center px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition">
          Selanjutnya <ArrowRight className="ml-2 h-4 w-4" />
        </button>
      </div>
    </form>
  );
};