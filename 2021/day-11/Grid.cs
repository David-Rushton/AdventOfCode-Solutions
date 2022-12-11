namespace Day11;

public class Grid
{
    private readonly int[] _cells = new int[100];
    private readonly Queue<(int row, int column)> _readyToFlash = new();


    public Grid(string content)
    {
        var rows = content.Replace("\r", "").Split("\n");

        // let's hope this is a 10x10 string.
        for (var row = 0; row < 10; row++)
            for (var column = 0; column < 10; column++)
                this[row, column] = int.Parse(rows[row][column].ToString());
    }


    public int this[int row, int column]
    {
        get => _cells[ToIndex(row, column)];
        set
        {
            _cells[ToIndex(row, column)] = value;
            if (value == 10)
                _readyToFlash.Enqueue((row, column));
        }
    }

    public bool AllFlashed => _cells.Where(v => v == 0).Count() == 100;

    public bool ReadyToFlash => _readyToFlash.Any();

    public void Flash()
    {
        while (_readyToFlash.Any())
        {
            var (row, column) = _readyToFlash.Dequeue();
            IncrementNeighbours(row, column);
        }
    }

    public void IncrementNeighbours(int row, int column)
    {
        foreach (var rowOffset in new[] {-1, 0, 1})
        {
            foreach (var columnsOffset in new[] {-1, 0, 1})
            {
                var r = row + rowOffset;
                var c = column + columnsOffset;

                if (IsValidAddress(r, c))
                    this[r, c]++;
            }
        }


        bool IsValidAddress(int r, int c) =>
            r is >= 0 and < 10
            && c is >= 0 and < 10
            && !(r == row && c == column);
    }

    public void PrettyPrint()
    {
        for (var row = 0; row < 10; row++)
        {
            for (var column = 0; column < 10; column++)
            {
                Console.Write(_cells[ToIndex(row, column)]);
            }

            Console.WriteLine("");
        }
    }

    private static int ToIndex(int row, int column) => row + (column * 10);
}
