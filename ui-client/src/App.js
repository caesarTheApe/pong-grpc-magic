import React from 'react';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import GameBoard from './components/GameBoard';
import LoginComponent from './components/LoginComponent';

function App() {
  return (
    <AuthProvider>
      <div className="App">
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

  return <GameBoard playerId={playerId} />;
};

export default App;
