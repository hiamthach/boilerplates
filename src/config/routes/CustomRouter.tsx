import React, { useLayoutEffect, useState } from 'react';
import { Router } from 'react-router-dom';

import { Action, History, Location } from 'history';

interface ICustomRouterProps {
  history: History;
  children: React.ReactNode;
}

interface ICustomRouterState {
  action: Action;
  location: Location;
}

const CustomRouter = ({ children, history, ...props }: ICustomRouterProps) => {
  const [state, setState] = useState<ICustomRouterState>({
    action: history.action,
    location: history.location,
  });

  useLayoutEffect(() => history.listen(setState), [history]);

  return (
    <Router {...props} location={state.location} navigationType={state.action} navigator={history}>
      {children}
    </Router>
  );
};

export default CustomRouter;
