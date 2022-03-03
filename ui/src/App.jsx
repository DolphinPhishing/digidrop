import React, { useState, useEffect } from 'react';
import {
  Outlet, Routes, Route,
} from 'react-router-dom';
import './App.scss';
import { gql, useQuery } from '@apollo/client';
import Home from './pages/home';
import Admin from './pages/admin';
import Drop from './pages/drop';
import Compromised from './pages/compromised';

const IS_COMPROMISED_QUERY = gql`
  query compromised {
    compromised
  }
`;

const App = () => {
  const { data, error, loading } = useQuery(IS_COMPROMISED_QUERY);
  const [user, setUser] = useState(null);
  const [isCompromised, setIsCompromised] = useState(false);

  useEffect(() => {
    fetch(`${process.env.REACT_APP_API_URI}/me`, {
      credentials: 'include',
    }).then((res) => {
      if (res.status === 410) setIsCompromised(true);
      return res.json();
    }).then((u) => {
      setUser(u);
    }).catch(console.error);
  }, []);

  if (isCompromised || (user && user.type !== 'ADMIN' && (!loading && ((data && data.compromised) || error)))) {
    return <Compromised />;
  }
  return (
    <Routes>
      <Route
        path="/"
        element={(
          <div className="app">
            <div className="content">
              <Outlet />
            </div>
          </div>
    )}
      >
        <Route index element={<Home />} />
      </Route>
      <Route path="/drop" element={<Drop />} />
      <Route path="/admin" element={<Admin />} />
    </Routes>
  );
};

export default App;
