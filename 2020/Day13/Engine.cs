using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;


namespace AoC
{
    public class Engine
    {
        public void Calculate(Dictionary<long, long> input)
        {
            var firstNumber = input.First();
            var maxNumber = input.Last();           // TODO: TEST ONLY
            var firstMaxDiff = Math.Abs(firstNumber.Value - maxNumber.Value);
            var candidate = firstNumber.Value;
            var candidatesChecked = 0;

            while(true)
            {
                Debug.Assert(candidate % firstNumber.Value == 0, "Bad candidate: {candidate}");
                candidatesChecked++;

                if(IsValidSolution() || candidate >= 1202161486)
                {
                    Console.WriteLine($"\nWinner!\n\n\nCandidate #{candidatesChecked}\nSolution: {candidate}\nProof: {GetProof()}\n");
                    Environment.Exit(0);
                }


                var maxMod = (candidate + maxNumber.Key) % maxNumber.Value;
                var maxOffset = (long)Math.Ceiling((decimal)(maxNumber.Value - maxMod) / firstMaxDiff);

                candidate += firstNumber.Value * maxOffset;
                Verbose.WriteLine($"Candidate #{candidatesChecked}: {candidate.ToString("#,#")}");
            }



            bool IsValidSolution()
            {
                foreach(var item in input.Skip(1))
                {
                    if((item.Key + candidate) % item.Value != 0)
                        return false;
                }

                // All checks passed!
                return true;
            }

            string GetProof() =>
                string.Join(", ", input.Select(kvp => $"{kvp.Value}: {(kvp.Key + candidate) % kvp.Value}"))
            ;
        }
    }
}
