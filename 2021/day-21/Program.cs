var player1 = new Player { Index = 0, Position = 1 };   // 4
var player2 = new Player { Index = 1, Position = 3 };   // 8
var playerWins = new long[] { 0, 0 };

// Rolling 3 3-sided dice --------------------------------------------------------------------------
// There are 27 outcomes, of which 7 are unique.  The smallest you can role is a 3 (1, 1, 1).  The
// largest is a 9 (3, 3, 3).  There are multiple ways to make each of the numbers in between.  The
// full list is:
//  Total face value: 3, 4,  5,  6,  7,  8,  9  ( 7)
//  Possible hands:   1, 3,  6,  7,  6,  3,  1  (27)
var handCache = new Dictionary<int, long>{ {3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1} };

// Scores ------------------------------------------------------------------------------------------
// There are 30 possible scores in the game.  The lowest score is the starting score of 0.  The
// highest is 29 (you can score up 20 and still have another go, scoring up to 9 on your last go).

// Moves -------------------------------------------------------------------------------------------
// There are 2,100 moves in the game.  10 starting positions x 30 possible scores x 7 possible dice
// roles.

var cache = new Dictionary<(int position, int score, int roll), (int newPosition, int newScore, long hands)>();
for (var position = 1; position <= 10; position++)
{
    for (var score = 0; score <= 29; score++)
    {
        for (var roll = 3; roll <= 9; roll++)
        {
            var newPosition = (position + roll) % 10;
            if (newPosition == 0)
                newPosition = 10;

            cache.Add((position, score, roll), (newPosition, score + newPosition, handCache[roll]));
        }
    }
}

PlayGame(new[] { player1, player2 }, 0, 1);
Console.WriteLine($"Player one wins {playerWins[0]} | Player two wins {playerWins[1]}");
Environment.Exit(0);



void PlayGame(Player[] players, int index, long games)
{
    var rolls = new[] { 3, 4, 5, 6, 7, 8, 9 };

    foreach (var roll in rolls)
    {
        (int newPosition, int newScore, long hands) = cache[(players[index].Position, players[index].Score, roll)];
        if (newScore >= 21)
        {
            playerWins[index] += games * hands;
        }
        else
        {
            var updatedPlayer = players[index] with { Position = newPosition, Score = newScore };
            PlayGame(
                new[]
                {
                    index == 0 ? updatedPlayer : players[0],
                    index == 1 ? updatedPlayer : players[1],
                },
                index == 0 ? 1 : 0,
                games * hands);
        }
    }
}


public readonly record struct Player(
    int Index,
    int Position,
    int Score);
