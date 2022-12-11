using Day22;

var isTest = args.Contains("--test");
var isStarTwo = args.Contains("--star-two");
var path = args[0];
var steps = File.ReadAllLines(path);
var stepCounter = 0;
var instructions = new List<Cube>();
var cubes = new List<Cube>();

foreach (var instruction in ReadCubes(path, isTest))
{
    var newCubes = new List<Cube>();
    foreach (var overlaps in cubes.Where(c => c.Intersects(instruction)))
    {
        newCubes.Add(overlaps.GetIntersection(instruction));
    }

    cubes.AddRange(newCubes);

    if (instruction.Polarity == Polarity.Positive)
        cubes.Add(instruction);

    Console.Write($"\rStep {stepCounter++} ({cubes.Count})");
}

foreach (var c in cubes)
    Console.WriteLine(c);

var cubesOn = cubes.Sum(c => c.GetVolume());
Console.WriteLine($"Cubes on {cubesOn} ({cubes.Count})");



IEnumerable<Cube> ReadCubes(string path, bool truncate)
{
    foreach (var cube in File.ReadAllLines(path))
    {
        var elements = cube.Split(" ");
        var dimensions = elements[1].Split(",");
        var ranges = new Dictionary<char, DimensionRange>();

        foreach (var dimension in dimensions)
        {
            var key = dimension[0];
            var minMax = dimension[2..].Split("..").Select(i => int.Parse(i)).ToArray();
            var min = GetValue(minMax[0]);
            var max = GetValue(minMax[1]);

            ranges.Add(key, new DimensionRange { Min = min, Max = max });
        }

        yield return new Cube
        {
            MinX = ranges['x'].Min,
            MaxX = ranges['x'].Max,
            MinY = ranges['y'].Min,
            MaxY = ranges['y'].Max,
            MinZ = ranges['z'].Min,
            MaxZ = ranges['z'].Max,
            Polarity = cube.StartsWith("on") ? Polarity.Positive : Polarity.Negative
        };



        int GetValue(int value)
        {
            if (!truncate)
                return value;

            if (value < -50)
                return -50;

            if (value > 50)
                return 50;

            return value;
        }
    }
}


public readonly struct DimensionRange
{
    // public DimensionRange(int min, int max) => (Min, )
    public int Min { get; init; }
    public int Max { get; init; }
}

public readonly struct Location
{
    public int X { get; init; }
    public int Y { get; init; }
    public int Z { get; init; }

    public override string ToString() => $"Location {{ X = {X}, Y = {Y}, Z = {Z} }}";
}
