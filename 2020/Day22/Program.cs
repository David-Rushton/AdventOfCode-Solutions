using System;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var playCrabCombat = args.Contains("--star-one");
            var components = Bootstrap(args.Contains("--test"));

            if(playCrabCombat)
                components.crabCombat.Play(components.playerOneHand, components.playerTwoHand);
            else
                components.recursiveCombat.Play(components.playerOneHand, components.playerTwoHand);
        }




        private static (CrabCombat crabCombat, RecursiveCombat recursiveCombat, int[] playerOneHand, int[] playerTwoHand) Bootstrap(bool useTest)
        {
            return
            (
                new CrabCombat(),
                new RecursiveCombat(),
                PlayerOneHand(),
                PlayerTwoHand()
            );


            int[] PlayerOneHand() =>
                useTest
                ? new int[] { 9, 2, 6, 3, 1 }
                : new int[] { 28, 3, 35, 27, 19, 40, 14, 15, 17, 22, 45, 47, 26, 13, 32, 38, 43, 24, 29, 5, 31, 48, 49, 41, 25 }
            ;

            int[] PlayerTwoHand() =>
                useTest
                ? new [] { 5, 8, 4, 7, 10 }
                : new [] { 34, 12, 2, 50, 16, 1, 44, 11, 36, 6, 10, 42, 20, 8, 46, 9, 37, 4, 7, 18, 23, 39, 30, 33, 21 }
            ;
        }
    }
}
