using System.Text;

namespace Day04;

public enum CheckNumberResult
{
    /// <summary>
    /// Number not found.
    /// </summary>
    NotFound,

    /// <summary>
    /// Number found, but did not complete a row or column.
    /// </summary>
    Found,

    /// <summary>
    /// Number completed a row or column.
    /// </summary>
    Winner
}


public class Board
{
    private readonly Dictionary<int, int> _rowScore = new()
    {
        { 0, 0 },
        { 1, 0 },
        { 2, 0 },
        { 3, 0 },
        { 4, 0 }
    };
    private readonly Dictionary<int, int> _columnScore = new()
    {
        { 0, 0 },
        { 1, 0 },
        { 2, 0 },
        { 3, 0 },
        { 4, 0 }
    };
    private readonly Dictionary<int, Cell> _cells = new();
    private int rowsAdded;


    public bool IsCompleted { get; private set; }


    public void AddRow(int[] numbers)
    {
        for (var i = 0; i < numbers.Length; i++)
        {
            // Assumption: numbers are not duplicated within any board 🤷.
            _cells.Add(numbers[i], new Cell() { Value = numbers[i], Row = rowsAdded, Column = i });
        }

        rowsAdded++;
    }

    public CheckNumberResult CheckForNumber(int number)
    {
        if (_cells.TryGetValue(number, out var cell))
        {
            _cells[number] = cell with { Marked = true };
            _rowScore[cell.Row] += 1;
            _columnScore[cell.Column] += 1;

            // US rules = 5 x 6 grid
            if (_rowScore[cell.Row] == 5 || _columnScore[cell.Column] == 5)
            {
                IsCompleted = true;
                return CheckNumberResult.Winner;
            }

            return CheckNumberResult.Found;
        }

        return CheckNumberResult.NotFound;
    }

    public int GetSumOfUnmarked() =>
        (
            from cell in _cells
            where !cell.Value.Marked
            select cell.Key
        ).Sum();


    public override string ToString()
    {
        StringBuilder sb = new();
        foreach (var cell in _cells)
        {
            sb.Append(cell.Value.ToString());

            if (cell.Value.Column == 4)
                sb.Append('\n');
        }

        return sb.ToString();
    }



    private readonly struct Cell
    {
        public int Value { get; init; }
        public int Row { get; init; }
        public int Column { get; init; }
        public bool Marked { get; init; }

        public override string ToString()
        {
            var paddedValue = Value.ToString("###").PadLeft(3);
            return Marked
                ? $"\u001b[36m{paddedValue}\u001b[0m"
                : paddedValue;
        }
    }
}
