using System;
using System.Linq;


namespace AoC
{
    public static class Verbose
    {
        readonly static bool _showVerobseOutput = Environment.GetCommandLineArgs().Contains("--verbose");


        public static void WriteLine(string message)
        {
            if(_showVerobseOutput)
                Console.WriteLine(message);
        }
    }
}
