using System;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var stopwatch = new Stopwatch();
            stopwatch.Start();

            new CupsStarTwo().Play
            (
                GetCupsStartingSequence(args.Contains("--test"))
            );

            stopwatch.Stop();
            Console.WriteLine($"Run time: {stopwatch.Elapsed.Seconds} seconds");
        }

        private static int[] GetCupsStartingSequence(bool useTestData) =>
            useTestData
            ? new [] { 3, 8, 9, 1, 2, 5, 4, 6, 7 }
            : new [] { 7, 8, 9, 4, 6, 5, 1, 2, 3 }
        ;
    }
}
