using System;
using System.Collections;
using System.IO;
using System.Linq;


namespace Day11
{
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
            var round = 0;

            VerboseWriteLine("Running simulation");

            while(true)
            {
                var nextGrid = GetNextRound(lastGrid);

                VerboseWriteLine($"  Round {round++}");

                if(lastGrid.FlatBoard == nextGrid.FlatBoard)
                {
                    VerboseWriteLine(nextGrid.Board);
                    Console.WriteLine
                    (
                        string.Format
                        (
                            "\nGrid stabilised\nRound: {0}\nOccupied seat count: {1}",
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

        static bool isSeat(char value) => value != _floor;


        static void WriteLineVerbose(Grid grid) => VerboseWriteLine(grid.Board);
        static void VerboseWriteLine(char[][] message) => VerboseWriteLine(string.Join('\n', message.Select(l => string.Join(string.Empty, l))));
        static void VerboseWriteLine(char[] message) => VerboseWriteLine(string.Join(string.Empty, message));
        static void VerboseWriteLine(char message) => VerboseWriteLine(message.ToString());
        static void VerboseWriteLine(string message)
        {
            if(_verboseOutputOn)
                Console.WriteLine(message);
        }
    }
}
