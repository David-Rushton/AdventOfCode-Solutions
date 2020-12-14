using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Input
    {
        static long _index = 0;


        public Dictionary<long, long> Get(string inputPath) =>
            File.ReadLines(inputPath)
                .ToArray()[1]
                .Split(',')
                .ToDictionary(k => _index++, v => v)
                .Where(kvp => kvp.Value != "x")
                .ToDictionary(k => k.Key, v => long.Parse(v.Value))
            ;
    }
}
