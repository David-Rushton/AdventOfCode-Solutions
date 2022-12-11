var starTwo = args.Contains("--star-two");

const string TestInput =
@"5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526";

const string Input =
@"8577245547
1654333653
5365633785
1333243226
4272385165
5688328432
3175634254
6775142227
6152721415
2678227325";

var content = args.Contains("--test")
    ? TestInput
    : Input;
var grid = new Day11.Grid(content);
var flashes = 0;
var step = 1;
var readyToExit = false;

Console.WriteLine("Before any steps:");
grid.PrettyPrint();
Console.WriteLine("");


while (!readyToExit)
{
    // the energy level of each octopus increases by 1
    for (var row = 0; row < 10; row++)
        for (var column = 0; column < 10; column++)
            grid[row, column]++;


    // any octopus with an energy level greater than 9 flashes
    while (grid.ReadyToFlash)
        grid.Flash();


    // any octopus that flashed during this step has its energy reset
    for (var row = 0; row < 10; row++)
    {
        for (var column = 0; column < 10; column++)
        {
            if (grid[row, column] > 9)
            {
                flashes++;
                grid[row, column] = 0;
            }
        }
    }

    if (step < 10 || step % 10 == 0)
    {
        Console.WriteLine($"After step {step}:");
        grid.PrettyPrint();
        Console.WriteLine("");
    }

    if ((step == 100 && !starTwo) || (starTwo && grid.AllFlashed))
        readyToExit = true;
    else
        step++;
}



Console.WriteLine($"Flashes {flashes} | Steps {step}");
