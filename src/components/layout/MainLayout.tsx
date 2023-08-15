import { Outlet } from 'react-router-dom';

import Footer from './Footer';
import Header from './Header';

const MainLayout = () => {
  return (
    <main>
      <Header />
      <Outlet />
      <Footer />
    </main>
  );
};

export default MainLayout;
