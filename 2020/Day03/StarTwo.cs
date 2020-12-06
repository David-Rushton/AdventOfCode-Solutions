using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public record RunResult(
        int Right,
        int Down,
        int TreesEncountered
    );


    public class StarTwo: IStar
    {
        public void Invoke(List<string> input)
        {
            var runResults = new List<RunResult>();
            (int right, int down)[] routes = new []
            {
                (1, 1),
                (3, 1),
                (5, 1),
                (7, 1),
                (1, 2)
            };

            foreach(var route in routes)
            {
                var treesEncountered = 0;
                var left = 0;

                for(var line = 0; line < input.Count; line = line + route.down)
                {
                    if(input[line][left] == '#')
                        treesEncountered++;

                    left = left + route.right;
                    if(left >= input[line].Length)
                        left = left - input[line].Length;
                }

                runResults.Add(new RunResult(route.right, route.down, treesEncountered));
            }


            Int64 answer = 1;
            foreach(var result in runResults)
            {
                Console.WriteLine(result);
                answer = answer * result.TreesEncountered;
            }
            Console.WriteLine($"\nAnwser: {answer}");
        }
    }
}
