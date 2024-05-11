import React from 'react';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import GameBoard from './components/GameBoard';
import LoginComponent from './components/LoginComponent';

function App() {
  return (
    <AuthProvider>
      <div className="App bg-gray-900 min-h-screen text-white">
        <AuthContent />
      </div>
    </AuthProvider>
  );
}

const AuthContent = () => {
  const { playerId } = useAuth();

  if (!playerId) {
    return <LoginComponent />;
  }

  return (
    <div className="flex justify-center items-center h-screen">
      <GameBoard playerId={playerId} />
    </div>
  );
};

export default App;
