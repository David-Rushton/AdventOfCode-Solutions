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


            Bootstrap();
        }


        private static void Bootstrap()
        {
            // no-op
        }
    }
}
