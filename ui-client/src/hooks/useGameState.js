import { useState, useEffect } from 'react';
import { PongGameClient } from '../pongrpc/pong_grpc_web_pb';
import { GameStreamRequest, PlayerInput, SignalReadyRequest, GameStartedRequest } from '../pongrpc/pong_pb';
import { useAuth } from '../contexts/AuthContext';
import { ethers } from 'ethers';

export const useGameState = (playerId) => {
  const [gameState, setGameState] = useState(null);
  const [gameStarted, setGameStarted] = useState(false);
  const [error, setError] = useState('');
  const client = new PongGameClient('http://localhost:8085', null, null);
  const metadata = { 'client-id': playerId };
  
  const { pongGameContract, signer } = useAuth();

  useEffect(() => {
    const streamRequest = new GameStreamRequest();
    const stream = client.streamUpdates(streamRequest, metadata);
    stream.on('data', (response) => {
      try {
        const bytes = response.getData();
        const json = new TextDecoder('utf-8').decode(bytes);
        const update = JSON.parse(json);

        setGameState({
          p1Score: update.p1Score,
          p2Score: update.p2Score,
          ballX: update.ballX,
          ballY: update.ballY,
          p1Y: update.p1Y,
          p2Y: update.p2Y,
          gameWidth: update.gameWidth,
          gameHeight: update.gameHeight,
          p1Height: update.p1Height,
          p2Height: update.p2Height
        });
      } catch (err) {
        console.error('Failed to decode game update:', err);
        setError('Failed to process game update.');
      }
    });

    stream.on('error', (err) => {
      console.error('Stream Error:', err);
      setError('Failed to connect to game updates stream.');
    });

    stream.on('end', () => {
      console.log('Stream ended by server');
      setError('Game stream ended unexpectedly.');
    });

    return () => stream.cancel();
  }, [gameStarted, playerId]);

  const sendInput = (input) => {
    const request = new PlayerInput();
    request.setInput(input);

    return new Promise((resolve, reject) => {
      client.sendInput(request, metadata, (err, response) => {
        if (err) {
          console.error('Send Input Error:', err);
          setError(`Failed to send input: ${err.message}`);
          reject(err);
        } else {
          console.log('Input sent successfully:', response.toObject());
          resolve(response.toObject());
        }
      });
    });
  };

  const signalReady = async (betAmount) => {
    if (!pongGameContract) {
      setError("Contract not loaded");
      return;
    }
  
    const transactionOptions = {
      from: signer.address,
      value: ethers.parseEther(betAmount)  // Convert the bet amount to Wei
    };
  
    try {
      console.log('Sending bet amount:', betAmount)
      console.log(transactionOptions)
      // await pongGame.lockOrJoinGame(betAmount, { from: accounts[0], value: betAmount });
      const transaction = await pongGameContract.connect(signer).lockOrJoinGame(ethers.parseEther(betAmount), transactionOptions);
      await transaction.wait();
  
      const req = new SignalReadyRequest();
      req.setClientId(playerId);
      return new Promise((resolve, reject) => {
        client.signalReady(req, null, (err, response) => {
          if (err) {
            console.error('Error signaling readiness:', err.message);
            setError(err.message); // Update the state to reflect the error
            reject(err.message);  // Reject the promise with the error message
          } else {
            console.log('Signal ready response received successfully:', response.toObject());
            resolve();
          }
        });
      });
    } catch (error) {
      console.error('Error signaling readiness or locking in game bet:', error);
      setError(error.message);
      throw error;
    }
  };

  useEffect(() => {
    const notifyRequest = new GameStartedRequest();
    notifyRequest.setClientId(playerId);
    const notifyStream = client.notifyGameStarted(notifyRequest, metadata);

    notifyStream.on('data', (response) => {
      console.log(response.toObject())
      setGameStarted(true);
      setError(null);
    });

    notifyStream.on('error', (err) => {
      console.error('Notify Stream Error:', err);
      setError('Failed to connect to game start notifications.');
    });

    notifyStream.on('end', () => {
      console.log('Notify Stream ended by server');
      setError('Game start notifications stream ended unexpectedly.');
    });

    return () => notifyStream.cancel();
  }, [playerId]);

  return { gameState, sendInput, error, signalReady, gameStarted };
};
