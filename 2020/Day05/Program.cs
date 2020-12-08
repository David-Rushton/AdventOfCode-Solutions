using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;


var inputPath = Path.Join(Directory.GetCurrentDirectory(), "input.txt");
var seatIds = new List<int>();


foreach(var input in File.ReadLines(inputPath))
{
    int upperRow = 127;
    int lowerRow = 0;
    int upperCol = 7;
    int lowerCol = 0;

    foreach(var ch in input)
    {
        // take the lower half
        if(ch == 'F')
            upperRow -= (upperRow - lowerRow + 1) / 2;

        // take the upper half
        if(ch == 'B')
            lowerRow += (upperRow - lowerRow + 1) / 2;

        // take the lower half
        if(ch == 'L')
            upperCol -= (upperCol - lowerCol + 1) / 2;

        // take the upper half
        if(ch == 'R')
            lowerCol += (upperCol - lowerCol + 1) / 2;
    }

    Debug.Assert(upperRow == lowerRow, $"Expected upper and lower rows to match: upper {upperRow} | lower {lowerRow}");
    Debug.Assert(upperCol == lowerCol, $"Expected upper and lower cols to match: upper {upperCol} | lower {lowerCol}");

    seatIds.Add(upperRow * 8 + upperCol);
}


seatIds.Sort();


var minSeatId = seatIds[0];
var maxSeatId = seatIds[seatIds.Count -1];
var mySeatId = -1;


// Seat ids are a continuous range of integers (1, 2, 3, ..., n).
// There is one break in the sequence.
// That is my id.
for(var index = 0; index < seatIds.Count; index++)
{
    var currentSeatId = minSeatId + index;
    if(currentSeatId != seatIds[index])
    {
        mySeatId = minSeatId + index;
        Console.WriteLine($"Min SeatId {minSeatId} : My Seat Id {currentSeatId} : Max SeatId {maxSeatId}");
        Environment.Exit(0);
    }
}
