using System;
using System.Collections;
using System.IO;
using System.Linq;


namespace AoC
{
    public enum NavigationDirection
    {
        North,
        South,
        East,
        West,
        Left,
        Right,
        Forward,
        Unknown
    }


    public record NavigationInstruction(
        NavigationDirection Direction,
        int Value
    );
}
