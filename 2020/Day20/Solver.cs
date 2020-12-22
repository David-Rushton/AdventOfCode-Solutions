using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Solver
    {
        public void FindCorners(IEnumerable<Tile> tiles)
        {
            var solutionFound = false;
            var numberOfAttempts = 1;
            var maximumNumberOfAttempts = 15;
            var tileSet = tiles.ToDictionary(k => k.Id, v => v);

            while(solutionFound == false)
            {
                Console.WriteLine($"Attempt #{numberOfAttempts}");

                var scoredTiles = ScoreTiles(tileSet.Values);
                var suspectTiles = scoredTiles
                    .Where(t => t.Score.Top is > 1 || t.Score.Left > 1 || t.Score.Bottom is > 1 || t.Score.Right > 1)
                    .Where(t => t.Score.Top + t.Score.Left + t.Score.Bottom + t.Score.Right > 0)
                    .Where(t => t.Tile.IsFlipped == false)
                    .OrderBy(t => t.Score.Top + t.Score.Left + t.Score.Bottom + t.Score.Right)
                    .ToList()
                ;

                foreach(var tile in scoredTiles)
                    Console.WriteLine($"  {tile.Tile.Id} {tile.Score.Top}x{tile.Score.Left}x{tile.Score.Bottom}x{tile.Score.Right} ({string.Join("", tile.Tile.Top.Pattern)}) Flipped: {tile.Tile.IsFlipped}");



                if(suspectTiles.Count == 0)
                {
                    Console.WriteLine("WINNER!!");
                    solutionFound = true;
                }




                var worstOffender = suspectTiles.First();
                tileSet[worstOffender.Tile.Id] = worstOffender.Tile.GetFlipped();


                ExitIfExceededMaximumNumberOfAttempts();
            }





            void ExitIfExceededMaximumNumberOfAttempts()
            {
                if(numberOfAttempts++ > maximumNumberOfAttempts)
                    Environment.Exit(1);
            }
        }



        private IEnumerable<(Tile Tile, (int Top, int Left, int Bottom, int Right) Score)> ScoreTiles(IEnumerable<Tile> tiles)
        {
            var countOfEdges = tiles
                .SelectMany(t => new [] { t.Top.GetValue(), t.Left.GetValue(), t.Bottom.GetValue(), t.Right.GetValue() } )
                .GroupBy(g => g)
                .Select(s => new { Key = s.Key, Value = s.Count() } )
                .ToDictionary(k => k.Key, v => v.Value)
            ;

            foreach(var tile in tiles)
                yield return
                (
                    tile,
                    (
                        countOfEdges[tile.Top.GetValue()] -1,
                        countOfEdges[tile.Left.GetValue()] -1,
                        countOfEdges[tile.Bottom.GetValue()] -1,
                        countOfEdges[tile.Right.GetValue()] -1
                    )
                );
        }



    }
}
