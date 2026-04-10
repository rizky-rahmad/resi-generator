import React from 'react';

export const PrintLayout = ({ invoice }: { invoice: any }) => {
  if (!invoice) return null;

  return (
    <div className="hidden print:block w-full bg-white text-black p-8 text-left">
      <div className="border-b-4 border-blue-800 pb-4 mb-6 flex justify-between items-end">
        <div>
          <h1 className="text-4xl font-black text-blue-800 tracking-tighter">FLEETIFY<span className="text-orange-500">.LOGISTICS</span></h1>
          <p className="text-sm mt-1 text-gray-600">Jl. Teknologi No. 99, Jakarta Selatan</p>
        </div>
        <div className="text-right">
          <h2 className="text-2xl font-bold text-gray-800">INVOICE / RESI</h2>
          <p className="text-lg font-mono text-gray-600">{invoice.invoice_number}</p>
        </div>
      </div>
      <div className="flex justify-between mb-8">
        <div className="w-1/2 pr-4">
          <h3 className="font-bold text-gray-700 uppercase text-xs mb-1">Pengirim:</h3>
          <p className="font-bold text-lg">{invoice.sender_name}</p>
          <p className="text-sm text-gray-600 whitespace-pre-wrap">{invoice.sender_address}</p>
        </div>
        <div className="w-1/2 pl-4 border-l border-gray-200">
          <h3 className="font-bold text-gray-700 uppercase text-xs mb-1">Penerima:</h3>
          <p className="font-bold text-lg">{invoice.receiver_name}</p>
          <p className="text-sm text-gray-600 whitespace-pre-wrap">{invoice.receiver_address}</p>
        </div>
      </div>
      <table className="w-full mb-8 border-collapse">
        <thead>
          <tr className="bg-gray-100 border-y-2 border-gray-800 text-sm">
            <th className="py-2 px-3 text-left font-bold">KODE</th>
            <th className="py-2 px-3 text-left font-bold">DESKRIPSI BARANG</th>
            <th className="py-2 px-3 text-center font-bold">QTY</th>
            <th className="py-2 px-3 text-right font-bold">HARGA (Rp)</th>
            <th className="py-2 px-3 text-right font-bold">JUMLAH (Rp)</th>
          </tr>
        </thead>
        <tbody className="text-sm">
          {invoice.details.map((detail: any, idx: number) => (
            <tr key={idx} className="border-b border-gray-200">
              <td className="py-3 px-3">{detail.item_code}</td>
              <td className="py-3 px-3">{detail.item_name}</td>
              <td className="py-3 px-3 text-center">{detail.quantity}</td>
              <td className="py-3 px-3 text-right">{detail.price.toLocaleString('id-ID')}</td>
              <td className="py-3 px-3 text-right">{detail.subtotal.toLocaleString('id-ID')}</td>
            </tr>
          ))}
        </tbody>
        <tfoot>
          <tr className="font-bold text-lg border-b-2 border-gray-800">
            <td colSpan={4} className="py-4 px-3 text-right">GRAND TOTAL</td>
            <td className="py-4 px-3 text-right">Rp {invoice.total_amount.toLocaleString('id-ID')}</td>
          </tr>
        </tfoot>
      </table>
      <div className="flex justify-between mt-16 text-sm">
        <div className="text-center w-48"><p className="mb-16">Penerima,</p><p className="border-t border-gray-400 pt-1">( Nama Terang & Ttd )</p></div>
        <div className="text-center w-48"><p className="mb-16">Petugas / Kerani,</p><p className="border-t border-gray-400 pt-1">{invoice.created_by_name}</p></div>
      </div>
    </div>
  );
};