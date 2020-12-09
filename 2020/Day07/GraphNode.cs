using System;
using System.Collections.Generic;


namespace AoC
{
    public class GraphNode
    {
        public GraphNode(string key) =>
            (Key, Neighbours, Relationships, Costs) = (key, new List<GraphNode>(), new List<GraphEdge>(), new List<int>())
        ;


        public (GraphNode Neighbour, GraphEdge Relationship, int Cost) this[int index] =>
            (Neighbours[index], Relationships[index], Costs[index])
        ;

        public string Key { get; init; }

        public List<GraphNode> Neighbours { get; init; }

        public List<GraphEdge> Relationships { get; init; }

        public List<int> Costs { get; init; }

        public int Count => Neighbours.Count;


        public IEnumerable<(GraphNode Neighbour, GraphEdge Relationship, int Cost)> GetNeighbours(GraphEdge relationship)
        {
            for(var i =0; i < Count; i++)
            {
                if(this[i].Relationship == relationship)
                    yield return this[i];
            }
        }
    }
}
