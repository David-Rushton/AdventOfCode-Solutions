using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    public static class Extensions
    {
        public static List<Instruction> CloneInstructions(this List<Instruction> listToClone) =>
            listToClone.Select(i => i with { }).ToList()
        ;
    }
}
