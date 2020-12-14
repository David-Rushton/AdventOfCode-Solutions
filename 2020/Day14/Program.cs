using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var useTestInput = args.Contains("--test");
            var showVerboseOutput = args.Contains("--verbose");
            var interpreter = Bootstrap(useTestInput, showVerboseOutput);

            interpreter.Invoke();
        }


        private static Interpreter Bootstrap(bool useTestInput, bool showVerboseMode)
        {
            Verbose.ShowVerboseOutput = showVerboseMode;

            var inputPath = Path.Join(Directory.GetCurrentDirectory(), useTestInput ? "Test-Input.txt" : "Input.txt");
            var tokeniser = new Tokeniser(inputPath);
            var converter = new Converter();
            var interpreter = new Interpreter(tokeniser, converter);

            return interpreter;
        }
    }
}
