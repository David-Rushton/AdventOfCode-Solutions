var isTest = args.Contains("--test");
var path = isTest ? "./input.test.txt" : "./input.txt";
var grid = new List<List<char>>();

// read input into a 2d grid.
foreach (var line in File.ReadAllLines(path))
{
    grid.Add(new (line.Select(c => c)));
}

// Let's do this.
RunProgram(grid);



void RunProgram(List<List<char>> grid)
{
    var updated = 0;
    var step = 0;

    do
    {
        Console.WriteLine($"\nStep {step}");
        PrettyPrint(grid);
        step++;
        (updated, grid) = NextStep(grid);

    } while (updated > 0);

    Console.WriteLine($"\nStep {step}");
    PrettyPrint(grid);
}


(int updated, List<List<char>> nextGrid) NextStep(List<List<char>> grid)
{
    var moves = 0;
    var last = CloneGrid(grid);
    var result = CloneGrid(grid);

    foreach (var direction in new[] { '>', 'v' })
    {
        for (var r = 0; r < last.Count; r++)
        {
            for (var c = 0; c < last[r].Count; c++)
            {
                var next = direction == '>' ? NextLeft(r, c) : NextDown(r, c);

                if (last[r][c] == direction && last[next.r][next.c] == '.')
                {
                    moves++;
                    result[r][c] = '.';
                    result[next.r][next.c] = direction;
                }
            }
        }

        last = CloneGrid(result);
    }

    return (moves, result);

    (int r, int c) NextLeft(int row, int column)
    {
        return column + 1 < grid[row].Count
            ? (row, column + 1)
            : (row, 0);
    }

    (int r, int c) NextDown(int row, int column)
    {
        return row + 1 < grid.Count
            ? (row + 1, column)
            : (0, column);
    }
}

List<List<char>> CloneGrid(List<List<char>> grid)
{
    var result = new List<List<char>>();

    for (var r = 0; r < grid.Count; r++)
    {
        result.Add(new List<char>());

        for (var c = 0; c < grid[r].Count; c++)
        {
            result[r].Add(grid[r][c]);
        }
    }

    return result;
}

void PrettyPrint(List<List<char>> grid)
{
    if (!isTest)
        return;

    for (var r = 0; r < grid.Count; r++)
    {
        for (var c = 0; c < grid[r].Count; c++)
        {
            Console.Write(grid[r][c]);
        }

        Console.WriteLine(string.Empty);
    }
}
