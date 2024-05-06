import React, { createContext, useContext, useState } from 'react';
import { ethers } from 'ethers';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [playerId, setPlayerId] = useState(null);

  const login = async () => {
    if (window.ethereum) {
      try {
        // Correctly instantiate the provider using the ethers package
        const provider = new ethers.BrowserProvider(window.ethereum);
        const accounts = await provider.send("eth_requestAccounts", []);
        setPlayerId(accounts[0]); // Set the first account as the playerId
      } catch (error) {
        console.error("Error connecting to MetaMask", error);
      }
    } else {
      alert("Please install MetaMask!");
    }
  };

  const logout = () => {
    setPlayerId(null);
  };

  return (
    <AuthContext.Provider value={{ playerId, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
