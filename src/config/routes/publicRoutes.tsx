import { RouteProps } from 'react-router-dom';

import ComingSoon from '@/pages/exeption/ComingSoon';
import NotFound from '@/pages/exeption/NotFound';
import HomePage from '@/pages/main/HomePage';

const publicRoutes: RouteProps[] = [
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: '/coming-soon',
    element: <ComingSoon />,
  },
  {
    path: '*',
    element: <NotFound />,
  },
];

export default publicRoutes;
