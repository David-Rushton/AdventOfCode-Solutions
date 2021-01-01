using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    public static class Program
    {
        static (long[] numbers, long[] remainders) _testInput =>
            (
                new long[] { 7, 13, 59, 31, 19 },
                new long[] { 0, 12, 55, 25, 12 }
            )
        ;

        static (long[] numbers, long[] remainders) _input =>
            (
                new long[] { 19, 41, 643,  17,  13,  23, 509,  37,  29 },
                new long[] { 0,  32, 624, -19, -24, -19, 459, -19, -50 }
            )
        ;


        public static void Main(string[] args)
        {
            var components = Bootstrap(args.Contains("--test"));
            var result = components.chineseRemainderTheorem.Solve(components.input.numbers, components.input.remainders);

            Console.WriteLine($"\nResult: {result}");
        }


        private static ((long[] numbers, long[] remainders) input, ChineseRemainderTheorem chineseRemainderTheorem) Bootstrap(bool useTestInput) =>
            (
                useTestInput ? _testInput : _input,
                new ChineseRemainderTheorem()
            )
        ;
    }
}
