namespace Day16;

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
/*
 1  function Dijkstra(Graph, source):
 2
 3      create vertex set Q
 4
 5      for each vertex v in Graph:
 6          dist[v] ← INFINITY
 7          prev[v] ← UNDEFINED
 8          add v to Q
 9      dist[source] ← 0
10
11      while Q is not empty:
12          u ← vertex in Q with min dist[u]
13
14          remove u from Q
*/
public static class PathFinder
{
    public static void DijkstraSearch(List<Cell> grid, bool test)
    {
        var unvisited = new List<Cell>(grid);
        var priorityQueue = new Queue<Cell>(grid.Where(c => c.IsDeparture));
        var iteration = 0;

        while (priorityQueue.Any())
        {
            var current = priorityQueue.Dequeue();
            var up = unvisited.Where(c => c.Row == current.Row - 1 && c.Column == current.Column).FirstOrDefault();
            var down = unvisited.Where(c => c.Row == current.Row + 1 && c.Column == current.Column).FirstOrDefault();
            var left = unvisited.Where(c => c.Row == current.Row && c.Column == current.Column + 1).FirstOrDefault();
            var right = unvisited.Where(c => c.Row == current.Row && c.Column == current.Column - 1).FirstOrDefault();


            if (up is not null)
            {
                up.Distance = GetDistance(current, up);
            }

            if (down is not null)
            {
                down.Distance = GetDistance(current, down);
            }

            if (left is not null)
            {
                left.Distance = GetDistance(current, left);
            }

            if (right is not null)
            {
                right.Distance = GetDistance(current, right);
            }

            current.Visited = true;

            if (test || iteration++ % 1000 == 0)
            {
                Console.WriteLine($"Visited cell {current.Row}x{current.Column} ({current.Distance})");
            }

            if (current.IsDestination)
            {
                Console.WriteLine($"Shortest path = {current.Distance}");
                Environment.Exit(0);
            }

            unvisited.Remove(current);

            var next = unvisited.Where(c => !c.Visited).OrderBy(c => c.Distance).FirstOrDefault();
            if (next is not null)
            {
                priorityQueue.Enqueue(next);
            }
        }


        static int GetDistance(Cell visited, Cell unvisited)
        {
            var distance = visited.Distance + unvisited.Value;
            return distance < unvisited.Distance ? distance : unvisited.Distance;
        }
    }
}


public class Cell
{
    public int Row { get; set; }
    public int Column { get; set; }
    public int Value { get; set; }
    public int Distance { get; set; } = int.MaxValue;
    public bool Visited { get; set; }
    public bool IsDeparture { get; set;  }
    public bool IsDestination { get; set; }
}
