var test = args.Contains("--test");
var path = test ? "./input.test.txt" : "./input.txt";
var starTwo = args.Contains("--star-two");
var rawGrid = new List<List<int>>();
var grid = new List<Day16.Cell>();


// read input
foreach (var line in File.ReadAllLines(path))
{
    if (line.Length > 0)
        rawGrid.Add(new(line.Select(c => int.Parse(c.ToString()))));
}

if (starTwo)
{
    rawGrid = EnlargeGrid(rawGrid);
}

for (var r = 0; r < rawGrid.Count; r++)
{
    for (var c = 0; c < rawGrid[r].Count; c++)
    {
        grid.Add(new Day16.Cell
        {
            Row = r,
            Column = c,
            Value = rawGrid[r][c],
            Distance = r == c && r == 0 ? 0 : int.MaxValue,
            IsDeparture = r == 0 && c == 0,
            IsDestination = r == rawGrid.Count - 1 && c == rawGrid[r].Count - 1
        });
    }
}


// let's search this thing
Day16.PathFinder.DijkstraSearch(grid, test);



// var bestScore = long.MaxValue;
// var visitCache = new Dictionary<(int row, int column), int>();


// // HACK: Value of starting cell does not contribute to total.
// GetSomething(totalScore: grid[0][0] * -1, row: 0, column: 0, depth: 0);
// Console.WriteLine($"Risk: {bestScore}");


// // plot our path
// void GetSomething(int totalScore, int row, int column, int depth)
// {
//     // Console.WriteLine($"New cell: {totalScore} ({row}x{column} @ {depth})");

//     if (totalScore > bestScore)
//         return;

//     if (row < 0 || row >= grid.Count)
//         return;

//     if (column < 0 || column >= grid[row].Count)
//         return;


//     totalScore += grid[row][column];

//     if (visitCache.TryGetValue((row, column), out var value))
//     {
//         // We've been before via a safer route.
//         if (totalScore > value)
//         {
//             return;
//         }

//         visitCache[(row, column)] = totalScore;
//     }
//     else
//     {
//         visitCache.Add((row, column), totalScore);
//     }


//     if (row == (grid.Count -1) && column == (grid[row].Count - 1))
//     {
//         if (totalScore < bestScore)
//         {
//             Console.WriteLine($"Candidate route found | Risk = {totalScore} | Depth = {depth}");
//             bestScore = totalScore;
//             return;
//         }
//         else
//         {
//             return;
//         }
//     }
//     else
//     {
//         GetSomething(totalScore, row, column + 1, depth + 1);
//         GetSomething(totalScore, row + 1, column, depth + 1);
//     }
// }

List<List<int>> EnlargeGrid(List<List<int>> original)
{
    var result = new List<List<int>>();

    for (var row = 0; row < original.Count; row++)
    {
        result.Add(new List<int>());

        for (var offset = 0; offset < 5; offset++)
        {
            result[row].AddRange(original[row].Select(i => i + offset > 9 ? i + offset -9 : i + offset));
        }
    }

    for (var offset = 1 ; offset < 5; offset++)
    {
        for (var row = 0; row < original.Count; row++)
        {
            result.Add(result[row].Select(i => i + offset > 9 ? i + offset - 9 : i + offset).ToList());
        }
    }

    return result;
}

// void PrintGrid()
// {
//     foreach (var row in grid)
//     {
//         foreach (var column in row)
//         {
//             Console.Write(column);
//         }

//         Console.WriteLine(string.Empty);
//     }
// }
