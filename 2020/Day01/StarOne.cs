using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace Day01
{
    public record Candidate (int x, int y);

    public class StarOne
    {
        public void Invoke()
        {
            var candidates = new List<Candidate>();
            var path = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");
            var lines = File.ReadLines(path).ToList<string>();
            var outerLineNumber = 0;
            var innerLineNumber = 0;


            Console.WriteLine("\nSeaching for candidates...");
            foreach(var outer in lines)
            {
                outerLineNumber++;
                innerLineNumber=0;

                foreach(var inner in lines)
                {
                    var x = int.Parse(outer);
                    var y = int.Parse(inner);

                    innerLineNumber++;
                    if(! (outerLineNumber == innerLineNumber) )
                        if(x + y == 2020)
                        {
                            Console.WriteLine($"Candidate found: {outerLineNumber}:{outer} x {innerLineNumber}:{inner}");

                            var newCandidate = new Candidate
                            (
                                x < y ? x : y,
                                x < y ? y : x
                            );

                            if(! candidates.Contains(newCandidate) )
                                candidates.Add(newCandidate);
                        }
                }
            }


            Console.WriteLine("\nResults:");
            foreach(var result in candidates)
            {
                Console.WriteLine($"{result.x} * {result.y} = {result.x * result.y}");
            }
        }
    }
}
