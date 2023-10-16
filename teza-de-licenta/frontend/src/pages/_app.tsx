import '@sumup/design-tokens/light.css';
import '@sumup/circuit-ui/styles.css';
import type { AppProps } from 'next/app';
import { light } from '@sumup/design-tokens';
import { cache } from '@emotion/css';
import { ThemeProvider, CacheProvider } from '@emotion/react';

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ThemeProvider theme={light}>
      <CacheProvider value={cache}>
        <Component {...pageProps} />
      </CacheProvider>
    </ThemeProvider>
  );
}