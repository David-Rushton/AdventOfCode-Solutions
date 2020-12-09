using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            // (\d\s)?\w+\s\w+\sbag
            // ^[a-z]+\s[a-z]+\sbags\scontain(\s\d\s[a-z]+\s[a-z]+\s(bags|bag)(,|.))+
            // posh green bags contain no other bags.
            var tokeniser = new Tokeniser();
            var graph = new Graph();


            // Populate graph.
            foreach(var token in tokeniser.GetTokens())
            {
                graph.CreateNodeIfNotExists(token.Container);

                if(token.Contained != string.Empty)
                {
                    graph.CreateNodeIfNotExists(token.Contained);
                    graph.CreateEdge(token.Contained, token.Container, token.ContainedCount);
                }
            }

            var sgb = graph.Nodes["shiny gold bag"];
            var containers = new List<string>();
            var contained = 0;

            WalkGraphContainedBy(sgb, ref containers);
            WalkGraphContains(sgb, 1, ref contained);


            // Count includes shiny gold bag
            Console.WriteLine($"{sgb.Key} can eventually be contained by {containers.Count - 1} different bag types");
            Console.WriteLine($"{sgb.Key} can contains {contained} other bags");
        }


        public static void WalkGraphContainedBy(GraphNode node, ref List<string> containers)
        {
            if( ! containers.Contains(node.Key) )
                containers.Add(node.Key);

            foreach(var neighbour in node.GetNeighbours(GraphEdge.ContainedBy))
            {
                WalkGraphContainedBy(neighbour.Neighbour, ref containers);
            }
        }

        public static void WalkGraphContains(GraphNode node, int multiplier, ref int contained)
        {
            foreach(var neighbour in node.GetNeighbours(GraphEdge.Container))
            {
                contained += neighbour.Cost * multiplier;
                WalkGraphContains(neighbour.Neighbour, multiplier + neighbour.Cost, ref contained);
            }
        }
    }
}
