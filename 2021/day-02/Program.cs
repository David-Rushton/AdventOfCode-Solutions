using System.Diagnostics;

Location location = new();
foreach (var instruction in File.ReadAllLines("./Input.txt"))
{
    var elements = instruction.Split(' ');

    // Console.WriteLine(elements.Length);
    Debug.Assert(elements.Length == 2);

    switch (elements[0])
    {
        case "forward":
            location.Horizontal += int.Parse(elements[1]);
            location.Depth += (location.Aim * int.Parse(elements[1]));
            break;

        case "up":
            location.Aim -= int.Parse(elements[1]);
            break;

        case "down":
            location.Aim += int.Parse(elements[1]);
            break;

        default:
            throw new Exception($"Unexpected direction: {elements[0]}");
    }
}

Console.WriteLine($"Location:");
Console.WriteLine($"\tHorizontal: {location.Horizontal}");
Console.WriteLine($"\tDepth: {location.Depth} ");
Console.WriteLine($"\tAim: {location.Aim} ");
Console.WriteLine($"\tFinal: {location.Horizontal * location.Depth} ");


struct Location
{
    public int Horizontal { get; set; }
    public int Depth { get; set; }
    public int Aim { get; set; }
}
