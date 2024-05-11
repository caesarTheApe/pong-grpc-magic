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
  const boardWidth = gameState.gameWidth;
  const boardHeight = gameState.gameHeight;
  const p1Height = gameState.p1Height;
  const p2Height = gameState.p2Height;

  // Ball position
  const ballStyle = {
    left: `${gameState.ballX}px`,
    top: `${gameState.ballY}px`,
    width: `${gameState.ballWidth}px`,
    height: `${gameState.ballHeight}px`
  };

  // Paddle positions
  const paddle1Style = {
    left: '5px', // 5px from the left edge
    top: `${gameState.p1Y}px`,
    width: `${gameState.p1Width}px`,
    height: `${p1Height}px`
  };
  const paddle2Style = {
    right: '5px', // 5px from the right edge
    top: `${gameState.p2Y}px`,
    width: `${gameState.p2Width}px`,
    height: `${p2Height}px`
  };

  return (
    <div className="flex flex-col items-center justify-center my-8">
      <div className="game-board relative mx-auto" style={{ width: `${boardWidth}px`, height: `${boardHeight}px`, backgroundColor: 'black' }}>
        <div className="paddle bg-white" style={paddle1Style}></div>
        <div className="paddle bg-white" style={paddle2Style}></div>
        <div className="ball bg-red-500 rounded-full" style={ballStyle}></div>
      </div>
      <div className="score text-white text-xl mt-4">
        Player 1 Score: {gameState.p1Score} - Player 2 Score: {gameState.p2Score}
      </div>
    </div>
  );
}
