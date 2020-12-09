using System;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var useTestInput = args.Contains("--use-test-input");
            var verbose = args.Contains("--verbose");
            var parser = new Parser();
            var interpreter = new Interpreter();

            var result = interpreter.Invoke(parser.Parse(useTestInput), verbose);
            Console.WriteLine($"Result: {result}");
        }
    }
}
