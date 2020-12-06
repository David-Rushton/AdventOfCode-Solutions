using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class StarOne: IStar
    {
        public void Invoke(List<string> input)
        {
            var treesEncountered = 0;
            var left = 0;

            foreach(var line in input)
            {
                Console.WriteLine($"{left} {line}");
                if(line[left] == '#')
                    treesEncountered++;


                // Next move is right 3, down 1.
                left = left + 3;
                if(left >= line.Length)
                    left = left - line.Length;
            }

            Console.WriteLine($"\nResult: {treesEncountered}");
        }
    }
}
