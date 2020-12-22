using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public record TileEdge(char[] Pattern)
    {
        public int GetValue()
        {
            int result = 0;

            for(var i = 0; i < this.Pattern.Length; i++)
                if(this.Pattern[i] == '#')
                    result += (int)Math.Pow(i + 1, 2);


            return result;
        }
    }


    public record Tile
    (
        int Id,
        TileEdge Top,
        TileEdge Left,
        TileEdge Bottom,
        TileEdge Right,
        bool IsFlipped
    )
    {
        public Tile(int id, char[] top, char[] left, char[] bottom, char[] right)
            :this(id, new TileEdge(top), new TileEdge(left), new TileEdge(bottom), new TileEdge(right), false)
        { }


        public Tile GetFlipped() =>
            this with
            {
                Top = new TileEdge(this.Bottom.Pattern.Reverse().ToArray()),
                // Left = new TileEdge(this.Right.Pattern.Reverse().ToArray()),
                Bottom = new TileEdge(this.Top.Pattern.Reverse().ToArray()),
                // Right = new TileEdge(this.Left.Pattern.Reverse().ToArray()),
                IsFlipped = ! this.IsFlipped
            }
        ;
    };
}
