string[] input = File.ReadAllLines("./Input.txt");
int[] items = input.Select(item => int.Parse(item)).ToArray();

// star one
int increases = 0;
int last = int.Parse(input[0]);
foreach (var item in input.Skip(1))
{
    if (int.Parse(item) > last)
        increases++;

    last = int.Parse(item);
}

Console.WriteLine($"start one: {increases}");


// star two
var current = 0;
var slidingWindow = new List<int>();

while (current + 2 < items.Length)
{

    slidingWindow.Add(items[current] + items[current + 1] + items[current + 2]);
    current++;
}

var windowIncrease = 0;
var lastWindow = slidingWindow[0];
foreach (var currentWindow in slidingWindow)
{
    if (currentWindow > lastWindow)
        windowIncrease++;

    lastWindow = currentWindow;
}

Console.WriteLine($"start two: {windowIncrease}");
