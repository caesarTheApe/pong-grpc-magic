import React from 'react';
import { useAuth } from '../contexts/AuthContext';

const LoginComponent = () => {
  const { login } = useAuth();

  return (
    <div className="flex h-screen items-center justify-center bg-gray-800">
      <button onClick={login} className="px-6 py-2 text-white font-semibold rounded-lg bg-green-500 hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-700">
        Login with Browser Wallet
      </button>
    </div>
  );
};

export default LoginComponent;
