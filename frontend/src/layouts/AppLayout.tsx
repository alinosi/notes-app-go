import React from 'react';
import { Navbar } from '../components/Navbar';

interface AppLayoutProps {
  children: React.ReactNode;
  onSearch?: (query: string) => void;
}

export const AppLayout: React.FC<AppLayoutProps> = ({ children, onSearch }) => {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar onSearch={onSearch} />
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {children}
      </main>
    </div>
  );
};
