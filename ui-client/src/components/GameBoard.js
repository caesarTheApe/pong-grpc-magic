import React, { useEffect } from 'react';
import { useGameState } from '../hooks/useGameState';

export default function GameBoard({ playerId }) {
  const { gameState, sendInput, error, signalReady, gameStarted } = useGameState(playerId);

  useEffect(() => {
    const handleKeyPress = (event) => {
      if (event.key === 'ArrowUp' || event.key === 'w') {
        sendInput('ArrowUp');
      } else if (event.key === 'ArrowDown' || event.key === 's') {
        sendInput('ArrowDown');
      } else if (event.key === ' ') {  // Handle space bar
        signalReady("0.01");
      }
    };

    window.addEventListener('keydown', handleKeyPress);
    return () => window.removeEventListener('keydown', handleKeyPress);
  }, [sendInput, signalReady, gameStarted]); // Include signalReady in the dependency array

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!gameState) {
    return <div>Loading game...</div>;
  }

  // Calculate positions based on the board dimensions
  const boardWidth = gameState.gameWidth*10; // Same as CSS
  const boardHeight = gameState.gameHeight*10; // Same as CSS
  const p1Height = gameState.p1Height; // Same as CSS
  const p2Height = gameState.p2Height;

  // Ball position
  const ballStyle = {
    left: `${(gameState.ballX / 100) * boardWidth}px`,
    top: `${(gameState.ballY / 100) * boardHeight}px`
  };

  // Paddle positions
  const paddle1Style = {
    left: '5px', // 5px from the left edge
    top: `${(gameState.p1Y / 100) * (boardHeight - p1Height)}px`
  };

  const paddle2Style = {
    right: '5px', // 5px from the right edge
    top: `${(gameState.p2Y / 100) * (boardHeight - p2Height)}px`
  };

  return (
    <div>
      <div className="game-board" style={{ width: `${boardWidth}px`, height: `${boardHeight}px` }}>
        <div className="paddle" style={paddle1Style}></div>
        <div className="paddle" style={paddle2Style}></div>
        <div className="ball" style={ballStyle}></div>
      </div>
      <div className="score">
        Player 1 Score: {gameState.p1Score} - Player 2 Score: {gameState.p2Score}
      </div>
    </div>
  );
}
