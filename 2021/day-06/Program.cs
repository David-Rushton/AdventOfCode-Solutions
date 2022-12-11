using System.Diagnostics;


var test = args.Contains("--test");
var starTwo = args.Contains("--star-two");
var path = test ? "./Input.Test.txt" : "./Input.txt" ;
var line = File.ReadAllLines(path).First();
var school = new List<int>(line.Split(',').Select(i => int.Parse(i)));
var day = starTwo ? 256 : 80;
var cache = new Dictionary<(int, int), long>();
long totalFish = 0;


foreach (var fish in school)
{
    var batch = GetSchoolSize2(fish, day);
    totalFish += batch;
    Console.WriteLine($"Day {day} | Initial state: {fish} | School size: {batch}");
}

Console.WriteLine($"Day {day} | School Size {totalFish}");


long GetSchoolSize2(int initialTimer, int days)
{
    if (cache.TryGetValue((initialTimer, days), out var result))
        return result;


    var day = days;
    var timer = initialTimer;
    long totalFish = 1;

    while (day > 0)
    {
        day--;

        if (timer == 0)
        {
            totalFish += GetSchoolSize2(8, day);
            timer = 6;
        }
        else
        {
            timer--;
        }
    }

    cache.Add((initialTimer, days), totalFish);
    return totalFish;
}
