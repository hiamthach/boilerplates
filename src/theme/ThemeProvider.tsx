import React from 'react';

import { MantineProvider } from '@mantine/core';

const ThemeProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <MantineProvider withNormalizeCSS withGlobalStyles>
      {children}
    </MantineProvider>
  );
};

export default ThemeProvider;
