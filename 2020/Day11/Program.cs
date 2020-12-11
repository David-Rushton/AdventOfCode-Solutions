using System;
using System.Collections;
using System.IO;
using System.Linq;


namespace Day11
{
    class Program
    {
        static bool _verboseOutputOn = Environment.GetCommandLineArgs().Contains("--verbose");
        static bool _useTestInput = Environment.GetCommandLineArgs().Contains("--test");
        static string _inputPath = Path.Join
        (
                Directory.GetCurrentDirectory(),
                _useTestInput ? "Test-Input.txt" : "Input.txt"
        );


        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }


        static void VerboseWriteLine(string message)
        {
            if(_verboseOutputOn)
                Console.WriteLine(message);
        }
    }
}
