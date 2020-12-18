using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using AoC.Tokeniser;

namespace AoC
{
    public class FieldDetector
    {
        const int _unsetValidation = 0;
        const int _passesValidation = 11;
        const int _failsValidation = 99;


        public void Invoke(TicketToken myTicket, List<TicketToken> passengerTickets, List<RuleToken> rules)
        {
            Debug.Assert(passengerTickets.Count > 0, "Where are the tickets!?");
            Debug.Assert(rules.Count > 0, "Where are the rules man!?");
            Debug.Assert(rules.Count == passengerTickets[0].Fields.Count(), "Count of rules and fields should always be equal");

            var fieldCount = rules.Count;
            var fieldMatrix = new bool[fieldCount, fieldCount];


            // Count of fields and rules will always match
            // Because assert above and
            // This is a field to rule mapping exercise
            for(int rule = 0; rule < fieldCount; rule++)
                for(int field = 0; field < fieldCount; field++)
                    foreach(var ticket in passengerTickets)
                    {
                        if( ! PassesFieldRules(ticket.Fields[field], rules[rule]) )
                        {
                            fieldMatrix[rule, field] = true;
                            break;
                        }
                    }

            // How did we do?
            var candidates = GetCandidates(fieldMatrix, rules);

            PrintFieldMatrix(fieldMatrix, rules);
            PrintCandidates(candidates);
            PrintMyTicket(myTicket, candidates);


            bool PassesFieldRules(int value, RuleToken rule) =>
                   ( value >= rule.FirstRangeLowerBound  && value <= rule.FirstRangeUpperBound  )
                || ( value >= rule.SecondRangeLowerBound && value <= rule.SecondRangeUpperBound )
            ;
        }


        private Dictionary<RuleToken, int> GetCandidates(bool[,] fieldMatrix, List<RuleToken> rules)
        {
            var rulesCount = rules.Count;
            var fieldMap = new Dictionary<RuleToken, int>();
            var rulePossibleMatches = GetRulePossibleMatches();
            var updates = 1;

            while(updates > 0)
            {
                updates = 0;

                foreach(var match in rulePossibleMatches.Where(r => r.Value.Count == 1))
                {
                    var fieldNumber = match.Value[0];
                    fieldMap.Add(match.Key, fieldNumber);

                    foreach(var ruledOut in rulePossibleMatches.Where(r => r.Value.Contains(fieldNumber)))
                        ruledOut.Value.Remove(fieldNumber);

                    rulePossibleMatches.Remove(match.Key);
                    updates++;
                }
            }


            return fieldMap;


            Dictionary<RuleToken, List<int>> GetRulePossibleMatches()
            {
                var result = new Dictionary<RuleToken, List<int>>();

                for(var rule = 0; rule < rulesCount; rule++)
                {
                    result.Add(rules[rule], new List<int>());

                    for(var field = 0; field < rulesCount; field++)
                        if(fieldMatrix[rule, field] == false)
                            result[rules[rule]].Add(field);
                }

                return result;
            }
        }

        private void PrintCandidates(Dictionary<RuleToken, int> candidates)
        {
            foreach(var candidate in candidates.OrderBy(c => c.Key.FieldName))
                Console.WriteLine($"{candidate.Key.FieldName.PadRight(20)}: {candidate.Value}");

            Console.WriteLine('\n');
        }


        private void PrintFieldMatrix(bool[,] fieldMatrix, List<RuleToken> rules)
        {
            PrintHeaderRow(true, rules.Count);
            PrintHeaderRow(false, rules.Count);
            PrintDetailRows();
            PrintHeaderRow(false, rules.Count);
            Console.WriteLine('\n');


            void PrintHeaderRow(bool useFieldNames, int fieldCount)
            {
                Console.Write($"{"".PadRight(20)} | ");
                for(var f = 0; f < fieldCount; f++)
                    Console.Write
                    (
                        string.Format
                        (
                            "{0}{1} | ",
                            useFieldNames ? "f" : "-",
                            useFieldNames ? f.ToString("000") : "---"
                        )
                    );

                Console.WriteLine();
            }

            void PrintDetailRows()
            {
                for(var row = 0; row < rules.Count; row++)
                {
                    Console.Write($"{ rules[row].FieldName.PadRight(20) } | ");

                    for(var col = 0; col < rules.Count; col++)
                        Console.Write
                        (
                            string.Format
                            (
                                " {0} | ",
                                fieldMatrix[row, col] ? "xx " : "   "
                            )
                        );

                    Console.WriteLine();
                }
            }
        }


        void PrintMyTicket(TicketToken myTicket, Dictionary<RuleToken, int> candidates)
        {
            var departureFields = candidates.Where(f => f.Key.FieldName.StartsWith("departure ")).ToDictionary(k => k.Key.FieldName, v => v.Value);
            var result = departureFields.Sum(f => myTicket.Fields[f.Value]);
            var ticket = @"
.--------------------------------------------------------.
| ????: 101    ?????: 102   ??????????: 103  result: {0} |
|                                                        |
| departure date    : {1}    departure station: {2}      |
| departure location: {3}    departure time   : {4}      |
| departure platform: {5}    departure track  : {6}      |
|                                                        |
| ??: 401  ??: 402           ???? ????: 403    ????????? |
'--------------------------------------------------------'
";

            Console.WriteLine(string.Format
            (
                ticket,
                result,
                departureFields["departure date"].ToString("000"),
                departureFields["departure station"].ToString("000"),
                departureFields["departure location"].ToString("000"),
                departureFields["departure time"].ToString("000"),
                departureFields["departure platform"].ToString("000"),
                departureFields["departure track"].ToString("000")
            ));
            Console.WriteLine('\n');
        }
    }
}
