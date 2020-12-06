using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class StarTwo: IStar
    {
        public void Invoke(List<string> input)
        {
            for(var x = 0; x < input.Count; x++)
            {
                for(var y = x + 1; y < input.Count; y++)
                {
                    for(var z = y + 1; z < input.Count; z ++)
                    {
                        var xInt = int.Parse(input[x]);
                        var yInt = int.Parse(input[y]);
                        var zInt = int.Parse(input[z]);

                        if(xInt + yInt + zInt == 2020)
                        {
                            Console.WriteLine($"Result: {xInt} * {yInt} * {zInt} = {xInt * yInt * zInt}");
                            Environment.Exit(0);
                        }
                    }
                }
            }
        }
    }
}
