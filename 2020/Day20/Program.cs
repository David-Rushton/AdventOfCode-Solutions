using System;
using System.IO;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var useTestData = args.Contains("--test");
            var app = Bootstrap(useTestData);


            app.Solver.FindCorners
            (
                app.TileReader.GetTiles(app.InputPath)
            );
        }


        private static (string InputPath, TileReader TileReader, Solver Solver) Bootstrap(bool useTestData)
        {
            var filename = useTestData ? "Test-Input.txt" : "Input.txt";

            return
            (
                Path.Join(Directory.GetCurrentDirectory(), filename),
                new TileReader(),
                new Solver()
            );
        }
    }
}
