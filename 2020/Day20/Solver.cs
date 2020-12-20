using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Solver
    {
        public IEnumerable<Tile> FindCorners(IEnumerable<Tile> tiles)
        {
            var countOfEdges = tiles
                .SelectMany
                (   t => new[]
                    {
                        t.Edge.Top, t.Edge.Left, t.Edge.Bottom, t.Edge.Right, t.FlippedEdge.Top,
                        t.FlippedEdge.Left, t.FlippedEdge.Bottom, t.FlippedEdge.Right
                    }
                )
                .GroupBy(g => g )
                .Select(s => new { Key = s.Key, Value = s.Count() } )
                .ToDictionary(k => k.Key, v => v.Value)
            ;

            foreach(var kvp in countOfEdges.OrderBy(kvp => kvp.Value))
                Console.WriteLine($" {kvp.Key} {kvp.Value}");





            return tiles;
        }
    }
}
