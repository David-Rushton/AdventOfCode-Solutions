using System;
using System.Collections.Generic;
using System.Linq;
using AoC;


var groupAnswered = new Dictionary<char, int>();
var groupMembership = 0;
var totalQuestionsEveryoneAnswered  = 0;
var totalQuestionsAnyoneAnswered = 0;

foreach(var token in new Tokeniser().GetTokens())
{
    if(token.Type == TokenTypes.EndOfGroup)
    {
        // Collect stats
        totalQuestionsAnyoneAnswered += groupAnswered.Count;
        totalQuestionsEveryoneAnswered += groupAnswered.Count(i => i.Value == groupMembership);

        // reset
        groupMembership = 0;
        groupAnswered.Clear();
    }

    if(token.Type == TokenTypes.Answer)
    {
        groupMembership++;

        foreach(var ch in token.Value)
        {
            if( ! groupAnswered.ContainsKey(ch) )
                groupAnswered.Add(ch, 1);
            else
                groupAnswered[ch]++;
        }
    }
}


Console.WriteLine($"Sum of group questions anyone answered: {totalQuestionsAnyoneAnswered}");
Console.WriteLine($"Sum of group questions everyone answered: {totalQuestionsEveryoneAnswered}");
