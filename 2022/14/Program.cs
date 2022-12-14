var isTest = args[0] == "test";
var isStarTwo = args[1] == "star2";
var path = isTest ? "input.txt.test" : "input.txt";
var rocks = ParseInput(path);
var settledSand = new Dictionary<Point, int>();
var minRight = rocks.Min(rock => rock.right) - 20;
var maxRight = rocks.Max(rock => rock.right) + 20;
var minBottom = rocks.Max(rock => rock.bottom);
var maxBottom = rocks.Max(rock => rock.bottom);
var origin = new Point(500, 0);
var iteration = 1;

Console.CursorVisible = false;
Console.CancelKeyPress += (_, _) => Console.CursorVisible = true;
Console.Clear();

PrintMap(iteration, rocks, settledSand, origin, minRight, maxRight, minBottom, maxBottom);

while (true)
{
    var sand = origin;

    while (TryMoveSand(sand, rocks, settledSand, out var newSand))
    {
        iteration++;

        if (iteration % 10000 == 0 || isTest)
            PrintMap(iteration, rocks, settledSand, newSand, minRight, maxRight, minBottom, maxBottom);

        sand = newSand;

        if (isStarTwo && newSand.bottom == maxBottom + 1)
        {
            break;
        }

        if (!isStarTwo && newSand.bottom > maxBottom)
        {
            PrintMap(iteration, rocks, settledSand, newSand, minRight, maxRight, minBottom, maxBottom);
            Environment.Exit(0);
        }
    }

    settledSand.Add(sand, iteration);

    if (isStarTwo && sand == origin)
    {
        PrintMap(iteration, rocks, settledSand, sand, minRight, maxRight, minBottom, maxBottom);
        Environment.Exit(0);
    }
}


bool TryMoveSand(Point sand, HashSet<Point> rocks, Dictionary<Point, int> settledSand, out Point movedSand)
{
    foreach (var rightOffset in new [] {0, -1, 1})
    {
        var candidatePoint = new Point(sand.right + rightOffset, sand.bottom + 1);

        if (! (rocks.Contains(candidatePoint) || settledSand.ContainsKey(candidatePoint)))
        {
            movedSand = candidatePoint;
            return true;
        }
    }

    movedSand = sand;
    return false;
}

void PrintMap(int iteration, HashSet<Point> rocks, Dictionary<Point, int> settledSand, Point sand, int minRight, int maxRight, int minBottom, int maxBottom)
{
    Console.SetCursorPosition(0, 0);

    Console.WriteLine("== Map ==");
    Console.WriteLine($"- Iteration {iteration}");
    Console.WriteLine($"- Last Settled Sand {settledSand.LastOrDefault().Value}");
    Console.WriteLine($"- Settled Grains of Sand {settledSand.Count()}");
    Console.WriteLine();

    var sandOrigin = new Point(500, 0);

    var startingBottom = maxBottom > Console.BufferHeight + 9
        ? maxBottom - Console.BufferHeight + 9
        : 0;

    for (var bottom = startingBottom; bottom < maxBottom + 3; bottom++)
    {
        Console.Write($"{bottom:000} ");
        for (var right = minRight - 2; right < maxRight + 2; right++)
        {
            var point = new Point(right, bottom);
            var cell = rocks.Contains(point) ? '#' : '.';

            if (point == sandOrigin)
                cell = '+';

            if (settledSand.ContainsKey(point) || point == sand)
                cell = 'o';

            if (point.bottom == maxBottom + 2)
                cell = '#';

            Console.Write(cell);
        }

        Console.WriteLine();
    }

    // Thread.Sleep(5);
}

HashSet<Point> ParseInput(string path)
{
    var rocks = new HashSet<Point>();

    foreach (var line in File.ReadAllLines(path))
    {
        if (line.Length == 0)
            continue;

        var points = line.Split(" -> ");
        Point? lastPoint = null;

        foreach (var point in points)
        {
            var currentPoint = ParsePoint(point);

            if (lastPoint is not null)
            {
                var minRight = Math.Min(lastPoint.right, currentPoint.right);
                var maxRight = Math.Max(lastPoint.right, currentPoint.right);
                var minBottom = Math.Min(lastPoint.bottom, currentPoint.bottom);
                var maxBottom = Math.Max(lastPoint.bottom, currentPoint.bottom);

                for (var bottom = minBottom; bottom <= maxBottom; bottom++)
                {
                    for (var right = minRight; right <= maxRight; right++)
                    {
                        var newPoint = new Point(right, bottom);
                        if (!rocks.Contains(newPoint))
                            rocks.Add(newPoint);
                    }
                }
            }

            lastPoint = currentPoint;
        }
    }

    return rocks;
}

Point ParsePoint(string point)
{
    var elements = point.Split(",");
    var right = int.Parse(elements[0]);
    var bottom = int.Parse(elements[1]);

    return new(right, bottom);
}

public record Point(int right, int bottom);
