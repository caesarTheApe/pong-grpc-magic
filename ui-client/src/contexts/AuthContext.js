import React, { createContext, useContext, useState, useEffect } from 'react';
import { ethers } from 'ethers';
import PongGameArtifact from '../contracts/PongGame.json';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [playerId, setPlayerId] = useState(null);
  const [provider, setProvider] = useState(null);
  const [pongGameContract, setPongGameContract] = useState(null);
  const [signer, setSigner] = useState(null);

  useEffect(() => {
    if (window.ethereum) {
      try {
        const newProvider = new ethers.BrowserProvider(window.ethereum);
        setProvider(newProvider);
        const signer = newProvider.getSigner();
        if (!signer) {
          console.error("No signer available. Make sure MetaMask is connected.");
        }
        // const networkId = '31337';  // Hardhat's default local network ID
        // console.log(PongGameArtifact.networks[networkId].address)
        const contractAddressLocalNetwork = "0xA51c1fc2f0D1a1b8494Ed1FE312d7C3a78Ed91C0";
        const contract = new ethers.Contract(
          contractAddressLocalNetwork, // Use your contract's network ID
          PongGameArtifact.abi,
          signer
        );
        setPongGameContract(contract);
      } catch (error) {
        console.error("Error setting up Ethereum provider:", error);
      }
    } else {
      console.log("Please install MetaMask to use this application.");
    }
  }, []);

  const login = async () => {
    if (!provider) {
        alert("Ethereum provider not initialized. Please refresh the page.");
        return;
    }

    try {
        const signer = await provider.getSigner();
        setSigner(signer);
        
        // Request accounts access
        const accounts = await provider.send("eth_requestAccounts", []);
        const account = accounts[0];
        setPlayerId(account);

        // Prepare and sign a message
        // const message = `Sign in on ${new Date().toISOString()}`;
        // const signature = await signer.signMessage(message);
        // console.log(`Signature: ${signature}`);
    } catch (error) {
        console.error("Error during login", error);
    }
  };

  const logout = () => {
    setPlayerId(null);
  };

  return (
    <AuthContext.Provider value={{ playerId, provider, signer, pongGameContract, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
