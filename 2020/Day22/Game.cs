using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    class Game
    {
        public void Play(int[] playerOneStartingHand, int[] playerTwoStartingHand)
        {
            var playerOneHand = new Queue<int>(playerOneStartingHand);
            var playerTwoHand = new Queue<int>(playerTwoStartingHand);
            var roundNumber = 1;

            while(playerOneHand.Count > 0 && playerTwoHand.Count > 0)
            {
                var playerOneCard = playerOneHand.Dequeue();
                var playerTwoCard = playerTwoHand.Dequeue();

                if(playerOneCard > playerTwoCard)
                {
                    playerOneHand.Enqueue(playerOneCard);
                    playerOneHand.Enqueue(playerTwoCard);
                }
                else
                {
                    playerTwoHand.Enqueue(playerTwoCard);
                    playerTwoHand.Enqueue(playerOneCard);
                }

                PrintRound();
                roundNumber++;
            }


            PrintWinner();
            return ;


            void PrintRound()
            {
                Console.WriteLine($"\nRound: {roundNumber}");
                Console.WriteLine($"  Player 1: { string.Join(", ", playerOneHand.Select(c => c)) }");
                Console.WriteLine($"  Player 2: { string.Join(", ", playerTwoHand.Select(c => c)) }");
            }

            void PrintWinner()
            {
                var winningHand = playerOneHand.Count > 0 ? playerOneHand : playerTwoHand;
                var winningPlayer = playerOneHand.Count > 0 ? "Player 1" : "Player 2";
                var winningScore = GetWinningScore(winningHand);

                Console.WriteLine($"\nWinner: {winningPlayer}\nScore: {winningScore}");
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
