import React from 'react';
import { css } from '@emotion/react';
import { SIDE_PANEL_WIDTH, SideNavigation, TOP_NAVIGATION_HEIGHT, TopNavigation } from '@sumup/circuit-ui';
import { Gauge, Search, SumUpCard, SumUpLogo } from '@sumup/icons';
import styled from '@emotion/styled';
import { useRouter } from 'next/router';

export const HEADER_HEIGHT = '60px';

const mainBaseStyles = () => css`
  display: flex;
  flex-direction: column;
  flex: 1;
  max-width: 100vw;
  min-height: calc(100vh - ${TOP_NAVIGATION_HEIGHT});
  padding: env(safe-area-inset-top) env(safe-area-inset-right) env(safe-area-inset-bottom) env(safe-area-inset-left);
  margin-left: 48px;
`;

const Main = styled.main(mainBaseStyles);

function Layout({ children }: { children: React.ReactNode }) {
  const router = useRouter();

  return (
    <>
      <SideNavigation
        isOpen={false}
        onClose={() => {}}
        primaryNavigationLabel="Reviews"
        closeButtonLabel="Close navigation"
        secondaryNavigationLabel="Settings"
        primaryLinks={[
          {
            label: 'Analysis',
            icon: Gauge,
            onClick: () => {
              router.push('/');
            },
          },
          {
            label: 'Reviews',
            icon: Search,
            onClick: () => {
              router.push('/reviews');
            },
          },
        ]}
      />
      <Main>{children}</Main>
    </>
  );
}

export default Layout;
