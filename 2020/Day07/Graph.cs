using System;
using System.Collections.Generic;


namespace AoC
{
    public class Graph
    {
        public Graph() => (Nodes) = (new Dictionary<string, GraphNode>());


        public Dictionary<string, GraphNode> Nodes { get; init; }


        public GraphNode CreateNodeIfNotExists(string key)
        {
            if( ! Nodes.ContainsKey(key) )
                Nodes.Add(key, new GraphNode(key));

            return Nodes[key];
        }


        public void CreateEdge(string fromKey, string toKey, int cost)
        {
            var from = Nodes[fromKey];
            var to = Nodes[toKey];

            CreateEdge(from, to, cost);
        }

        // Edges are undirected and weighted.
        public void CreateEdge(GraphNode from, GraphNode to, int cost)
        {
            from.Neighbours.Add(to);
            from.Costs.Add(cost);
            from.Relationships.Add(GraphEdge.ContainedBy);

            to.Neighbours.Add(from);
            to.Costs.Add(cost);
            to.Relationships.Add(GraphEdge.Container);
        }
    }
}
