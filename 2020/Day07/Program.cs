using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var graph = BuildGraph();

            var sgb = graph.Nodes["shiny gold bag"];
            var containers = new List<string>();
            var contained = 0;

            WalkGraphContainedBy(sgb, ref containers);
            WalkGraphContains(sgb, 1, ref contained);


            // Count includes shiny gold bag
            Console.WriteLine($"{sgb.Key} can eventually be contained by {containers.Count - 1} different bag types");
            Console.WriteLine($"{sgb.Key} can contains {contained} other bags");
        }


        private static Graph BuildGraph()
        {
            var graph = new Graph();
            var tokeniser = new Tokeniser();

            foreach(var token in tokeniser.GetTokens())
            {
                graph.CreateNodeIfNotExists(token.Container);

                if(token.Contained != string.Empty)
                {
                    graph.CreateNodeIfNotExists(token.Contained);
                    graph.CreateEdge(token.Contained, token.Container, token.ContainedCount);
                }
            }

            return graph;
        }

        private static void WalkGraphContainedBy(GraphNode node, ref List<string> containers)
        {
            if( ! containers.Contains(node.Key) )
                containers.Add(node.Key);

            foreach(var neighbour in node.GetNeighbours(GraphEdge.ContainedBy))
            {
                WalkGraphContainedBy(neighbour.Neighbour, ref containers);
            }
        }

        private static void WalkGraphContains(GraphNode node, int multiplier, ref int contained, int level = 0)
        {
            foreach(var neighbour in node.GetNeighbours(GraphEdge.Container))
            {
                var cost = neighbour.Cost * multiplier;
                contained += cost;

                Console.WriteLine($"{new string(' ', level)} {neighbour.Neighbour.Key} ({neighbour.Cost} x {multiplier} = {cost}) ({contained})");
                WalkGraphContains(neighbour.Neighbour, cost, ref contained, level + 1);
            }
        }
    }
}
