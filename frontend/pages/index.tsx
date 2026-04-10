import React from 'react';
import Head from 'next/head';
import { Package, LogOut } from 'lucide-react';
import { useStore } from '@/store/useStore';
import { useHasHydrated } from '@/hooks/useHasHydrated';
import { LoginForm } from '@/components/auth/LoginForm';
import { ProgressBar } from '@/components/wizard/ProgressBar';
import { Step1 } from '@/components/wizard/Step1';
import { Step2 } from '@/components/wizard/Step2';
import { Step3 } from '@/components/wizard/Step3';
import { PrintLayout } from '@/components/invoice/PrintLayout';

export default function Home() {
  const hasHydrated = useHasHydrated();
  const { user, currentStep, setUser, resetForm, invoiceData } = useStore();

  if (!hasHydrated) return null; // Mencegah Hydration Mismatch Error

  if (!user) return <LoginForm />;

  return (
    <>
      <Head>
        <title>Resi Generator - Fleetify</title>
      </Head>

      {/* Kontainer Utama - Disembunyikan saat Print */}
      <div className="min-h-screen bg-gray-100 print:hidden">
        <header className="bg-white shadow-sm">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4 flex justify-between items-center">
            <div className="flex items-center space-x-2">
              <Package className="h-6 w-6 text-blue-600" />
              <h1 className="text-xl font-bold text-gray-900">Resi Generator</h1>
            </div>
            <div className="flex items-center space-x-4">
              <div className="text-sm text-gray-600">
                Halo, <span className="font-semibold text-gray-900">{user.name}</span> 
                <span className="ml-2 px-2 py-0.5 rounded-full bg-blue-100 text-blue-800 text-xs capitalize">{user.role}</span>
              </div>
              <button onClick={() => { resetForm(); setUser(null); }} className="text-gray-500 hover:text-red-600 transition">
                <LogOut className="h-5 w-5" />
              </button>
            </div>
          </div>
        </header>

        <main className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <ProgressBar currentStep={currentStep} />
          <div className="bg-white shadow rounded-lg p-6 sm:p-8 mt-12">
            {currentStep === 1 && <Step1 />}
            {currentStep === 2 && <Step2 />}
            {currentStep === 3 && <Step3 />}
          </div>
        </main>
      </div>

      {/* Layout Cetak (Hanya Muncul Saat di Print Window) */}
      <PrintLayout invoice={invoiceData} />
    </>
  );
}