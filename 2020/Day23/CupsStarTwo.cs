using System;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    class CupsStarTwo
    {
        //const int _numberOfMovesInGame = 10_000_000;
        const int _numberOfMovesInGame = 100;


        public void Play(int[] cups)
        {
            Console.WriteLine("Welcome to cups!");

            (Index index, int value) currentCup = (0, cups[0]);
            cups = SupersizeCups(cups);


            for(var move = 1; move <= _numberOfMovesInGame; move++)
            {
                int[] buffer;

                (cups, buffer)  = PickUpThreeCups(cups, currentCup.index.GetOffset(1));
                var destinationCupIndex = GetDestinationCupIndex(cups, currentCup.value);
                cups = InsertCupsAfter(cups, destinationCupIndex, buffer);
                currentCup = IncrementCurrentCup(cups, currentCup);
                cups = AlignCurrentCupIndexAndValue(cups, currentCup);

                if(move % 1_000 == 0)
                    Console.WriteLine($"  Move # {move.ToString("#,#")}");
            }

            var indexOfOne = Array.IndexOf(cups, 1);
            var valueOne = cups[indexOfOne + 1];
            var valueTwo = cups[indexOfOne + 2];

            Console.WriteLine($"\nResult: {valueOne}x{valueTwo}={valueOne * valueTwo}");
        }


        private int[] SupersizeCups(int[] cups)
        {
            var newCups = new int[1_000_000];

            for(var i = 0; i < 1_000_000; i++)
                newCups[i] = i + 1;

            for(var idx = 0; idx < cups.Length; idx++)
                newCups[idx] = cups[idx];


            return newCups;
        }

        private (int[]cups, int[] buffer) PickUpThreeCups(int[] cups, Index pickUpAfter)
        {
            var extractRange = new Range(pickUpAfter.Value + 1, pickUpAfter.Value + 4);
            var buffer = cups.Concat(cups).ToArray()[extractRange];

            return
            (
                cups.Where(i => ! buffer.Contains(i)).ToArray(),
                buffer
            );
        }

        private Index GetDestinationCupIndex(int[] cups, int valueLessThan)
        {
            var min = cups.Where(i => i < valueLessThan).OrderByDescending(i => i).ToArray();
            var max = cups.Max();
            int value =
                min.Count() > 0
                ? min[0]
                : max
            ;

            return Array.IndexOf(cups, value);
        }

        private int[] InsertCupsAfter(int[] cups, Index insertAfterIndex, int[] toInsert)
        {
            var index = insertAfterIndex.Value + 1;

            return
                cups[..index]
                    .Union(toInsert)
                    .Union(cups[index..])
                    .ToArray()
            ;
        }

        private (Index index, int value) IncrementCurrentCup(int[] cups, (Index index, int value) currentCup)
        {
            var nextIndex =
                currentCup.index.Value >= cups.Length - 1
                ? 0
                : currentCup.index.Value + 1
            ;
            var nextValue = cups.Concat(cups).ToArray()[Array.IndexOf(cups, currentCup.value) + 1];

            return
            (
                nextIndex,
                nextValue
            );
        }

        private int[] AlignCurrentCupIndexAndValue(int[] cups, (Index index, int value) currentCup)
        {
            while(cups[currentCup.index] != currentCup.value)
            {
                var firstItem = new int[] { cups[0] };
                cups = cups[1..].Union(firstItem).ToArray();
            }

            return cups;
        }

        private void PrintMove(int[] cups, Index currentCupIndex, int moveNumber) =>
            Console.WriteLine($"  Current cup index: {currentCupIndex.Value}\n  {moveNumber}: { string.Join(", ", cups) } \n")
        ;
    }
}
