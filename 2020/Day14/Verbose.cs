using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public static class Verbose
    {
        public static bool ShowVerboseOutput { get; set; }


        public static void WriteLine(string message)
        {
            if(ShowVerboseOutput)
                Console.WriteLine(message);
        }
    }
}
