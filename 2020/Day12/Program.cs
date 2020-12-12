using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    public static class Program
    {
        public static void Main(string[] args)
        {
            var useWaypoint = Environment.GetCommandLineArgs().Contains("--waypoint");
            var useTestData = Environment.GetCommandLineArgs().Contains("--test");
            var inputPath = Path.Join(Directory.GetCurrentDirectory(), useTestData ? "Test-Input.txt" : "Input.txt");
            var app = Bootstrap(inputPath, useWaypoint);
            var result = app.Ferry.PlotCourse(app.Instructions);


            Console.WriteLine
            (
                string.Format
                (
                    "\nResult:\n  North|South: {0}\n  East|West: {1}\n\nManhattan distance: {2}\n",
                    result.NorthSouth,
                    result.EastWest,
                    Math.Abs(result.NorthSouth) + Math.Abs(result.EastWest)
                )
            );
        }


        private static (IFerry Ferry, IEnumerable<NavigationInstruction> Instructions) Bootstrap(string inputPath, bool useWaypoint)
        {
            var navigationInstructions = new NavigationInstructions().GetInstructions(inputPath);
            IFerry ferry = useWaypoint ? new Ferry2() : new Ferry();

            return (ferry, navigationInstructions);
        }
    }
}
