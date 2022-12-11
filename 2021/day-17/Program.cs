using Day17;

//test:     target area: x=20..30,  y=-10..-5
//puzzle:   target area: x=70..125, y=-159..-121

var isTest = args.Contains("--test");
var testZone = new Zone
{
    TopLeft = new(X: 20, Y: -5),
    BottomRight = new(X: 30, Y: -10)
};
var liveZone = new Zone
{
    TopLeft     = new(X: 70, Y: -121),
    BottomRight = new(X: 125, Y: -159)
};
var targetZone = isTest ? testZone : liveZone;
var firingZone = RangeFinder.GetFiringZone(targetZone);
var hits = new List<Velocity>();


for (var x = firingZone.TopLeft.X; x <= firingZone.BottomRight.X; x++)
{
    for (var y = firingZone.TopLeft.Y; y >= firingZone.BottomRight.Y; y--)
    {
        var velocity = new Velocity(x, y);
        if (IsHit(targetZone, velocity))
        {
            Console.WriteLine($"Direct Hit! {velocity}");
            hits.Add(velocity);
        }
    }
}

Console.WriteLine($"Total hits: {hits.Count}");




static bool IsHit(Zone targetZone, Velocity velocity)
{
    Point probe = new(0, 0);
    bool inRange = true;
    int stepCount = 0;

    while (inRange)
    {
        stepCount++;

        probe.X += velocity.X;
        probe.Y += velocity.Y;

        if (IsPointWithinZone(targetZone, probe))
        {
            return true;
        }

        if (probe.X > targetZone.BottomRight.X || probe.Y < targetZone.BottomRight.Y)
        {
            return false;
        }

        if (velocity.X > 0)
        {
            velocity.X--;
        }

        if (velocity.X < 0)
        {
            velocity.X++;
        }

        velocity.Y--;
    }

    // we will never get here, but it makes the compiler feel better about my code.
    return false;
}

static void PrettyPrint(Zone targetRange, List<Point> steps)
{
    var minX = Math.Min(steps.Min(i => i.X), targetRange.TopLeft.X);
    var maxX = Math.Max(steps.Max(i => i.X), targetRange.BottomRight.X);
    var minY = Math.Min(steps.Min(i => i.Y), targetRange.BottomRight.Y);
    var maxY = Math.Max(steps.Max(i => i.Y), targetRange.TopLeft.Y);

    for (var y = maxY; y >= minY; y--)
    {
        for (var x = minX; x <= maxX; x++)
        {
            if (steps.Exists(s => s.X == x && s.Y == y))
            {
                if (x == 0 && y == 0)
                {
                    Console.Write("S");
                }
                else
                {
                    Console.Write("#");
                }
            }
            else
            {
                if (IsPointWithinZone(targetRange, new Point(x, y)))
                {
                    Console.Write("T");
                }
                else
                {
                    Console.Write(".");
                }
            }
        }

        Console.Write("\n");
    }
}

static bool IsPointWithinZone(Zone zone, Point point)
{
    if (point.X >= zone.TopLeft.X && point.X <= zone.BottomRight.X)
    {
        if (point.Y <= zone.TopLeft.Y && point.Y >= zone.BottomRight.Y)
        {
            return true;
        }
    }

    return false;
}
