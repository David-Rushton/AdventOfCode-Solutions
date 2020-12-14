using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.IO;
using System.Text.RegularExpressions;


namespace AoC
{
    public record Token(
        string Mask,
        int Address,
        long Value
    );


    public class Tokeniser
    {
        readonly Regex _maskRegex = new Regex(@"^mask\s=\s(X|1|0)+$");
        readonly Regex _memAddressValueRegex = new Regex(@"^mem\[\d+]\s=\s\d+$");
        readonly Regex _numbersRegex = new Regex(@"\d+");
        readonly string _inputPath;


        public Tokeniser(string inputPath) => (_inputPath) = (inputPath);


        public IEnumerable<Token> GetTokens()
        {
            var mask = string.Empty;

            foreach(var line in File.ReadLines(_inputPath))
            {
                if(_maskRegex.IsMatch(line))
                {
                    mask = line.Split('=')[1].Trim();
                    continue;
                }


                if(_memAddressValueRegex.IsMatch(line))
                {
                    var numbers = _numbersRegex.Matches(line);

                    Debug.Assert(mask != string.Empty, "Cannot create token:\n\tMask not found");
                    Debug.Assert(numbers.Count == 2, $"Cannot create token:\n\tExpected: mem[x] = y\n\tActual: {line}");

                    yield return new Token(mask, int.Parse(numbers[0].Value), int.Parse(numbers[1].Value));
                    continue;
                }


                // Above nested continues should prevent this block from running
                Debug.Assert(false, $"Cannot create token:\n\tFormat not recognised: {line}");
            }
        }
    }
}
