import React from 'react';
import { useAuth } from '../contexts/AuthContext';

const LoginComponent = () => {
  const { login } = useAuth();

  return (
    <div>
      <button onClick={login}>Login with Browser Wallet</button>
    </div>
  );
};

export default LoginComponent;
