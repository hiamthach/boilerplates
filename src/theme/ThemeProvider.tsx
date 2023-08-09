import React from 'react';

import { ConfigProvider } from 'antd';

const ThemeProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <ConfigProvider
      theme={{
        token: {
          borderRadius: 8,
        },
      }}
    >
      {children}
    </ConfigProvider>
  );
};

export default ThemeProvider;
