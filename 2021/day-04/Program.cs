

var (numbers, boards) = Day04.InputReader.Read(args.Contains("-t"));
var starTwo = args.Contains("--star-two");
var completed = 0;

foreach (var number in numbers)
{
    // Console.WriteLine($"\nNext Number: {number} -----------------------------------\n");

    foreach (var board in boards)
    {
        if (board.IsCompleted)
            continue;

        if (board.CheckForNumber(number) == Day04.CheckNumberResult.Winner)
        {
            var unmarked = board.GetSumOfUnmarked();
            var result = unmarked * number;

            Console.WriteLine($"Winner! {unmarked} * {number} = {result}\n");
            Console.WriteLine(board.ToString());

            if (starTwo)
            {
                completed++;
                if (completed == boards.Count)
                    Environment.Exit(0);
            }
            else
            {
                Environment.Exit(0);
            }
        }

        // Console.WriteLine(board);
    }
}


Console.WriteLine("I guess something went wrong");
