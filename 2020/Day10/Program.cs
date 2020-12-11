using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace Day10
{
    class Program
    {
        static void Main(string[] args)
        {
            var testMode = args.Contains("--test");
            var verboseMode = args.Contains("--verbose");
            var inputPath = Path.Join(Directory.GetCurrentDirectory(), testMode ? "Test-Input.txt" : "Input.txt");
            var sortedAdapters = GetInput(inputPath).OrderBy(a => a);
            var lastAdapter = 0;
            var consecutiveOneJoltJumps = 0;
            var joltJumps = new List<int>();
            var joltJumpsCount = new Dictionary<int, int>()
            {
                {0, 0},
                {1, 0},
                {2, 0},
                {3, 0}
            };


            foreach(var adapter in sortedAdapters)
            {
                var joltJump = adapter - lastAdapter;

                RecordJump(joltJump);
                lastAdapter = adapter;
            }


            // Final adapter is always a +3 step up, and not included in input array
            joltJumpsCount[3]++;
            joltJumps.Add(3);


            Console.WriteLine(string.Format
            (
                @"
Results

Jolt steps:
  0 step jumps {0}
  1 step jumps {1}
  2 step jumps {2}
  3 step jumps {3}

Possible adapter combinations:
{4}

Steps {5}:
1 x 3 step jumps = {6}
                ",
                joltJumpsCount[0],
                joltJumpsCount[1],
                joltJumpsCount[2],
                joltJumpsCount[3],
                CalculateDistinctValidAdapterCombinations(),
                joltJumps.Count(),
                joltJumpsCount[1] * joltJumpsCount[3]
            ));

            Environment.Exit(0);

            IEnumerable<int> GetInput(string inputPath) => File.ReadAllLines(inputPath).Select(i => int.Parse(i));

            void RecordJump(int joltJump)
            {
                // records number of 1 and 3 jolt jumps
                joltJumpsCount[joltJump]++;

                // records consecutive jumps of 1
                if(joltJump == 1)
                {
                    consecutiveOneJoltJumps++;
                }
                else
                {
                    if(consecutiveOneJoltJumps > 0)
                        joltJumps.Add(consecutiveOneJoltJumps);

                    consecutiveOneJoltJumps = 0;
                }
            }

            int CalculateDistinctValidAdapterCombinations()
            {
                // Output is the count of every valid combination.  A combination is valid if it every jump <= 3 jolts.
                // Consecutive sequences of 1 jumps can be compressed, resulting in multiple valid combinations.
                // However consecutive sequences greater than 3 also include invalid combinations.
                // EX: (0), 1, 2, 3, 7 (10): We can remove any of these items, but not all of them.  Because 0 -> 7 is > 3.
                // Invalid combinations follow Bernoulli's Triangle (1, 3, 8, 20, 48, 112...).
                // Where:
                //
                //  | Consecutive 1s  | Possible Combinations | Invalid Combinations | Valid Combinations |
                //  |              1  |                     2 |                    0 |                  2 |
                //  |              2  |                     4 |                    0 |                  4 |
                //  |              3  |                     8 |                    1 |                  7 |
                //  |              4  |                    16 |                    3 |                 13 |
                //  |              5  |                    32 |                    8 |                 24 |

                var validCombinations = 1;

                foreach(var item in joltJumps)
                {
                    // Don't include final jump in calculations
                    // Removing this would create a jump of 4 (invalid)
                    var itemLessOne = item - 1;

                    if(itemLessOne > 0)
                    {
                        var possibleCombinations = (int)Math.Pow(2, itemLessOne);
                        var invalidCombinations = BernoullisTriangle(itemLessOne);
                        validCombinations *= (possibleCombinations - invalidCombinations);

                        if(verboseMode)
                            Console.WriteLine($"-> {itemLessOne}: Possible {possibleCombinations} |Invalid {invalidCombinations} ({validCombinations})");
                    }
                }

                return validCombinations;
            }

            int BernoullisTriangle(int sequenceNumber)
            {
                // Bernoulli's triangle: https://en.wikipedia.org/wiki/Bernoulli%27s_triangle
                // = (n+2)*2^(n-1)
                sequenceNumber -= 3;
                if(sequenceNumber >= 0)
                    return (int)((sequenceNumber + 2) * Math.Pow(2, (sequenceNumber - 1)));

                return 0;
            }
        }
    }
}
