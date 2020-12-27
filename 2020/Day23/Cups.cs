using System;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    class Cups
    {
        const int _numberOfMovesInGame = 100;


        public void Play(int[] cups)
        {
            Console.WriteLine("Welcome to cups!");

            (Index index, int value) currentCup = (0, cups[0]);

            for(var move = 1; move <= _numberOfMovesInGame; move++)
            {
                int[] buffer;

                PrintMove(cups, currentCup.index, move);
                (cups, buffer)  = PickUpThreeCups(cups, currentCup.index.GetOffset(1));
                var destinationCupIndex = GetDestinationCupIndex(cups, currentCup.value);
                cups = InsertCupsAfter(cups, destinationCupIndex, buffer);
                currentCup = IncrementCurrentCup(cups, currentCup);
                cups = AlignCurrentCupIndexAndValue(cups, currentCup);
            }

            var indexOfOne = Array.IndexOf(cups, 1);
            if(indexOfOne < cups.Length - 1)
                cups = AlignCurrentCupIndexAndValue(cups, (new Index(0), cups[indexOfOne + 1]));

            Console.WriteLine($"\nResult: { string.Join("", cups[..^1]) }");
        }


        private (int[]cups, int[] buffer) PickUpThreeCups(int[] cups, Index pickUpAfter)
        {
            var extractRange = new Range(pickUpAfter.Value + 1, pickUpAfter.Value + 4);
            var buffer = cups.Concat(cups).ToArray()[extractRange];

            Console.WriteLine($"  Picked up: { string.Join(", ", buffer) }");

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

            Console.WriteLine($"  Destinatation cup: {value}");

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
                var lastItem = new int[] { cups[^1] };
                cups = lastItem.Union(cups[0..^1]).ToArray();
            }

            return cups;
        }

        private void PrintMove(int[] cups, Index currentCupIndex, int moveNumber) =>
            Console.WriteLine($"  Current cup index: {currentCupIndex.Value}\n  {moveNumber}: { string.Join(", ", cups) } \n")
        ;
    }
}
