import '@sumup/design-tokens/light.css';
import '@sumup/circuit-ui/styles.css';
import type { AppProps } from 'next/app';
import { light } from '@sumup/design-tokens';
import { cache } from '@emotion/css';
import { ThemeProvider, CacheProvider } from '@emotion/react';
import { ModalProvider } from '@sumup/circuit-ui';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ThemeProvider theme={light}>
      <CacheProvider value={cache}>
        <ModalProvider>
          <Component {...pageProps} />
        </ModalProvider>
      </CacheProvider>
    </ThemeProvider>
  );
}
