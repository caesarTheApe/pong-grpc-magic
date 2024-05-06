import { useState, useEffect } from 'react';
import { PongGameClient } from '../pongrpc/pong_grpc_web_pb';
import { GameStreamRequest, PlayerInput, SignalReadyRequest, GameStartedRequest } from '../pongrpc/pong_pb';

export const useGameState = (playerId) => {
  const [gameState, setGameState] = useState(null);
  const [gameStarted, setGameStarted] = useState(false);
  const [error, setError] = useState('');
  const client = new PongGameClient('http://localhost:8085', null, null);
  const metadata = { 'client-id': playerId };

  useEffect(() => {
    const streamRequest = new GameStreamRequest();
    // streamRequest.setPlayerId(playerId); // Ensure you're setting the player ID
    const stream = client.streamUpdates(streamRequest, metadata);
    // const stream = client.streamUpdates(streamRequest, {});
    stream.on('data', (response) => {
      try {
        // Assuming response.getData() gives you the binary data
        const bytes = response.getData();  
        const json = new TextDecoder('utf-8').decode(bytes);
        const update = JSON.parse(json);  // Parse the JSON to get the GameUpdate object

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

    // Handle stream end or cancel
    stream.on('end', () => {
      console.log('Stream ended by server');
      setError('Game stream ended unexpectedly.');
    });

    // Cleanup on unmount
    return () => {
      console.log('Cancelling game updates stream');
      stream.cancel();
    };
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

  const signalReady = () => {
    const req = new SignalReadyRequest();
    req.setClientId(playerId);
  
    // Creating a new promise to handle the async call
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
  };

  useEffect(() => {
    const notifyRequest = new GameStartedRequest();
    notifyRequest.setClientId(playerId);
    const notifyStream = client.notifyGameStarted(notifyRequest, metadata);

    notifyStream.on('data', (response) => {
      console.log(response.toObject())
      setGameStarted(true)
      setError(null)
      // Handle game start notifications
    });

    notifyStream.on('error', (err) => {
      // Handle errors
    });

    notifyStream.on('end', () => {
      // Handle stream end
    });

    return () => notifyStream.cancel();
  }, [playerId]);

  return { gameState, sendInput, error, signalReady, gameStarted };
};