import { Route, RouteProps, Routes } from 'react-router-dom';

import MainLayout from '@/components/layout/MainLayout';
import CustomRouter from '@/config/routes/CustomRouter';
import history from '@/config/routes/history';
import publicRoutes from '@/config/routes/publicRoutes';

function App() {
  const renderRouter = (routes: RouteProps[]) => {
    return routes.map((route, index) => <Route path={route.path} element={route.element} key={index} />);
  };

  return (
    <CustomRouter history={history}>
      <Routes>
        <Route path="/" element={<MainLayout />}>
          {renderRouter(publicRoutes)}
        </Route>
      </Routes>
    </CustomRouter>
  );
}

export default App;
