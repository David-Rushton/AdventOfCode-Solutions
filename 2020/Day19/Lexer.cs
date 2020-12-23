using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;


namespace AoC
{
    public class Lexer
    {
        readonly IEnumerable<string> _rulesInput;

        readonly IEnumerable<string> _imagesInput;


        public Lexer(string inputPath) => (_rulesInput, _imagesInput) = ReadRaw(inputPath);


        public IEnumerable<RuleToken> GetRules()
        {
            return
                from rule in _rulesInput
                let containsPipe = rule.Contains('|')
                let numbers = ExtractNumbersFromText(rule).ToArray()
                let isBaseRule = numbers.Length == 1
                let totalRuleCount = numbers.Count() - 1
                let subsetCount = containsPipe ? 2 : 1
                let subsetRuleCount = containsPipe ? totalRuleCount / 2 : totalRuleCount
                select new RuleToken
                (
                    Id: numbers[0],
                    SubRules: GetSubRules(subsetCount, subsetRuleCount, numbers).ToArray(),
                    Value: isBaseRule ? GetBaseRule(rule) : string.Empty
                )
            ;


            IEnumerable<int> ExtractNumbersFromText(string text) =>
                new Regex(@"\d+").Matches(text).Select(i => int.Parse(i.Value))
            ;

            IEnumerable<SubRules> GetSubRules(int subsetCount, int subsetRuleCount, int[] subsetRules)
            {
                yield return new SubRules(subsetRules.Skip(1).Take(subsetRuleCount).ToArray());

                if(subsetRuleCount == 2)
                    yield return new SubRules(subsetRules.Skip(subsetRuleCount + 1).Take(subsetRuleCount).ToArray());
            }

            string GetBaseRule(string rule) => (rule.Contains('a') ? 'a' : 'b').ToString();
        }

        public IEnumerable<string> GetImages() => _imagesInput;


        private (IEnumerable<string> Rules, IEnumerable<String> Images) ReadRaw(string inputPath)
        {
            var input = File.ReadLines(inputPath);

            return
            (
                input.Where(line => TextStartsWithNumber(line)),
                input.Where(line => TextStartsWithAorB(line))
            );


            bool TextStartsWithNumber(string text) => text.Length > 0 && int.TryParse(text.Substring(0, 1), out _);

            bool TextStartsWithAorB(string text) => text.StartsWith('a') || text.StartsWith('b');
        }
    }
}
