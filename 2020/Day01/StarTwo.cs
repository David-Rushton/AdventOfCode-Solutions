using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace Day01
{
    public record CandidateTwo (int x, int y, int z);

    public class StarTwo
    {
        public void Invoke()
        {
            var candidates = new List<CandidateTwo>();
            var path = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");
            var lines = File.ReadLines(path).ToList<string>();
            var xLineNumber = 0;
            var yLineNumber = 0;
            var zLineNumber = 0;


            Console.WriteLine("\nSeaching for candidates...");
            foreach(var x in lines)
            {
                xLineNumber++;
                yLineNumber=0;

                foreach(var y in lines)
                {
                    yLineNumber++;
                    zLineNumber=0;

                    foreach(var z in lines)
                    {
                        var xInt = int.Parse(x);
                        var yInt = int.Parse(y);
                        var zInt = int.Parse(z);

                        zLineNumber++;

                        if((xLineNumber != yLineNumber) && (xLineNumber != zLineNumber))
                        {
                            if(xInt + yInt + zInt == 2020)
                            {
                                var values = new [] {xInt, yInt, zInt};
                                Array.Sort(values);

                                Console.WriteLine($"Candidate found: {x} {y} {z}");
                                var newCandidate = new CandidateTwo
                                (
                                    values[0],
                                    values[1],
                                    values[2]
                                );

                                if( ! candidates.Contains(newCandidate) )
                                    candidates.Add(newCandidate);
                            }
                        }

                    }
                }
            }

            Console.WriteLine("\nResults...");
            foreach(var result in candidates)
            {
                Console.Write($"{result.x} {result.y} {result.z} = {result.x * result.y * result.z}");
            }
        }
    }
}
