import React from 'react';
import { useMutation } from '@tanstack/react-query';
import { ArrowLeft, CheckCircle, Printer, Plus, Package } from 'lucide-react';
import { useStore } from '@/store/useStore';
import { apiFetch } from '@/lib/api';

export const Step3 = () => {
  const { form, user, setStep, resetForm, invoiceData, setInvoiceData } = useStore();
  const totalAmount = form.items.reduce((acc, curr) => acc + curr.subtotal, 0);

  const submitMutation = useMutation({
    mutationFn: async () => {
      if (!user) throw new Error("Unauthorized");
      let formattedItems;
      if (user.role === 'kerani') {
        formattedItems = form.items.map(item => ({ item_id: item.id, quantity: item.quantity }));
      } else {
        formattedItems = form.items.map(item => ({ item_id: item.id, quantity: item.quantity, price: item.price, subtotal: item.subtotal }));
      }

      const payload = {
        sender_name: form.sender_name, sender_address: form.sender_address,
        receiver_name: form.receiver_name, receiver_address: form.receiver_address,
        items: formattedItems
      };

      const data = await apiFetch('/invoices', {
        method: 'POST',
        headers: { Authorization: `Bearer ${user.token}` },
        body: JSON.stringify(payload),
      });
      return data.data;
    },
    onSuccess: (data) => setInvoiceData(data),
  });

  if (invoiceData) {
    return (
      <div className="text-center py-10 animate-in zoom-in duration-300">
        <CheckCircle className="h-16 w-16 text-green-500 mx-auto mb-4" />
        <h2 className="text-2xl font-bold text-gray-900 mb-2">Invoice Berhasil Dibuat!</h2>
        <p className="text-gray-500 mb-6">Nomor Resi: <span className="font-semibold text-gray-900">{invoiceData.invoice_number}</span></p>
        <div className="flex justify-center space-x-4">
          <button onClick={() => window.print()} className="flex items-center px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition"><Printer className="mr-2 h-4 w-4" /> Cetak Invoice</button>
          <button onClick={resetForm} className="flex items-center px-4 py-2 bg-blue-100 text-blue-700 rounded-md hover:bg-blue-200 transition"><Plus className="mr-2 h-4 w-4" /> Buat Baru</button>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6 animate-in fade-in slide-in-from-bottom-4 duration-500">
      <div className="bg-white rounded-lg border border-gray-200 overflow-hidden">
        <div className="px-4 py-5 sm:px-6 bg-gray-50 border-b">
          <h3 className="text-lg leading-6 font-medium text-gray-900">Review Data</h3>
        </div>
        <div className="border-t border-gray-200 px-4 py-5 sm:p-0">
          <dl className="sm:divide-y sm:divide-gray-200">
            <div className="py-4 sm:py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500">Pengirim</dt>
              <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2"><span className="font-semibold block">{form.sender_name}</span>{form.sender_address}</dd>
            </div>
            <div className="py-4 sm:py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500">Penerima</dt>
              <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2"><span className="font-semibold block">{form.receiver_name}</span>{form.receiver_address}</dd>
            </div>
            <div className="py-4 sm:py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500">Daftar Barang</dt>
              <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                <ul className="border border-gray-200 rounded-md divide-y divide-gray-200">
                  {form.items.map((item) => (
                    <li key={item.id} className="pl-3 pr-4 py-3 flex items-center justify-between text-sm">
                      <div className="w-0 flex-1 flex items-center"><Package className="flex-shrink-0 h-5 w-5 text-gray-400" /><span className="ml-2 flex-1 w-0 truncate">{item.quantity}x {item.name}</span></div>
                      <div className="ml-4 flex-shrink-0 font-medium">Rp {item.subtotal.toLocaleString('id-ID')}</div>
                    </li>
                  ))}
                  <li className="pl-3 pr-4 py-3 flex justify-between text-sm bg-gray-50 font-bold"><span>Total Estimasi</span><span>Rp {totalAmount.toLocaleString('id-ID')}</span></li>
                </ul>
              </dd>
            </div>
          </dl>
        </div>
      </div>
      {submitMutation.isError && <div className="p-3 bg-red-100 text-red-700 rounded-md text-sm">Gagal menyimpan invoice: {submitMutation.error.message}</div>}
      <div className="flex justify-between pt-4 border-t">
        <button onClick={() => setStep(2)} type="button" className="flex items-center px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50"><ArrowLeft className="mr-2 h-4 w-4" /> Kembali</button>
        <button onClick={() => submitMutation.mutate()} disabled={submitMutation.isPending} className="flex items-center px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50">
          {submitMutation.isPending ? 'Memproses...' : 'Submit Transaksi'}
        </button>
      </div>
    </div>
  );
};