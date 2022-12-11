using System.Diagnostics;


// 159 -> 173 14
// x 👉, y 👇
// fold along x = vertical
// folder along y = horizontal

using System.Diagnostics.CodeAnalysis;

var isTestMode = args.Contains("--test");
var isStarTwo = args.Contains("--star-two");
var path = isTestMode ? "./input.test.txt" : "./input.txt";
var coordinates = new List<Coordinate>();
var folds = new List<Fold>();


// Parse input
foreach (var line in File.ReadAllLines(path))
{
    if (string.IsNullOrWhiteSpace(line))
        continue;

    if (line.StartsWith("fold along"))
    {
        var elements = line.Replace("fold along ", "").Split('=');
        var fold = new Fold { Axis = elements[0], Value = int.Parse(elements[1]) };
        folds.Add(fold);
    }
    else
    {
        var elements = line.Split(',');
        coordinates.Add(new Coordinate { X = int.Parse(elements[0]), Y = int.Parse(elements[1])});
    }
}


foreach (var fold in folds)
{
    if (fold.Axis == "x")
    {
        coordinates = FoldLeft(fold.Value, coordinates);
    }
    else
    {
        coordinates = FoldUp(fold.Value, coordinates);
    }

    Console.WriteLine($"Applied fold {fold.Axis} {fold.Value}");
    if (isTestMode)
        PrettyPrint(coordinates);
    Console.WriteLine($"Distinct Dots: {coordinates.Distinct().Count()}\n");

    if (!isStarTwo)
        Environment.Exit(0);
}

PrettyPrint(coordinates);



List<Coordinate> FoldUp(int lineNumber, List<Coordinate> original)
{
    var result = original.Where(c => c.Y < lineNumber).ToList();

    for (var newRow = 0; newRow < lineNumber; newRow++)
    {
        var originalRow = (lineNumber * 2) - newRow;
        var foldedRow = original
            .Where(c => c.Y == originalRow)
            .Select(c => new Coordinate { X = c.X, Y = newRow });

        result.AddRange(foldedRow);
    }

    return result;
}

List<Coordinate> FoldLeft(int lineNumber, List<Coordinate> original)
{
    var result = original.Where(c => c.X < lineNumber).ToList();

    for (var newColumn = 0; newColumn < lineNumber; newColumn++)
    {
        var originalColumn = (lineNumber * 2) - newColumn;
        var foldedColumn = original
            .Where(c => c.X == originalColumn)
            .Select(c => new Coordinate { X = newColumn, Y = c.Y });

        result.AddRange(foldedColumn);
    }

    return result;
}

void PrettyPrint(List<Coordinate> list)
{
    Debug.Assert(list is not null);

    var maxX = list.Select(c => c.X).Max();
    var maxY = list.Select(c => c.Y).Max();

    for (var y = 0; y <= maxY; y++)
    {
        for (var x = 0; x <= maxX; x++)
        {
            var cellOccupied = list.Where(c => c.X == x && c.Y == y).Any();
            Console.Write(cellOccupied ? "#" : ".");
        }

        Console.WriteLine(string.Empty);
    }
}


public struct Coordinate
{
    public int X { get; set; }
    public int Y { get; set; }
}

public struct Fold
{
    public string Axis { get; set; }
    public int Value { get; set; }
}
