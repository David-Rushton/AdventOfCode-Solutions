using System.IO;

var setIndex = 0;
var locations = new Dictionary<int, List<(string Address, int Steps, int Distance)>>
{
    { 0, new() },
    { 1, new() }
};
foreach (var line in File.ReadAllLines("./Input.txt"))
{
    var items = line.Split(",");
    var location = new Location();

    foreach (var item in items)
    {
        var direction = item[..1].ToLower();
        var units = int.Parse(item[1..]);

        for (var i = 0; i < units; i++)
        {
            switch (direction)
            {
                case "u":
                    location.Top++;
                    break;

                case "r":
                    location.Right++;
                    break;

                case "d":
                    location.Top--;
                    break;

                case "l":
                    location.Right--;
                    break;

                default:
                    throw new Exception($"Unexpected direction: {direction}");
            }

            location.Steps++;
            locations[setIndex].Add((location.Address, location.Steps, location.Distance));
        }
    }

    setIndex++;
}


var closest =
(
    from setOne in locations[0]
    where setOne.Distance != 0
    join setTwo in locations[1] on setOne.Address equals setTwo.Address
    orderby Math.Abs(setOne.Distance)
    select setOne
).First();
Console.WriteLine($"Star one || Closest overlap: {closest.Address} {closest.Distance}");

var shortest =
(
    from setOne in locations[0]
    where setOne.Distance != 0
    join setTwo in locations[1] on setOne.Address equals setTwo.Address
    let totalSteps = setOne.Steps + setTwo.Steps
    orderby totalSteps
    select totalSteps
).First();
Console.WriteLine($"Star two || Shortest overlap: {shortest}");



struct Location
{
    public int Steps { get; set; }
    public int Top { get; set; }
    public int Right { get; set; }
    public int Distance => Math.Abs(Top) + Math.Abs(Right);
    public string Address => $"{Top}x{Right}";
}
