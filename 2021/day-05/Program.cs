using System.Diagnostics;
using System.Text.RegularExpressions;


var test = args.Contains("-t");
var starTwo = args.Contains("--star-two");
var path = test ? "./Input.Test.txt" : "Input.txt";
var lines = File.ReadAllLines(path);
var plottedPoints = new Dictionary<Point, int>();

foreach (var line in lines)
{
    var terminals = GetTerminals(line);
    if (starTwo || terminals.Start.X == terminals.End.X || terminals.Start.Y == terminals.End.Y)
    {
        foreach (var point in GetConnectingPoints(terminals))
        {
            if (!plottedPoints.ContainsKey(point))
            {
                plottedPoints.Add(point, 1);
            }
            else
            {
                plottedPoints[point]++;
            }
        }
    }
}


if (test)
    PlotTest(plottedPoints);

var overlappingPoints = plottedPoints.Where(kvp => kvp.Value > 1).Count();
Console.WriteLine($"Star one || Overlapping points: {overlappingPoints}");


static IEnumerable<Point> GetConnectingPoints(Terminals terminals)
{
    var point = terminals.Start;
    while (point != terminals.End)
    {
        yield return point;
        point = point with
        {
            X = GetNextValue(point.X, terminals.End.X),
            Y = GetNextValue(point.Y, terminals.End.Y)
        };
    }

    yield return terminals.End;

    int GetNextValue(int from, int until)
    {
        if (from == until)
            return from;

        return from > until ? from - 1 : from + 1;
    }
}

static Terminals GetTerminals(string segmentLine)
{
    var matches = new Regex(@"\d+").Matches(segmentLine);
    Debug.Assert(matches.Count == 4);

    return new Terminals(
        new Point(int.Parse(matches[0].Value), int.Parse(matches[1].Value)),
        new Point(int.Parse(matches[2].Value), int.Parse(matches[3].Value))
    );
}


void PlotTest(Dictionary<Point, int> points)
{
    Console.WriteLine("Plot ------------------\n");

    for (var y = 0; y < 10; y++)
    {
        for (var x = 0; x < 10; x++)
        {
            var point = new Point(x, y);
            var value = points.ContainsKey(point)
                ? points[point].ToString()
                : ".";
            Console.Write(value);
        }

        Console.Write("\n\n");
    }
}


public readonly record struct Point(int X, int Y);

public readonly record struct Terminals(Point Start, Point End);
