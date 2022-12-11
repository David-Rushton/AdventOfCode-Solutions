namespace Day12;

public readonly record struct Cave
{
    public Cave(string name)
    {
        Name = name;
        AtMostOnce = name.ToLower() == name;
        MaybeTwice = name.ToLower() == name && name != "start" && name != "end";
    }


    public string Name { get; init; }
    public bool AtMostOnce { get; init; }
    public bool MaybeTwice { get; init; }
}
