'use client';

import { Toaster } from 'react-hot-toast';

import { TOAST_DEFAULT_OPTIONS } from '@/config/helpers/toast.helper';

const AppProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      {children}
      <Toaster {...TOAST_DEFAULT_OPTIONS} />
    </>
  );
};

export default AppProvider;
