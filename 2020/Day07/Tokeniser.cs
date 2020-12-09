using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;


namespace AoC
{
    public record Input(
        string Parent,
        IEnumerable<string> Children
    );


    public class Tokeniser
    {
        const string _nothingContainedText = " contain no other bags.";

        readonly string _inputPath = Path.Join(Directory.GetCurrentDirectory(), "input.txt");

        readonly Regex _regex = new Regex(@"(\d\s)?\w+\s\w+\sbag");


        public IEnumerable<Token> GetTokens()
        {
            foreach(var input in GetInput())
            {
                yield return new Token(input.container, string.Empty, 0);

                var contained = input.contained.Select
                (
                    c =>
                    {
                        var kv = GetKeyValue(c);
                        return new Token(input.container, kv.key, kv.value);
                    }
                );

                foreach(var item in contained)
                {
                    yield return item;
                }
            }

            (string key, int value) GetKeyValue(string item)
            {
                // Before first space is numeric value
                // After is key
                // Ex: 1 bright white bag
                var firstSpace = item.IndexOf(' ') + 1;
                var key = item.Substring(firstSpace);
                var value = int.Parse(item.Substring(0, firstSpace));

                return (key, value);
            }
        }

        private IEnumerable<(string container, IEnumerable<string> contained)> GetInput()
        {
            // First entry in each line is the parent
            // Subsequent entries are direct children of the parent
            foreach(var line in File.ReadLines(_inputPath))
            {
                var lineNothingContainedTextRemoved = line.Replace(_nothingContainedText, string.Empty);
                var matches = _regex.Matches(lineNothingContainedTextRemoved);

                yield return (matches[0].Value, matches.Skip(1).Select(m => m.Value));
            }
        }
    }
}
