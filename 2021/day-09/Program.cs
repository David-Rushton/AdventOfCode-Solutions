var test = args.Contains("--test");
var starTwo = args.Contains("--star-two");
var path = test ? "./input.test.txt" : "./input.txt";
var cells = File.ReadAllText(path).Replace("\r", "").Split('\n');
var sumOfRisk = 0;
var basinsFound = 0;
var cellsMapped = new Dictionary<(int row, int column), int>();

for (var r = 0; r < cells.Length; r++)
{
    for (var c = 0; c < cells[r].Length; c++)
    {
        var cell = GetCell(r, c);
        var minAdjacent = GetAdjacentCells(r, c).Min();

        if (cell < minAdjacent)
        {
            Console.WriteLine($"We got one | row = {r} | column = {c} | value = {cell}");
            sumOfRisk += cell + 1;
        }

        if (!cellsMapped.ContainsKey((r, c)) && cell != 9)
        {
            MapSurroundingCells(r, c);
            basinsFound++;
        }
    }
}

var basins =
    from basin in cellsMapped
    group basin by basin.Value into basinGroup
    orderby basinGroup.Count() descending
    select basinGroup.Count();
var countPfBasins = basins.Take(3).Aggregate((x, y) => x * y);

Console.WriteLine($"Star one | {sumOfRisk}");
Console.WriteLine($"Star two | {countPfBasins}");



int GetCell(int row, int column)
{
    if (cells is null || row < 0 || row >= cells.Length || column < 0 || column >= cells[row].Length)
        return int.MaxValue;

    return int.Parse(cells[row][column].ToString());
}




void MapSurroundingCells(int row, int column)
{
    if (cellsMapped is null)
        throw new NullReferenceException(nameof(cellsMapped));

    foreach (var cell in GetAdjacentIndexes(row, column).Union(new[] {(row, column)}))
    {
        var cellValue = GetCell(cell.row, cell.column);

        if (cellValue != int.MaxValue && cellValue != 9 && !cellsMapped.ContainsKey((cell.row, cell.column)))
        {
            cellsMapped.Add((cell.row, cell.column), basinsFound);
            MapSurroundingCells(cell.row, cell.column);
        }
    }
}

IEnumerable<int> GetAdjacentCells(int row, int column)
{
    yield return GetCell(row + 1, column);
    yield return GetCell(row - 1, column);
    yield return GetCell(row, column + 1);
    yield return GetCell(row, column - 1);
}

IEnumerable<(int row, int column)> GetAdjacentIndexes(int row, int column)
{
    yield return (row + 1, column);
    yield return (row - 1, column);
    yield return (row, column + 1);
    yield return (row, column - 1);
}
