using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    public class Input
    {
        private string _isActive = "#";

        private string _isInactive = ".";


        public IEnumerable<ConwayCube> GetInput(bool useTestInput) =>
            ConvertFromRawInput
            (
                useTestInput
                ? GetRawTestInput()
                : GetRawInput()
            )
        ;


        private IEnumerable<ConwayCube> ConvertFromRawInput(string[] rawInput)
        {
            for(var row = 0; row < rawInput.Length; row++)
                for(var col = 0; col < rawInput[row].Length; col++)
                    yield return new ConwayCube(col, row, 0, 0, getIsActive(rawInput[row], col));


            bool getIsActive(string row, int col) => row.Substring(col, 1) == _isActive;
        }

        private string[] GetRawInput() => new []
            {
                "......##",
                "####.#..",
                ".##....#",
                ".##.#..#",
                "........",
                ".#.#.###",
                "#.##....",
                "####.#.."
            }
        ;

        private string[] GetRawTestInput() => new []
            {
                ".#.",
                "..#",
                "###"
            }
        ;
    }
}
