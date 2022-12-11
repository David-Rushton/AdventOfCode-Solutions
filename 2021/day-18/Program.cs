using System.Diagnostics;
using Day18;

var (isStarTwo, trees) = ArgsParser.Parse(args);

Console.WriteLine($"Config | Tress = {trees.Count} | StarTwo = {isStarTwo}");

if (isStarTwo)
{
    StarTwo(trees);
}
else
{
    StarOne(trees);
}


void StarOne(List<string> trees)
{
    Tree? last = null;
    foreach (var rawTree in trees)
    {
        var current = InputParser.Parse(rawTree);

        if (last is not null)
            current = AddTrees(last, current);

        ProcessPairs(current);
        last = current;
    }

    Debug.Assert(last is not null);
    GetMagnitude(last);
    CliPrinter.Print(last);
}

void StarTwo(List<string> trees)
{
    long maxMagnitude = 0;

    foreach (var left in trees)
    {
        foreach (var right in trees)
        {
            if (left != right)
            {
                var tree = InputParser.Parse($"[{left},{right}]");

                ProcessPairs(tree);
                var magnitude = GetMagnitude(tree);
                CliPrinter.Print(tree);

                if (maxMagnitude < magnitude)
                    maxMagnitude = magnitude;
            }
        }
    }

    Console.WriteLine($"\nMax Magnitude = {maxMagnitude}");
}

void ProcessPairs(Tree tree)
{
    while (ExplodePairs(tree) > 0)
    {
        SplitPairs(tree);
    }
}

int ExplodePairs(Tree tree)
{
    var nodesToProcess = new Queue<Node>(new[] { tree.Root });
    var exploded = 0;

    while (nodesToProcess.Any())
    {
        var node = nodesToProcess.Dequeue();

        if (node.Level > 5)
        {
            exploded++;
            node.ExplodeValue();
            tree.DeleteNode(node);
        }
        else
        {
            // we favour left most, which is a requirement to reach the result
            if (node.Left is not null)
                nodesToProcess.Enqueue(node.Left);

            if (node.Right is not null)
                nodesToProcess.Enqueue(node.Right);
        }
    }

    return exploded;
}

int SplitPairs(Tree tree)
{
    var node = GetNextNode();
    var split = 0;

    while (node is not null)
    {
        split++;

        tree.AddLevel(node);
        node.Left!.Value = node.Value / 2;
        node.Right!.Value = node.Value - node.Left!.Value;
        node.Value = 0;

        if (tree.Nodes.Max(n => n.Level) > 5)
            return split;

        node = GetNextNode();
    }

    return split;

    Node? GetNextNode() =>
        tree.Nodes.Where(n => n.IsLeaf && n.Value > 9).OrderBy(n => n.GetPositionScore()).FirstOrDefault();
}

Tree AddTrees(Tree left, Tree? right)
{
    if (right is null)
        return left;

    left.Nodes.ForEach(n => n.Level++);
    right.Nodes.ForEach(n => n.Level++);

    var result = new Tree { Expression = right.Expression };
    result.Nodes.AddRange(left.Nodes);
    result.Nodes.AddRange(right.Nodes);
    result.Root.Left = left.Root;
    result.Root.Right = right.Root;

    foreach (var node in result.Nodes.Where(n => n.Level == 2))
    {
        node.Parent = result.Root;
    }


    return result;
}


long GetMagnitude(Tree tree)
{
    foreach (var node in tree.Nodes.Where(n => n.IsLeaf))
    {
        node.Magnitude = node.Value;
    }


    var processed = 0;
    do
    {
        processed = 0;
        foreach (var node in tree.Nodes.Where(n => n.Left?.Magnitude is not null && n.Right?.Magnitude is not null && n.Magnitude is null))
        {
            node.Magnitude = (3 * node.Left!.Magnitude) + (2 * node.Right!.Magnitude);
            // if (node.Parent is null)
            // else
            //     node.Magnitude = (node.Left!.Magnitude + node.Right!.Magnitude);

            processed++;
        }
    } while (processed > 0);


    return tree.Root.Magnitude ?? throw new Exception("Cannot calculate magnitude");
}
