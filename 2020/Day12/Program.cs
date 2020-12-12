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

            var useTestData = Environment.GetCommandLineArgs().Contains("--test");
            var inputPath = Path.Join(Directory.GetCurrentDirectory(), useTestData ? "Test-Input.txt" : "Input.txt");
            var app = Bootstrap(inputPath);
            var result = app.Ferry.PlotCourse(app.Instructions);


            Console.WriteLine(result);
        }


        private static (Ferry Ferry, IEnumerable<NavigationInstruction> Instructions) Bootstrap(string inputPath)
        {
            var navigationInstructions = new NavigationInstructions().GetInstructions(inputPath);
            var ferry = new Ferry();

            return (ferry, navigationInstructions);
        }
    }
}
