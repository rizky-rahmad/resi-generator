import React from 'react';
import { CheckCircle } from 'lucide-react';

interface Props {
  currentStep: number;
}

export const ProgressBar: React.FC<Props> = ({ currentStep }) => {
  const steps = [
    { id: 1, title: 'Data Klien' },
    { id: 2, title: 'Data Barang' },
    { id: 3, title: 'Review & Submit' },
  ];

  return (
    <nav aria-label="Progress" className="mb-8">
      <ol role="list" className="flex items-center justify-between">
        {steps.map((step, stepIdx) => (
          <li key={step.id} className={`relative ${stepIdx !== steps.length - 1 ? 'pr-8 sm:pr-20 w-full' : ''}`}>
            <div className="absolute inset-0 flex items-center" aria-hidden="true">
              {stepIdx !== steps.length - 1 && (
                <div className={`h-1 w-full rounded ${currentStep > step.id ? 'bg-blue-600' : 'bg-gray-200'}`} />
              )}
            </div>
            <div
              className={`relative flex h-8 w-8 items-center justify-center rounded-full border-2 bg-white ${
                currentStep === step.id
                  ? 'border-blue-600 text-blue-600 font-bold'
                  : currentStep > step.id
                  ? 'border-blue-600 bg-blue-600 text-white'
                  : 'border-gray-300 text-gray-500'
              }`}
            >
              {currentStep > step.id ? <CheckCircle className="h-5 w-5" /> : <span>{step.id}</span>}
            </div>
            <span className={`absolute -bottom-6 left-1/2 -translate-x-1/2 text-xs font-medium whitespace-nowrap ${currentStep >= step.id ? 'text-blue-600' : 'text-gray-500'}`}>
              {step.title}
            </span>
          </li>
        ))}
      </ol>
    </nav>
  );
};