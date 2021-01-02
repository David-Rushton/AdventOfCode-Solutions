using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    class RecursiveCombat
    {
        Dictionary<string, int> _previousRoundsAllGames = new();
        Dictionary<string, int> _previousGames = new();


        public int Play(int[] playerOneStartingHand, int[] playerTwoStartingHand, int game = 1)
        {
            var previousRoundsInThisGame = new Dictionary<string, int>();
            var playerOneHand = new Queue<int>(playerOneStartingHand);
            var playerTwoHand = new Queue<int>(playerTwoStartingHand);
            var roundNumber = 0;
            var roundKey = string.Empty;
            var gameKey = GetRoundKey();


            Debug.Assert(playerOneStartingHand.Count() > 0 && playerTwoStartingHand.Count() > 0, "Cannot play game");


            if(_previousGames.ContainsKey(gameKey))
            {
                Console.WriteLine($"Game level matchy matchy! {_previousGames.Count}");
                return _previousGames[gameKey];
            }


            while(playerOneHand.Count > 0 && playerTwoHand.Count > 0)
            {
                roundKey = GetRoundKey();
                roundNumber++;
                PrintRound();


                if(CanShortcutResult(out var winner))
                    return winner;


                var playerOneCard = playerOneHand.Dequeue();
                var playerTwoCard = playerTwoHand.Dequeue();
                ConsoleEx.WriteLine($"Player one plays: {playerOneCard}", game);
                ConsoleEx.WriteLine($"Player two plays: {playerTwoCard}", game);


                if(PlaySubGame(playerOneCard, playerTwoCard))
                {
                    ConsoleEx.WriteLine("Entering sub game...", game);
                    var subGameWinner = Play(playerOneHand.Take(playerOneCard).ToArray(), playerTwoHand.Take(playerTwoCard).ToArray(), game + 1);
                    SetRoundWinner(subGameWinner, playerOneCard, playerTwoCard);
                }
                else
                    SetRoundWinner(playerOneCard > playerTwoCard ? 1 : 2, playerOneCard, playerTwoCard);
            }


            var gameWinner = playerOneHand.Count == 0 ? 2 : 1;
            if( ! _previousGames.ContainsKey(gameKey) )
                _previousGames.Add(gameKey, gameWinner);


            if(game == 1)
                PrintWinner();


            return gameWinner;



            string GetRoundKey() =>  $"{ string.Join(',', playerOneHand) }:{ string.Join(',', playerTwoHand) }";

            bool CanShortcutResult(out int winner)
            {
                if(previousRoundsInThisGame.ContainsKey(roundKey))
                {
                    winner = 1;
                    ConsoleEx.WriteLine("Matchy matchy!  Recursive result is player 1 wins");
                    return true;
                }
                previousRoundsInThisGame.Add(roundKey, 1);

                if(_previousRoundsAllGames.ContainsKey(roundKey))
                {
                    winner = _previousRoundsAllGames[roundKey];
                    Debug.Assert(winner is 1 or 2, "Cannot shortcut round - winner not recorded");
                    ConsoleEx.WriteLine($"Matchy matchy!  Recursive result is player {winner} wins");
                    return true;
                }


                winner = 0;
                return false;
            }

            bool PlaySubGame(int playerOneCard, int playerTwoCard) =>
                playerOneHand.Count >= playerOneCard && playerTwoHand.Count >= playerTwoCard
            ;

            void SetRoundWinner(int winningPlayer, int playerOneCard, int playerTwoCard)
            {
                var roundKey = GetRoundKey();

                Debug.Assert(winningPlayer is 1 or 2, $"Unexpected winner {winningPlayer}");
                if(winningPlayer == 1)
                {
                    playerOneHand.Enqueue(playerOneCard);
                    playerOneHand.Enqueue(playerTwoCard);
                }
                else
                {
                    playerTwoHand.Enqueue(playerTwoCard);
                    playerTwoHand.Enqueue(playerOneCard);
                }

                if( ! _previousRoundsAllGames.ContainsKey(roundKey) )
                    _previousRoundsAllGames.Add(roundKey, winningPlayer);

                ConsoleEx.WriteLine($"Player {( winningPlayer == 1 ? "one" : "two" )} wins round {roundNumber} of game {game}\n", game);
            }

            void PrintRound()
            {
                ConsoleEx.WriteLine($"\nRound {roundNumber} of game {game}---------------------------", game);
                ConsoleEx.WriteLine($"  Player 1: { string.Join(", ", playerOneHand.Select(c => c)) }", game);
                ConsoleEx.WriteLine($"  Player 2: { string.Join(", ", playerTwoHand.Select(c => c)) }", game);
            }

            void PrintWinner(int overrideWinner = 0)
            {
                var winningHand = overrideWinner == 1 || playerOneHand.Count > 0 ? playerOneHand : playerTwoHand;
                var winningPlayer = overrideWinner == 1 || playerOneHand.Count > 0 ? "Player 1" : "Player 2";
                var winningScore = GetWinningScore(winningHand);

                ConsoleEx.WriteLine($"\nWinner: {winningPlayer}\nScore: {winningScore}", game);
            }

            int GetWinningScore(Queue<int> winningHand)
            {
                var cardMultiplier = 1;
                var cards = winningHand.ToArray().Reverse();
                var score = 0;

                foreach(var card in cards)
                {
                    score += (card * cardMultiplier);
                    cardMultiplier++;
                }

                return score;
            }
        }
    }
}
