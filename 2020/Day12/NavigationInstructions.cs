using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    public class NavigationInstructions
    {
        public IEnumerable<NavigationInstruction> GetInstructions(string instructionsPath)
        {
            foreach(var instruction in File.ReadLines(instructionsPath))
            {
                yield return new NavigationInstruction
                (
                    ConvertToDirection(instruction.Substring(0, 1).ToCharArray()[0]),
                    int.Parse(instruction.Substring(1))
                );
            }
        }


        private NavigationDirection ConvertToDirection(char direction) =>
            direction switch
            {
                'N' =>  NavigationDirection.North,
                'S' =>  NavigationDirection.South,
                'E' =>  NavigationDirection.East,
                'W' =>  NavigationDirection.West,
                'L' =>  NavigationDirection.Left,
                'R' =>  NavigationDirection.Right,
                'F' =>  NavigationDirection.Forward,
                _   =>  NavigationDirection.Unknown
            }
        ;
    }
}
