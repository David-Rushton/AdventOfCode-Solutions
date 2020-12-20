using System;
using System.Collections.Generic;
using System.Linq;

namespace AoC
{
    class Program
    {
        const int _bootSequenceCycleCount = 6;


        static void Main(string[] args)
        {
            var testMode = args.Contains("--test");
            var pocketDimension = Bootstrap(testMode);

            pocketDimension.ExecuteBootSequence(_bootSequenceCycleCount);
        }


        private static PocketDimension Bootstrap(bool useTestInput) =>
            new PocketDimension(new Input().GetInput(useTestInput))
        ;
    }
}
