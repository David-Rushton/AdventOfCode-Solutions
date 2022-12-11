using System.IO;

namespace Day04;

public class InputReader
{
    private const string Path = "./Input.txt";
    private const string TestPath = "./Input.Test.txt";


    public static (int[] numbers, List<Board> boards) Read(bool testInput)
    {
        var lines = File.ReadAllLines(testInput ? TestPath : Path);
        var numbers = lines[0].Split(',').Select(i => int.Parse(i)).ToArray();

        List<Board> boards = new();
        bool startNewBoard = true;
        Board currentBoard = new Board();
        foreach(var line in lines.Skip(2))
        {
            // first line contains numbers in order they wil be drawn.
            // Boards are then returned with blank lines between them.

            if (string.IsNullOrWhiteSpace(line))
            {
                startNewBoard = true;
                boards.Add(currentBoard);
            }
            else
            {
                if (startNewBoard)
                {
                    startNewBoard = false;
                    currentBoard = new Board();
                }

                int[] row = line
                    .Split(' ', StringSplitOptions.RemoveEmptyEntries)
                    .Select(i => int.Parse(i))
                    .ToArray();
                currentBoard.AddRow(row);
            }
        }

        boards.Add(currentBoard);

        return (numbers, boards);
    }
}
