using System.Diagnostics;

namespace Day22;

public enum Polarity
{
    Positive,
    Negative
}

public class Cube
{
    public int MinX { get; init; }
    public int MaxX { get; init; }
    public int MinY { get; init; }
    public int MaxY { get; init; }
    public int MinZ { get; init; }
    public int MaxZ { get; init; }
    public Polarity Polarity { get; init; }
    public long GetVolume()
    {
        var x = MaxX - MinX + 1;
        var y = MaxY - MinY + 1;
        var z = MaxZ - MinZ + 1;
        var volume = Polarity == Polarity.Negative
            ? x * y * z * -1
            : x * y * z;

        Debug.Assert(MinX <= MaxX, $"{MinX}..{MaxX} X");
        Debug.Assert(MinY <= MaxY, $"{MinY}..{MaxY} Y");
        Debug.Assert(MinZ <= MaxZ, $"{MinZ}..{MaxZ} Z");

        return volume;
    }

    public bool Intersects(Cube other)
    {
        if ((other.MinX >= MinX && other.MinX <= MaxX) || (MinX >= other.MinX && MinX <= other.MaxX))
            if ((other.MinY >= MinY && other.MinY <= MaxY) || (MinY >= other.MinY && MinY <= other.MaxY))
                if ((other.MinZ >= MinZ && other.MinZ <= MaxZ) || (MinZ >= other.MinZ && MinZ <= other.MaxZ))
                    return true;

        return false;
    }

    public Cube GetIntersection(Cube other)
    {
        return new Cube
        {
            MinX = Math.Max(MinX, other.MinX),
            MaxX = Math.Min(MaxX, other.MaxX),
            MinY = Math.Max(MinY, other.MinY),
            MaxY = Math.Min(MaxY, other.MaxY),
            MinZ = Math.Max(MinZ, other.MinZ),
            MaxZ = Math.Min(MaxZ, other.MaxZ),
            Polarity = Polarity == Polarity.Positive ? Polarity.Negative : Polarity.Positive
        };
    }

    public override string ToString()
    {
        return $"Cube = {{ x={MinX}..{MaxX}, y={MinY}..{MaxY}, z={MinZ}..{MaxZ}, polarity={Polarity}, volume={GetVolume()} }}";
    }
}
