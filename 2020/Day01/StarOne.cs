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
