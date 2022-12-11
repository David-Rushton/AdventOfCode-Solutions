using System.Diagnostics;
using System.Text.RegularExpressions;

var test = args.Contains("--test");
var starTwo = args.Contains("--star-two");
Day12.Cave start = new Day12.Cave("start");
Day12.Cave end = new Day12.Cave("end");
var caves = new Dictionary<Day12.Cave, List<Day12.Cave>>
{
    { start, new List<Day12.Cave>() },
    { end, new List<Day12.Cave>() }
};
List<string> paths = new();


foreach (var line in File.ReadAllLines(test? "./input.test.txt" : "./input.txt"))
{
    var elements = line.Split('-');
    var from = new Day12.Cave(elements[0]);
    var to = new Day12.Cave(elements[1]);

    if (!caves.ContainsKey(from))
        caves.Add(from, new List<Day12.Cave>());

    if (!caves.ContainsKey(to))
        caves.Add(to, new List<Day12.Cave>());

    caves[from].Add(to);
    caves[to].Add(from);
}

var specialCaves = starTwo
    ? caves.Where(c => c.Key.MaybeTwice).Select(c => c.Key)
    : new[] { start };

foreach (var specialCave in specialCaves)
    PlotPathsSpecial(start, specialCave, string.Empty);

var completePaths = paths.Where(p => p.EndsWith("end")).ToArray();
Console.WriteLine(string.Join("\n", completePaths.OrderBy(p => p)));
Console.WriteLine($"\nComplete paths: {completePaths.Length}");


void PlotPathsSpecial(Day12.Cave startFrom, Day12.Cave specialCave, string path)
{
    Debug.Assert(caves is not null);
    Debug.Assert(paths is not null);

    path += path.Length > 0 ? $",{startFrom.Name}" : startFrom.Name;
    if (! paths.Contains(path))
        paths.Add(path);

    if (startFrom.Name == "end")
    {
        Console.WriteLine($"found: {specialCave.Name} - {path}");
        return;
    }

    foreach (var cave in caves[startFrom])
    {
        if (cave.AtMostOnce && !cave.MaybeTwice && path.Contains(cave.Name))
            continue;

        if (cave.MaybeTwice)
        {
            var removedLength = path.Replace(cave.Name, string.Empty).Length;
            var instances = (path.Length - removedLength) / cave.Name.Length;

            if ((cave.Name != specialCave.Name && instances == 1) || (cave.Name == specialCave.Name && instances == 2))
                continue;
        }

        PlotPathsSpecial(cave, specialCave, path);
    }
}
