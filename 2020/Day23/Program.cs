using System;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args) =>
            new Cups().Play
            (
                GetCupsStartingSequence(args.Contains("--test"))
            )
        ;


        private static int[] GetCupsStartingSequence(bool useTestData) =>
            useTestData
            ? new [] { 3, 8, 9, 1, 2, 5, 4, 6, 7 }
            : new [] { 7, 8, 9, 4, 6, 5, 1, 2, 3 }
        ;
    }
}
