using System;
using System.Collections;
using System.Diagnostics;
using System.IO;
using System.Linq;


namespace Day11
{
    public enum BoardDirection
    {
        Up,
        UpRight,
        Right,
        DownRight,
        Down,
        DownLeft,
        Left,
        UpLeft
    }

    public record Grid(
        char[][] Board,
        string FlatBoard,
        int Rows,
        int Cols
    );


    class Program
    {
        static bool _verboseOutputOn = Environment.GetCommandLineArgs().Contains("--verbose");
        static bool _useTestInput = Environment.GetCommandLineArgs().Contains("--test");
        static bool _useExtendedAdjacency = Environment.GetCommandLineArgs().Contains("--extend--adjacency");
        static int _adjacencyThreshold = _useExtendedAdjacency ? 5 : 4;
        static string _inputPath = Path.Join
        (
                Directory.GetCurrentDirectory(),
                _useTestInput ? "Test-Input.txt" : "Input.txt"
        );
        static char _seatOccupied = '#';
        static char _seatEmpty = 'L';
        static char _floor = '.';


        static void Main(string[] args)
        {
            var lastGrid = GetInput();
            var round = 1;

            WriteLineVerbose("Running simulation");

            while(true)
            {
                var nextGrid = GetNextRound(lastGrid);

                WriteLineVerbose($"  Round {round++}");
                WriteLineVerbose(nextGrid.Board);

                if(lastGrid.FlatBoard == nextGrid.FlatBoard)
                {
                    Console.WriteLine
                    (
                        string.Format
                        (
                            "\nGrid stabilised\nRound: {0}\nOccupied seat count: {1}\n",
                            round,
                            nextGrid.FlatBoard.Count(c => c == _seatOccupied)
                        )
                    );

                    Environment.Exit(0);
                }

                lastGrid = nextGrid;
            }
        }


        static Grid GetInput()
        {
            var grid = File.ReadAllLines(_inputPath).Select(line => line.ToCharArray()).ToArray();
            var lines = File.ReadAllLines(_inputPath);


            return new Grid
            (
                lines.Select(line => line.ToCharArray()).ToArray(),
                string.Join(string.Empty, lines),
                grid.Count(),
                grid[0].Count()
            );
        }


        static Grid GetNextRound(Grid grid)
        {
            var nextBoard = grid.Board.Select(a => a.ToArray()).ToArray();
            var flatBoard = string.Empty;

            for(var r = 0; r < grid.Rows; r++)
            {
                for(var c = 0; c < grid.Cols; c++)
                {
                    var cell = grid.Board[r][c];

                    nextBoard[r][c] = ShouldBeFlipped(grid, r, c) ? FlipSeat(cell) : cell;
                    flatBoard += nextBoard[r][c];
                }
            }


            return new Grid
            (
                nextBoard,
                flatBoard,
                grid.Rows,
                grid.Cols
            );


            char FlipSeat(char value) => value == _seatOccupied ? _seatEmpty : _seatOccupied;
        }

        static bool ShouldBeFlipped(Grid grid, int rowIndex, int colIndex)
        {
            if(isSeat(grid.Board[rowIndex][colIndex]))
            {
                return
                (
                       IsEmptyWithoutOccupiedAdjacentSeats(grid, rowIndex, colIndex)
                    || IsOccupiedAndExceedsOccupiedAdjacencyThreshold(grid, rowIndex, colIndex)
                );
            }

            return false;
        }

        static bool IsEmptyWithoutOccupiedAdjacentSeats(Grid grid, int rowIndex, int colIndex) =>
            grid.Board[rowIndex][colIndex] == _seatEmpty && CountOfOccupiedAdjacentSeats(grid, rowIndex, colIndex) == 0
        ;

        static bool IsOccupiedAndExceedsOccupiedAdjacencyThreshold(Grid grid, int rowIndex, int colIndex) =>
           grid.Board[rowIndex][colIndex] ==
                    _seatOccupied
                &&  CountOfOccupiedAdjacentSeats(grid, rowIndex, colIndex) >= _adjacencyThreshold
        ;

        static int CountOfOccupiedAdjacentSeats(Grid grid, int rowIndex, int colIndex) =>
            _useExtendedAdjacency
                ? CountOfOccupiedAdjacentSeatsExtended(grid, rowIndex, colIndex)
                : CountOfOccupiedAdjacentSeatsSimple(grid, rowIndex, colIndex)
        ;

        static int CountOfOccupiedAdjacentSeatsSimple(Grid grid, int rowIndex, int colIndex)
        {
            var startRow = rowIndex - 1 < 0 ? 0 : rowIndex - 1;
            var endRow = rowIndex + 1 > grid.Rows - 1 ? grid.Rows - 1 : rowIndex + 1;
            var startCol = colIndex - 1 < 0 ? 0 : colIndex - 1;
            var endCol = colIndex + 1 > grid.Cols - 1 ? grid.Cols - 1 : colIndex + 1;
            var occupiedCount = 0;

            for(var r = startRow; r <= endRow; r++)
            {
                for(var c = startCol; c <= endCol; c++)
                {
                    if(! (r == rowIndex && c == colIndex) )
                        if(grid.Board[r][c] ==_seatOccupied)
                            occupiedCount++;
                }
            }


            return occupiedCount;
        }

        static int CountOfOccupiedAdjacentSeatsExtended(Grid grid, int rowIndex, int colIndex)
        {
            var occupiedCount = 0;

            foreach(var direction in (BoardDirection[])Enum.GetValues(typeof(BoardDirection)))
            {
                occupiedCount += ReadFirstSeat(grid, direction, rowIndex, colIndex) == _seatOccupied ? 1 : 0;
            }


            Debug.Assert(occupiedCount <= 8, $"Adjacent count cannot exceed the number of directions checks: {occupiedCount}");
            return occupiedCount;
        }

        static char ReadFirstSeat(Grid grid, BoardDirection direction, int row, int col)
        {
            var value = string.Empty;
            int offset = direction switch
            {
                BoardDirection.Up           => grid.Cols * -1,
                BoardDirection.UpRight      => (grid.Cols - 1) * -1,
                BoardDirection.Right        => 1,
                BoardDirection.DownRight    => grid.Cols + 1,
                BoardDirection.Down         => grid.Cols,
                BoardDirection.DownLeft     => grid.Cols - 1,
                BoardDirection.Left         => -1,
                BoardDirection.UpLeft       => (grid.Cols + 1) * -1,
                _                           => -999
            };
            var lastIndex = ConvertAddressToIndex(grid.Rows - 1, grid.Cols - 1);
            var index = ConvertAddressToIndex(row, col) + offset;
            var address = (row, col);

            while(index >= 0 && index <= lastIndex && isValidMove(direction, address.row, address.col))
            {
                address = ConvertIndexToAddress(index);
                var cell = grid.Board[address.row][address.col];

                if(isSeat(cell))
                    return cell;

                index += offset;
            }


            // If we didn't find a seat then the default value (_seatEmpty).
            return _seatEmpty;


            int ConvertAddressToIndex(int row, int col) => (row * (grid.Cols)) + col;

            (int row, int col) ConvertIndexToAddress(int index) => (index / grid.Cols, index % grid.Cols);

            bool isValidMove(BoardDirection direction, int rowIdx, int colIdx)
            {
                if(rowIdx == 0 && new [] {BoardDirection.Up, BoardDirection.UpRight, BoardDirection.UpLeft}.Contains(direction))
                    return false;

                if(colIdx == grid.Cols - 1 && new [] {BoardDirection.UpRight, BoardDirection.Right, BoardDirection.DownRight}.Contains(direction))
                    return false;

                if(rowIdx == grid.Rows -1 && new [] {BoardDirection.DownRight, BoardDirection.Down, BoardDirection.DownLeft}.Contains(direction))
                    return false;

                if(colIdx == 0 && new [] {BoardDirection.DownLeft, BoardDirection.Left, BoardDirection.UpLeft}.Contains(direction))
                    return false;

                return true;
            }
        }

        static bool isSeat(char value) => value != _floor;


        static void WriteLineVerbose(Grid grid) => WriteLineVerbose(grid.Board);
        static void WriteLineVerbose(char[][] message) => WriteLineVerbose(string.Join('\n', message.Select(l => string.Join(string.Empty, l))));
        static void WriteLineVerbose(char[] message) => WriteLineVerbose(string.Join(string.Empty, message.Append('\n')));
        static void WriteLineVerbose(char message) => WriteLineVerbose(message.ToString());
        static void WriteLineVerbose(string message)
        {
            if(_verboseOutputOn)
                Console.WriteLine(message);
        }
    }
}
