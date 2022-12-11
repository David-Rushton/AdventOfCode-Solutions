using System.Diagnostics;

namespace Day18;

public static class ArgsParser
{
    public static (bool isStarTwo, List<string> trees) Parse(string[] args)
    {
        var isStarTwo = false;
        var trees = new List<string>();

        foreach (var arg in args)
        {
            if (arg == "--star-two")
            {
                isStarTwo = true;
                break;
            }

            if (arg.StartsWith('['))
            {
                trees.Add(arg);
            }

            if (File.Exists(arg))
            {
                foreach (var line in File.ReadAllLines(arg))
                {
                    trees.Add(line);
                }
            }
        }

        return (isStarTwo, trees);
    }
}

public static class InputParser
{
    public static Tree Parse(string input)
    {
        var tree = new Tree { Expression = input };
        var current = tree.Root;
        current.Left = new Node();
        current.Right = new Node();

        foreach (var line in input.Replace("\r", string.Empty).Split("\n"))
        {
            Debug.Assert(line.StartsWith('['));
            Debug.Assert(current is not null);

            foreach (var character in line)
            {
                Debug.Assert(current is not null);

                switch (character)
                {
                    case '[':
                        current = tree.AddLevel(current);
                        break;

                    case ']':
                        current = current!.Parent;
                        break;

                    case ',':
                        current = current!.Parent!.Right;
                        break;

                    default:
                        if (int.TryParse(character.ToString(), out var number))
                        {
                            current!.Value = number;
                        }
                        else
                        {
                            throw new Exception($"Unexpected character: {character}");
                        }
                        break;
                }
            }
        }

        return tree;
    }
}

public class Tree
{
    public Tree()
    {
        Root = new() { Level = 1 };
        Nodes = new(new[] { Root });
    }

    public Node Root { get; set; }
    public List<Node> Nodes { get; set; }
    public string Expression { get; set; } = string.Empty;


    public IEnumerable<Node> GetLeafNodes()
    {
        Queue<Node> nodes = new(new[] { Root });

        while (nodes.Any())
        {
            var node = nodes.Dequeue();

            if (node.IsLeaf)
                yield return node;

            if (node.Left is not null && !node.Left.IsLeaf)
                nodes.Enqueue(node.Left);

            if (node.Right is not null && !node.Right.IsLeaf)
                nodes.Enqueue(node.Right);
        }
    }

    public Node AddLevel(Node parent)
    {
        parent.Left = new Node { Level = parent.Level + 1, Parent = parent };
        Nodes.Add(parent.Left);

        parent.Right = new Node { Level = parent.Level + 1, Parent = parent };
        Nodes.Add(parent.Right);

        return parent.Left;
    }

    public void DeleteNode(Node node)
    {
        Debug.Assert(node.IsLeaf);
        Debug.Assert(node.Parent is not null);

        Nodes.Remove(node);

        if (node.Parent.Left == node)
            node.Parent.Left = null;

        if (node.Parent.Right == node)
            node.Parent.Right = null;
    }
}

public class Node
{
    public int Level { get; set; }
    public bool IsLeaf => Left is null && Right is null;
    public int Value { get; set; }
    public long? Magnitude { get; set; } = null;
    public Node? Parent { get; set; }
    public bool IsLeft => this.Parent is null ? false : this.Parent.Left == this;
    public Node? Left { get; set; }
    public Node? Right { get; set; }


    public void ExplodeValue()
    {
        var other = this.IsLeft
            ? FindFirstLeft()
            : FindFirstRight();

        if (other is not null)
        {
            other.Value += this.Value;
        }
    }

    private Node? FindFirstLeft()
    {
        Debug.Assert(this.IsLeaf);
        Debug.Assert(this.Parent is not null);

        Node? current = this.Parent!;
        Node previous = this;

        while (current is not null)
        {
            // keep searching
            if (current.Left == previous)
            {
                previous = current;
                current = current.Parent;
            }
            // we've found a branch further left that the starting point
            else
            {
                Debug.Assert(current is not null);

                current = current.Left;

                while (current is not null)
                {
                    if (current!.IsLeaf)
                        return current;

                    current = current.Right ?? current.Left;
                }
            }
        }

        return null;
    }

    private Node? FindFirstRight()
    {
        Debug.Assert(this.IsLeaf);
        Debug.Assert(this.Parent is not null);

        Node? current = this.Parent!;
        Node previous = this;

        while (current is not null)
        {
            // keep searching
            if (current.Right == previous)
            {
                previous = current;
                current = current.Parent;
            }
            // we've found a branch further left that the starting point
            else
            {
                Debug.Assert(current is not null);

                current = current.Right;

                while (current is not null)
                {
                    if (current!.IsLeaf)
                        return current;

                    current = current.Left ?? current.Right;
                }
            }
        }

        return null;
    }

    public int GetPositionScore()
    {
        var result = 0;
        Node? current = this;

        while (current is not null)
        {
            var multiplier = (current.Level) switch
            {
                1 => 100000000,
                2 => 10000000,
                3 => 1000000,
                4 => 100000,
                5 => 10000,
                6 => 1000,
                7 => 100,
                8 => 10,
                _ => throw new Exception($"Level not supported {current.Level}")
            };

            if (!current.IsLeft)
                multiplier *= 2;

            result += multiplier;
            current = current.Parent;
        }

        return result;
    }
}
