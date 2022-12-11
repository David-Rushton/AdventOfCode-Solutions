using System.Text;
using Spectre.Console;

namespace Day18;

public static class CliPrinter
{
    public static void Print(Day18.Tree tree)
    {
        var outTree = new Spectre.Console.Tree("snailfish");
        var root = outTree.AddNode(GetPrettyExpression($"{tree.Root.Magnitude ?? 0}"));
        var toProcess = new Queue<(Day18.Node inNode, Spectre.Console.TreeNode outNode)>(new[] { (tree.Root, root) });

        while (toProcess.Any())
        {
            var (inNode, outNode) = toProcess.Dequeue();

            if (inNode.Left is not null)
            {
                if (inNode.Left.IsLeaf)
                {
                    outNode.AddNode($"[blue]L{inNode.Level}: {PrettyValue(inNode.Left!.Value)} ({inNode.Left!.Magnitude})[/]");
                }
                else
                {
                    var next = outNode.AddNode($"[yellow]L{inNode.Level} ({inNode.Left!.Magnitude})[/]");
                    toProcess.Enqueue((inNode.Left!, next));
                }
            }

            if (inNode.Right is not null)
            {
                if (inNode.Right!.IsLeaf)
                {
                    outNode.AddNode($"[blue]R{inNode.Level}: {PrettyValue(inNode.Right!.Value)} ({inNode.Right!.Magnitude})[/]");
                }
                else
                {
                    var next = outNode.AddNode($"[yellow]R{inNode.Level} ({inNode.Right!.Magnitude})[/]");
                    toProcess.Enqueue((inNode.Right!, next));
                }
            }
        }


        AnsiConsole.Write(outTree);
        // Console.ReadLine();

        string PrettyValue(int value) =>
            value > 9
                ? $"[white on red]{value}[/]"
                : value.ToString();
    }

    public static string GetPrettyExpression(string expression)
    {
        var level = 0;
        var levelColours = new List<string>(new string[] {"white", "green", "blue", "fuchsia", "slowblink white on red" });
        var result = string.Empty;

        foreach (var character in expression)
        {
            switch (character)
            {
                case '[':
                    var printLevel = level++;
                    printLevel = level > 4 ? 4 : printLevel;
                    result += $"[{levelColours[printLevel]}][[";
                    break;

                case ']':
                    level--;
                    result += "]][/]";
                    break;

                default:
                    result += character;
                    break;
            }
        }

        return result;
    }
}



public record QueueItem(Day18.Node inNode, Spectre.Console.TreeNode outNode);
