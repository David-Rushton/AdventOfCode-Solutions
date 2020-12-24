using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;


namespace AoC
{
    public class Parser
    {
        public string GetRuleZero(IEnumerable<RuleToken> rules)
        {
            var rulesDictionary = rules.ToDictionary(k => k.Id, v => v);
            var ruleZero = rulesDictionary[0];
            var ruleDependants = GetDependantRules(rulesDictionary, ruleZero);
            var ruleValues = PopulateWithLeafNodes();

            foreach(var rule in ExtractSortedNonLeafNodes())
            {
                var subRuleCount = rule.SubRules.Length;
                var template = subRuleCount == 1 ? "{0}" : "({0}|{1})";
                var subSet1Value = string.Empty;
                var subSet2Value = string.Empty;



                ruleValues.Add
                (
                    rule.Id,
                    string.Format(template, GetSubSetRuleValues(rule, 0), GetSubSetRuleValues(rule, 1))
                );
            }


            foreach(var entry in ruleValues)
                Console.WriteLine($"{entry.Key} {entry.Value}");

            return $"^{ruleValues[0]}$";


            string GetSubSetRuleValues(RuleToken rule, int subRuleIndex)
            {
                var result = string.Empty;

                if(rule.SubRules.Length > subRuleIndex)
                    foreach(var subRuleId in rule.SubRules[subRuleIndex].Value)
                        result += ruleValues[subRuleId];


                return result;
            }

            IEnumerable<RuleToken> ExtractSortedNonLeafNodes() =>
                from rule in ruleDependants
                where rule.Value.rule.IsLeaf == false
                orderby rule.Value.level descending
                select rule.Value.rule
            ;

            Dictionary<int, string> PopulateWithLeafNodes() =>
                rules.Where(r => r.IsLeaf).ToDictionary(k => k.Id, v => v.Value)
            ;
        }


        private Dictionary<int, (int level, RuleToken rule)> GetDependantRules(
            Dictionary<int, RuleToken> rules,
            RuleToken rule
        )
        {
            var dependantRules = new Dictionary<int, (int level, RuleToken rule)>();

            PopulateDependantRules(rules[0], 0);
            PrintDependantRules(dependantRules);
            return dependantRules;


            void PopulateDependantRules(RuleToken rule, int level)
            {
                // if(level > 100)
                //     return;

                AddIfNotExitsToDependantRules(rule, level);

                foreach(var subRuleSet in rule.SubRules)
                    foreach(var ruleId in subRuleSet.Value)
                        PopulateDependantRules(rules[ruleId], level + 1);
            }

            void AddIfNotExitsToDependantRules(RuleToken rule, int level)
            {
                if( ! dependantRules.ContainsKey(rule.Id) )
                    dependantRules.Add(rule.Id, (level, rule));
            }

        }
        private void PrintDependantRules(Dictionary<int, (int level, RuleToken rule)> dependantRules)
        {
            Console.WriteLine("\nRule dependency tree");
            foreach(var rule in dependantRules)
            {
                var ruleV = rule.Value;
                Console.WriteLine(
                    string.Format
                    (
                        "|{0}> {1}: {2}",
                        "".PadLeft(ruleV.level).Replace(' ', '-'),
                        ruleV.rule.Id,
                        ruleV.rule.Value
                    )
                );
            }
            Console.WriteLine();
        }
    }
}
