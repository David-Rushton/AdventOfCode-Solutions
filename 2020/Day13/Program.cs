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

            app.Engine.Calculate(app.Input);
        }


        private static (Engine Engine, Dictionary<long, long> Input) Bootstrap(string inputPath)
        {
            var input = new Input().Get(inputPath);
            var engine = new Engine();

            return (engine, input);
        }
    }
}
