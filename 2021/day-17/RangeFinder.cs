namespace Day17;

public static class RangeFinder
{
    public static Zone GetFiringZone(Zone targetZone)
    {
        var minX = 1;
        var next = 0;
        while (minX < targetZone.TopLeft.X)
        {
            minX += next;
            next++;
        }


        return new Zone(
            TopLeft: new Point(
                next - 1,
                Math.Abs(targetZone.BottomRight.Y) - 1),
            BottomRight: new Point(
                targetZone.BottomRight.X,
                targetZone.BottomRight.Y)
        );
    }
}



public record struct Point(int X, int Y);

public record struct Velocity(int X, int Y);

public record struct Zone(Point TopLeft, Point BottomRight);
