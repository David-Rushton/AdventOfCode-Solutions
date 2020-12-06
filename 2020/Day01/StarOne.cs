using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace Day01
{
    public class StarOne: IStar
    {
        public void Invoke()
        {
            var path = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");
            var input = File.ReadLines(path).ToList<string>();


            for(var x = 0; x < input.Count; x++)
            {
                for(var y = x + 1; y < input.Count; y++)
                {
                    var xInt = int.Parse(input[x]);
                    var yInt = int.Parse(input[y]);

                    if(xInt + yInt == 2020)
                    {
                        Console.WriteLine($"Result: {xInt} * {yInt} = {xInt * yInt}");
                        Environment.Exit(0);
                    }
                }
            }
        }
    }
}
